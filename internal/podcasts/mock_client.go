// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/podcasts/client.go

// Package podcasts is a generated GoMock package.
package podcasts

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPodcastClient is a mock of PodcastClient interface.
type MockPodcastClient struct {
	ctrl     *gomock.Controller
	recorder *MockPodcastClientMockRecorder
}

// MockPodcastClientMockRecorder is the mock recorder for MockPodcastClient.
type MockPodcastClientMockRecorder struct {
	mock *MockPodcastClient
}

// NewMockPodcastClient creates a new mock instance.
func NewMockPodcastClient(ctrl *gomock.Controller) *MockPodcastClient {
	mock := &MockPodcastClient{ctrl: ctrl}
	mock.recorder = &MockPodcastClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodcastClient) EXPECT() *MockPodcastClientMockRecorder {
	return m.recorder
}

// ListPodcasts mocks base method.
func (m *MockPodcastClient) ListPodcasts(ctx context.Context, params *ListPodcastsParams) (*ListPodcastsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPodcasts", ctx, params)
	ret0, _ := ret[0].(*ListPodcastsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPodcasts indicates an expected call of ListPodcasts.
func (mr *MockPodcastClientMockRecorder) ListPodcasts(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPodcasts", reflect.TypeOf((*MockPodcastClient)(nil).ListPodcasts), ctx, params)
}
