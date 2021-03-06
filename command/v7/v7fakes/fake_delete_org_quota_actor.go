// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDeleteOrgQuotaActor struct {
	DeleteOrganizationQuotaStub        func(string) (v7action.Warnings, error)
	deleteOrganizationQuotaMutex       sync.RWMutex
	deleteOrganizationQuotaArgsForCall []struct {
		arg1 string
	}
	deleteOrganizationQuotaReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteOrganizationQuotaReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrgQuotaActor) DeleteOrganizationQuota(arg1 string) (v7action.Warnings, error) {
	fake.deleteOrganizationQuotaMutex.Lock()
	ret, specificReturn := fake.deleteOrganizationQuotaReturnsOnCall[len(fake.deleteOrganizationQuotaArgsForCall)]
	fake.deleteOrganizationQuotaArgsForCall = append(fake.deleteOrganizationQuotaArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteOrganizationQuota", []interface{}{arg1})
	fake.deleteOrganizationQuotaMutex.Unlock()
	if fake.DeleteOrganizationQuotaStub != nil {
		return fake.DeleteOrganizationQuotaStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteOrganizationQuotaReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteOrgQuotaActor) DeleteOrganizationQuotaCallCount() int {
	fake.deleteOrganizationQuotaMutex.RLock()
	defer fake.deleteOrganizationQuotaMutex.RUnlock()
	return len(fake.deleteOrganizationQuotaArgsForCall)
}

func (fake *FakeDeleteOrgQuotaActor) DeleteOrganizationQuotaCalls(stub func(string) (v7action.Warnings, error)) {
	fake.deleteOrganizationQuotaMutex.Lock()
	defer fake.deleteOrganizationQuotaMutex.Unlock()
	fake.DeleteOrganizationQuotaStub = stub
}

func (fake *FakeDeleteOrgQuotaActor) DeleteOrganizationQuotaArgsForCall(i int) string {
	fake.deleteOrganizationQuotaMutex.RLock()
	defer fake.deleteOrganizationQuotaMutex.RUnlock()
	argsForCall := fake.deleteOrganizationQuotaArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDeleteOrgQuotaActor) DeleteOrganizationQuotaReturns(result1 v7action.Warnings, result2 error) {
	fake.deleteOrganizationQuotaMutex.Lock()
	defer fake.deleteOrganizationQuotaMutex.Unlock()
	fake.DeleteOrganizationQuotaStub = nil
	fake.deleteOrganizationQuotaReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrgQuotaActor) DeleteOrganizationQuotaReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deleteOrganizationQuotaMutex.Lock()
	defer fake.deleteOrganizationQuotaMutex.Unlock()
	fake.DeleteOrganizationQuotaStub = nil
	if fake.deleteOrganizationQuotaReturnsOnCall == nil {
		fake.deleteOrganizationQuotaReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteOrganizationQuotaReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrgQuotaActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteOrganizationQuotaMutex.RLock()
	defer fake.deleteOrganizationQuotaMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteOrgQuotaActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DeleteOrgQuotaActor = new(FakeDeleteOrgQuotaActor)
