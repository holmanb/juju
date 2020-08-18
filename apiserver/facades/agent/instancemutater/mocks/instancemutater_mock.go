// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/instancemutater (interfaces: InstanceMutaterState,Machine,Unit,Application,Charm)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	charm_v6 "github.com/juju/charm/v8"
	instancemutater "github.com/juju/juju/apiserver/facades/agent/instancemutater"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
	status "github.com/juju/juju/core/status"
	state "github.com/juju/juju/state"
	names_v3 "github.com/juju/names/v4"
)

// MockInstanceMutaterState is a mock of InstanceMutaterState interface
type MockInstanceMutaterState struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMutaterStateMockRecorder
}

// MockInstanceMutaterStateMockRecorder is the mock recorder for MockInstanceMutaterState
type MockInstanceMutaterStateMockRecorder struct {
	mock *MockInstanceMutaterState
}

// NewMockInstanceMutaterState creates a new mock instance
func NewMockInstanceMutaterState(ctrl *gomock.Controller) *MockInstanceMutaterState {
	mock := &MockInstanceMutaterState{ctrl: ctrl}
	mock.recorder = &MockInstanceMutaterStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInstanceMutaterState) EXPECT() *MockInstanceMutaterStateMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockInstanceMutaterState) Application(arg0 string) (instancemutater.Application, error) {
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(instancemutater.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockInstanceMutaterStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockInstanceMutaterState)(nil).Application), arg0)
}

// Charm mocks base method
func (m *MockInstanceMutaterState) Charm(arg0 *charm_v6.URL) (instancemutater.Charm, error) {
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(instancemutater.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm
func (mr *MockInstanceMutaterStateMockRecorder) Charm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockInstanceMutaterState)(nil).Charm), arg0)
}

// ControllerTimestamp mocks base method
func (m *MockInstanceMutaterState) ControllerTimestamp() (*time.Time, error) {
	ret := m.ctrl.Call(m, "ControllerTimestamp")
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerTimestamp indicates an expected call of ControllerTimestamp
func (mr *MockInstanceMutaterStateMockRecorder) ControllerTimestamp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTimestamp", reflect.TypeOf((*MockInstanceMutaterState)(nil).ControllerTimestamp))
}

// FindEntity mocks base method
func (m *MockInstanceMutaterState) FindEntity(arg0 names_v3.Tag) (state.Entity, error) {
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity
func (mr *MockInstanceMutaterStateMockRecorder) FindEntity(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockInstanceMutaterState)(nil).FindEntity), arg0)
}

// MockMachine is a mock of Machine interface
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// CharmProfiles mocks base method
func (m *MockMachine) CharmProfiles() ([]string, error) {
	ret := m.ctrl.Call(m, "CharmProfiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmProfiles indicates an expected call of CharmProfiles
func (mr *MockMachineMockRecorder) CharmProfiles() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmProfiles", reflect.TypeOf((*MockMachine)(nil).CharmProfiles))
}

// InstanceId mocks base method
func (m *MockMachine) InstanceId() (instance.Id, error) {
	ret := m.ctrl.Call(m, "InstanceId")
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceId indicates an expected call of InstanceId
func (mr *MockMachineMockRecorder) InstanceId() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceId", reflect.TypeOf((*MockMachine)(nil).InstanceId))
}

// SetCharmProfiles mocks base method
func (m *MockMachine) SetCharmProfiles(arg0 []string) error {
	ret := m.ctrl.Call(m, "SetCharmProfiles", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharmProfiles indicates an expected call of SetCharmProfiles
func (mr *MockMachineMockRecorder) SetCharmProfiles(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharmProfiles", reflect.TypeOf((*MockMachine)(nil).SetCharmProfiles), arg0)
}

// SetModificationStatus mocks base method
func (m *MockMachine) SetModificationStatus(arg0 status.StatusInfo) error {
	ret := m.ctrl.Call(m, "SetModificationStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModificationStatus indicates an expected call of SetModificationStatus
func (mr *MockMachineMockRecorder) SetModificationStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModificationStatus", reflect.TypeOf((*MockMachine)(nil).SetModificationStatus), arg0)
}

// Units mocks base method
func (m *MockMachine) Units() ([]instancemutater.Unit, error) {
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]instancemutater.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockMachineMockRecorder) Units() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockMachine)(nil).Units))
}

// MockUnit is a mock of Unit interface
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockUnit) Application() string {
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(string)
	return ret0
}

// Application indicates an expected call of Application
func (mr *MockUnitMockRecorder) Application() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUnit)(nil).Application))
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// CharmURL mocks base method
func (m *MockApplication) CharmURL() *charm_v6.URL {
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*charm_v6.URL)
	return ret0
}

// CharmURL indicates an expected call of CharmURL
func (mr *MockApplicationMockRecorder) CharmURL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockApplication)(nil).CharmURL))
}

// MockCharm is a mock of Charm interface
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method
func (m *MockCharm) LXDProfile() lxdprofile.Profile {
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(lxdprofile.Profile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile
func (mr *MockCharmMockRecorder) LXDProfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockCharm)(nil).LXDProfile))
}
