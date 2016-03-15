// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/exec"
)

type FakeTaskDelegate struct {
	InitializingStub        func(atc.TaskConfig)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 atc.TaskConfig
	}
	StartedStub         func()
	startedMutex        sync.RWMutex
	startedArgsForCall  []struct{}
	FinishedStub        func(exec.ExitStatus)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 exec.ExitStatus
	}
	FailedStub        func(error)
	failedMutex       sync.RWMutex
	failedArgsForCall []struct {
		arg1 error
	}
	ImageVersionDeterminedStub        func(db.VolumeIdentifier) error
	imageVersionDeterminedMutex       sync.RWMutex
	imageVersionDeterminedArgsForCall []struct {
		arg1 db.VolumeIdentifier
	}
	imageVersionDeterminedReturns struct {
		result1 error
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct{}
	stdoutReturns     struct {
		result1 io.Writer
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct{}
	stderrReturns     struct {
		result1 io.Writer
	}
}

func (fake *FakeTaskDelegate) Initializing(arg1 atc.TaskConfig) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 atc.TaskConfig
	}{arg1})
	fake.initializingMutex.Unlock()
	if fake.InitializingStub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakeTaskDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakeTaskDelegate) InitializingArgsForCall(i int) atc.TaskConfig {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return fake.initializingArgsForCall[i].arg1
}

func (fake *FakeTaskDelegate) Started() {
	fake.startedMutex.Lock()
	fake.startedArgsForCall = append(fake.startedArgsForCall, struct{}{})
	fake.startedMutex.Unlock()
	if fake.StartedStub != nil {
		fake.StartedStub()
	}
}

func (fake *FakeTaskDelegate) StartedCallCount() int {
	fake.startedMutex.RLock()
	defer fake.startedMutex.RUnlock()
	return len(fake.startedArgsForCall)
}

func (fake *FakeTaskDelegate) Finished(arg1 exec.ExitStatus) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 exec.ExitStatus
	}{arg1})
	fake.finishedMutex.Unlock()
	if fake.FinishedStub != nil {
		fake.FinishedStub(arg1)
	}
}

func (fake *FakeTaskDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeTaskDelegate) FinishedArgsForCall(i int) exec.ExitStatus {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return fake.finishedArgsForCall[i].arg1
}

func (fake *FakeTaskDelegate) Failed(arg1 error) {
	fake.failedMutex.Lock()
	fake.failedArgsForCall = append(fake.failedArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.failedMutex.Unlock()
	if fake.FailedStub != nil {
		fake.FailedStub(arg1)
	}
}

func (fake *FakeTaskDelegate) FailedCallCount() int {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return len(fake.failedArgsForCall)
}

func (fake *FakeTaskDelegate) FailedArgsForCall(i int) error {
	fake.failedMutex.RLock()
	defer fake.failedMutex.RUnlock()
	return fake.failedArgsForCall[i].arg1
}

func (fake *FakeTaskDelegate) ImageVersionDetermined(arg1 db.VolumeIdentifier) error {
	fake.imageVersionDeterminedMutex.Lock()
	fake.imageVersionDeterminedArgsForCall = append(fake.imageVersionDeterminedArgsForCall, struct {
		arg1 db.VolumeIdentifier
	}{arg1})
	fake.imageVersionDeterminedMutex.Unlock()
	if fake.ImageVersionDeterminedStub != nil {
		return fake.ImageVersionDeterminedStub(arg1)
	} else {
		return fake.imageVersionDeterminedReturns.result1
	}
}

func (fake *FakeTaskDelegate) ImageVersionDeterminedCallCount() int {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return len(fake.imageVersionDeterminedArgsForCall)
}

func (fake *FakeTaskDelegate) ImageVersionDeterminedArgsForCall(i int) db.VolumeIdentifier {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return fake.imageVersionDeterminedArgsForCall[i].arg1
}

func (fake *FakeTaskDelegate) ImageVersionDeterminedReturns(result1 error) {
	fake.ImageVersionDeterminedStub = nil
	fake.imageVersionDeterminedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	} else {
		return fake.stdoutReturns.result1
	}
}

func (fake *FakeTaskDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeTaskDelegate) StdoutReturns(result1 io.Writer) {
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeTaskDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	} else {
		return fake.stderrReturns.result1
	}
}

func (fake *FakeTaskDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeTaskDelegate) StderrReturns(result1 io.Writer) {
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

var _ exec.TaskDelegate = new(FakeTaskDelegate)
