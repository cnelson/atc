// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/scheduler"
)

type FakeBuildScheduler struct {
	TryNextPendingBuildStub        func(atc.JobConfig) error
	tryNextPendingBuildMutex       sync.RWMutex
	tryNextPendingBuildArgsForCall []struct {
		arg1 atc.JobConfig
	}
	tryNextPendingBuildReturns struct {
		result1 error
	}
	BuildLatestInputsStub        func(atc.JobConfig) error
	buildLatestInputsMutex       sync.RWMutex
	buildLatestInputsArgsForCall []struct {
		arg1 atc.JobConfig
	}
	buildLatestInputsReturns struct {
		result1 error
	}
	TrackInFlightBuildsStub        func() error
	trackInFlightBuildsMutex       sync.RWMutex
	trackInFlightBuildsArgsForCall []struct{}
	trackInFlightBuildsReturns     struct {
		result1 error
	}
}

func (fake *FakeBuildScheduler) TryNextPendingBuild(arg1 atc.JobConfig) error {
	fake.tryNextPendingBuildMutex.Lock()
	fake.tryNextPendingBuildArgsForCall = append(fake.tryNextPendingBuildArgsForCall, struct {
		arg1 atc.JobConfig
	}{arg1})
	fake.tryNextPendingBuildMutex.Unlock()
	if fake.TryNextPendingBuildStub != nil {
		return fake.TryNextPendingBuildStub(arg1)
	} else {
		return fake.tryNextPendingBuildReturns.result1
	}
}

func (fake *FakeBuildScheduler) TryNextPendingBuildCallCount() int {
	fake.tryNextPendingBuildMutex.RLock()
	defer fake.tryNextPendingBuildMutex.RUnlock()
	return len(fake.tryNextPendingBuildArgsForCall)
}

func (fake *FakeBuildScheduler) TryNextPendingBuildArgsForCall(i int) atc.JobConfig {
	fake.tryNextPendingBuildMutex.RLock()
	defer fake.tryNextPendingBuildMutex.RUnlock()
	return fake.tryNextPendingBuildArgsForCall[i].arg1
}

func (fake *FakeBuildScheduler) TryNextPendingBuildReturns(result1 error) {
	fake.TryNextPendingBuildStub = nil
	fake.tryNextPendingBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildScheduler) BuildLatestInputs(arg1 atc.JobConfig) error {
	fake.buildLatestInputsMutex.Lock()
	fake.buildLatestInputsArgsForCall = append(fake.buildLatestInputsArgsForCall, struct {
		arg1 atc.JobConfig
	}{arg1})
	fake.buildLatestInputsMutex.Unlock()
	if fake.BuildLatestInputsStub != nil {
		return fake.BuildLatestInputsStub(arg1)
	} else {
		return fake.buildLatestInputsReturns.result1
	}
}

func (fake *FakeBuildScheduler) BuildLatestInputsCallCount() int {
	fake.buildLatestInputsMutex.RLock()
	defer fake.buildLatestInputsMutex.RUnlock()
	return len(fake.buildLatestInputsArgsForCall)
}

func (fake *FakeBuildScheduler) BuildLatestInputsArgsForCall(i int) atc.JobConfig {
	fake.buildLatestInputsMutex.RLock()
	defer fake.buildLatestInputsMutex.RUnlock()
	return fake.buildLatestInputsArgsForCall[i].arg1
}

func (fake *FakeBuildScheduler) BuildLatestInputsReturns(result1 error) {
	fake.BuildLatestInputsStub = nil
	fake.buildLatestInputsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildScheduler) TrackInFlightBuilds() error {
	fake.trackInFlightBuildsMutex.Lock()
	fake.trackInFlightBuildsArgsForCall = append(fake.trackInFlightBuildsArgsForCall, struct{}{})
	fake.trackInFlightBuildsMutex.Unlock()
	if fake.TrackInFlightBuildsStub != nil {
		return fake.TrackInFlightBuildsStub()
	} else {
		return fake.trackInFlightBuildsReturns.result1
	}
}

func (fake *FakeBuildScheduler) TrackInFlightBuildsCallCount() int {
	fake.trackInFlightBuildsMutex.RLock()
	defer fake.trackInFlightBuildsMutex.RUnlock()
	return len(fake.trackInFlightBuildsArgsForCall)
}

func (fake *FakeBuildScheduler) TrackInFlightBuildsReturns(result1 error) {
	fake.TrackInFlightBuildsStub = nil
	fake.trackInFlightBuildsReturns = struct {
		result1 error
	}{result1}
}

var _ scheduler.BuildScheduler = new(FakeBuildScheduler)
