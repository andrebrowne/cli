// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeSpacesActor struct {
	GetOrganizationSpacesStub        func(string) ([]v2action.Space, v2action.Warnings, error)
	getOrganizationSpacesMutex       sync.RWMutex
	getOrganizationSpacesArgsForCall []struct {
		arg1 string
	}
	getOrganizationSpacesReturns struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationSpacesReturnsOnCall map[int]struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpacesActor) GetOrganizationSpaces(arg1 string) ([]v2action.Space, v2action.Warnings, error) {
	fake.getOrganizationSpacesMutex.Lock()
	ret, specificReturn := fake.getOrganizationSpacesReturnsOnCall[len(fake.getOrganizationSpacesArgsForCall)]
	fake.getOrganizationSpacesArgsForCall = append(fake.getOrganizationSpacesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationSpaces", []interface{}{arg1})
	fake.getOrganizationSpacesMutex.Unlock()
	if fake.GetOrganizationSpacesStub != nil {
		return fake.GetOrganizationSpacesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationSpacesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpacesActor) GetOrganizationSpacesCallCount() int {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	return len(fake.getOrganizationSpacesArgsForCall)
}

func (fake *FakeSpacesActor) GetOrganizationSpacesCalls(stub func(string) ([]v2action.Space, v2action.Warnings, error)) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = stub
}

func (fake *FakeSpacesActor) GetOrganizationSpacesArgsForCall(i int) string {
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	argsForCall := fake.getOrganizationSpacesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpacesActor) GetOrganizationSpacesReturns(result1 []v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = nil
	fake.getOrganizationSpacesReturns = struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpacesActor) GetOrganizationSpacesReturnsOnCall(i int, result1 []v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.getOrganizationSpacesMutex.Lock()
	defer fake.getOrganizationSpacesMutex.Unlock()
	fake.GetOrganizationSpacesStub = nil
	if fake.getOrganizationSpacesReturnsOnCall == nil {
		fake.getOrganizationSpacesReturnsOnCall = make(map[int]struct {
			result1 []v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationSpacesReturnsOnCall[i] = struct {
		result1 []v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpacesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationSpacesMutex.RLock()
	defer fake.getOrganizationSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpacesActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.SpacesActor = new(FakeSpacesActor)
