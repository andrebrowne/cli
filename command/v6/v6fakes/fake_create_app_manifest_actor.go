// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2v3action "code.cloudfoundry.org/cli/actor/v2v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
	manifest "code.cloudfoundry.org/cli/util/manifest"
)

type FakeCreateAppManifestActor struct {
	CreateApplicationManifestByNameAndSpaceStub        func(string, string) (manifest.Application, v2v3action.Warnings, error)
	createApplicationManifestByNameAndSpaceMutex       sync.RWMutex
	createApplicationManifestByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createApplicationManifestByNameAndSpaceReturns struct {
		result1 manifest.Application
		result2 v2v3action.Warnings
		result3 error
	}
	createApplicationManifestByNameAndSpaceReturnsOnCall map[int]struct {
		result1 manifest.Application
		result2 v2v3action.Warnings
		result3 error
	}
	WriteApplicationManifestStub        func(manifest.Application, string) error
	writeApplicationManifestMutex       sync.RWMutex
	writeApplicationManifestArgsForCall []struct {
		arg1 manifest.Application
		arg2 string
	}
	writeApplicationManifestReturns struct {
		result1 error
	}
	writeApplicationManifestReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateAppManifestActor) CreateApplicationManifestByNameAndSpace(arg1 string, arg2 string) (manifest.Application, v2v3action.Warnings, error) {
	fake.createApplicationManifestByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.createApplicationManifestByNameAndSpaceReturnsOnCall[len(fake.createApplicationManifestByNameAndSpaceArgsForCall)]
	fake.createApplicationManifestByNameAndSpaceArgsForCall = append(fake.createApplicationManifestByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateApplicationManifestByNameAndSpace", []interface{}{arg1, arg2})
	fake.createApplicationManifestByNameAndSpaceMutex.Unlock()
	if fake.CreateApplicationManifestByNameAndSpaceStub != nil {
		return fake.CreateApplicationManifestByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createApplicationManifestByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCreateAppManifestActor) CreateApplicationManifestByNameAndSpaceCallCount() int {
	fake.createApplicationManifestByNameAndSpaceMutex.RLock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.RUnlock()
	return len(fake.createApplicationManifestByNameAndSpaceArgsForCall)
}

func (fake *FakeCreateAppManifestActor) CreateApplicationManifestByNameAndSpaceCalls(stub func(string, string) (manifest.Application, v2v3action.Warnings, error)) {
	fake.createApplicationManifestByNameAndSpaceMutex.Lock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.Unlock()
	fake.CreateApplicationManifestByNameAndSpaceStub = stub
}

func (fake *FakeCreateAppManifestActor) CreateApplicationManifestByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.createApplicationManifestByNameAndSpaceMutex.RLock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.createApplicationManifestByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreateAppManifestActor) CreateApplicationManifestByNameAndSpaceReturns(result1 manifest.Application, result2 v2v3action.Warnings, result3 error) {
	fake.createApplicationManifestByNameAndSpaceMutex.Lock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.Unlock()
	fake.CreateApplicationManifestByNameAndSpaceStub = nil
	fake.createApplicationManifestByNameAndSpaceReturns = struct {
		result1 manifest.Application
		result2 v2v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateAppManifestActor) CreateApplicationManifestByNameAndSpaceReturnsOnCall(i int, result1 manifest.Application, result2 v2v3action.Warnings, result3 error) {
	fake.createApplicationManifestByNameAndSpaceMutex.Lock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.Unlock()
	fake.CreateApplicationManifestByNameAndSpaceStub = nil
	if fake.createApplicationManifestByNameAndSpaceReturnsOnCall == nil {
		fake.createApplicationManifestByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 manifest.Application
			result2 v2v3action.Warnings
			result3 error
		})
	}
	fake.createApplicationManifestByNameAndSpaceReturnsOnCall[i] = struct {
		result1 manifest.Application
		result2 v2v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCreateAppManifestActor) WriteApplicationManifest(arg1 manifest.Application, arg2 string) error {
	fake.writeApplicationManifestMutex.Lock()
	ret, specificReturn := fake.writeApplicationManifestReturnsOnCall[len(fake.writeApplicationManifestArgsForCall)]
	fake.writeApplicationManifestArgsForCall = append(fake.writeApplicationManifestArgsForCall, struct {
		arg1 manifest.Application
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("WriteApplicationManifest", []interface{}{arg1, arg2})
	fake.writeApplicationManifestMutex.Unlock()
	if fake.WriteApplicationManifestStub != nil {
		return fake.WriteApplicationManifestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.writeApplicationManifestReturns
	return fakeReturns.result1
}

func (fake *FakeCreateAppManifestActor) WriteApplicationManifestCallCount() int {
	fake.writeApplicationManifestMutex.RLock()
	defer fake.writeApplicationManifestMutex.RUnlock()
	return len(fake.writeApplicationManifestArgsForCall)
}

func (fake *FakeCreateAppManifestActor) WriteApplicationManifestCalls(stub func(manifest.Application, string) error) {
	fake.writeApplicationManifestMutex.Lock()
	defer fake.writeApplicationManifestMutex.Unlock()
	fake.WriteApplicationManifestStub = stub
}

func (fake *FakeCreateAppManifestActor) WriteApplicationManifestArgsForCall(i int) (manifest.Application, string) {
	fake.writeApplicationManifestMutex.RLock()
	defer fake.writeApplicationManifestMutex.RUnlock()
	argsForCall := fake.writeApplicationManifestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreateAppManifestActor) WriteApplicationManifestReturns(result1 error) {
	fake.writeApplicationManifestMutex.Lock()
	defer fake.writeApplicationManifestMutex.Unlock()
	fake.WriteApplicationManifestStub = nil
	fake.writeApplicationManifestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreateAppManifestActor) WriteApplicationManifestReturnsOnCall(i int, result1 error) {
	fake.writeApplicationManifestMutex.Lock()
	defer fake.writeApplicationManifestMutex.Unlock()
	fake.WriteApplicationManifestStub = nil
	if fake.writeApplicationManifestReturnsOnCall == nil {
		fake.writeApplicationManifestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeApplicationManifestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreateAppManifestActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createApplicationManifestByNameAndSpaceMutex.RLock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.RUnlock()
	fake.writeApplicationManifestMutex.RLock()
	defer fake.writeApplicationManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateAppManifestActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.CreateAppManifestActor = new(FakeCreateAppManifestActor)
