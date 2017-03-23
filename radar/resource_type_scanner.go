package radar

import (
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type resourceTypeScanner struct {
	resourceFactory resource.ResourceFactory
	defaultInterval time.Duration
	db              RadarDB
	dbPipeline      dbng.Pipeline
	externalURL     string
}

func NewResourceTypeScanner(
	resourceFactory resource.ResourceFactory,
	defaultInterval time.Duration,
	db RadarDB,
	dbPipeline dbng.Pipeline,
	externalURL string,
) Scanner {
	return &resourceTypeScanner{
		resourceFactory: resourceFactory,
		defaultInterval: defaultInterval,
		db:              db,
		dbPipeline:      dbPipeline,
		externalURL:     externalURL,
	}
}

func (scanner *resourceTypeScanner) Run(logger lager.Logger, resourceTypeName string) (time.Duration, error) {
	pipelinePaused, err := scanner.db.IsPaused()
	if err != nil {
		logger.Error("failed-to-check-if-pipeline-paused", err)
		return 0, err
	}

	if pipelinePaused {
		logger.Debug("pipeline-paused")
		return scanner.defaultInterval, nil
	}

	savedResourceType, found, err := scanner.db.GetResourceType(resourceTypeName)
	if err != nil {
		logger.Error("failed-to-get-current-version", err)
		return 0, err
	}

	if !found {
		return 0, db.ResourceTypeNotFoundError{Name: resourceTypeName}
	}

	lockLogger := logger.Session("lock", lager.Data{
		"resource-type": resourceTypeName,
	})

	lock, acquired, err := scanner.db.AcquireResourceTypeCheckingLock(logger, savedResourceType, scanner.defaultInterval, false)
	if err != nil {
		lockLogger.Error("failed-to-get-lock", err, lager.Data{
			"resource-type": resourceTypeName,
		})
		return scanner.defaultInterval, ErrFailedToAcquireLock
	}

	if !acquired {
		lockLogger.Debug("did-not-get-lock")
		return scanner.defaultInterval, ErrFailedToAcquireLock
	}

	defer lock.Release()

	err = scanner.resourceTypeScan(logger.Session("tick"), savedResourceType)
	if err != nil {
		return 0, err
	}

	return scanner.defaultInterval, nil
}

func (scanner *resourceTypeScanner) Scan(logger lager.Logger, resourceTypeName string) error {
	return nil
}

func (scanner *resourceTypeScanner) ScanFromVersion(logger lager.Logger, resourceTypeName string, fromVersion atc.Version) error {
	return nil
}

func (scanner *resourceTypeScanner) resourceTypeScan(logger lager.Logger, savedResourceType db.SavedResourceType) error {
	resourceTypes, err := scanner.dbPipeline.ResourceTypes()
	if err != nil {
		logger.Error("failed-to-get-resource-types", err)
		return err
	}

	versionedResourceTypes := deserializeVersionedResourceTypes(resourceTypes)

	resourceSpec := worker.ContainerSpec{
		ImageSpec: worker.ImageSpec{
			ResourceType: savedResourceType.Config.Type,
			Privileged:   true,
		},
		Ephemeral: true,
		Tags:      []string{},
		TeamID:    scanner.dbPipeline.TeamID(),
	}

	res, err := scanner.resourceFactory.NewCheckResource(
		logger,
		dbng.ForResourceType{
			ResourceTypeID: savedResourceType.ID,
		},
		dbng.ContainerMetadata{
			Type:           dbng.ContainerTypeCheck,
			PipelineID:     scanner.dbPipeline.ID(),
			ResourceTypeID: savedResourceType.ID,
		},
		resourceSpec,
		versionedResourceTypes,
		worker.NoopImageFetchingDelegate{},
		atc.ResourceConfig{
			Type:   savedResourceType.Config.Type,
			Source: savedResourceType.Config.Source,
		},
	)
	if err != nil {
		logger.Error("failed-to-initialize-new-container", err)
		return err
	}

	newVersions, err := res.Check(savedResourceType.Config.Source, atc.Version(savedResourceType.Version))
	if err != nil {
		if rErr, ok := err.(resource.ErrResourceScriptFailed); ok {
			logger.Info("check-failed", lager.Data{"exit-status": rErr.ExitStatus})
			return nil
		}

		logger.Error("failed-to-check", err)
		return err
	}

	if len(newVersions) == 0 {
		logger.Debug("no-new-versions")
		return nil
	}

	logger.Info("versions-found", lager.Data{
		"versions": newVersions,
		"total":    len(newVersions),
	})

	version := newVersions[len(newVersions)-1]
	err = scanner.db.SaveResourceTypeVersion(savedResourceType.Config, version)
	if err != nil {
		logger.Error("failed-to-save-resource-type-version", err, lager.Data{
			"version": version,
		})
		return err
	}

	return nil
}
