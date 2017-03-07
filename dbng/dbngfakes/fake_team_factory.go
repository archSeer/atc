// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeTeamFactory struct {
	CreateTeamStub        func(name string) (dbng.Team, error)
	createTeamMutex       sync.RWMutex
	createTeamArgsForCall []struct {
		name string
	}
	createTeamReturns struct {
		result1 dbng.Team
		result2 error
	}
	createTeamReturnsOnCall map[int]struct {
		result1 dbng.Team
		result2 error
	}
	FindTeamStub        func(name string) (dbng.Team, bool, error)
	findTeamMutex       sync.RWMutex
	findTeamArgsForCall []struct {
		name string
	}
	findTeamReturns struct {
		result1 dbng.Team
		result2 bool
		result3 error
	}
	findTeamReturnsOnCall map[int]struct {
		result1 dbng.Team
		result2 bool
		result3 error
	}
	GetByIDStub        func(teamID int) dbng.Team
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		teamID int
	}
	getByIDReturns struct {
		result1 dbng.Team
	}
	getByIDReturnsOnCall map[int]struct {
		result1 dbng.Team
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamFactory) CreateTeam(name string) (dbng.Team, error) {
	fake.createTeamMutex.Lock()
	ret, specificReturn := fake.createTeamReturnsOnCall[len(fake.createTeamArgsForCall)]
	fake.createTeamArgsForCall = append(fake.createTeamArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("CreateTeam", []interface{}{name})
	fake.createTeamMutex.Unlock()
	if fake.CreateTeamStub != nil {
		return fake.CreateTeamStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTeamReturns.result1, fake.createTeamReturns.result2
}

func (fake *FakeTeamFactory) CreateTeamCallCount() int {
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return len(fake.createTeamArgsForCall)
}

func (fake *FakeTeamFactory) CreateTeamArgsForCall(i int) string {
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return fake.createTeamArgsForCall[i].name
}

func (fake *FakeTeamFactory) CreateTeamReturns(result1 dbng.Team, result2 error) {
	fake.CreateTeamStub = nil
	fake.createTeamReturns = struct {
		result1 dbng.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamFactory) CreateTeamReturnsOnCall(i int, result1 dbng.Team, result2 error) {
	fake.CreateTeamStub = nil
	if fake.createTeamReturnsOnCall == nil {
		fake.createTeamReturnsOnCall = make(map[int]struct {
			result1 dbng.Team
			result2 error
		})
	}
	fake.createTeamReturnsOnCall[i] = struct {
		result1 dbng.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamFactory) FindTeam(name string) (dbng.Team, bool, error) {
	fake.findTeamMutex.Lock()
	ret, specificReturn := fake.findTeamReturnsOnCall[len(fake.findTeamArgsForCall)]
	fake.findTeamArgsForCall = append(fake.findTeamArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindTeam", []interface{}{name})
	fake.findTeamMutex.Unlock()
	if fake.FindTeamStub != nil {
		return fake.FindTeamStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findTeamReturns.result1, fake.findTeamReturns.result2, fake.findTeamReturns.result3
}

func (fake *FakeTeamFactory) FindTeamCallCount() int {
	fake.findTeamMutex.RLock()
	defer fake.findTeamMutex.RUnlock()
	return len(fake.findTeamArgsForCall)
}

func (fake *FakeTeamFactory) FindTeamArgsForCall(i int) string {
	fake.findTeamMutex.RLock()
	defer fake.findTeamMutex.RUnlock()
	return fake.findTeamArgsForCall[i].name
}

func (fake *FakeTeamFactory) FindTeamReturns(result1 dbng.Team, result2 bool, result3 error) {
	fake.FindTeamStub = nil
	fake.findTeamReturns = struct {
		result1 dbng.Team
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamFactory) FindTeamReturnsOnCall(i int, result1 dbng.Team, result2 bool, result3 error) {
	fake.FindTeamStub = nil
	if fake.findTeamReturnsOnCall == nil {
		fake.findTeamReturnsOnCall = make(map[int]struct {
			result1 dbng.Team
			result2 bool
			result3 error
		})
	}
	fake.findTeamReturnsOnCall[i] = struct {
		result1 dbng.Team
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamFactory) GetByID(teamID int) dbng.Team {
	fake.getByIDMutex.Lock()
	ret, specificReturn := fake.getByIDReturnsOnCall[len(fake.getByIDArgsForCall)]
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		teamID int
	}{teamID})
	fake.recordInvocation("GetByID", []interface{}{teamID})
	fake.getByIDMutex.Unlock()
	if fake.GetByIDStub != nil {
		return fake.GetByIDStub(teamID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getByIDReturns.result1
}

func (fake *FakeTeamFactory) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeTeamFactory) GetByIDArgsForCall(i int) int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return fake.getByIDArgsForCall[i].teamID
}

func (fake *FakeTeamFactory) GetByIDReturns(result1 dbng.Team) {
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 dbng.Team
	}{result1}
}

func (fake *FakeTeamFactory) GetByIDReturnsOnCall(i int, result1 dbng.Team) {
	fake.GetByIDStub = nil
	if fake.getByIDReturnsOnCall == nil {
		fake.getByIDReturnsOnCall = make(map[int]struct {
			result1 dbng.Team
		})
	}
	fake.getByIDReturnsOnCall[i] = struct {
		result1 dbng.Team
	}{result1}
}

func (fake *FakeTeamFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	fake.findTeamMutex.RLock()
	defer fake.findTeamMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeamFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.TeamFactory = new(FakeTeamFactory)
