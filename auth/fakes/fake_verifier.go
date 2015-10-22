// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/concourse/atc/auth"
)

type FakeVerifier struct {
	VerifyStub        func(*http.Client) (bool, error)
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		arg1 *http.Client
	}
	verifyReturns struct {
		result1 bool
		result2 error
	}
}

func (fake *FakeVerifier) Verify(arg1 *http.Client) (bool, error) {
	fake.verifyMutex.Lock()
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		arg1 *http.Client
	}{arg1})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(arg1)
	} else {
		return fake.verifyReturns.result1, fake.verifyReturns.result2
	}
}

func (fake *FakeVerifier) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeVerifier) VerifyArgsForCall(i int) *http.Client {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.verifyArgsForCall[i].arg1
}

func (fake *FakeVerifier) VerifyReturns(result1 bool, result2 error) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

var _ auth.Verifier = new(FakeVerifier)