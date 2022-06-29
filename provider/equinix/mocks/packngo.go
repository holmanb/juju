// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/packethost/packngo (interfaces: DeviceService,OSService,PlanService,ProjectIPService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	packngo "github.com/packethost/packngo"
)

// MockDeviceService is a mock of DeviceService interface.
type MockDeviceService struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceMockRecorder
}

// MockDeviceServiceMockRecorder is the mock recorder for MockDeviceService.
type MockDeviceServiceMockRecorder struct {
	mock *MockDeviceService
}

// NewMockDeviceService creates a new mock instance.
func NewMockDeviceService(ctrl *gomock.Controller) *MockDeviceService {
	mock := &MockDeviceService{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceService) EXPECT() *MockDeviceServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDeviceService) Create(arg0 *packngo.DeviceCreateRequest) (*packngo.Device, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*packngo.Device)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockDeviceServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockDeviceService) Delete(arg0 string, arg1 bool) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDeviceServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceService)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockDeviceService) Get(arg0 string, arg1 *packngo.GetOptions) (*packngo.Device, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*packngo.Device)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockDeviceServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceService)(nil).Get), arg0, arg1)
}

// GetBandwidth mocks base method.
func (m *MockDeviceService) GetBandwidth(arg0 string, arg1 *packngo.BandwidthOpts) (*packngo.BandwidthIO, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBandwidth", arg0, arg1)
	ret0, _ := ret[0].(*packngo.BandwidthIO)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBandwidth indicates an expected call of GetBandwidth.
func (mr *MockDeviceServiceMockRecorder) GetBandwidth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBandwidth", reflect.TypeOf((*MockDeviceService)(nil).GetBandwidth), arg0, arg1)
}

// List mocks base method.
func (m *MockDeviceService) List(arg0 string, arg1 *packngo.GetOptions) ([]packngo.Device, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]packngo.Device)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockDeviceServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeviceService)(nil).List), arg0, arg1)
}

// ListBGPNeighbors mocks base method.
func (m *MockDeviceService) ListBGPNeighbors(arg0 string, arg1 *packngo.GetOptions) ([]packngo.BGPNeighbor, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBGPNeighbors", arg0, arg1)
	ret0, _ := ret[0].([]packngo.BGPNeighbor)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBGPNeighbors indicates an expected call of ListBGPNeighbors.
func (mr *MockDeviceServiceMockRecorder) ListBGPNeighbors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBGPNeighbors", reflect.TypeOf((*MockDeviceService)(nil).ListBGPNeighbors), arg0, arg1)
}

// ListBGPSessions mocks base method.
func (m *MockDeviceService) ListBGPSessions(arg0 string, arg1 *packngo.GetOptions) ([]packngo.BGPSession, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBGPSessions", arg0, arg1)
	ret0, _ := ret[0].([]packngo.BGPSession)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBGPSessions indicates an expected call of ListBGPSessions.
func (mr *MockDeviceServiceMockRecorder) ListBGPSessions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBGPSessions", reflect.TypeOf((*MockDeviceService)(nil).ListBGPSessions), arg0, arg1)
}

// ListEvents mocks base method.
func (m *MockDeviceService) ListEvents(arg0 string, arg1 *packngo.GetOptions) ([]packngo.Event, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEvents", arg0, arg1)
	ret0, _ := ret[0].([]packngo.Event)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListEvents indicates an expected call of ListEvents.
func (mr *MockDeviceServiceMockRecorder) ListEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEvents", reflect.TypeOf((*MockDeviceService)(nil).ListEvents), arg0, arg1)
}

// Lock mocks base method.
func (m *MockDeviceService) Lock(arg0 string) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockDeviceServiceMockRecorder) Lock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockDeviceService)(nil).Lock), arg0)
}

// PowerOff mocks base method.
func (m *MockDeviceService) PowerOff(arg0 string) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOff", arg0)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOff indicates an expected call of PowerOff.
func (mr *MockDeviceServiceMockRecorder) PowerOff(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOff", reflect.TypeOf((*MockDeviceService)(nil).PowerOff), arg0)
}

// PowerOn mocks base method.
func (m *MockDeviceService) PowerOn(arg0 string) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOn", arg0)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PowerOn indicates an expected call of PowerOn.
func (mr *MockDeviceServiceMockRecorder) PowerOn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOn", reflect.TypeOf((*MockDeviceService)(nil).PowerOn), arg0)
}

// Reboot mocks base method.
func (m *MockDeviceService) Reboot(arg0 string) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reboot", arg0)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reboot indicates an expected call of Reboot.
func (mr *MockDeviceServiceMockRecorder) Reboot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reboot", reflect.TypeOf((*MockDeviceService)(nil).Reboot), arg0)
}

// Reinstall mocks base method.
func (m *MockDeviceService) Reinstall(arg0 string, arg1 *packngo.DeviceReinstallFields) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reinstall", arg0, arg1)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reinstall indicates an expected call of Reinstall.
func (mr *MockDeviceServiceMockRecorder) Reinstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reinstall", reflect.TypeOf((*MockDeviceService)(nil).Reinstall), arg0, arg1)
}

// Unlock mocks base method.
func (m *MockDeviceService) Unlock(arg0 string) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unlock indicates an expected call of Unlock.
func (mr *MockDeviceServiceMockRecorder) Unlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockDeviceService)(nil).Unlock), arg0)
}

// Update mocks base method.
func (m *MockDeviceService) Update(arg0 string, arg1 *packngo.DeviceUpdateRequest) (*packngo.Device, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*packngo.Device)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockDeviceServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeviceService)(nil).Update), arg0, arg1)
}

// MockOSService is a mock of OSService interface.
type MockOSService struct {
	ctrl     *gomock.Controller
	recorder *MockOSServiceMockRecorder
}

// MockOSServiceMockRecorder is the mock recorder for MockOSService.
type MockOSServiceMockRecorder struct {
	mock *MockOSService
}

// NewMockOSService creates a new mock instance.
func NewMockOSService(ctrl *gomock.Controller) *MockOSService {
	mock := &MockOSService{ctrl: ctrl}
	mock.recorder = &MockOSServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOSService) EXPECT() *MockOSServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockOSService) List() ([]packngo.OS, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]packngo.OS)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockOSServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOSService)(nil).List))
}

// MockPlanService is a mock of PlanService interface.
type MockPlanService struct {
	ctrl     *gomock.Controller
	recorder *MockPlanServiceMockRecorder
}

// MockPlanServiceMockRecorder is the mock recorder for MockPlanService.
type MockPlanServiceMockRecorder struct {
	mock *MockPlanService
}

// NewMockPlanService creates a new mock instance.
func NewMockPlanService(ctrl *gomock.Controller) *MockPlanService {
	mock := &MockPlanService{ctrl: ctrl}
	mock.recorder = &MockPlanServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanService) EXPECT() *MockPlanServiceMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockPlanService) List(arg0 *packngo.GetOptions) ([]packngo.Plan, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]packngo.Plan)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockPlanServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlanService)(nil).List), arg0)
}

// OrganizationList mocks base method.
func (m *MockPlanService) OrganizationList(arg0 string, arg1 *packngo.GetOptions) ([]packngo.Plan, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationList", arg0, arg1)
	ret0, _ := ret[0].([]packngo.Plan)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OrganizationList indicates an expected call of OrganizationList.
func (mr *MockPlanServiceMockRecorder) OrganizationList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationList", reflect.TypeOf((*MockPlanService)(nil).OrganizationList), arg0, arg1)
}

// ProjectList mocks base method.
func (m *MockPlanService) ProjectList(arg0 string, arg1 *packngo.GetOptions) ([]packngo.Plan, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProjectList", arg0, arg1)
	ret0, _ := ret[0].([]packngo.Plan)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProjectList indicates an expected call of ProjectList.
func (mr *MockPlanServiceMockRecorder) ProjectList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectList", reflect.TypeOf((*MockPlanService)(nil).ProjectList), arg0, arg1)
}

// MockProjectIPService is a mock of ProjectIPService interface.
type MockProjectIPService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectIPServiceMockRecorder
}

// MockProjectIPServiceMockRecorder is the mock recorder for MockProjectIPService.
type MockProjectIPServiceMockRecorder struct {
	mock *MockProjectIPService
}

// NewMockProjectIPService creates a new mock instance.
func NewMockProjectIPService(ctrl *gomock.Controller) *MockProjectIPService {
	mock := &MockProjectIPService{ctrl: ctrl}
	mock.recorder = &MockProjectIPServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectIPService) EXPECT() *MockProjectIPServiceMockRecorder {
	return m.recorder
}

// AvailableAddresses mocks base method.
func (m *MockProjectIPService) AvailableAddresses(arg0 string, arg1 *packngo.AvailableRequest) ([]string, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableAddresses", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AvailableAddresses indicates an expected call of AvailableAddresses.
func (mr *MockProjectIPServiceMockRecorder) AvailableAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableAddresses", reflect.TypeOf((*MockProjectIPService)(nil).AvailableAddresses), arg0, arg1)
}

// Get mocks base method.
func (m *MockProjectIPService) Get(arg0 string, arg1 *packngo.GetOptions) (*packngo.IPAddressReservation, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*packngo.IPAddressReservation)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockProjectIPServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProjectIPService)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockProjectIPService) List(arg0 string, arg1 *packngo.GetOptions) ([]packngo.IPAddressReservation, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]packngo.IPAddressReservation)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockProjectIPServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectIPService)(nil).List), arg0, arg1)
}

// Remove mocks base method.
func (m *MockProjectIPService) Remove(arg0 string) (*packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(*packngo.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove.
func (mr *MockProjectIPServiceMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockProjectIPService)(nil).Remove), arg0)
}

// Request mocks base method.
func (m *MockProjectIPService) Request(arg0 string, arg1 *packngo.IPReservationRequest) (*packngo.IPAddressReservation, *packngo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", arg0, arg1)
	ret0, _ := ret[0].(*packngo.IPAddressReservation)
	ret1, _ := ret[1].(*packngo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Request indicates an expected call of Request.
func (mr *MockProjectIPServiceMockRecorder) Request(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockProjectIPService)(nil).Request), arg0, arg1)
}
