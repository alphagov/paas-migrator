// Code generated by counterfeiter. DO NOT EDIT.
package credhubfakes

import (
	"sync"

	"github.com/ishustava/migrator/credhub"
)

type FakeBulkSetObserver struct {
	BeginBulkSetStub        func(numPasswords, numCertificates, numRsaKeys, numSshKeys int)
	beginBulkSetMutex       sync.RWMutex
	beginBulkSetArgsForCall []struct {
		numPasswords    int
		numCertificates int
		numRsaKeys      int
		numSshKeys      int
	}
	FailPasswordSetStub        func(name string, err error)
	failPasswordSetMutex       sync.RWMutex
	failPasswordSetArgsForCall []struct {
		name string
		err  error
	}
	EndPasswordsSetStub           func()
	endPasswordsSetMutex          sync.RWMutex
	endPasswordsSetArgsForCall    []struct{}
	FailCertificateSetStub        func(name string, err error)
	failCertificateSetMutex       sync.RWMutex
	failCertificateSetArgsForCall []struct {
		name string
		err  error
	}
	EndCertificatesSetStub        func()
	endCertificatesSetMutex       sync.RWMutex
	endCertificatesSetArgsForCall []struct{}
	FailRsaKeySetStub             func(name string, err error)
	failRsaKeySetMutex            sync.RWMutex
	failRsaKeySetArgsForCall      []struct {
		name string
		err  error
	}
	EndRsaKeysSetStub        func()
	endRsaKeysSetMutex       sync.RWMutex
	endRsaKeysSetArgsForCall []struct{}
	FailSshKeySetStub        func(name string, err error)
	failSshKeySetMutex       sync.RWMutex
	failSshKeySetArgsForCall []struct {
		name string
		err  error
	}
	EndSshKeysSetStub        func()
	endSshKeysSetMutex       sync.RWMutex
	endSshKeysSetArgsForCall []struct{}
	EndBulkSetStub           func() error
	endBulkSetMutex          sync.RWMutex
	endBulkSetArgsForCall    []struct{}
	endBulkSetReturns        struct {
		result1 error
	}
	endBulkSetReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBulkSetObserver) BeginBulkSet(numPasswords int, numCertificates int, numRsaKeys int, numSshKeys int) {
	fake.beginBulkSetMutex.Lock()
	fake.beginBulkSetArgsForCall = append(fake.beginBulkSetArgsForCall, struct {
		numPasswords    int
		numCertificates int
		numRsaKeys      int
		numSshKeys      int
	}{numPasswords, numCertificates, numRsaKeys, numSshKeys})
	fake.recordInvocation("BeginBulkSet", []interface{}{numPasswords, numCertificates, numRsaKeys, numSshKeys})
	fake.beginBulkSetMutex.Unlock()
	if fake.BeginBulkSetStub != nil {
		fake.BeginBulkSetStub(numPasswords, numCertificates, numRsaKeys, numSshKeys)
	}
}

func (fake *FakeBulkSetObserver) BeginBulkSetCallCount() int {
	fake.beginBulkSetMutex.RLock()
	defer fake.beginBulkSetMutex.RUnlock()
	return len(fake.beginBulkSetArgsForCall)
}

func (fake *FakeBulkSetObserver) BeginBulkSetArgsForCall(i int) (int, int, int, int) {
	fake.beginBulkSetMutex.RLock()
	defer fake.beginBulkSetMutex.RUnlock()
	return fake.beginBulkSetArgsForCall[i].numPasswords, fake.beginBulkSetArgsForCall[i].numCertificates, fake.beginBulkSetArgsForCall[i].numRsaKeys, fake.beginBulkSetArgsForCall[i].numSshKeys
}

func (fake *FakeBulkSetObserver) FailPasswordSet(name string, err error) {
	fake.failPasswordSetMutex.Lock()
	fake.failPasswordSetArgsForCall = append(fake.failPasswordSetArgsForCall, struct {
		name string
		err  error
	}{name, err})
	fake.recordInvocation("FailPasswordSet", []interface{}{name, err})
	fake.failPasswordSetMutex.Unlock()
	if fake.FailPasswordSetStub != nil {
		fake.FailPasswordSetStub(name, err)
	}
}

func (fake *FakeBulkSetObserver) FailPasswordSetCallCount() int {
	fake.failPasswordSetMutex.RLock()
	defer fake.failPasswordSetMutex.RUnlock()
	return len(fake.failPasswordSetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailPasswordSetArgsForCall(i int) (string, error) {
	fake.failPasswordSetMutex.RLock()
	defer fake.failPasswordSetMutex.RUnlock()
	return fake.failPasswordSetArgsForCall[i].name, fake.failPasswordSetArgsForCall[i].err
}

func (fake *FakeBulkSetObserver) EndPasswordsSet() {
	fake.endPasswordsSetMutex.Lock()
	fake.endPasswordsSetArgsForCall = append(fake.endPasswordsSetArgsForCall, struct{}{})
	fake.recordInvocation("EndPasswordsSet", []interface{}{})
	fake.endPasswordsSetMutex.Unlock()
	if fake.EndPasswordsSetStub != nil {
		fake.EndPasswordsSetStub()
	}
}

func (fake *FakeBulkSetObserver) EndPasswordsSetCallCount() int {
	fake.endPasswordsSetMutex.RLock()
	defer fake.endPasswordsSetMutex.RUnlock()
	return len(fake.endPasswordsSetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailCertificateSet(name string, err error) {
	fake.failCertificateSetMutex.Lock()
	fake.failCertificateSetArgsForCall = append(fake.failCertificateSetArgsForCall, struct {
		name string
		err  error
	}{name, err})
	fake.recordInvocation("FailCertificateSet", []interface{}{name, err})
	fake.failCertificateSetMutex.Unlock()
	if fake.FailCertificateSetStub != nil {
		fake.FailCertificateSetStub(name, err)
	}
}

func (fake *FakeBulkSetObserver) FailCertificateSetCallCount() int {
	fake.failCertificateSetMutex.RLock()
	defer fake.failCertificateSetMutex.RUnlock()
	return len(fake.failCertificateSetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailCertificateSetArgsForCall(i int) (string, error) {
	fake.failCertificateSetMutex.RLock()
	defer fake.failCertificateSetMutex.RUnlock()
	return fake.failCertificateSetArgsForCall[i].name, fake.failCertificateSetArgsForCall[i].err
}

func (fake *FakeBulkSetObserver) EndCertificatesSet() {
	fake.endCertificatesSetMutex.Lock()
	fake.endCertificatesSetArgsForCall = append(fake.endCertificatesSetArgsForCall, struct{}{})
	fake.recordInvocation("EndCertificatesSet", []interface{}{})
	fake.endCertificatesSetMutex.Unlock()
	if fake.EndCertificatesSetStub != nil {
		fake.EndCertificatesSetStub()
	}
}

func (fake *FakeBulkSetObserver) EndCertificatesSetCallCount() int {
	fake.endCertificatesSetMutex.RLock()
	defer fake.endCertificatesSetMutex.RUnlock()
	return len(fake.endCertificatesSetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailRsaKeySet(name string, err error) {
	fake.failRsaKeySetMutex.Lock()
	fake.failRsaKeySetArgsForCall = append(fake.failRsaKeySetArgsForCall, struct {
		name string
		err  error
	}{name, err})
	fake.recordInvocation("FailRsaKeySet", []interface{}{name, err})
	fake.failRsaKeySetMutex.Unlock()
	if fake.FailRsaKeySetStub != nil {
		fake.FailRsaKeySetStub(name, err)
	}
}

func (fake *FakeBulkSetObserver) FailRsaKeySetCallCount() int {
	fake.failRsaKeySetMutex.RLock()
	defer fake.failRsaKeySetMutex.RUnlock()
	return len(fake.failRsaKeySetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailRsaKeySetArgsForCall(i int) (string, error) {
	fake.failRsaKeySetMutex.RLock()
	defer fake.failRsaKeySetMutex.RUnlock()
	return fake.failRsaKeySetArgsForCall[i].name, fake.failRsaKeySetArgsForCall[i].err
}

func (fake *FakeBulkSetObserver) EndRsaKeysSet() {
	fake.endRsaKeysSetMutex.Lock()
	fake.endRsaKeysSetArgsForCall = append(fake.endRsaKeysSetArgsForCall, struct{}{})
	fake.recordInvocation("EndRsaKeysSet", []interface{}{})
	fake.endRsaKeysSetMutex.Unlock()
	if fake.EndRsaKeysSetStub != nil {
		fake.EndRsaKeysSetStub()
	}
}

func (fake *FakeBulkSetObserver) EndRsaKeysSetCallCount() int {
	fake.endRsaKeysSetMutex.RLock()
	defer fake.endRsaKeysSetMutex.RUnlock()
	return len(fake.endRsaKeysSetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailSshKeySet(name string, err error) {
	fake.failSshKeySetMutex.Lock()
	fake.failSshKeySetArgsForCall = append(fake.failSshKeySetArgsForCall, struct {
		name string
		err  error
	}{name, err})
	fake.recordInvocation("FailSshKeySet", []interface{}{name, err})
	fake.failSshKeySetMutex.Unlock()
	if fake.FailSshKeySetStub != nil {
		fake.FailSshKeySetStub(name, err)
	}
}

func (fake *FakeBulkSetObserver) FailSshKeySetCallCount() int {
	fake.failSshKeySetMutex.RLock()
	defer fake.failSshKeySetMutex.RUnlock()
	return len(fake.failSshKeySetArgsForCall)
}

func (fake *FakeBulkSetObserver) FailSshKeySetArgsForCall(i int) (string, error) {
	fake.failSshKeySetMutex.RLock()
	defer fake.failSshKeySetMutex.RUnlock()
	return fake.failSshKeySetArgsForCall[i].name, fake.failSshKeySetArgsForCall[i].err
}

func (fake *FakeBulkSetObserver) EndSshKeysSet() {
	fake.endSshKeysSetMutex.Lock()
	fake.endSshKeysSetArgsForCall = append(fake.endSshKeysSetArgsForCall, struct{}{})
	fake.recordInvocation("EndSshKeysSet", []interface{}{})
	fake.endSshKeysSetMutex.Unlock()
	if fake.EndSshKeysSetStub != nil {
		fake.EndSshKeysSetStub()
	}
}

func (fake *FakeBulkSetObserver) EndSshKeysSetCallCount() int {
	fake.endSshKeysSetMutex.RLock()
	defer fake.endSshKeysSetMutex.RUnlock()
	return len(fake.endSshKeysSetArgsForCall)
}

func (fake *FakeBulkSetObserver) EndBulkSet() error {
	fake.endBulkSetMutex.Lock()
	ret, specificReturn := fake.endBulkSetReturnsOnCall[len(fake.endBulkSetArgsForCall)]
	fake.endBulkSetArgsForCall = append(fake.endBulkSetArgsForCall, struct{}{})
	fake.recordInvocation("EndBulkSet", []interface{}{})
	fake.endBulkSetMutex.Unlock()
	if fake.EndBulkSetStub != nil {
		return fake.EndBulkSetStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.endBulkSetReturns.result1
}

func (fake *FakeBulkSetObserver) EndBulkSetCallCount() int {
	fake.endBulkSetMutex.RLock()
	defer fake.endBulkSetMutex.RUnlock()
	return len(fake.endBulkSetArgsForCall)
}

func (fake *FakeBulkSetObserver) EndBulkSetReturns(result1 error) {
	fake.EndBulkSetStub = nil
	fake.endBulkSetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBulkSetObserver) EndBulkSetReturnsOnCall(i int, result1 error) {
	fake.EndBulkSetStub = nil
	if fake.endBulkSetReturnsOnCall == nil {
		fake.endBulkSetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.endBulkSetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBulkSetObserver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginBulkSetMutex.RLock()
	defer fake.beginBulkSetMutex.RUnlock()
	fake.failPasswordSetMutex.RLock()
	defer fake.failPasswordSetMutex.RUnlock()
	fake.endPasswordsSetMutex.RLock()
	defer fake.endPasswordsSetMutex.RUnlock()
	fake.failCertificateSetMutex.RLock()
	defer fake.failCertificateSetMutex.RUnlock()
	fake.endCertificatesSetMutex.RLock()
	defer fake.endCertificatesSetMutex.RUnlock()
	fake.failRsaKeySetMutex.RLock()
	defer fake.failRsaKeySetMutex.RUnlock()
	fake.endRsaKeysSetMutex.RLock()
	defer fake.endRsaKeysSetMutex.RUnlock()
	fake.failSshKeySetMutex.RLock()
	defer fake.failSshKeySetMutex.RUnlock()
	fake.endSshKeysSetMutex.RLock()
	defer fake.endSshKeysSetMutex.RUnlock()
	fake.endBulkSetMutex.RLock()
	defer fake.endBulkSetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBulkSetObserver) recordInvocation(key string, args []interface{}) {
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

var _ credhub.BulkSetObserver = new(FakeBulkSetObserver)