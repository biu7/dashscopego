// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/devinyf/dashscopego/httpclient (interfaces: IHttpClient)
//
// Generated by this command:
//
//	mockgen -destination=http_client_mock.go -package=httpclient . IHttpClient
//
// Package httpclient is a generated GoMock package.
package httpclient

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIHttpClient is a mock of IHttpClient interface.
type MockIHttpClient struct {
	ctrl     *gomock.Controller
	recorder *MockIHttpClientMockRecorder
}

// MockIHttpClientMockRecorder is the mock recorder for MockIHttpClient.
type MockIHttpClientMockRecorder struct {
	mock *MockIHttpClient
}

// NewMockIHttpClient creates a new mock instance.
func NewMockIHttpClient(ctrl *gomock.Controller) *MockIHttpClient {
	mock := &MockIHttpClient{ctrl: ctrl}
	mock.recorder = &MockIHttpClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHttpClient) EXPECT() *MockIHttpClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIHttpClient) Get(arg0 context.Context, arg1 string, arg2 map[string]string, arg3 any, arg4 ...HTTPOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIHttpClientMockRecorder) Get(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIHttpClient)(nil).Get), varargs...)
}

// GetImage mocks base method.
func (m *MockIHttpClient) GetImage(arg0 context.Context, arg1 string, arg2 ...HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetImage", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImage indicates an expected call of GetImage.
func (mr *MockIHttpClientMockRecorder) GetImage(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockIHttpClient)(nil).GetImage), varargs...)
}

// Post mocks base method.
func (m *MockIHttpClient) Post(arg0 context.Context, arg1 string, arg2, arg3 any, arg4 ...HTTPOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Post", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockIHttpClientMockRecorder) Post(arg0, arg1, arg2, arg3 any, arg4 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockIHttpClient)(nil).Post), varargs...)
}

// PostSSE mocks base method.
func (m *MockIHttpClient) PostSSE(arg0 context.Context, arg1 string, arg2 any, arg3 ...HTTPOption) (chan string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostSSE", varargs...)
	ret0, _ := ret[0].(chan string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSSE indicates an expected call of PostSSE.
func (mr *MockIHttpClientMockRecorder) PostSSE(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSSE", reflect.TypeOf((*MockIHttpClient)(nil).PostSSE), varargs...)
}
