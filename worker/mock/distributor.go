// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/OmSingh2003/nimbus/worker (interfaces: TaskDistributor)
//
// Generated by this command:
//
//	mockgen -package mockwk -destination worker/mock/distributor.go github.com/OmSingh2003/nimbus/worker TaskDistributor
//

// Package mockwk is a generated GoMock package.
package mockwk

import (
	context "context"
	reflect "reflect"

	worker "github.com/OmSingh2003/nimbus/worker"
	asynq "github.com/hibiken/asynq"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskDistributor is a mock of TaskDistributor interface.
type MockTaskDistributor struct {
	ctrl     *gomock.Controller
	recorder *MockTaskDistributorMockRecorder
	isgomock struct{}
}

// MockTaskDistributorMockRecorder is the mock recorder for MockTaskDistributor.
type MockTaskDistributorMockRecorder struct {
	mock *MockTaskDistributor
}

// NewMockTaskDistributor creates a new mock instance.
func NewMockTaskDistributor(ctrl *gomock.Controller) *MockTaskDistributor {
	mock := &MockTaskDistributor{ctrl: ctrl}
	mock.recorder = &MockTaskDistributorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskDistributor) EXPECT() *MockTaskDistributorMockRecorder {
	return m.recorder
}

// DistributeTaskDemoResponse mocks base method.
func (m *MockTaskDistributor) DistributeTaskDemoResponse(ctx context.Context, payload *worker.PayloadDemoResponse, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskDemoResponse", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskDemoResponse indicates an expected call of DistributeTaskDemoResponse.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskDemoResponse(ctx, payload any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskDemoResponse", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskDemoResponse), varargs...)
}

// DistributeTaskSendVerifyEmail mocks base method.
func (m *MockTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *worker.PayloadSendVerifyEmail, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskSendVerifyEmail", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskSendVerifyEmail indicates an expected call of DistributeTaskSendVerifyEmail.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskSendVerifyEmail(ctx, payload any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskSendVerifyEmail", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskSendVerifyEmail), varargs...)
}
