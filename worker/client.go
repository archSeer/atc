package worker

import (
	"os"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/baggageclaim"
)

//go:generate counterfeiter . Client

type Client interface {
	FindOrCreateBuildContainer(
		lager.Logger,
		<-chan os.Signal,
		ImageFetchingDelegate,
		Identifier,
		Metadata,
		ContainerSpec,
		atc.VersionedResourceTypes,
		map[string]string,
	) (Container, error)

	FindBuildContainer(
		lager.Logger,
		<-chan os.Signal,
		ImageFetchingDelegate,
		Identifier,
		Metadata,
		ContainerSpec,
		atc.VersionedResourceTypes,
		map[string]string,
	) (Container, error)

	CreateResourceGetContainer(
		logger lager.Logger,
		resourceUser dbng.ResourceUser,
		cancel <-chan os.Signal,
		delegate ImageFetchingDelegate,
		id Identifier,
		metadata Metadata,
		spec ContainerSpec,
		resourceTypes atc.VersionedResourceTypes,
		outputPaths map[string]string,
		resourceType string,
		version atc.Version,
		source atc.Source,
		params atc.Params,
	) (Container, error)

	FindOrCreateResourceCheckContainer(
		logger lager.Logger,
		resourceUser dbng.ResourceUser,
		cancel <-chan os.Signal,
		delegate ImageFetchingDelegate,
		id Identifier,
		metadata Metadata,
		spec ContainerSpec,
		resourceTypes atc.VersionedResourceTypes,
		resourceType string,
		source atc.Source,
	) (Container, error)

	CreateVolumeForResourceCache(
		logger lager.Logger,
		vs VolumeSpec,
		resourceCache *dbng.UsedResourceCache,
	) (Volume, error)

	FindInitializedVolumeForResourceCache(
		logger lager.Logger,
		resourceCache *dbng.UsedResourceCache,
	) (Volume, bool, error)

	FindContainerForIdentifier(lager.Logger, Identifier) (Container, bool, error)
	FindContainerByHandle(lager.Logger, string, int) (Container, bool, error)
	FindResourceTypeByPath(path string) (atc.WorkerResourceType, bool)
	LookupVolume(lager.Logger, string) (Volume, bool, error)

	Satisfying(WorkerSpec, atc.VersionedResourceTypes) (Worker, error)
	AllSatisfying(WorkerSpec, atc.VersionedResourceTypes) ([]Worker, error)
	RunningWorkers() ([]Worker, error)
	GetWorker(workerName string) (Worker, error)
}

type VolumeSpec struct {
	Strategy   Strategy
	Properties VolumeProperties
	Privileged bool
	TTL        time.Duration
}

func (spec VolumeSpec) baggageclaimVolumeSpec() baggageclaim.VolumeSpec {
	return baggageclaim.VolumeSpec{
		Strategy:   spec.Strategy.baggageclaimStrategy(),
		Privileged: spec.Privileged,
		Properties: baggageclaim.VolumeProperties(spec.Properties),
	}
}

type Strategy interface {
	baggageclaimStrategy() baggageclaim.Strategy
}

type ResourceCacheStrategy struct {
	ResourceHash    string
	ResourceVersion atc.Version
}

func (ResourceCacheStrategy) baggageclaimStrategy() baggageclaim.Strategy {
	return baggageclaim.EmptyStrategy{}
}

type OutputStrategy struct {
	Name string
}

func (OutputStrategy) baggageclaimStrategy() baggageclaim.Strategy {
	return baggageclaim.EmptyStrategy{}
}

type ImageArtifactReplicationStrategy struct {
	Name string
}

func (ImageArtifactReplicationStrategy) baggageclaimStrategy() baggageclaim.Strategy {
	return baggageclaim.EmptyStrategy{}
}

type ContainerRootFSStrategy struct {
	Parent Volume
}

func (strategy ContainerRootFSStrategy) baggageclaimStrategy() baggageclaim.Strategy {
	return strategy.Parent.COWStrategy()
}

type HostRootFSStrategy struct {
	Path       string
	WorkerName string
	Version    *string
}

func (strategy HostRootFSStrategy) baggageclaimStrategy() baggageclaim.Strategy {
	return baggageclaim.ImportStrategy{
		Path: strategy.Path,
	}
}

//go:generate counterfeiter . Container

type Container interface {
	garden.Container

	Destroy() error

	VolumeMounts() []VolumeMount

	WorkerName() string

	MarkAsHijacked() error
}

type ResourceCacheIdentifier db.ResourceCacheIdentifier
type VolumeProperties map[string]string

type Identifier db.ContainerIdentifier
type Metadata db.ContainerMetadata
