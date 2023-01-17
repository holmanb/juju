// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/instancemutater (interfaces: InstanceMutatorWatcher,InstanceMutaterState,Machine,Unit,Application,Charm)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v10"
	instancemutater "github.com/juju/juju/apiserver/facades/agent/instancemutater"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
	status "github.com/juju/juju/core/status"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
)

// MockInstanceMutatorWatcher is a mock of InstanceMutatorWatcher interface.
type MockInstanceMutatorWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMutatorWatcherMockRecorder
}

// MockInstanceMutatorWatcherMockRecorder is the mock recorder for MockInstanceMutatorWatcher.
type MockInstanceMutatorWatcherMockRecorder struct {
	mock *MockInstanceMutatorWatcher
}

// NewMockInstanceMutatorWatcher creates a new mock instance.
func NewMockInstanceMutatorWatcher(ctrl *gomock.Controller) *MockInstanceMutatorWatcher {
	mock := &MockInstanceMutatorWatcher{ctrl: ctrl}
	mock.recorder = &MockInstanceMutatorWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceMutatorWatcher) EXPECT() *MockInstanceMutatorWatcherMockRecorder {
	return m.recorder
}

// WatchLXDProfileVerificationForMachine mocks base method.
func (m *MockInstanceMutatorWatcher) WatchLXDProfileVerificationForMachine(arg0 instancemutater.Machine) (state.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLXDProfileVerificationForMachine", arg0)
	ret0, _ := ret[0].(state.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileVerificationForMachine indicates an expected call of WatchLXDProfileVerificationForMachine.
func (mr *MockInstanceMutatorWatcherMockRecorder) WatchLXDProfileVerificationForMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileVerificationForMachine", reflect.TypeOf((*MockInstanceMutatorWatcher)(nil).WatchLXDProfileVerificationForMachine), arg0)
}

// MockInstanceMutaterState is a mock of InstanceMutaterState interface.
type MockInstanceMutaterState struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMutaterStateMockRecorder
}

// MockInstanceMutaterStateMockRecorder is the mock recorder for MockInstanceMutaterState.
type MockInstanceMutaterStateMockRecorder struct {
	mock *MockInstanceMutaterState
}

// NewMockInstanceMutaterState creates a new mock instance.
func NewMockInstanceMutaterState(ctrl *gomock.Controller) *MockInstanceMutaterState {
	mock := &MockInstanceMutaterState{ctrl: ctrl}
	mock.recorder = &MockInstanceMutaterStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceMutaterState) EXPECT() *MockInstanceMutaterStateMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockInstanceMutaterState) Application(arg0 string) (instancemutater.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(instancemutater.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockInstanceMutaterStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockInstanceMutaterState)(nil).Application), arg0)
}

// Charm mocks base method.
func (m *MockInstanceMutaterState) Charm(arg0 *charm.URL) (instancemutater.Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(instancemutater.Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockInstanceMutaterStateMockRecorder) Charm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockInstanceMutaterState)(nil).Charm), arg0)
}

// ControllerTimestamp mocks base method.
func (m *MockInstanceMutaterState) ControllerTimestamp() (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTimestamp")
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerTimestamp indicates an expected call of ControllerTimestamp.
func (mr *MockInstanceMutaterStateMockRecorder) ControllerTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTimestamp", reflect.TypeOf((*MockInstanceMutaterState)(nil).ControllerTimestamp))
}

// FindEntity mocks base method.
func (m *MockInstanceMutaterState) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockInstanceMutaterStateMockRecorder) FindEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockInstanceMutaterState)(nil).FindEntity), arg0)
}

// Machine mocks base method.
func (m *MockInstanceMutaterState) Machine(arg0 string) (instancemutater.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(instancemutater.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockInstanceMutaterStateMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockInstanceMutaterState)(nil).Machine), arg0)
}

// ModelName mocks base method.
func (m *MockInstanceMutaterState) ModelName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelName indicates an expected call of ModelName.
func (mr *MockInstanceMutaterStateMockRecorder) ModelName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelName", reflect.TypeOf((*MockInstanceMutaterState)(nil).ModelName))
}

// Unit mocks base method.
func (m *MockInstanceMutaterState) Unit(arg0 string) (instancemutater.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(instancemutater.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockInstanceMutaterStateMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockInstanceMutaterState)(nil).Unit), arg0)
}

// WatchApplicationCharms mocks base method.
func (m *MockInstanceMutaterState) WatchApplicationCharms() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplicationCharms")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchApplicationCharms indicates an expected call of WatchApplicationCharms.
func (mr *MockInstanceMutaterStateMockRecorder) WatchApplicationCharms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplicationCharms", reflect.TypeOf((*MockInstanceMutaterState)(nil).WatchApplicationCharms))
}

// WatchCharms mocks base method.
func (m *MockInstanceMutaterState) WatchCharms() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchCharms")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchCharms indicates an expected call of WatchCharms.
func (mr *MockInstanceMutaterStateMockRecorder) WatchCharms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchCharms", reflect.TypeOf((*MockInstanceMutaterState)(nil).WatchCharms))
}

// WatchMachines mocks base method.
func (m *MockInstanceMutaterState) WatchMachines() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchMachines")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchMachines indicates an expected call of WatchMachines.
func (mr *MockInstanceMutaterStateMockRecorder) WatchMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchMachines", reflect.TypeOf((*MockInstanceMutaterState)(nil).WatchMachines))
}

// WatchModelMachines mocks base method.
func (m *MockInstanceMutaterState) WatchModelMachines() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModelMachines")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchModelMachines indicates an expected call of WatchModelMachines.
func (mr *MockInstanceMutaterStateMockRecorder) WatchModelMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModelMachines", reflect.TypeOf((*MockInstanceMutaterState)(nil).WatchModelMachines))
}

// WatchUnits mocks base method.
func (m *MockInstanceMutaterState) WatchUnits() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUnits")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchUnits indicates an expected call of WatchUnits.
func (mr *MockInstanceMutaterStateMockRecorder) WatchUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnits", reflect.TypeOf((*MockInstanceMutaterState)(nil).WatchUnits))
}

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// CharmProfiles mocks base method.
func (m *MockMachine) CharmProfiles() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmProfiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmProfiles indicates an expected call of CharmProfiles.
func (mr *MockMachineMockRecorder) CharmProfiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmProfiles", reflect.TypeOf((*MockMachine)(nil).CharmProfiles))
}

// ContainerType mocks base method.
func (m *MockMachine) ContainerType() instance.ContainerType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockMachineMockRecorder) ContainerType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockMachine)(nil).ContainerType))
}

// Id mocks base method.
func (m *MockMachine) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockMachineMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMachine)(nil).Id))
}

// InstanceId mocks base method.
func (m *MockMachine) InstanceId() (instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceId")
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceId indicates an expected call of InstanceId.
func (mr *MockMachineMockRecorder) InstanceId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceId", reflect.TypeOf((*MockMachine)(nil).InstanceId))
}

// IsManual mocks base method.
func (m *MockMachine) IsManual() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManual")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManual indicates an expected call of IsManual.
func (mr *MockMachineMockRecorder) IsManual() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManual", reflect.TypeOf((*MockMachine)(nil).IsManual))
}

// SetCharmProfiles mocks base method.
func (m *MockMachine) SetCharmProfiles(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharmProfiles", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharmProfiles indicates an expected call of SetCharmProfiles.
func (mr *MockMachineMockRecorder) SetCharmProfiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharmProfiles", reflect.TypeOf((*MockMachine)(nil).SetCharmProfiles), arg0)
}

// SetModificationStatus mocks base method.
func (m *MockMachine) SetModificationStatus(arg0 status.StatusInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModificationStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModificationStatus indicates an expected call of SetModificationStatus.
func (mr *MockMachineMockRecorder) SetModificationStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModificationStatus", reflect.TypeOf((*MockMachine)(nil).SetModificationStatus), arg0)
}

// Units mocks base method.
func (m *MockMachine) Units() ([]instancemutater.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]instancemutater.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockMachineMockRecorder) Units() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockMachine)(nil).Units))
}

// WatchContainers mocks base method.
func (m *MockMachine) WatchContainers(arg0 instance.ContainerType) state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchContainers", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchContainers indicates an expected call of WatchContainers.
func (mr *MockMachineMockRecorder) WatchContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchContainers", reflect.TypeOf((*MockMachine)(nil).WatchContainers), arg0)
}

// WatchInstanceData mocks base method.
func (m *MockMachine) WatchInstanceData() state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchInstanceData")
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// WatchInstanceData indicates an expected call of WatchInstanceData.
func (mr *MockMachineMockRecorder) WatchInstanceData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchInstanceData", reflect.TypeOf((*MockMachine)(nil).WatchInstanceData))
}

// MockUnit is a mock of Unit interface.
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit.
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance.
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockUnit) Application() (instancemutater.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(instancemutater.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockUnitMockRecorder) Application() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUnit)(nil).Application))
}

// ApplicationName mocks base method.
func (m *MockUnit) ApplicationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ApplicationName indicates an expected call of ApplicationName.
func (mr *MockUnitMockRecorder) ApplicationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationName", reflect.TypeOf((*MockUnit)(nil).ApplicationName))
}

// AssignedMachineId mocks base method.
func (m *MockUnit) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockUnitMockRecorder) AssignedMachineId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockUnit)(nil).AssignedMachineId))
}

// CharmURL mocks base method.
func (m *MockUnit) CharmURL() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*string)
	return ret0
}

// CharmURL indicates an expected call of CharmURL.
func (mr *MockUnitMockRecorder) CharmURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockUnit)(nil).CharmURL))
}

// Life mocks base method.
func (m *MockUnit) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockUnitMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockUnit)(nil).Life))
}

// Name mocks base method.
func (m *MockUnit) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockUnitMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnit)(nil).Name))
}

// PrincipalName mocks base method.
func (m *MockUnit) PrincipalName() (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrincipalName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// PrincipalName indicates an expected call of PrincipalName.
func (mr *MockUnitMockRecorder) PrincipalName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrincipalName", reflect.TypeOf((*MockUnit)(nil).PrincipalName))
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// CharmURL mocks base method.
func (m *MockApplication) CharmURL() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*string)
	return ret0
}

// CharmURL indicates an expected call of CharmURL.
func (mr *MockApplicationMockRecorder) CharmURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockApplication)(nil).CharmURL))
}

// MockCharm is a mock of Charm interface.
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm.
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance.
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method.
func (m *MockCharm) LXDProfile() lxdprofile.Profile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(lxdprofile.Profile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockCharmMockRecorder) LXDProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockCharm)(nil).LXDProfile))
}

// Revision mocks base method.
func (m *MockCharm) Revision() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision.
func (mr *MockCharmMockRecorder) Revision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockCharm)(nil).Revision))
}
