// Code generated by mockery v2.39.1. DO NOT EDIT.

package management

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// mockApiKeyProvider is an autogenerated mock type for the apiKeyProvider type
type mockApiKeyProvider struct {
	mock.Mock
}

// CreateAdminAPIKey provides a mock function with given fields: ctx, name
func (_m *mockApiKeyProvider) CreateAdminAPIKey(ctx context.Context, name string) (int64, string, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for CreateAdminAPIKey")
	}

	var r0 int64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, string, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) string); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// IsAPIKeyAuth provides a mock function with given fields: headers
func (_m *mockApiKeyProvider) IsAPIKeyAuth(headers http.Header) bool {
	ret := _m.Called(headers)

	if len(ret) == 0 {
		panic("no return value specified for IsAPIKeyAuth")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(http.Header) bool); ok {
		r0 = rf(headers)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// newMockApiKeyProvider creates a new instance of mockApiKeyProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockApiKeyProvider(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockApiKeyProvider {
	mock := &mockApiKeyProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
