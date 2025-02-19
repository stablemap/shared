// Code generated by MockGen. DO NOT EDIT.
// Source: redis_cache_api_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	redis "github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockCacheService is a mock of CacheService interface.
type MockCacheService struct {
	ctrl     *gomock.Controller
	recorder *MockCacheServiceMockRecorder
}

// MockCacheServiceMockRecorder is the mock recorder for MockCacheService.
type MockCacheServiceMockRecorder struct {
	mock *MockCacheService
}

// NewMockCacheService creates a new mock instance.
func NewMockCacheService(ctrl *gomock.Controller) *MockCacheService {
	mock := &MockCacheService{ctrl: ctrl}
	mock.recorder = &MockCacheServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheService) EXPECT() *MockCacheServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCacheService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCacheServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCacheService)(nil).Close))
}

// Del mocks base method.
func (m *MockCacheService) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockCacheServiceMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockCacheService)(nil).Del), varargs...)
}

// FlushAll mocks base method.
func (m *MockCacheService) FlushAll(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockCacheServiceMockRecorder) FlushAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockCacheService)(nil).FlushAll), ctx)
}

// Get mocks base method.
func (m *MockCacheService) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheServiceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheService)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockCacheService) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheServiceMockRecorder) Set(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheService)(nil).Set), ctx, key, value, expiration)
}
