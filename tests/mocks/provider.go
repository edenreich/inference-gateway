// Code generated by MockGen. DO NOT EDIT.
// Source: common.go
//
// Generated by this command:
//
//	mockgen -source=common.go -destination=../tests/mocks/provider.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	providers "github.com/inference-gateway/inference-gateway/providers"
	gomock "go.uber.org/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
	isgomock struct{}
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// BuildGenTokensRequest mocks base method.
func (m *MockProvider) BuildGenTokensRequest(model string, messages []providers.GenerateMessage) any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildGenTokensRequest", model, messages)
	ret0, _ := ret[0].(any)
	return ret0
}

// BuildGenTokensRequest indicates an expected call of BuildGenTokensRequest.
func (mr *MockProviderMockRecorder) BuildGenTokensRequest(model, messages any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildGenTokensRequest", reflect.TypeOf((*MockProvider)(nil).BuildGenTokensRequest), model, messages)
}

// BuildGenTokensResponse mocks base method.
func (m *MockProvider) BuildGenTokensResponse(model string, responseBody any) (providers.GenerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildGenTokensResponse", model, responseBody)
	ret0, _ := ret[0].(providers.GenerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildGenTokensResponse indicates an expected call of BuildGenTokensResponse.
func (mr *MockProviderMockRecorder) BuildGenTokensResponse(model, responseBody any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildGenTokensResponse", reflect.TypeOf((*MockProvider)(nil).BuildGenTokensResponse), model, responseBody)
}

// GetAuthType mocks base method.
func (m *MockProvider) GetAuthType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAuthType indicates an expected call of GetAuthType.
func (mr *MockProviderMockRecorder) GetAuthType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthType", reflect.TypeOf((*MockProvider)(nil).GetAuthType))
}

// GetExtraXHeaders mocks base method.
func (m *MockProvider) GetExtraXHeaders() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtraXHeaders")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetExtraXHeaders indicates an expected call of GetExtraXHeaders.
func (mr *MockProviderMockRecorder) GetExtraXHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtraXHeaders", reflect.TypeOf((*MockProvider)(nil).GetExtraXHeaders))
}

// GetID mocks base method.
func (m *MockProvider) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockProviderMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockProvider)(nil).GetID))
}

// GetName mocks base method.
func (m *MockProvider) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockProviderMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockProvider)(nil).GetName))
}

// GetProxyURL mocks base method.
func (m *MockProvider) GetProxyURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProxyURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetProxyURL indicates an expected call of GetProxyURL.
func (mr *MockProviderMockRecorder) GetProxyURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProxyURL", reflect.TypeOf((*MockProvider)(nil).GetProxyURL))
}

// GetToken mocks base method.
func (m *MockProvider) GetToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetToken indicates an expected call of GetToken.
func (mr *MockProviderMockRecorder) GetToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockProvider)(nil).GetToken))
}

// GetURL mocks base method.
func (m *MockProvider) GetURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURL indicates an expected call of GetURL.
func (mr *MockProviderMockRecorder) GetURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockProvider)(nil).GetURL))
}
