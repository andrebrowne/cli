// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeServicesSummariesActor struct {
	GetServiceSummaryByNameStub        func(string) (v2action.ServiceSummary, v2action.Warnings, error)
	getServiceSummaryByNameMutex       sync.RWMutex
	getServiceSummaryByNameArgsForCall []struct {
		arg1 string
	}
	getServiceSummaryByNameReturns struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	getServiceSummaryByNameReturnsOnCall map[int]struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	GetServiceSummaryForSpaceByNameStub        func(string, string) (v2action.ServiceSummary, v2action.Warnings, error)
	getServiceSummaryForSpaceByNameMutex       sync.RWMutex
	getServiceSummaryForSpaceByNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceSummaryForSpaceByNameReturns struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	getServiceSummaryForSpaceByNameReturnsOnCall map[int]struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	GetServicesSummariesStub        func() ([]v2action.ServiceSummary, v2action.Warnings, error)
	getServicesSummariesMutex       sync.RWMutex
	getServicesSummariesArgsForCall []struct {
	}
	getServicesSummariesReturns struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	getServicesSummariesReturnsOnCall map[int]struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	GetServicesSummariesForSpaceStub        func(string) ([]v2action.ServiceSummary, v2action.Warnings, error)
	getServicesSummariesForSpaceMutex       sync.RWMutex
	getServicesSummariesForSpaceArgsForCall []struct {
		arg1 string
	}
	getServicesSummariesForSpaceReturns struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	getServicesSummariesForSpaceReturnsOnCall map[int]struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryByName(arg1 string) (v2action.ServiceSummary, v2action.Warnings, error) {
	fake.getServiceSummaryByNameMutex.Lock()
	ret, specificReturn := fake.getServiceSummaryByNameReturnsOnCall[len(fake.getServiceSummaryByNameArgsForCall)]
	fake.getServiceSummaryByNameArgsForCall = append(fake.getServiceSummaryByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServiceSummaryByName", []interface{}{arg1})
	fake.getServiceSummaryByNameMutex.Unlock()
	if fake.GetServiceSummaryByNameStub != nil {
		return fake.GetServiceSummaryByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServiceSummaryByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryByNameCallCount() int {
	fake.getServiceSummaryByNameMutex.RLock()
	defer fake.getServiceSummaryByNameMutex.RUnlock()
	return len(fake.getServiceSummaryByNameArgsForCall)
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryByNameCalls(stub func(string) (v2action.ServiceSummary, v2action.Warnings, error)) {
	fake.getServiceSummaryByNameMutex.Lock()
	defer fake.getServiceSummaryByNameMutex.Unlock()
	fake.GetServiceSummaryByNameStub = stub
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryByNameArgsForCall(i int) string {
	fake.getServiceSummaryByNameMutex.RLock()
	defer fake.getServiceSummaryByNameMutex.RUnlock()
	argsForCall := fake.getServiceSummaryByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryByNameReturns(result1 v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServiceSummaryByNameMutex.Lock()
	defer fake.getServiceSummaryByNameMutex.Unlock()
	fake.GetServiceSummaryByNameStub = nil
	fake.getServiceSummaryByNameReturns = struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryByNameReturnsOnCall(i int, result1 v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServiceSummaryByNameMutex.Lock()
	defer fake.getServiceSummaryByNameMutex.Unlock()
	fake.GetServiceSummaryByNameStub = nil
	if fake.getServiceSummaryByNameReturnsOnCall == nil {
		fake.getServiceSummaryByNameReturnsOnCall = make(map[int]struct {
			result1 v2action.ServiceSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceSummaryByNameReturnsOnCall[i] = struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryForSpaceByName(arg1 string, arg2 string) (v2action.ServiceSummary, v2action.Warnings, error) {
	fake.getServiceSummaryForSpaceByNameMutex.Lock()
	ret, specificReturn := fake.getServiceSummaryForSpaceByNameReturnsOnCall[len(fake.getServiceSummaryForSpaceByNameArgsForCall)]
	fake.getServiceSummaryForSpaceByNameArgsForCall = append(fake.getServiceSummaryForSpaceByNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceSummaryForSpaceByName", []interface{}{arg1, arg2})
	fake.getServiceSummaryForSpaceByNameMutex.Unlock()
	if fake.GetServiceSummaryForSpaceByNameStub != nil {
		return fake.GetServiceSummaryForSpaceByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServiceSummaryForSpaceByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryForSpaceByNameCallCount() int {
	fake.getServiceSummaryForSpaceByNameMutex.RLock()
	defer fake.getServiceSummaryForSpaceByNameMutex.RUnlock()
	return len(fake.getServiceSummaryForSpaceByNameArgsForCall)
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryForSpaceByNameCalls(stub func(string, string) (v2action.ServiceSummary, v2action.Warnings, error)) {
	fake.getServiceSummaryForSpaceByNameMutex.Lock()
	defer fake.getServiceSummaryForSpaceByNameMutex.Unlock()
	fake.GetServiceSummaryForSpaceByNameStub = stub
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryForSpaceByNameArgsForCall(i int) (string, string) {
	fake.getServiceSummaryForSpaceByNameMutex.RLock()
	defer fake.getServiceSummaryForSpaceByNameMutex.RUnlock()
	argsForCall := fake.getServiceSummaryForSpaceByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryForSpaceByNameReturns(result1 v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServiceSummaryForSpaceByNameMutex.Lock()
	defer fake.getServiceSummaryForSpaceByNameMutex.Unlock()
	fake.GetServiceSummaryForSpaceByNameStub = nil
	fake.getServiceSummaryForSpaceByNameReturns = struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServiceSummaryForSpaceByNameReturnsOnCall(i int, result1 v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServiceSummaryForSpaceByNameMutex.Lock()
	defer fake.getServiceSummaryForSpaceByNameMutex.Unlock()
	fake.GetServiceSummaryForSpaceByNameStub = nil
	if fake.getServiceSummaryForSpaceByNameReturnsOnCall == nil {
		fake.getServiceSummaryForSpaceByNameReturnsOnCall = make(map[int]struct {
			result1 v2action.ServiceSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceSummaryForSpaceByNameReturnsOnCall[i] = struct {
		result1 v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServicesSummaries() ([]v2action.ServiceSummary, v2action.Warnings, error) {
	fake.getServicesSummariesMutex.Lock()
	ret, specificReturn := fake.getServicesSummariesReturnsOnCall[len(fake.getServicesSummariesArgsForCall)]
	fake.getServicesSummariesArgsForCall = append(fake.getServicesSummariesArgsForCall, struct {
	}{})
	fake.recordInvocation("GetServicesSummaries", []interface{}{})
	fake.getServicesSummariesMutex.Unlock()
	if fake.GetServicesSummariesStub != nil {
		return fake.GetServicesSummariesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServicesSummariesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesCallCount() int {
	fake.getServicesSummariesMutex.RLock()
	defer fake.getServicesSummariesMutex.RUnlock()
	return len(fake.getServicesSummariesArgsForCall)
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesCalls(stub func() ([]v2action.ServiceSummary, v2action.Warnings, error)) {
	fake.getServicesSummariesMutex.Lock()
	defer fake.getServicesSummariesMutex.Unlock()
	fake.GetServicesSummariesStub = stub
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesReturns(result1 []v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServicesSummariesMutex.Lock()
	defer fake.getServicesSummariesMutex.Unlock()
	fake.GetServicesSummariesStub = nil
	fake.getServicesSummariesReturns = struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesReturnsOnCall(i int, result1 []v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServicesSummariesMutex.Lock()
	defer fake.getServicesSummariesMutex.Unlock()
	fake.GetServicesSummariesStub = nil
	if fake.getServicesSummariesReturnsOnCall == nil {
		fake.getServicesSummariesReturnsOnCall = make(map[int]struct {
			result1 []v2action.ServiceSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServicesSummariesReturnsOnCall[i] = struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesForSpace(arg1 string) ([]v2action.ServiceSummary, v2action.Warnings, error) {
	fake.getServicesSummariesForSpaceMutex.Lock()
	ret, specificReturn := fake.getServicesSummariesForSpaceReturnsOnCall[len(fake.getServicesSummariesForSpaceArgsForCall)]
	fake.getServicesSummariesForSpaceArgsForCall = append(fake.getServicesSummariesForSpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicesSummariesForSpace", []interface{}{arg1})
	fake.getServicesSummariesForSpaceMutex.Unlock()
	if fake.GetServicesSummariesForSpaceStub != nil {
		return fake.GetServicesSummariesForSpaceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServicesSummariesForSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesForSpaceCallCount() int {
	fake.getServicesSummariesForSpaceMutex.RLock()
	defer fake.getServicesSummariesForSpaceMutex.RUnlock()
	return len(fake.getServicesSummariesForSpaceArgsForCall)
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesForSpaceCalls(stub func(string) ([]v2action.ServiceSummary, v2action.Warnings, error)) {
	fake.getServicesSummariesForSpaceMutex.Lock()
	defer fake.getServicesSummariesForSpaceMutex.Unlock()
	fake.GetServicesSummariesForSpaceStub = stub
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesForSpaceArgsForCall(i int) string {
	fake.getServicesSummariesForSpaceMutex.RLock()
	defer fake.getServicesSummariesForSpaceMutex.RUnlock()
	argsForCall := fake.getServicesSummariesForSpaceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesForSpaceReturns(result1 []v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServicesSummariesForSpaceMutex.Lock()
	defer fake.getServicesSummariesForSpaceMutex.Unlock()
	fake.GetServicesSummariesForSpaceStub = nil
	fake.getServicesSummariesForSpaceReturns = struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) GetServicesSummariesForSpaceReturnsOnCall(i int, result1 []v2action.ServiceSummary, result2 v2action.Warnings, result3 error) {
	fake.getServicesSummariesForSpaceMutex.Lock()
	defer fake.getServicesSummariesForSpaceMutex.Unlock()
	fake.GetServicesSummariesForSpaceStub = nil
	if fake.getServicesSummariesForSpaceReturnsOnCall == nil {
		fake.getServicesSummariesForSpaceReturnsOnCall = make(map[int]struct {
			result1 []v2action.ServiceSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServicesSummariesForSpaceReturnsOnCall[i] = struct {
		result1 []v2action.ServiceSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServicesSummariesActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceSummaryByNameMutex.RLock()
	defer fake.getServiceSummaryByNameMutex.RUnlock()
	fake.getServiceSummaryForSpaceByNameMutex.RLock()
	defer fake.getServiceSummaryForSpaceByNameMutex.RUnlock()
	fake.getServicesSummariesMutex.RLock()
	defer fake.getServicesSummariesMutex.RUnlock()
	fake.getServicesSummariesForSpaceMutex.RLock()
	defer fake.getServicesSummariesForSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServicesSummariesActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.ServicesSummariesActor = new(FakeServicesSummariesActor)
