// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/exec"
)

type FakeFactory struct {
	GetStub        func(exec.SessionID, exec.IOConfig, atc.ResourceConfig, atc.Params, atc.Version) exec.Step
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 atc.ResourceConfig
		arg4 atc.Params
		arg5 atc.Version
	}
	getReturns struct {
		result1 exec.Step
	}
	PutStub        func(exec.SessionID, exec.IOConfig, atc.ResourceConfig, atc.Params) exec.Step
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 atc.ResourceConfig
		arg4 atc.Params
	}
	putReturns struct {
		result1 exec.Step
	}
	ExecuteStub        func(exec.SessionID, exec.IOConfig, exec.Privileged, exec.BuildConfigSource) exec.Step
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 exec.Privileged
		arg4 exec.BuildConfigSource
	}
	executeReturns struct {
		result1 exec.Step
	}
	HijackStub        func(exec.SessionID, exec.IOConfig, atc.HijackProcessSpec) (exec.HijackedProcess, error)
	hijackMutex       sync.RWMutex
	hijackArgsForCall []struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 atc.HijackProcessSpec
	}
	hijackReturns struct {
		result1 exec.HijackedProcess
		result2 error
	}
}

func (fake *FakeFactory) Get(arg1 exec.SessionID, arg2 exec.IOConfig, arg3 atc.ResourceConfig, arg4 atc.Params, arg5 atc.Version) exec.Step {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 atc.ResourceConfig
		arg4 atc.Params
		arg5 atc.Version
	}{arg1, arg2, arg3, arg4, arg5})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeFactory) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFactory) GetArgsForCall(i int) (exec.SessionID, exec.IOConfig, atc.ResourceConfig, atc.Params, atc.Version) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3, fake.getArgsForCall[i].arg4, fake.getArgsForCall[i].arg5
}

func (fake *FakeFactory) GetReturns(result1 exec.Step) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeFactory) Put(arg1 exec.SessionID, arg2 exec.IOConfig, arg3 atc.ResourceConfig, arg4 atc.Params) exec.Step {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 atc.ResourceConfig
		arg4 atc.Params
	}{arg1, arg2, arg3, arg4})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.putReturns.result1
	}
}

func (fake *FakeFactory) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeFactory) PutArgsForCall(i int) (exec.SessionID, exec.IOConfig, atc.ResourceConfig, atc.Params) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3, fake.putArgsForCall[i].arg4
}

func (fake *FakeFactory) PutReturns(result1 exec.Step) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeFactory) Execute(arg1 exec.SessionID, arg2 exec.IOConfig, arg3 exec.Privileged, arg4 exec.BuildConfigSource) exec.Step {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 exec.Privileged
		arg4 exec.BuildConfigSource
	}{arg1, arg2, arg3, arg4})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeFactory) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeFactory) ExecuteArgsForCall(i int) (exec.SessionID, exec.IOConfig, exec.Privileged, exec.BuildConfigSource) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].arg1, fake.executeArgsForCall[i].arg2, fake.executeArgsForCall[i].arg3, fake.executeArgsForCall[i].arg4
}

func (fake *FakeFactory) ExecuteReturns(result1 exec.Step) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 exec.Step
	}{result1}
}

func (fake *FakeFactory) Hijack(arg1 exec.SessionID, arg2 exec.IOConfig, arg3 atc.HijackProcessSpec) (exec.HijackedProcess, error) {
	fake.hijackMutex.Lock()
	fake.hijackArgsForCall = append(fake.hijackArgsForCall, struct {
		arg1 exec.SessionID
		arg2 exec.IOConfig
		arg3 atc.HijackProcessSpec
	}{arg1, arg2, arg3})
	fake.hijackMutex.Unlock()
	if fake.HijackStub != nil {
		return fake.HijackStub(arg1, arg2, arg3)
	} else {
		return fake.hijackReturns.result1, fake.hijackReturns.result2
	}
}

func (fake *FakeFactory) HijackCallCount() int {
	fake.hijackMutex.RLock()
	defer fake.hijackMutex.RUnlock()
	return len(fake.hijackArgsForCall)
}

func (fake *FakeFactory) HijackArgsForCall(i int) (exec.SessionID, exec.IOConfig, atc.HijackProcessSpec) {
	fake.hijackMutex.RLock()
	defer fake.hijackMutex.RUnlock()
	return fake.hijackArgsForCall[i].arg1, fake.hijackArgsForCall[i].arg2, fake.hijackArgsForCall[i].arg3
}

func (fake *FakeFactory) HijackReturns(result1 exec.HijackedProcess, result2 error) {
	fake.HijackStub = nil
	fake.hijackReturns = struct {
		result1 exec.HijackedProcess
		result2 error
	}{result1, result2}
}

var _ exec.Factory = new(FakeFactory)