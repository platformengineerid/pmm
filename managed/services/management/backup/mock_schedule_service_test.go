// Code generated by mockery v2.36.0. DO NOT EDIT.

package backup

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
	scheduler "github.com/percona/pmm/managed/services/scheduler"
)

// mockScheduleService is an autogenerated mock type for the scheduleService type
type mockScheduleService struct {
	mock.Mock
}

// Add provides a mock function with given fields: task, params
func (_m *mockScheduleService) Add(task scheduler.Task, params scheduler.AddParams) (*models.ScheduledTask, error) {
	ret := _m.Called(task, params)

	var r0 *models.ScheduledTask
	var r1 error
	if rf, ok := ret.Get(0).(func(scheduler.Task, scheduler.AddParams) (*models.ScheduledTask, error)); ok {
		return rf(task, params)
	}
	if rf, ok := ret.Get(0).(func(scheduler.Task, scheduler.AddParams) *models.ScheduledTask); ok {
		r0 = rf(task, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ScheduledTask)
		}
	}

	if rf, ok := ret.Get(1).(func(scheduler.Task, scheduler.AddParams) error); ok {
		r1 = rf(task, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *mockScheduleService) Remove(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: ctx
func (_m *mockScheduleService) Run(ctx context.Context) {
	_m.Called(ctx)
}

// Update provides a mock function with given fields: id, params
func (_m *mockScheduleService) Update(id string, params models.ChangeScheduledTaskParams) error {
	ret := _m.Called(id, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.ChangeScheduledTaskParams) error); ok {
		r0 = rf(id, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockScheduleService creates a new instance of mockScheduleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockScheduleService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockScheduleService {
	mock := &mockScheduleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
