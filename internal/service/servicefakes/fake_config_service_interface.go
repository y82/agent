// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/nginx/agent/v3/api/grpc/instances"
	v1 "github.com/nginx/agent/v3/api/grpc/mpi/v1"
	"github.com/nginx/agent/v3/internal/service"
)

type FakeConfigServiceInterface struct {
	ParseInstanceConfigurationStub        func(context.Context) (any, error)
	parseInstanceConfigurationMutex       sync.RWMutex
	parseInstanceConfigurationArgsForCall []struct {
		arg1 context.Context
	}
	parseInstanceConfigurationReturns struct {
		result1 any
		result2 error
	}
	parseInstanceConfigurationReturnsOnCall map[int]struct {
		result1 any
		result2 error
	}
	RollbackStub        func(context.Context, map[string]*v1.FileMeta, string, string, string) error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct {
		arg1 context.Context
		arg2 map[string]*v1.FileMeta
		arg3 string
		arg4 string
		arg5 string
	}
	rollbackReturns struct {
		result1 error
	}
	rollbackReturnsOnCall map[int]struct {
		result1 error
	}
	SetConfigContextStub        func(any)
	setConfigContextMutex       sync.RWMutex
	setConfigContextArgsForCall []struct {
		arg1 any
	}
	UpdateInstanceConfigurationStub        func(context.Context, string) (map[string]*v1.FileMeta, *instances.ConfigurationStatus)
	updateInstanceConfigurationMutex       sync.RWMutex
	updateInstanceConfigurationArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	updateInstanceConfigurationReturns struct {
		result1 map[string]*v1.FileMeta
		result2 *instances.ConfigurationStatus
	}
	updateInstanceConfigurationReturnsOnCall map[int]struct {
		result1 map[string]*v1.FileMeta
		result2 *instances.ConfigurationStatus
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigServiceInterface) ParseInstanceConfiguration(arg1 context.Context) (any, error) {
	fake.parseInstanceConfigurationMutex.Lock()
	ret, specificReturn := fake.parseInstanceConfigurationReturnsOnCall[len(fake.parseInstanceConfigurationArgsForCall)]
	fake.parseInstanceConfigurationArgsForCall = append(fake.parseInstanceConfigurationArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ParseInstanceConfigurationStub
	fakeReturns := fake.parseInstanceConfigurationReturns
	fake.recordInvocation("ParseInstanceConfiguration", []interface{}{arg1})
	fake.parseInstanceConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigServiceInterface) ParseInstanceConfigurationCallCount() int {
	fake.parseInstanceConfigurationMutex.RLock()
	defer fake.parseInstanceConfigurationMutex.RUnlock()
	return len(fake.parseInstanceConfigurationArgsForCall)
}

func (fake *FakeConfigServiceInterface) ParseInstanceConfigurationCalls(stub func(context.Context) (any, error)) {
	fake.parseInstanceConfigurationMutex.Lock()
	defer fake.parseInstanceConfigurationMutex.Unlock()
	fake.ParseInstanceConfigurationStub = stub
}

func (fake *FakeConfigServiceInterface) ParseInstanceConfigurationArgsForCall(i int) context.Context {
	fake.parseInstanceConfigurationMutex.RLock()
	defer fake.parseInstanceConfigurationMutex.RUnlock()
	argsForCall := fake.parseInstanceConfigurationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigServiceInterface) ParseInstanceConfigurationReturns(result1 any, result2 error) {
	fake.parseInstanceConfigurationMutex.Lock()
	defer fake.parseInstanceConfigurationMutex.Unlock()
	fake.ParseInstanceConfigurationStub = nil
	fake.parseInstanceConfigurationReturns = struct {
		result1 any
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigServiceInterface) ParseInstanceConfigurationReturnsOnCall(i int, result1 any, result2 error) {
	fake.parseInstanceConfigurationMutex.Lock()
	defer fake.parseInstanceConfigurationMutex.Unlock()
	fake.ParseInstanceConfigurationStub = nil
	if fake.parseInstanceConfigurationReturnsOnCall == nil {
		fake.parseInstanceConfigurationReturnsOnCall = make(map[int]struct {
			result1 any
			result2 error
		})
	}
	fake.parseInstanceConfigurationReturnsOnCall[i] = struct {
		result1 any
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigServiceInterface) Rollback(arg1 context.Context, arg2 map[string]*v1.FileMeta, arg3 string, arg4 string, arg5 string) error {
	fake.rollbackMutex.Lock()
	ret, specificReturn := fake.rollbackReturnsOnCall[len(fake.rollbackArgsForCall)]
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct {
		arg1 context.Context
		arg2 map[string]*v1.FileMeta
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.RollbackStub
	fakeReturns := fake.rollbackReturns
	fake.recordInvocation("Rollback", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.rollbackMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigServiceInterface) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeConfigServiceInterface) RollbackCalls(stub func(context.Context, map[string]*v1.FileMeta, string, string, string) error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = stub
}

func (fake *FakeConfigServiceInterface) RollbackArgsForCall(i int) (context.Context, map[string]*v1.FileMeta, string, string, string) {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	argsForCall := fake.rollbackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeConfigServiceInterface) RollbackReturns(result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigServiceInterface) RollbackReturnsOnCall(i int, result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	if fake.rollbackReturnsOnCall == nil {
		fake.rollbackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rollbackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigServiceInterface) SetConfigContext(arg1 any) {
	fake.setConfigContextMutex.Lock()
	fake.setConfigContextArgsForCall = append(fake.setConfigContextArgsForCall, struct {
		arg1 any
	}{arg1})
	stub := fake.SetConfigContextStub
	fake.recordInvocation("SetConfigContext", []interface{}{arg1})
	fake.setConfigContextMutex.Unlock()
	if stub != nil {
		fake.SetConfigContextStub(arg1)
	}
}

func (fake *FakeConfigServiceInterface) SetConfigContextCallCount() int {
	fake.setConfigContextMutex.RLock()
	defer fake.setConfigContextMutex.RUnlock()
	return len(fake.setConfigContextArgsForCall)
}

func (fake *FakeConfigServiceInterface) SetConfigContextCalls(stub func(any)) {
	fake.setConfigContextMutex.Lock()
	defer fake.setConfigContextMutex.Unlock()
	fake.SetConfigContextStub = stub
}

func (fake *FakeConfigServiceInterface) SetConfigContextArgsForCall(i int) any {
	fake.setConfigContextMutex.RLock()
	defer fake.setConfigContextMutex.RUnlock()
	argsForCall := fake.setConfigContextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigServiceInterface) UpdateInstanceConfiguration(arg1 context.Context, arg2 string) (map[string]*v1.FileMeta, *instances.ConfigurationStatus) {
	fake.updateInstanceConfigurationMutex.Lock()
	ret, specificReturn := fake.updateInstanceConfigurationReturnsOnCall[len(fake.updateInstanceConfigurationArgsForCall)]
	fake.updateInstanceConfigurationArgsForCall = append(fake.updateInstanceConfigurationArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.UpdateInstanceConfigurationStub
	fakeReturns := fake.updateInstanceConfigurationReturns
	fake.recordInvocation("UpdateInstanceConfiguration", []interface{}{arg1, arg2})
	fake.updateInstanceConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigServiceInterface) UpdateInstanceConfigurationCallCount() int {
	fake.updateInstanceConfigurationMutex.RLock()
	defer fake.updateInstanceConfigurationMutex.RUnlock()
	return len(fake.updateInstanceConfigurationArgsForCall)
}

func (fake *FakeConfigServiceInterface) UpdateInstanceConfigurationCalls(stub func(context.Context, string) (map[string]*v1.FileMeta, *instances.ConfigurationStatus)) {
	fake.updateInstanceConfigurationMutex.Lock()
	defer fake.updateInstanceConfigurationMutex.Unlock()
	fake.UpdateInstanceConfigurationStub = stub
}

func (fake *FakeConfigServiceInterface) UpdateInstanceConfigurationArgsForCall(i int) (context.Context, string) {
	fake.updateInstanceConfigurationMutex.RLock()
	defer fake.updateInstanceConfigurationMutex.RUnlock()
	argsForCall := fake.updateInstanceConfigurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigServiceInterface) UpdateInstanceConfigurationReturns(result1 map[string]*v1.FileMeta, result2 *instances.ConfigurationStatus) {
	fake.updateInstanceConfigurationMutex.Lock()
	defer fake.updateInstanceConfigurationMutex.Unlock()
	fake.UpdateInstanceConfigurationStub = nil
	fake.updateInstanceConfigurationReturns = struct {
		result1 map[string]*v1.FileMeta
		result2 *instances.ConfigurationStatus
	}{result1, result2}
}

func (fake *FakeConfigServiceInterface) UpdateInstanceConfigurationReturnsOnCall(i int, result1 map[string]*v1.FileMeta, result2 *instances.ConfigurationStatus) {
	fake.updateInstanceConfigurationMutex.Lock()
	defer fake.updateInstanceConfigurationMutex.Unlock()
	fake.UpdateInstanceConfigurationStub = nil
	if fake.updateInstanceConfigurationReturnsOnCall == nil {
		fake.updateInstanceConfigurationReturnsOnCall = make(map[int]struct {
			result1 map[string]*v1.FileMeta
			result2 *instances.ConfigurationStatus
		})
	}
	fake.updateInstanceConfigurationReturnsOnCall[i] = struct {
		result1 map[string]*v1.FileMeta
		result2 *instances.ConfigurationStatus
	}{result1, result2}
}

func (fake *FakeConfigServiceInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseInstanceConfigurationMutex.RLock()
	defer fake.parseInstanceConfigurationMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	fake.setConfigContextMutex.RLock()
	defer fake.setConfigContextMutex.RUnlock()
	fake.updateInstanceConfigurationMutex.RLock()
	defer fake.updateInstanceConfigurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfigServiceInterface) recordInvocation(key string, args []interface{}) {
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

var _ service.ConfigServiceInterface = new(FakeConfigServiceInterface)
