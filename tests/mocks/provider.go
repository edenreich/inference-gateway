// Code generated by MockGen. DO NOT EDIT.
// Source: management.go
//
// Generated by this command:
//
//	mockgen -source=management.go -destination=../tests/mocks/provider.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
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

// GenerateTokens mocks base method.
func (m *MockProvider) GenerateTokens(ctx context.Context, model string, messages []providers.Message) (providers.GenerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateTokens", ctx, model, messages)
	ret0, _ := ret[0].(providers.GenerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateTokens indicates an expected call of GenerateTokens.
func (mr *MockProviderMockRecorder) GenerateTokens(ctx, model, messages any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateTokens", reflect.TypeOf((*MockProvider)(nil).GenerateTokens), ctx, model, messages)
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

// GetExtraHeaders mocks base method.
func (m *MockProvider) GetExtraHeaders() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtraHeaders")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetExtraHeaders indicates an expected call of GetExtraHeaders.
func (mr *MockProviderMockRecorder) GetExtraHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtraHeaders", reflect.TypeOf((*MockProvider)(nil).GetExtraHeaders))
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

// ListModels mocks base method.
func (m *MockProvider) ListModels(ctx context.Context) (providers.ListModelsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListModels", ctx)
	ret0, _ := ret[0].(providers.ListModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListModels indicates an expected call of ListModels.
func (mr *MockProviderMockRecorder) ListModels(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModels", reflect.TypeOf((*MockProvider)(nil).ListModels), ctx)
}

// StreamTokens mocks base method.
func (m *MockProvider) StreamTokens(ctx context.Context, model string, messages []providers.Message) (<-chan providers.GenerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamTokens", ctx, model, messages)
	ret0, _ := ret[0].(<-chan providers.GenerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamTokens indicates an expected call of StreamTokens.
func (mr *MockProviderMockRecorder) StreamTokens(ctx, model, messages any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamTokens", reflect.TypeOf((*MockProvider)(nil).StreamTokens), ctx, model, messages)
}
