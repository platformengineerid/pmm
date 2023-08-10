// Code generated by mockery v2.32.0. DO NOT EDIT.

package ia

import (
	mock "github.com/stretchr/testify/mock"

	models "github.com/percona/pmm/managed/models"
)

// mockTemplatesService is an autogenerated mock type for the templatesService type
type mockTemplatesService struct {
	mock.Mock
}

// GetTemplates provides a mock function with given fields:
func (_m *mockTemplatesService) GetTemplates() map[string]models.Template {
	ret := _m.Called()

	var r0 map[string]models.Template
	if rf, ok := ret.Get(0).(func() map[string]models.Template); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]models.Template)
		}
	}

	return r0
}

// newMockTemplatesService creates a new instance of mockTemplatesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockTemplatesService(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockTemplatesService {
	mock := &mockTemplatesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
