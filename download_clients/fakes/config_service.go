// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	download_clients "github.com/pivotal-cf/om/download_clients"
)

type Config struct {
	ConfigStub        func(string) (string, bool)
	configMutex       sync.RWMutex
	configArgsForCall []struct {
		arg1 string
	}
	configReturns struct {
		result1 string
		result2 bool
	}
	configReturnsOnCall map[int]struct {
		result1 string
		result2 bool
	}
	SetStub        func(string, string)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 string
		arg2 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Config) Config(arg1 string) (string, bool) {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Config", []interface{}{arg1})
	fake.configMutex.Unlock()
	if fake.ConfigStub != nil {
		return fake.ConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.configReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Config) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *Config) ConfigCalls(stub func(string) (string, bool)) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = stub
}

func (fake *Config) ConfigArgsForCall(i int) string {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	argsForCall := fake.configArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Config) ConfigReturns(result1 string, result2 bool) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *Config) ConfigReturnsOnCall(i int, result1 string, result2 bool) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *Config) Set(arg1 string, arg2 string) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Set", []interface{}{arg1, arg2})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(arg1, arg2)
	}
}

func (fake *Config) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *Config) SetCalls(stub func(string, string)) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *Config) SetArgsForCall(i int) (string, string) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Config) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Config) recordInvocation(key string, args []interface{}) {
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

var _ download_clients.Config = new(Config)
