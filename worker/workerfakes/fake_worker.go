// This file was generated by counterfeiter
package workerfakes

import (
	"os"
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeWorker struct {
	CreateBuildContainerStub        func(lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, map[string]string) (worker.Container, error)
	createBuildContainerMutex       sync.RWMutex
	createBuildContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 worker.Identifier
		arg5 worker.Metadata
		arg6 worker.ContainerSpec
		arg7 atc.ResourceTypes
		arg8 map[string]string
	}
	createBuildContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	CreateResourceGetContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, outputPaths map[string]string, resourceType string, version atc.Version, source atc.Source, params atc.Params) (worker.Container, error)
	createResourceGetContainerMutex       sync.RWMutex
	createResourceGetContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		outputPaths   map[string]string
		resourceType  string
		version       atc.Version
		source        atc.Source
		params        atc.Params
	}
	createResourceGetContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	CreateResourceCheckContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error)
	createResourceCheckContainerMutex       sync.RWMutex
	createResourceCheckContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}
	createResourceCheckContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	CreateResourceTypeCheckContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error)
	createResourceTypeCheckContainerMutex       sync.RWMutex
	createResourceTypeCheckContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}
	createResourceTypeCheckContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindContainerForIdentifierStub        func(lager.Logger, worker.Identifier) (worker.Container, bool, error)
	findContainerForIdentifierMutex       sync.RWMutex
	findContainerForIdentifierArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Identifier
	}
	findContainerForIdentifierReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	LookupContainerStub        func(lager.Logger, string) (worker.Container, bool, error)
	lookupContainerMutex       sync.RWMutex
	lookupContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupContainerReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	ValidateResourceCheckVersionStub        func(container db.SavedContainer) (bool, error)
	validateResourceCheckVersionMutex       sync.RWMutex
	validateResourceCheckVersionArgsForCall []struct {
		container db.SavedContainer
	}
	validateResourceCheckVersionReturns struct {
		result1 bool
		result2 error
	}
	FindResourceTypeByPathStub        func(path string) (atc.WorkerResourceType, bool)
	findResourceTypeByPathMutex       sync.RWMutex
	findResourceTypeByPathArgsForCall []struct {
		path string
	}
	findResourceTypeByPathReturns struct {
		result1 atc.WorkerResourceType
		result2 bool
	}
	FindVolumeStub        func(lager.Logger, worker.VolumeSpec) (worker.Volume, bool, error)
	findVolumeMutex       sync.RWMutex
	findVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
	}
	findVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	CreateVolumeStub        func(logger lager.Logger, vs worker.VolumeSpec, teamID int) (worker.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		logger lager.Logger
		vs     worker.VolumeSpec
		teamID int
	}
	createVolumeReturns struct {
		result1 worker.Volume
		result2 error
	}
	ListVolumesStub        func(lager.Logger, worker.VolumeProperties) ([]worker.Volume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeProperties
	}
	listVolumesReturns struct {
		result1 []worker.Volume
		result2 error
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	SatisfyingStub        func(worker.WorkerSpec, atc.ResourceTypes) (worker.Worker, error)
	satisfyingMutex       sync.RWMutex
	satisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}
	satisfyingReturns struct {
		result1 worker.Worker
		result2 error
	}
	AllSatisfyingStub        func(worker.WorkerSpec, atc.ResourceTypes) ([]worker.Worker, error)
	allSatisfyingMutex       sync.RWMutex
	allSatisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}
	allSatisfyingReturns struct {
		result1 []worker.Worker
		result2 error
	}
	WorkersStub        func() ([]worker.Worker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct{}
	workersReturns     struct {
		result1 []worker.Worker
		result2 error
	}
	GetWorkerStub        func(workerName string) (worker.Worker, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		workerName string
	}
	getWorkerReturns struct {
		result1 worker.Worker
		result2 error
	}
	ActiveContainersStub        func() int
	activeContainersMutex       sync.RWMutex
	activeContainersArgsForCall []struct{}
	activeContainersReturns     struct {
		result1 int
	}
	DescriptionStub        func() string
	descriptionMutex       sync.RWMutex
	descriptionArgsForCall []struct{}
	descriptionReturns     struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	UptimeStub        func() time.Duration
	uptimeMutex       sync.RWMutex
	uptimeArgsForCall []struct{}
	uptimeReturns     struct {
		result1 time.Duration
	}
	IsOwnedByTeamStub        func() bool
	isOwnedByTeamMutex       sync.RWMutex
	isOwnedByTeamArgsForCall []struct{}
	isOwnedByTeamReturns     struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) CreateBuildContainer(arg1 lager.Logger, arg2 <-chan os.Signal, arg3 worker.ImageFetchingDelegate, arg4 worker.Identifier, arg5 worker.Metadata, arg6 worker.ContainerSpec, arg7 atc.ResourceTypes, arg8 map[string]string) (worker.Container, error) {
	fake.createBuildContainerMutex.Lock()
	fake.createBuildContainerArgsForCall = append(fake.createBuildContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 worker.Identifier
		arg5 worker.Metadata
		arg6 worker.ContainerSpec
		arg7 atc.ResourceTypes
		arg8 map[string]string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("CreateBuildContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.createBuildContainerMutex.Unlock()
	if fake.CreateBuildContainerStub != nil {
		return fake.CreateBuildContainerStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	} else {
		return fake.createBuildContainerReturns.result1, fake.createBuildContainerReturns.result2
	}
}

func (fake *FakeWorker) CreateBuildContainerCallCount() int {
	fake.createBuildContainerMutex.RLock()
	defer fake.createBuildContainerMutex.RUnlock()
	return len(fake.createBuildContainerArgsForCall)
}

func (fake *FakeWorker) CreateBuildContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, map[string]string) {
	fake.createBuildContainerMutex.RLock()
	defer fake.createBuildContainerMutex.RUnlock()
	return fake.createBuildContainerArgsForCall[i].arg1, fake.createBuildContainerArgsForCall[i].arg2, fake.createBuildContainerArgsForCall[i].arg3, fake.createBuildContainerArgsForCall[i].arg4, fake.createBuildContainerArgsForCall[i].arg5, fake.createBuildContainerArgsForCall[i].arg6, fake.createBuildContainerArgsForCall[i].arg7, fake.createBuildContainerArgsForCall[i].arg8
}

func (fake *FakeWorker) CreateBuildContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateBuildContainerStub = nil
	fake.createBuildContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) CreateResourceGetContainer(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, outputPaths map[string]string, resourceType string, version atc.Version, source atc.Source, params atc.Params) (worker.Container, error) {
	fake.createResourceGetContainerMutex.Lock()
	fake.createResourceGetContainerArgsForCall = append(fake.createResourceGetContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		outputPaths   map[string]string
		resourceType  string
		version       atc.Version
		source        atc.Source
		params        atc.Params
	}{logger, cancel, delegate, id, metadata, spec, resourceTypes, outputPaths, resourceType, version, source, params})
	fake.recordInvocation("CreateResourceGetContainer", []interface{}{logger, cancel, delegate, id, metadata, spec, resourceTypes, outputPaths, resourceType, version, source, params})
	fake.createResourceGetContainerMutex.Unlock()
	if fake.CreateResourceGetContainerStub != nil {
		return fake.CreateResourceGetContainerStub(logger, cancel, delegate, id, metadata, spec, resourceTypes, outputPaths, resourceType, version, source, params)
	} else {
		return fake.createResourceGetContainerReturns.result1, fake.createResourceGetContainerReturns.result2
	}
}

func (fake *FakeWorker) CreateResourceGetContainerCallCount() int {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return len(fake.createResourceGetContainerArgsForCall)
}

func (fake *FakeWorker) CreateResourceGetContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, map[string]string, string, atc.Version, atc.Source, atc.Params) {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return fake.createResourceGetContainerArgsForCall[i].logger, fake.createResourceGetContainerArgsForCall[i].cancel, fake.createResourceGetContainerArgsForCall[i].delegate, fake.createResourceGetContainerArgsForCall[i].id, fake.createResourceGetContainerArgsForCall[i].metadata, fake.createResourceGetContainerArgsForCall[i].spec, fake.createResourceGetContainerArgsForCall[i].resourceTypes, fake.createResourceGetContainerArgsForCall[i].outputPaths, fake.createResourceGetContainerArgsForCall[i].resourceType, fake.createResourceGetContainerArgsForCall[i].version, fake.createResourceGetContainerArgsForCall[i].source, fake.createResourceGetContainerArgsForCall[i].params
}

func (fake *FakeWorker) CreateResourceGetContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateResourceGetContainerStub = nil
	fake.createResourceGetContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) CreateResourceCheckContainer(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error) {
	fake.createResourceCheckContainerMutex.Lock()
	fake.createResourceCheckContainerArgsForCall = append(fake.createResourceCheckContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.recordInvocation("CreateResourceCheckContainer", []interface{}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.createResourceCheckContainerMutex.Unlock()
	if fake.CreateResourceCheckContainerStub != nil {
		return fake.CreateResourceCheckContainerStub(logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source)
	} else {
		return fake.createResourceCheckContainerReturns.result1, fake.createResourceCheckContainerReturns.result2
	}
}

func (fake *FakeWorker) CreateResourceCheckContainerCallCount() int {
	fake.createResourceCheckContainerMutex.RLock()
	defer fake.createResourceCheckContainerMutex.RUnlock()
	return len(fake.createResourceCheckContainerArgsForCall)
}

func (fake *FakeWorker) CreateResourceCheckContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, string, atc.Source) {
	fake.createResourceCheckContainerMutex.RLock()
	defer fake.createResourceCheckContainerMutex.RUnlock()
	return fake.createResourceCheckContainerArgsForCall[i].logger, fake.createResourceCheckContainerArgsForCall[i].cancel, fake.createResourceCheckContainerArgsForCall[i].delegate, fake.createResourceCheckContainerArgsForCall[i].id, fake.createResourceCheckContainerArgsForCall[i].metadata, fake.createResourceCheckContainerArgsForCall[i].spec, fake.createResourceCheckContainerArgsForCall[i].resourceTypes, fake.createResourceCheckContainerArgsForCall[i].resourceType, fake.createResourceCheckContainerArgsForCall[i].source
}

func (fake *FakeWorker) CreateResourceCheckContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateResourceCheckContainerStub = nil
	fake.createResourceCheckContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) CreateResourceTypeCheckContainer(logger lager.Logger, cancel <-chan os.Signal, delegate worker.ImageFetchingDelegate, id worker.Identifier, metadata worker.Metadata, spec worker.ContainerSpec, resourceTypes atc.ResourceTypes, resourceType string, source atc.Source) (worker.Container, error) {
	fake.createResourceTypeCheckContainerMutex.Lock()
	fake.createResourceTypeCheckContainerArgsForCall = append(fake.createResourceTypeCheckContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		delegate      worker.ImageFetchingDelegate
		id            worker.Identifier
		metadata      worker.Metadata
		spec          worker.ContainerSpec
		resourceTypes atc.ResourceTypes
		resourceType  string
		source        atc.Source
	}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.recordInvocation("CreateResourceTypeCheckContainer", []interface{}{logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source})
	fake.createResourceTypeCheckContainerMutex.Unlock()
	if fake.CreateResourceTypeCheckContainerStub != nil {
		return fake.CreateResourceTypeCheckContainerStub(logger, cancel, delegate, id, metadata, spec, resourceTypes, resourceType, source)
	} else {
		return fake.createResourceTypeCheckContainerReturns.result1, fake.createResourceTypeCheckContainerReturns.result2
	}
}

func (fake *FakeWorker) CreateResourceTypeCheckContainerCallCount() int {
	fake.createResourceTypeCheckContainerMutex.RLock()
	defer fake.createResourceTypeCheckContainerMutex.RUnlock()
	return len(fake.createResourceTypeCheckContainerArgsForCall)
}

func (fake *FakeWorker) CreateResourceTypeCheckContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes, string, atc.Source) {
	fake.createResourceTypeCheckContainerMutex.RLock()
	defer fake.createResourceTypeCheckContainerMutex.RUnlock()
	return fake.createResourceTypeCheckContainerArgsForCall[i].logger, fake.createResourceTypeCheckContainerArgsForCall[i].cancel, fake.createResourceTypeCheckContainerArgsForCall[i].delegate, fake.createResourceTypeCheckContainerArgsForCall[i].id, fake.createResourceTypeCheckContainerArgsForCall[i].metadata, fake.createResourceTypeCheckContainerArgsForCall[i].spec, fake.createResourceTypeCheckContainerArgsForCall[i].resourceTypes, fake.createResourceTypeCheckContainerArgsForCall[i].resourceType, fake.createResourceTypeCheckContainerArgsForCall[i].source
}

func (fake *FakeWorker) CreateResourceTypeCheckContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateResourceTypeCheckContainerStub = nil
	fake.createResourceTypeCheckContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) FindContainerForIdentifier(arg1 lager.Logger, arg2 worker.Identifier) (worker.Container, bool, error) {
	fake.findContainerForIdentifierMutex.Lock()
	fake.findContainerForIdentifierArgsForCall = append(fake.findContainerForIdentifierArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Identifier
	}{arg1, arg2})
	fake.recordInvocation("FindContainerForIdentifier", []interface{}{arg1, arg2})
	fake.findContainerForIdentifierMutex.Unlock()
	if fake.FindContainerForIdentifierStub != nil {
		return fake.FindContainerForIdentifierStub(arg1, arg2)
	} else {
		return fake.findContainerForIdentifierReturns.result1, fake.findContainerForIdentifierReturns.result2, fake.findContainerForIdentifierReturns.result3
	}
}

func (fake *FakeWorker) FindContainerForIdentifierCallCount() int {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return len(fake.findContainerForIdentifierArgsForCall)
}

func (fake *FakeWorker) FindContainerForIdentifierArgsForCall(i int) (lager.Logger, worker.Identifier) {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return fake.findContainerForIdentifierArgsForCall[i].arg1, fake.findContainerForIdentifierArgsForCall[i].arg2
}

func (fake *FakeWorker) FindContainerForIdentifierReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerForIdentifierStub = nil
	fake.findContainerForIdentifierReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) LookupContainer(arg1 lager.Logger, arg2 string) (worker.Container, bool, error) {
	fake.lookupContainerMutex.Lock()
	fake.lookupContainerArgsForCall = append(fake.lookupContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupContainer", []interface{}{arg1, arg2})
	fake.lookupContainerMutex.Unlock()
	if fake.LookupContainerStub != nil {
		return fake.LookupContainerStub(arg1, arg2)
	} else {
		return fake.lookupContainerReturns.result1, fake.lookupContainerReturns.result2, fake.lookupContainerReturns.result3
	}
}

func (fake *FakeWorker) LookupContainerCallCount() int {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return len(fake.lookupContainerArgsForCall)
}

func (fake *FakeWorker) LookupContainerArgsForCall(i int) (lager.Logger, string) {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return fake.lookupContainerArgsForCall[i].arg1, fake.lookupContainerArgsForCall[i].arg2
}

func (fake *FakeWorker) LookupContainerReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.LookupContainerStub = nil
	fake.lookupContainerReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) ValidateResourceCheckVersion(container db.SavedContainer) (bool, error) {
	fake.validateResourceCheckVersionMutex.Lock()
	fake.validateResourceCheckVersionArgsForCall = append(fake.validateResourceCheckVersionArgsForCall, struct {
		container db.SavedContainer
	}{container})
	fake.recordInvocation("ValidateResourceCheckVersion", []interface{}{container})
	fake.validateResourceCheckVersionMutex.Unlock()
	if fake.ValidateResourceCheckVersionStub != nil {
		return fake.ValidateResourceCheckVersionStub(container)
	} else {
		return fake.validateResourceCheckVersionReturns.result1, fake.validateResourceCheckVersionReturns.result2
	}
}

func (fake *FakeWorker) ValidateResourceCheckVersionCallCount() int {
	fake.validateResourceCheckVersionMutex.RLock()
	defer fake.validateResourceCheckVersionMutex.RUnlock()
	return len(fake.validateResourceCheckVersionArgsForCall)
}

func (fake *FakeWorker) ValidateResourceCheckVersionArgsForCall(i int) db.SavedContainer {
	fake.validateResourceCheckVersionMutex.RLock()
	defer fake.validateResourceCheckVersionMutex.RUnlock()
	return fake.validateResourceCheckVersionArgsForCall[i].container
}

func (fake *FakeWorker) ValidateResourceCheckVersionReturns(result1 bool, result2 error) {
	fake.ValidateResourceCheckVersionStub = nil
	fake.validateResourceCheckVersionReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) FindResourceTypeByPath(path string) (atc.WorkerResourceType, bool) {
	fake.findResourceTypeByPathMutex.Lock()
	fake.findResourceTypeByPathArgsForCall = append(fake.findResourceTypeByPathArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("FindResourceTypeByPath", []interface{}{path})
	fake.findResourceTypeByPathMutex.Unlock()
	if fake.FindResourceTypeByPathStub != nil {
		return fake.FindResourceTypeByPathStub(path)
	} else {
		return fake.findResourceTypeByPathReturns.result1, fake.findResourceTypeByPathReturns.result2
	}
}

func (fake *FakeWorker) FindResourceTypeByPathCallCount() int {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return len(fake.findResourceTypeByPathArgsForCall)
}

func (fake *FakeWorker) FindResourceTypeByPathArgsForCall(i int) string {
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	return fake.findResourceTypeByPathArgsForCall[i].path
}

func (fake *FakeWorker) FindResourceTypeByPathReturns(result1 atc.WorkerResourceType, result2 bool) {
	fake.FindResourceTypeByPathStub = nil
	fake.findResourceTypeByPathReturns = struct {
		result1 atc.WorkerResourceType
		result2 bool
	}{result1, result2}
}

func (fake *FakeWorker) FindVolume(arg1 lager.Logger, arg2 worker.VolumeSpec) (worker.Volume, bool, error) {
	fake.findVolumeMutex.Lock()
	fake.findVolumeArgsForCall = append(fake.findVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeSpec
	}{arg1, arg2})
	fake.recordInvocation("FindVolume", []interface{}{arg1, arg2})
	fake.findVolumeMutex.Unlock()
	if fake.FindVolumeStub != nil {
		return fake.FindVolumeStub(arg1, arg2)
	} else {
		return fake.findVolumeReturns.result1, fake.findVolumeReturns.result2, fake.findVolumeReturns.result3
	}
}

func (fake *FakeWorker) FindVolumeCallCount() int {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	return len(fake.findVolumeArgsForCall)
}

func (fake *FakeWorker) FindVolumeArgsForCall(i int) (lager.Logger, worker.VolumeSpec) {
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	return fake.findVolumeArgsForCall[i].arg1, fake.findVolumeArgsForCall[i].arg2
}

func (fake *FakeWorker) FindVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindVolumeStub = nil
	fake.findVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) CreateVolume(logger lager.Logger, vs worker.VolumeSpec, teamID int) (worker.Volume, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		logger lager.Logger
		vs     worker.VolumeSpec
		teamID int
	}{logger, vs, teamID})
	fake.recordInvocation("CreateVolume", []interface{}{logger, vs, teamID})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(logger, vs, teamID)
	} else {
		return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
	}
}

func (fake *FakeWorker) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeWorker) CreateVolumeArgsForCall(i int) (lager.Logger, worker.VolumeSpec, int) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].logger, fake.createVolumeArgsForCall[i].vs, fake.createVolumeArgsForCall[i].teamID
}

func (fake *FakeWorker) CreateVolumeReturns(result1 worker.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) ListVolumes(arg1 lager.Logger, arg2 worker.VolumeProperties) ([]worker.Volume, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeProperties
	}{arg1, arg2})
	fake.recordInvocation("ListVolumes", []interface{}{arg1, arg2})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(arg1, arg2)
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeWorker) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeWorker) ListVolumesArgsForCall(i int) (lager.Logger, worker.VolumeProperties) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].arg1, fake.listVolumesArgsForCall[i].arg2
}

func (fake *FakeWorker) ListVolumesReturns(result1 []worker.Volume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
	}
}

func (fake *FakeWorker) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeWorker) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeWorker) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorker) Satisfying(arg1 worker.WorkerSpec, arg2 atc.ResourceTypes) (worker.Worker, error) {
	fake.satisfyingMutex.Lock()
	fake.satisfyingArgsForCall = append(fake.satisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}{arg1, arg2})
	fake.recordInvocation("Satisfying", []interface{}{arg1, arg2})
	fake.satisfyingMutex.Unlock()
	if fake.SatisfyingStub != nil {
		return fake.SatisfyingStub(arg1, arg2)
	} else {
		return fake.satisfyingReturns.result1, fake.satisfyingReturns.result2
	}
}

func (fake *FakeWorker) SatisfyingCallCount() int {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return len(fake.satisfyingArgsForCall)
}

func (fake *FakeWorker) SatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.ResourceTypes) {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return fake.satisfyingArgsForCall[i].arg1, fake.satisfyingArgsForCall[i].arg2
}

func (fake *FakeWorker) SatisfyingReturns(result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	fake.satisfyingReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) AllSatisfying(arg1 worker.WorkerSpec, arg2 atc.ResourceTypes) ([]worker.Worker, error) {
	fake.allSatisfyingMutex.Lock()
	fake.allSatisfyingArgsForCall = append(fake.allSatisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}{arg1, arg2})
	fake.recordInvocation("AllSatisfying", []interface{}{arg1, arg2})
	fake.allSatisfyingMutex.Unlock()
	if fake.AllSatisfyingStub != nil {
		return fake.AllSatisfyingStub(arg1, arg2)
	} else {
		return fake.allSatisfyingReturns.result1, fake.allSatisfyingReturns.result2
	}
}

func (fake *FakeWorker) AllSatisfyingCallCount() int {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return len(fake.allSatisfyingArgsForCall)
}

func (fake *FakeWorker) AllSatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.ResourceTypes) {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return fake.allSatisfyingArgsForCall[i].arg1, fake.allSatisfyingArgsForCall[i].arg2
}

func (fake *FakeWorker) AllSatisfyingReturns(result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	fake.allSatisfyingReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) Workers() ([]worker.Worker, error) {
	fake.workersMutex.Lock()
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct{}{})
	fake.recordInvocation("Workers", []interface{}{})
	fake.workersMutex.Unlock()
	if fake.WorkersStub != nil {
		return fake.WorkersStub()
	} else {
		return fake.workersReturns.result1, fake.workersReturns.result2
	}
}

func (fake *FakeWorker) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeWorker) WorkersReturns(result1 []worker.Worker, result2 error) {
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) GetWorker(workerName string) (worker.Worker, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		workerName string
	}{workerName})
	fake.recordInvocation("GetWorker", []interface{}{workerName})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(workerName)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2
	}
}

func (fake *FakeWorker) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeWorker) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].workerName
}

func (fake *FakeWorker) GetWorkerReturns(result1 worker.Worker, result2 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) ActiveContainers() int {
	fake.activeContainersMutex.Lock()
	fake.activeContainersArgsForCall = append(fake.activeContainersArgsForCall, struct{}{})
	fake.recordInvocation("ActiveContainers", []interface{}{})
	fake.activeContainersMutex.Unlock()
	if fake.ActiveContainersStub != nil {
		return fake.ActiveContainersStub()
	} else {
		return fake.activeContainersReturns.result1
	}
}

func (fake *FakeWorker) ActiveContainersCallCount() int {
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	return len(fake.activeContainersArgsForCall)
}

func (fake *FakeWorker) ActiveContainersReturns(result1 int) {
	fake.ActiveContainersStub = nil
	fake.activeContainersReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) Description() string {
	fake.descriptionMutex.Lock()
	fake.descriptionArgsForCall = append(fake.descriptionArgsForCall, struct{}{})
	fake.recordInvocation("Description", []interface{}{})
	fake.descriptionMutex.Unlock()
	if fake.DescriptionStub != nil {
		return fake.DescriptionStub()
	} else {
		return fake.descriptionReturns.result1
	}
}

func (fake *FakeWorker) DescriptionCallCount() int {
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	return len(fake.descriptionArgsForCall)
}

func (fake *FakeWorker) DescriptionReturns(result1 string) {
	fake.DescriptionStub = nil
	fake.descriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeWorker) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeWorker) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Uptime() time.Duration {
	fake.uptimeMutex.Lock()
	fake.uptimeArgsForCall = append(fake.uptimeArgsForCall, struct{}{})
	fake.recordInvocation("Uptime", []interface{}{})
	fake.uptimeMutex.Unlock()
	if fake.UptimeStub != nil {
		return fake.UptimeStub()
	} else {
		return fake.uptimeReturns.result1
	}
}

func (fake *FakeWorker) UptimeCallCount() int {
	fake.uptimeMutex.RLock()
	defer fake.uptimeMutex.RUnlock()
	return len(fake.uptimeArgsForCall)
}

func (fake *FakeWorker) UptimeReturns(result1 time.Duration) {
	fake.UptimeStub = nil
	fake.uptimeReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeWorker) IsOwnedByTeam() bool {
	fake.isOwnedByTeamMutex.Lock()
	fake.isOwnedByTeamArgsForCall = append(fake.isOwnedByTeamArgsForCall, struct{}{})
	fake.recordInvocation("IsOwnedByTeam", []interface{}{})
	fake.isOwnedByTeamMutex.Unlock()
	if fake.IsOwnedByTeamStub != nil {
		return fake.IsOwnedByTeamStub()
	} else {
		return fake.isOwnedByTeamReturns.result1
	}
}

func (fake *FakeWorker) IsOwnedByTeamCallCount() int {
	fake.isOwnedByTeamMutex.RLock()
	defer fake.isOwnedByTeamMutex.RUnlock()
	return len(fake.isOwnedByTeamArgsForCall)
}

func (fake *FakeWorker) IsOwnedByTeamReturns(result1 bool) {
	fake.IsOwnedByTeamStub = nil
	fake.isOwnedByTeamReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBuildContainerMutex.RLock()
	defer fake.createBuildContainerMutex.RUnlock()
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	fake.createResourceCheckContainerMutex.RLock()
	defer fake.createResourceCheckContainerMutex.RUnlock()
	fake.createResourceTypeCheckContainerMutex.RLock()
	defer fake.createResourceTypeCheckContainerMutex.RUnlock()
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	fake.validateResourceCheckVersionMutex.RLock()
	defer fake.validateResourceCheckVersionMutex.RUnlock()
	fake.findResourceTypeByPathMutex.RLock()
	defer fake.findResourceTypeByPathMutex.RUnlock()
	fake.findVolumeMutex.RLock()
	defer fake.findVolumeMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	fake.descriptionMutex.RLock()
	defer fake.descriptionMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.uptimeMutex.RLock()
	defer fake.uptimeMutex.RUnlock()
	fake.isOwnedByTeamMutex.RLock()
	defer fake.isOwnedByTeamMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.Worker = new(FakeWorker)
