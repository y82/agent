// Code generated by counterfeiter. DO NOT EDIT.
package nginx

import (
	"sync"

	"github.com/google/uuid"
	"github.com/nginx/agent/v3/api/grpc/instances"
)

type FakeNginxConfigInterface struct {
	CompleteStub        func() error
	completeMutex       sync.RWMutex
	completeArgsForCall []struct {
	}
	completeReturns struct {
		result1 error
	}
	completeReturnsOnCall map[int]struct {
		result1 error
	}
	ReloadStub        func() error
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct {
	}
	reloadReturns struct {
		result1 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func(map[string]*instances.File, string, uuid.UUID) (map[string]*instances.File, map[string]struct{}, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 map[string]*instances.File
		arg2 string
		arg3 uuid.UUID
	}
	writeReturns struct {
		result1 map[string]*instances.File
		result2 map[string]struct{}
		result3 error
	}
	writeReturnsOnCall map[int]struct {
		result1 map[string]*instances.File
		result2 map[string]struct{}
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNginxConfigInterface) Complete() error {
	fake.completeMutex.Lock()
	ret, specificReturn := fake.completeReturnsOnCall[len(fake.completeArgsForCall)]
	fake.completeArgsForCall = append(fake.completeArgsForCall, struct {
	}{})
	stub := fake.CompleteStub
	fakeReturns := fake.completeReturns
	fake.recordInvocation("Complete", []interface{}{})
	fake.completeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNginxConfigInterface) CompleteCallCount() int {
	fake.completeMutex.RLock()
	defer fake.completeMutex.RUnlock()
	return len(fake.completeArgsForCall)
}

func (fake *FakeNginxConfigInterface) CompleteCalls(stub func() error) {
	fake.completeMutex.Lock()
	defer fake.completeMutex.Unlock()
	fake.CompleteStub = stub
}

func (fake *FakeNginxConfigInterface) CompleteReturns(result1 error) {
	fake.completeMutex.Lock()
	defer fake.completeMutex.Unlock()
	fake.CompleteStub = nil
	fake.completeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNginxConfigInterface) CompleteReturnsOnCall(i int, result1 error) {
	fake.completeMutex.Lock()
	defer fake.completeMutex.Unlock()
	fake.CompleteStub = nil
	if fake.completeReturnsOnCall == nil {
		fake.completeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.completeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNginxConfigInterface) Reload() error {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct {
	}{})
	stub := fake.ReloadStub
	fakeReturns := fake.reloadReturns
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNginxConfigInterface) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeNginxConfigInterface) ReloadCalls(stub func() error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = stub
}

func (fake *FakeNginxConfigInterface) ReloadReturns(result1 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNginxConfigInterface) ReloadReturnsOnCall(i int, result1 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNginxConfigInterface) Validate() error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
	}{})
	stub := fake.ValidateStub
	fakeReturns := fake.validateReturns
	fake.recordInvocation("Validate", []interface{}{})
	fake.validateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeNginxConfigInterface) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeNginxConfigInterface) ValidateCalls(stub func() error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *FakeNginxConfigInterface) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNginxConfigInterface) ValidateReturnsOnCall(i int, result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNginxConfigInterface) Write(arg1 map[string]*instances.File, arg2 string, arg3 uuid.UUID) (map[string]*instances.File, map[string]struct{}, error) {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 map[string]*instances.File
		arg2 string
		arg3 uuid.UUID
	}{arg1, arg2, arg3})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1, arg2, arg3})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeNginxConfigInterface) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeNginxConfigInterface) WriteCalls(stub func(map[string]*instances.File, string, uuid.UUID) (map[string]*instances.File, map[string]struct{}, error)) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeNginxConfigInterface) WriteArgsForCall(i int) (map[string]*instances.File, string, uuid.UUID) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNginxConfigInterface) WriteReturns(result1 map[string]*instances.File, result2 map[string]struct{}, result3 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 map[string]*instances.File
		result2 map[string]struct{}
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNginxConfigInterface) WriteReturnsOnCall(i int, result1 map[string]*instances.File, result2 map[string]struct{}, result3 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 map[string]*instances.File
			result2 map[string]struct{}
			result3 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 map[string]*instances.File
		result2 map[string]struct{}
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNginxConfigInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.completeMutex.RLock()
	defer fake.completeMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNginxConfigInterface) recordInvocation(key string, args []interface{}) {
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

var _ NginxConfigInterface = new(FakeNginxConfigInterface)
