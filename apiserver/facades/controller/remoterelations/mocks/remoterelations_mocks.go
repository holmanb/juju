// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/remoterelations (interfaces: RemoteRelationsState,ControllerConfigAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	crossmodel "github.com/juju/juju/apiserver/common/crossmodel"
	crossmodel0 "github.com/juju/juju/core/crossmodel"
	permission "github.com/juju/juju/core/permission"
	secrets "github.com/juju/juju/core/secrets"
	config "github.com/juju/juju/environs/config"
	params "github.com/juju/juju/rpc/params"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	macaroon "gopkg.in/macaroon.v2"
)

// MockRemoteRelationsState is a mock of RemoteRelationsState interface.
type MockRemoteRelationsState struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteRelationsStateMockRecorder
}

// MockRemoteRelationsStateMockRecorder is the mock recorder for MockRemoteRelationsState.
type MockRemoteRelationsStateMockRecorder struct {
	mock *MockRemoteRelationsState
}

// NewMockRemoteRelationsState creates a new mock instance.
func NewMockRemoteRelationsState(ctrl *gomock.Controller) *MockRemoteRelationsState {
	mock := &MockRemoteRelationsState{ctrl: ctrl}
	mock.recorder = &MockRemoteRelationsStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteRelationsState) EXPECT() *MockRemoteRelationsStateMockRecorder {
	return m.recorder
}

// AddRelation mocks base method.
func (m *MockRemoteRelationsState) AddRelation(arg0 ...state.Endpoint) (crossmodel.Relation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddRelation", varargs...)
	ret0, _ := ret[0].(crossmodel.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRelation indicates an expected call of AddRelation.
func (mr *MockRemoteRelationsStateMockRecorder) AddRelation(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRelation", reflect.TypeOf((*MockRemoteRelationsState)(nil).AddRelation), arg0...)
}

// AddRemoteApplication mocks base method.
func (m *MockRemoteRelationsState) AddRemoteApplication(arg0 state.AddRemoteApplicationParams) (crossmodel.RemoteApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteApplication", arg0)
	ret0, _ := ret[0].(crossmodel.RemoteApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRemoteApplication indicates an expected call of AddRemoteApplication.
func (mr *MockRemoteRelationsStateMockRecorder) AddRemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteApplication", reflect.TypeOf((*MockRemoteRelationsState)(nil).AddRemoteApplication), arg0)
}

// AllModelUUIDs mocks base method.
func (m *MockRemoteRelationsState) AllModelUUIDs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelUUIDs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelUUIDs indicates an expected call of AllModelUUIDs.
func (mr *MockRemoteRelationsStateMockRecorder) AllModelUUIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelUUIDs", reflect.TypeOf((*MockRemoteRelationsState)(nil).AllModelUUIDs))
}

// AppNameForOffer mocks base method.
func (m *MockRemoteRelationsState) AppNameForOffer(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppNameForOffer", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppNameForOffer indicates an expected call of AppNameForOffer.
func (mr *MockRemoteRelationsStateMockRecorder) AppNameForOffer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppNameForOffer", reflect.TypeOf((*MockRemoteRelationsState)(nil).AppNameForOffer), arg0)
}

// Application mocks base method.
func (m *MockRemoteRelationsState) Application(arg0 string) (crossmodel.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(crossmodel.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockRemoteRelationsStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockRemoteRelationsState)(nil).Application), arg0)
}

// ApplicationOfferForUUID mocks base method.
func (m *MockRemoteRelationsState) ApplicationOfferForUUID(arg0 string) (*crossmodel0.ApplicationOffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationOfferForUUID", arg0)
	ret0, _ := ret[0].(*crossmodel0.ApplicationOffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationOfferForUUID indicates an expected call of ApplicationOfferForUUID.
func (mr *MockRemoteRelationsStateMockRecorder) ApplicationOfferForUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationOfferForUUID", reflect.TypeOf((*MockRemoteRelationsState)(nil).ApplicationOfferForUUID), arg0)
}

// ApplyOperation mocks base method.
func (m *MockRemoteRelationsState) ApplyOperation(arg0 state.ModelOperation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyOperation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyOperation indicates an expected call of ApplyOperation.
func (mr *MockRemoteRelationsStateMockRecorder) ApplyOperation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyOperation", reflect.TypeOf((*MockRemoteRelationsState)(nil).ApplyOperation), arg0)
}

// ControllerTag mocks base method.
func (m *MockRemoteRelationsState) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockRemoteRelationsStateMockRecorder) ControllerTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockRemoteRelationsState)(nil).ControllerTag))
}

// EndpointsRelation mocks base method.
func (m *MockRemoteRelationsState) EndpointsRelation(arg0 ...state.Endpoint) (crossmodel.Relation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EndpointsRelation", varargs...)
	ret0, _ := ret[0].(crossmodel.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndpointsRelation indicates an expected call of EndpointsRelation.
func (mr *MockRemoteRelationsStateMockRecorder) EndpointsRelation(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointsRelation", reflect.TypeOf((*MockRemoteRelationsState)(nil).EndpointsRelation), arg0...)
}

// ExportLocalEntity mocks base method.
func (m *MockRemoteRelationsState) ExportLocalEntity(arg0 names.Tag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportLocalEntity", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportLocalEntity indicates an expected call of ExportLocalEntity.
func (mr *MockRemoteRelationsStateMockRecorder) ExportLocalEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportLocalEntity", reflect.TypeOf((*MockRemoteRelationsState)(nil).ExportLocalEntity), arg0)
}

// GetOfferAccess mocks base method.
func (m *MockRemoteRelationsState) GetOfferAccess(arg0 string, arg1 names.UserTag) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOfferAccess", arg0, arg1)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOfferAccess indicates an expected call of GetOfferAccess.
func (mr *MockRemoteRelationsStateMockRecorder) GetOfferAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfferAccess", reflect.TypeOf((*MockRemoteRelationsState)(nil).GetOfferAccess), arg0, arg1)
}

// GetRemoteEntity mocks base method.
func (m *MockRemoteRelationsState) GetRemoteEntity(arg0 string) (names.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteEntity", arg0)
	ret0, _ := ret[0].(names.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteEntity indicates an expected call of GetRemoteEntity.
func (mr *MockRemoteRelationsStateMockRecorder) GetRemoteEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteEntity", reflect.TypeOf((*MockRemoteRelationsState)(nil).GetRemoteEntity), arg0)
}

// GetToken mocks base method.
func (m *MockRemoteRelationsState) GetToken(arg0 names.Tag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockRemoteRelationsStateMockRecorder) GetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockRemoteRelationsState)(nil).GetToken), arg0)
}

// ImportRemoteEntity mocks base method.
func (m *MockRemoteRelationsState) ImportRemoteEntity(arg0 names.Tag, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportRemoteEntity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportRemoteEntity indicates an expected call of ImportRemoteEntity.
func (mr *MockRemoteRelationsStateMockRecorder) ImportRemoteEntity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRemoteEntity", reflect.TypeOf((*MockRemoteRelationsState)(nil).ImportRemoteEntity), arg0, arg1)
}

// IngressNetworks mocks base method.
func (m *MockRemoteRelationsState) IngressNetworks(arg0 string) (state.RelationNetworks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngressNetworks", arg0)
	ret0, _ := ret[0].(state.RelationNetworks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngressNetworks indicates an expected call of IngressNetworks.
func (mr *MockRemoteRelationsStateMockRecorder) IngressNetworks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngressNetworks", reflect.TypeOf((*MockRemoteRelationsState)(nil).IngressNetworks), arg0)
}

// KeyRelation mocks base method.
func (m *MockRemoteRelationsState) KeyRelation(arg0 string) (crossmodel.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyRelation", arg0)
	ret0, _ := ret[0].(crossmodel.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeyRelation indicates an expected call of KeyRelation.
func (mr *MockRemoteRelationsStateMockRecorder) KeyRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyRelation", reflect.TypeOf((*MockRemoteRelationsState)(nil).KeyRelation), arg0)
}

// ModelConfig mocks base method.
func (m *MockRemoteRelationsState) ModelConfig() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockRemoteRelationsStateMockRecorder) ModelConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockRemoteRelationsState)(nil).ModelConfig))
}

// ModelTag mocks base method.
func (m *MockRemoteRelationsState) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockRemoteRelationsStateMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockRemoteRelationsState)(nil).ModelTag))
}

// ModelUUID mocks base method.
func (m *MockRemoteRelationsState) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockRemoteRelationsStateMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockRemoteRelationsState)(nil).ModelUUID))
}

// OfferConnectionForRelation mocks base method.
func (m *MockRemoteRelationsState) OfferConnectionForRelation(arg0 string) (crossmodel.OfferConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferConnectionForRelation", arg0)
	ret0, _ := ret[0].(crossmodel.OfferConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OfferConnectionForRelation indicates an expected call of OfferConnectionForRelation.
func (mr *MockRemoteRelationsStateMockRecorder) OfferConnectionForRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferConnectionForRelation", reflect.TypeOf((*MockRemoteRelationsState)(nil).OfferConnectionForRelation), arg0)
}

// OfferNameForRelation mocks base method.
func (m *MockRemoteRelationsState) OfferNameForRelation(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferNameForRelation", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OfferNameForRelation indicates an expected call of OfferNameForRelation.
func (mr *MockRemoteRelationsStateMockRecorder) OfferNameForRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferNameForRelation", reflect.TypeOf((*MockRemoteRelationsState)(nil).OfferNameForRelation), arg0)
}

// RemoteApplication mocks base method.
func (m *MockRemoteRelationsState) RemoteApplication(arg0 string) (crossmodel.RemoteApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteApplication", arg0)
	ret0, _ := ret[0].(crossmodel.RemoteApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteApplication indicates an expected call of RemoteApplication.
func (mr *MockRemoteRelationsStateMockRecorder) RemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplication", reflect.TypeOf((*MockRemoteRelationsState)(nil).RemoteApplication), arg0)
}

// RemoveRemoteEntity mocks base method.
func (m *MockRemoteRelationsState) RemoveRemoteEntity(arg0 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRemoteEntity", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRemoteEntity indicates an expected call of RemoveRemoteEntity.
func (mr *MockRemoteRelationsStateMockRecorder) RemoveRemoteEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteEntity", reflect.TypeOf((*MockRemoteRelationsState)(nil).RemoveRemoteEntity), arg0)
}

// RemoveSecretConsumer mocks base method.
func (m *MockRemoteRelationsState) RemoveSecretConsumer(arg0 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretConsumer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretConsumer indicates an expected call of RemoveSecretConsumer.
func (mr *MockRemoteRelationsStateMockRecorder) RemoveSecretConsumer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretConsumer", reflect.TypeOf((*MockRemoteRelationsState)(nil).RemoveSecretConsumer), arg0)
}

// SaveIngressNetworks mocks base method.
func (m *MockRemoteRelationsState) SaveIngressNetworks(arg0 string, arg1 []string) (state.RelationNetworks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveIngressNetworks", arg0, arg1)
	ret0, _ := ret[0].(state.RelationNetworks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveIngressNetworks indicates an expected call of SaveIngressNetworks.
func (mr *MockRemoteRelationsStateMockRecorder) SaveIngressNetworks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveIngressNetworks", reflect.TypeOf((*MockRemoteRelationsState)(nil).SaveIngressNetworks), arg0, arg1)
}

// SaveMacaroon mocks base method.
func (m *MockRemoteRelationsState) SaveMacaroon(arg0 names.Tag, arg1 *macaroon.Macaroon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveMacaroon", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveMacaroon indicates an expected call of SaveMacaroon.
func (mr *MockRemoteRelationsStateMockRecorder) SaveMacaroon(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveMacaroon", reflect.TypeOf((*MockRemoteRelationsState)(nil).SaveMacaroon), arg0, arg1)
}

// UpdateControllerForModel mocks base method.
func (m *MockRemoteRelationsState) UpdateControllerForModel(arg0 crossmodel0.ControllerInfo, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateControllerForModel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateControllerForModel indicates an expected call of UpdateControllerForModel.
func (mr *MockRemoteRelationsStateMockRecorder) UpdateControllerForModel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateControllerForModel", reflect.TypeOf((*MockRemoteRelationsState)(nil).UpdateControllerForModel), arg0, arg1)
}

// UpdateSecretConsumerOperation mocks base method.
func (m *MockRemoteRelationsState) UpdateSecretConsumerOperation(arg0 *secrets.URI, arg1 int) (state.ModelOperation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretConsumerOperation", arg0, arg1)
	ret0, _ := ret[0].(state.ModelOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecretConsumerOperation indicates an expected call of UpdateSecretConsumerOperation.
func (mr *MockRemoteRelationsStateMockRecorder) UpdateSecretConsumerOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretConsumerOperation", reflect.TypeOf((*MockRemoteRelationsState)(nil).UpdateSecretConsumerOperation), arg0, arg1)
}

// UserPermission mocks base method.
func (m *MockRemoteRelationsState) UserPermission(arg0 names.UserTag, arg1 names.Tag) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserPermission", arg0, arg1)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserPermission indicates an expected call of UserPermission.
func (mr *MockRemoteRelationsStateMockRecorder) UserPermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserPermission", reflect.TypeOf((*MockRemoteRelationsState)(nil).UserPermission), arg0, arg1)
}

// WatchOffer mocks base method.
func (m *MockRemoteRelationsState) WatchOffer(arg0 string) state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOffer", arg0)
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// WatchOffer indicates an expected call of WatchOffer.
func (mr *MockRemoteRelationsStateMockRecorder) WatchOffer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOffer", reflect.TypeOf((*MockRemoteRelationsState)(nil).WatchOffer), arg0)
}

// WatchOfferStatus mocks base method.
func (m *MockRemoteRelationsState) WatchOfferStatus(arg0 string) (state.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOfferStatus", arg0)
	ret0, _ := ret[0].(state.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOfferStatus indicates an expected call of WatchOfferStatus.
func (mr *MockRemoteRelationsStateMockRecorder) WatchOfferStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOfferStatus", reflect.TypeOf((*MockRemoteRelationsState)(nil).WatchOfferStatus), arg0)
}

// WatchRemoteApplicationRelations mocks base method.
func (m *MockRemoteRelationsState) WatchRemoteApplicationRelations(arg0 string) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRemoteApplicationRelations", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRemoteApplicationRelations indicates an expected call of WatchRemoteApplicationRelations.
func (mr *MockRemoteRelationsStateMockRecorder) WatchRemoteApplicationRelations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRemoteApplicationRelations", reflect.TypeOf((*MockRemoteRelationsState)(nil).WatchRemoteApplicationRelations), arg0)
}

// WatchRemoteApplications mocks base method.
func (m *MockRemoteRelationsState) WatchRemoteApplications() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRemoteApplications")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchRemoteApplications indicates an expected call of WatchRemoteApplications.
func (mr *MockRemoteRelationsStateMockRecorder) WatchRemoteApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRemoteApplications", reflect.TypeOf((*MockRemoteRelationsState)(nil).WatchRemoteApplications))
}

// WatchRemoteRelations mocks base method.
func (m *MockRemoteRelationsState) WatchRemoteRelations() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRemoteRelations")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchRemoteRelations indicates an expected call of WatchRemoteRelations.
func (mr *MockRemoteRelationsStateMockRecorder) WatchRemoteRelations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRemoteRelations", reflect.TypeOf((*MockRemoteRelationsState)(nil).WatchRemoteRelations))
}

// MockControllerConfigAPI is a mock of ControllerConfigAPI interface.
type MockControllerConfigAPI struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigAPIMockRecorder
}

// MockControllerConfigAPIMockRecorder is the mock recorder for MockControllerConfigAPI.
type MockControllerConfigAPIMockRecorder struct {
	mock *MockControllerConfigAPI
}

// NewMockControllerConfigAPI creates a new mock instance.
func NewMockControllerConfigAPI(ctrl *gomock.Controller) *MockControllerConfigAPI {
	mock := &MockControllerConfigAPI{ctrl: ctrl}
	mock.recorder = &MockControllerConfigAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigAPI) EXPECT() *MockControllerConfigAPIMockRecorder {
	return m.recorder
}

// ControllerAPIInfoForModels mocks base method.
func (m *MockControllerConfigAPI) ControllerAPIInfoForModels(arg0 params.Entities) (params.ControllerAPIInfoResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAPIInfoForModels", arg0)
	ret0, _ := ret[0].(params.ControllerAPIInfoResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerAPIInfoForModels indicates an expected call of ControllerAPIInfoForModels.
func (mr *MockControllerConfigAPIMockRecorder) ControllerAPIInfoForModels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAPIInfoForModels", reflect.TypeOf((*MockControllerConfigAPI)(nil).ControllerAPIInfoForModels), arg0)
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigAPI) ControllerConfig() (params.ControllerConfigResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(params.ControllerConfigResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigAPIMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigAPI)(nil).ControllerConfig))
}
