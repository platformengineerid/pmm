// Code generated by mockery v2.33.0. DO NOT EDIT.

package telemetry

import (
	context "context"

	reporterv1 "github.com/percona-platform/saas/gen/telemetry/reporter"
	mock "github.com/stretchr/testify/mock"
)

// mockSender is an autogenerated mock type for the sender type
type mockSender struct {
	mock.Mock
}

// SendTelemetry provides a mock function with given fields: ctx, report
func (_m *mockSender) SendTelemetry(ctx context.Context, report *reporterv1.ReportRequest) error {
	ret := _m.Called(ctx, report)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *reporterv1.ReportRequest) error); ok {
		r0 = rf(ctx, report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockSender creates a new instance of mockSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSender(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockSender {
	mock := &mockSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
