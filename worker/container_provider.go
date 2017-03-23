package worker

import (
	"errors"
	"fmt"
	"os"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/baggageclaim"
)

var ErrCreatedContainerNotFound = errors.New("container in created-state-not-found-in-garden")

const creatingContainerRetryDelay = 1 * time.Second

//go:generate counterfeiter . ContainerProviderFactory

type ContainerProviderFactory interface {
	ContainerProviderFor(Worker) ContainerProvider
}

type containerProviderFactory struct {
	gardenClient            garden.Client
	baggageclaimClient      baggageclaim.Client
	volumeClient            VolumeClient
	imageFactory            ImageFactory
	dbVolumeFactory         dbng.VolumeFactory
	dbResourceCacheFactory  dbng.ResourceCacheFactory
	dbResourceConfigFactory dbng.ResourceConfigFactory
	dbTeamFactory           dbng.TeamFactory

	db GardenWorkerDB

	httpProxyURL  string
	httpsProxyURL string
	noProxy       string

	clock clock.Clock
}

func NewContainerProviderFactory(
	gardenClient garden.Client,
	baggageclaimClient baggageclaim.Client,
	volumeClient VolumeClient,
	imageFactory ImageFactory,
	dbVolumeFactory dbng.VolumeFactory,
	dbResourceCacheFactory dbng.ResourceCacheFactory,
	dbResourceConfigFactory dbng.ResourceConfigFactory,
	dbTeamFactory dbng.TeamFactory,
	db GardenWorkerDB,
	httpProxyURL string,
	httpsProxyURL string,
	noProxy string,
	clock clock.Clock,
) ContainerProviderFactory {
	return &containerProviderFactory{
		gardenClient:            gardenClient,
		baggageclaimClient:      baggageclaimClient,
		volumeClient:            volumeClient,
		imageFactory:            imageFactory,
		dbVolumeFactory:         dbVolumeFactory,
		dbResourceCacheFactory:  dbResourceCacheFactory,
		dbResourceConfigFactory: dbResourceConfigFactory,
		dbTeamFactory:           dbTeamFactory,
		db:                      db,
		httpProxyURL:            httpProxyURL,
		httpsProxyURL:           httpsProxyURL,
		noProxy:                 noProxy,
		clock:                   clock,
	}
}

func (f *containerProviderFactory) ContainerProviderFor(worker Worker) ContainerProvider {
	return &containerProvider{
		gardenClient:            f.gardenClient,
		baggageclaimClient:      f.baggageclaimClient,
		volumeClient:            f.volumeClient,
		imageFactory:            f.imageFactory,
		dbVolumeFactory:         f.dbVolumeFactory,
		dbResourceCacheFactory:  f.dbResourceCacheFactory,
		dbResourceConfigFactory: f.dbResourceConfigFactory,
		dbTeamFactory:           f.dbTeamFactory,
		db:                      f.db,
		httpProxyURL:            f.httpProxyURL,
		httpsProxyURL:           f.httpsProxyURL,
		noProxy:                 f.noProxy,
		clock:                   f.clock,
		worker:                  worker,
	}
}

//go:generate counterfeiter . ContainerProvider

type ContainerProvider interface {
	FindContainerByHandle(
		logger lager.Logger,
		handle string,
		teamID int,
	) (Container, bool, error)

	FindOrCreateBuildContainer(
		logger lager.Logger,
		cancel <-chan os.Signal,
		delegate ImageFetchingDelegate,
		buildID int,
		planID atc.PlanID,
		metadata dbng.ContainerMetadata,
		spec ContainerSpec,
		resourceTypes atc.VersionedResourceTypes,
	) (Container, error)

	FindOrCreateResourceCheckContainer(
		logger lager.Logger,
		resourceUser dbng.ResourceUser,
		cancel <-chan os.Signal,
		delegate ImageFetchingDelegate,
		metadata dbng.ContainerMetadata,
		spec ContainerSpec,
		resourceTypes atc.VersionedResourceTypes,
		resourceType string,
		source atc.Source,
	) (Container, error)

	CreateResourceGetContainer(
		logger lager.Logger,
		resourceUser dbng.ResourceUser,
		cancel <-chan os.Signal,
		delegate ImageFetchingDelegate,
		metadata dbng.ContainerMetadata,
		spec ContainerSpec,
		resourceTypes atc.VersionedResourceTypes,
		resourceTypeName string,
		version atc.Version,
		source atc.Source,
		params atc.Params,
	) (Container, error)
}

type containerProvider struct {
	gardenClient            garden.Client
	baggageclaimClient      baggageclaim.Client
	volumeClient            VolumeClient
	imageFactory            ImageFactory
	dbVolumeFactory         dbng.VolumeFactory
	dbResourceCacheFactory  dbng.ResourceCacheFactory
	dbResourceConfigFactory dbng.ResourceConfigFactory
	dbTeamFactory           dbng.TeamFactory

	db       GardenWorkerDB
	provider WorkerProvider

	worker        Worker
	httpProxyURL  string
	httpsProxyURL string
	noProxy       string

	clock clock.Clock
}

func (p *containerProvider) FindOrCreateBuildContainer(
	logger lager.Logger,
	cancel <-chan os.Signal,
	delegate ImageFetchingDelegate,
	buildID int,
	planID atc.PlanID,
	metadata dbng.ContainerMetadata,
	spec ContainerSpec,
	resourceTypes atc.VersionedResourceTypes,
) (Container, error) {
	return p.findOrCreateContainer(
		logger,
		dbng.ForBuild{BuildID: buildID},
		cancel,
		delegate,
		metadata,
		spec,
		resourceTypes,
		func() (dbng.CreatingContainer, dbng.CreatedContainer, error) {
			return p.dbTeamFactory.GetByID(spec.TeamID).FindBuildContainerOnWorker(
				p.worker.Name(),
				buildID,
				planID,
			)
		},
		func() (dbng.CreatingContainer, error) {
			return p.dbTeamFactory.GetByID(spec.TeamID).CreateBuildContainer(
				p.worker.Name(),
				buildID,
				planID,
				metadata,
			)
		},
	)
}

func (p *containerProvider) FindOrCreateResourceCheckContainer(
	logger lager.Logger,
	resourceUser dbng.ResourceUser,
	cancel <-chan os.Signal,
	delegate ImageFetchingDelegate,
	metadata dbng.ContainerMetadata,
	spec ContainerSpec,
	resourceTypes atc.VersionedResourceTypes,
	resourceType string,
	source atc.Source,
) (Container, error) {
	resourceConfig, err := p.dbResourceConfigFactory.FindOrCreateResourceConfig(
		logger,
		resourceUser,
		resourceType,
		source,
		resourceTypes,
	)
	if err != nil {
		logger.Error("failed-to-get-resource-config", err)
		return nil, err
	}

	return p.findOrCreateContainer(
		logger,
		resourceUser,
		cancel,
		delegate,
		metadata,
		spec,
		resourceTypes,
		func() (dbng.CreatingContainer, dbng.CreatedContainer, error) {
			logger.Debug("looking-for-container-in-db", lager.Data{
				"team-id":            spec.TeamID,
				"worker-name":        p.worker.Name(),
				"resource-config-id": resourceConfig.ID,
			})
			return p.dbTeamFactory.GetByID(spec.TeamID).FindResourceCheckContainerOnWorker(
				p.worker.Name(),
				resourceConfig,
			)
		},
		func() (dbng.CreatingContainer, error) {
			logger.Debug("creating-container-in-db", lager.Data{
				"team-id":            spec.TeamID,
				"worker-name":        p.worker.Name(),
				"resource-config-id": resourceConfig.ID,
			})
			return p.dbTeamFactory.GetByID(spec.TeamID).CreateResourceCheckContainer(
				p.worker.Name(),
				resourceConfig,
				metadata,
			)
		},
	)
}

func (p *containerProvider) CreateResourceGetContainer(
	logger lager.Logger,
	resourceUser dbng.ResourceUser,
	cancel <-chan os.Signal,
	delegate ImageFetchingDelegate,
	metadata dbng.ContainerMetadata,
	spec ContainerSpec,
	resourceTypes atc.VersionedResourceTypes,
	resourceTypeName string,
	version atc.Version,
	source atc.Source,
	params atc.Params,
) (Container, error) {
	resourceCache, err := p.dbResourceCacheFactory.FindOrCreateResourceCache(
		logger,
		resourceUser,
		resourceTypeName,
		version,
		source,
		params,
		resourceTypes,
	)
	if err != nil {
		logger.Error("failed-to-get-resource-cache", err, lager.Data{"user": resourceUser})
		return nil, err
	}

	return p.findOrCreateContainer(
		logger,
		resourceUser,
		cancel,
		delegate,
		metadata,
		spec,
		resourceTypes,
		func() (dbng.CreatingContainer, dbng.CreatedContainer, error) {
			return nil, nil, nil
		},
		func() (dbng.CreatingContainer, error) {
			return p.dbTeamFactory.GetByID(spec.TeamID).CreateResourceGetContainer(
				p.worker.Name(),
				resourceCache,
				metadata,
			)
		},
	)
}

func (p *containerProvider) FindContainerByHandle(
	logger lager.Logger,
	handle string,
	teamID int,
) (Container, bool, error) {
	gardenContainer, err := p.gardenClient.Lookup(handle)
	if err != nil {
		if _, ok := err.(garden.ContainerNotFoundError); ok {
			logger.Info("container-not-found")
			return nil, false, nil
		}

		logger.Error("failed-to-lookup-on-garden", err)
		return nil, false, err
	}

	createdContainer, found, err := p.dbTeamFactory.GetByID(teamID).FindContainerByHandle(handle)
	if err != nil {
		logger.Error("failed-to-lookup-in-db", err)
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	createdVolumes, err := p.dbVolumeFactory.FindVolumesForContainer(createdContainer)
	if err != nil {
		return nil, false, err
	}

	container, err := newGardenWorkerContainer(
		logger,
		gardenContainer,
		createdContainer,
		createdVolumes,
		p.gardenClient,
		p.baggageclaimClient,
		p.db,
		p.worker.Name(),
	)

	if err != nil {
		logger.Error("failed-to-construct-container", err)
		return nil, false, err
	}

	return container, true, nil
}

func (p *containerProvider) findOrCreateContainer(
	logger lager.Logger,
	resourceUser dbng.ResourceUser,
	cancel <-chan os.Signal,
	delegate ImageFetchingDelegate,
	metadata dbng.ContainerMetadata,
	spec ContainerSpec,
	resourceTypes atc.VersionedResourceTypes,
	findContainerFunc func() (dbng.CreatingContainer, dbng.CreatedContainer, error),
	createContainerFunc func() (dbng.CreatingContainer, error),
) (Container, error) {
	var gardenContainer garden.Container

	creatingContainer, createdContainer, err := findContainerFunc()
	if err != nil {
		logger.Error("failed-to-find-container-in-db", err)
		return nil, err
	}

	if createdContainer != nil {
		logger.Debug("found-created-container-in-db")
		gardenContainer, err = p.gardenClient.Lookup(createdContainer.Handle())
		if err != nil {
			logger.Error("failed-to-lookup-created-container-in-garden", err)
			return nil, err
		}

		return p.constructGardenWorkerContainer(
			logger,
			createdContainer,
			gardenContainer,
		)
	}

	if creatingContainer != nil {
		logger.Debug("found-creating-container-in-db")
		gardenContainer, err = p.gardenClient.Lookup(creatingContainer.Handle())
		if err != nil {
			if _, ok := err.(garden.ContainerNotFoundError); !ok {
				logger.Error("failed-to-lookup-creating-container-in-garden", err)
				return nil, err
			}
		}
	}

	if gardenContainer == nil {
		image, err := p.imageFactory.GetImage(
			logger,
			p.worker,
			p.volumeClient,
			spec.ImageSpec,
			spec.TeamID,
			cancel,
			delegate,
			resourceUser,
			metadata,
			resourceTypes,
		)
		if err != nil {
			return nil, err
		}

		if creatingContainer == nil {
			logger.Debug("creating-container-in-db")
			creatingContainer, err = createContainerFunc()
			if err != nil {
				logger.Error("failed-to-create-container-in-db", err)
				return nil, err
			}
		}

		lock, acquired, err := p.db.AcquireContainerCreatingLock(logger, creatingContainer.ID())
		if err != nil {
			logger.Error("failed-to-acquire-volume-creating-lock", err)
			return nil, err
		}

		if !acquired {
			p.clock.Sleep(creatingContainerRetryDelay)
			return p.findOrCreateContainer(
				logger,
				resourceUser,
				cancel,
				delegate,
				metadata,
				spec,
				resourceTypes,
				findContainerFunc,
				createContainerFunc,
			)
		}

		defer lock.Release()

		fetchedImage, err := image.FetchForContainer(logger, creatingContainer)
		if err != nil {
			logger.Error("failed-to-fetch-image-for-container", err)
			return nil, err
		}

		gardenContainer, err = p.createGardenContainer(
			logger,
			creatingContainer,
			metadata,
			spec,
			fetchedImage.Metadata,
			fetchedImage.URL,
		)
		if err != nil {
			logger.Error("failed-to-create-container-in-garden", err)
			return nil, err
		}

		metadata.User = fetchedImage.Metadata.User
		if spec.User != "" {
			metadata.User = spec.User
		}

		// XXX
		// 		err = p.db.PutTheRestOfThisCrapInTheDatabaseButPleaseRemoveMeLater(
		// 			gardenContainer.Handle(),
		// 			db.ContainerMetadata(metadata),
		// 			p.maxContainerLifetime(metadata),
		// 		)
		// 		if err != nil {
		// 			logger.Error("failed-to-update-container-ttl", err)
		// 			return nil, err
		// 		}
	}

	createdContainer, err = creatingContainer.Created()
	if err != nil {
		logger.Error("failed-to-mark-container-as-created", err)
		return nil, err
	}

	return p.constructGardenWorkerContainer(
		logger,
		createdContainer,
		gardenContainer,
	)
}

func (p *containerProvider) constructGardenWorkerContainer(
	logger lager.Logger,
	createdContainer dbng.CreatedContainer,
	gardenContainer garden.Container,
) (Container, error) {
	createdVolumes, err := p.dbVolumeFactory.FindVolumesForContainer(createdContainer)
	if err != nil {
		logger.Error("failed-to-find-container-volumes", err)
		return nil, err
	}

	return newGardenWorkerContainer(
		logger,
		gardenContainer,
		createdContainer,
		createdVolumes,
		p.gardenClient,
		p.baggageclaimClient,
		p.db,
		p.worker.Name(),
	)
}

func (p *containerProvider) createGardenContainer(
	logger lager.Logger,
	creatingContainer dbng.CreatingContainer,
	metadata dbng.ContainerMetadata,
	spec ContainerSpec,
	imageMetadata ImageMetadata,
	imageURL string,
) (garden.Container, error) {
	volumeMounts := []VolumeMount{}

	if spec.Dir != "" && !p.anyMountTo(spec.Dir, spec.Inputs) {
		workdirVolume, volumeErr := p.volumeClient.FindOrCreateVolumeForContainer(
			logger,
			VolumeSpec{
				Strategy:   OutputStrategy{Name: "work-dir"}, // XXX: this is silly
				Privileged: bool(spec.ImageSpec.Privileged),
			},
			creatingContainer,
			spec.TeamID,
			spec.Dir,
		)
		if volumeErr != nil {
			return nil, volumeErr
		}

		volumeMounts = append(volumeMounts, VolumeMount{
			Volume:    workdirVolume,
			MountPath: spec.Dir,
		})
	}

	for _, inputSource := range spec.Inputs {
		var inputVolume Volume

		localVolume, found, err := inputSource.Source().VolumeOn(p.worker)
		if err != nil {
			return nil, err
		}

		if found {
			inputVolume, err = p.volumeClient.FindOrCreateVolumeForContainer(
				logger,
				VolumeSpec{
					Strategy: ContainerRootFSStrategy{
						Parent: localVolume,
					},
					Privileged: spec.ImageSpec.Privileged,
				},
				creatingContainer,
				spec.TeamID,
				inputSource.DestinationPath(),
			)
			if err != nil {
				return nil, err
			}
		} else {
			inputVolume, err = p.volumeClient.FindOrCreateVolumeForContainer(
				logger,
				VolumeSpec{
					Strategy:   OutputStrategy{Name: string(inputSource.Name())}, // XXX: this is silly
					Privileged: spec.ImageSpec.Privileged,
				},
				creatingContainer,
				spec.TeamID,
				inputSource.DestinationPath(),
			)
			if err != nil {
				return nil, err
			}

			err = inputSource.Source().StreamTo(inputVolume)
			if err != nil {
				return nil, err
			}
		}

		volumeMounts = append(volumeMounts, VolumeMount{
			Volume:    inputVolume,
			MountPath: inputSource.DestinationPath(),
		})
	}

	for name, outputPath := range spec.Outputs {
		outVolume, volumeErr := p.volumeClient.FindOrCreateVolumeForContainer(
			logger,
			VolumeSpec{
				Strategy:   OutputStrategy{Name: name},
				Privileged: bool(spec.ImageSpec.Privileged),
			},
			creatingContainer,
			spec.TeamID,
			outputPath,
		)
		if volumeErr != nil {
			return nil, volumeErr
		}

		volumeMounts = append(volumeMounts, VolumeMount{
			Volume:    outVolume,
			MountPath: outputPath,
		})
	}

	if spec.ResourceCache != nil {
		volumeMounts = append(volumeMounts, *spec.ResourceCache)
	}

	bindMounts := []garden.BindMount{}

	volumeHandleMounts := map[string]string{}
	for _, mount := range volumeMounts {
		bindMounts = append(bindMounts, garden.BindMount{
			SrcPath: mount.Volume.Path(),
			DstPath: mount.MountPath,
			Mode:    garden.BindMountModeRW,
		})
		volumeHandleMounts[mount.Volume.Handle()] = mount.MountPath
	}

	gardenProperties := garden.Properties{}

	if spec.User != "" {
		gardenProperties[userPropertyName] = spec.User
	} else {
		gardenProperties[userPropertyName] = imageMetadata.User
	}

	if spec.Ephemeral {
		gardenProperties[ephemeralPropertyName] = "true"
	}

	env := append(imageMetadata.Env, spec.Env...)

	if p.httpProxyURL != "" {
		env = append(env, fmt.Sprintf("http_proxy=%s", p.httpProxyURL))
	}

	if p.httpsProxyURL != "" {
		env = append(env, fmt.Sprintf("https_proxy=%s", p.httpsProxyURL))
	}

	if p.noProxy != "" {
		env = append(env, fmt.Sprintf("no_proxy=%s", p.noProxy))
	}

	return p.gardenClient.Create(garden.ContainerSpec{
		BindMounts: bindMounts,
		Privileged: spec.ImageSpec.Privileged,
		Properties: gardenProperties,
		RootFSPath: imageURL,
		Env:        env,
		Handle:     creatingContainer.Handle(),
	})
}

func (p *containerProvider) anyMountTo(path string, inputs []InputSource) bool {
	for _, input := range inputs {
		if input.DestinationPath() == path {
			return true
		}
	}

	return false
}

// func (p *containerProvider) maxContainerLifetime(metadata Metadata) time.Duration {
// 	if metadata.Type == db.ContainerTypeCheck {
// 		uptime := p.worker.Uptime()
// 		switch {
// 		case uptime < 5*time.Minute:
// 			return 5 * time.Minute
// 		case uptime > 1*time.Hour:
// 			return 1 * time.Hour
// 		default:
// 			return uptime
// 		}
// 	}

// 	return time.Duration(0)
// }
