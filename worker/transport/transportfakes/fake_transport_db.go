// This file was generated by counterfeiter
package transportfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker/transport"
)

type FakeTransportDB struct {
	GetWorkerStub        func(name string) (dbng.Worker, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		name string
	}
	getWorkerReturns struct {
		result1 dbng.Worker
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTransportDB) GetWorker(name string) (dbng.Worker, bool, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetWorker", []interface{}{name})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(name)
	}
	return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2, fake.getWorkerReturns.result3
}

func (fake *FakeTransportDB) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeTransportDB) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].name
}

func (fake *FakeTransportDB) GetWorkerReturns(result1 dbng.Worker, result2 bool, result3 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 dbng.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTransportDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTransportDB) recordInvocation(key string, args []interface{}) {
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

var _ transport.TransportDB = new(FakeTransportDB)
