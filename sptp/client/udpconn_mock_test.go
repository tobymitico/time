// Code generated by MockGen. DO NOT EDIT.
// Source: time/sptp/client/client.go

// Package client is a generated GoMock package.
package client

import (
	gomock "github.com/golang/mock/gomock"
	unix "golang.org/x/sys/unix"
	net "net"
	reflect "reflect"
	time "time"
)

// MockUDPConn is a mock of UDPConn interface
type MockUDPConn struct {
	ctrl     *gomock.Controller
	recorder *MockUDPConnMockRecorder
}

// MockUDPConnMockRecorder is the mock recorder for MockUDPConn
type MockUDPConnMockRecorder struct {
	mock *MockUDPConn
}

// NewMockUDPConn creates a new mock instance
func NewMockUDPConn(ctrl *gomock.Controller) *MockUDPConn {
	mock := &MockUDPConn{ctrl: ctrl}
	mock.recorder = &MockUDPConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUDPConn) EXPECT() *MockUDPConnMockRecorder {
	return m.recorder
}

// ReadFromUDP mocks base method
func (m *MockUDPConn) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromUDP", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*net.UDPAddr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFromUDP indicates an expected call of ReadFromUDP
func (mr *MockUDPConnMockRecorder) ReadFromUDP(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromUDP", reflect.TypeOf((*MockUDPConn)(nil).ReadFromUDP), b)
}

// WriteTo mocks base method
func (m *MockUDPConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", b, addr)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo
func (mr *MockUDPConnMockRecorder) WriteTo(b, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockUDPConn)(nil).WriteTo), b, addr)
}

// Close mocks base method
func (m *MockUDPConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockUDPConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUDPConn)(nil).Close))
}

// MockUDPConnWithTS is a mock of UDPConnWithTS interface
type MockUDPConnWithTS struct {
	ctrl     *gomock.Controller
	recorder *MockUDPConnWithTSMockRecorder
}

// MockUDPConnWithTSMockRecorder is the mock recorder for MockUDPConnWithTS
type MockUDPConnWithTSMockRecorder struct {
	mock *MockUDPConnWithTS
}

// NewMockUDPConnWithTS creates a new mock instance
func NewMockUDPConnWithTS(ctrl *gomock.Controller) *MockUDPConnWithTS {
	mock := &MockUDPConnWithTS{ctrl: ctrl}
	mock.recorder = &MockUDPConnWithTSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUDPConnWithTS) EXPECT() *MockUDPConnWithTSMockRecorder {
	return m.recorder
}

// ReadFromUDP mocks base method
func (m *MockUDPConnWithTS) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromUDP", b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*net.UDPAddr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFromUDP indicates an expected call of ReadFromUDP
func (mr *MockUDPConnWithTSMockRecorder) ReadFromUDP(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromUDP", reflect.TypeOf((*MockUDPConnWithTS)(nil).ReadFromUDP), b)
}

// WriteTo mocks base method
func (m *MockUDPConnWithTS) WriteTo(b []byte, addr net.Addr) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", b, addr)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo
func (mr *MockUDPConnWithTSMockRecorder) WriteTo(b, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockUDPConnWithTS)(nil).WriteTo), b, addr)
}

// Close mocks base method
func (m *MockUDPConnWithTS) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockUDPConnWithTSMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUDPConnWithTS)(nil).Close))
}

// WriteToWithTS mocks base method
func (m *MockUDPConnWithTS) WriteToWithTS(b []byte, addr net.Addr) (int, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToWithTS", b, addr)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WriteToWithTS indicates an expected call of WriteToWithTS
func (mr *MockUDPConnWithTSMockRecorder) WriteToWithTS(b, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToWithTS", reflect.TypeOf((*MockUDPConnWithTS)(nil).WriteToWithTS), b, addr)
}

// ReadPacketWithRXTimestamp mocks base method
func (m *MockUDPConnWithTS) ReadPacketWithRXTimestamp() ([]byte, unix.Sockaddr, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPacketWithRXTimestamp")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(unix.Sockaddr)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ReadPacketWithRXTimestamp indicates an expected call of ReadPacketWithRXTimestamp
func (mr *MockUDPConnWithTSMockRecorder) ReadPacketWithRXTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPacketWithRXTimestamp", reflect.TypeOf((*MockUDPConnWithTS)(nil).ReadPacketWithRXTimestamp))
}
