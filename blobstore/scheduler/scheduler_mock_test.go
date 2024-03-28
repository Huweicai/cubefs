// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/scheduler (interfaces: ITaskRunner,IVolumeCache,MMigrator,IVolumeInspector,IClusterTopology)

// Package scheduler is a generated GoMock package.
package scheduler

import (
	context "context"
	reflect "reflect"

	scheduler "github.com/cubefs/cubefs/blobstore/api/scheduler"
	proto "github.com/cubefs/cubefs/blobstore/common/proto"
	client "github.com/cubefs/cubefs/blobstore/scheduler/client"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskRunner is a mock of ITaskRunner interface.
type MockTaskRunner struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRunnerMockRecorder
}

// MockTaskRunnerMockRecorder is the mock recorder for MockTaskRunner.
type MockTaskRunnerMockRecorder struct {
	mock *MockTaskRunner
}

// NewMockTaskRunner creates a new mock instance.
func NewMockTaskRunner(ctrl *gomock.Controller) *MockTaskRunner {
	mock := &MockTaskRunner{ctrl: ctrl}
	mock.recorder = &MockTaskRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRunner) EXPECT() *MockTaskRunnerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTaskRunner) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockTaskRunnerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTaskRunner)(nil).Close))
}

// Enabled mocks base method.
func (m *MockTaskRunner) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockTaskRunnerMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockTaskRunner)(nil).Enabled))
}

// GetErrorStats mocks base method.
func (m *MockTaskRunner) GetErrorStats() ([]string, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErrorStats")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// GetErrorStats indicates an expected call of GetErrorStats.
func (mr *MockTaskRunnerMockRecorder) GetErrorStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorStats", reflect.TypeOf((*MockTaskRunner)(nil).GetErrorStats))
}

// GetTaskStats mocks base method.
func (m *MockTaskRunner) GetTaskStats() ([20]int, [20]int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskStats")
	ret0, _ := ret[0].([20]int)
	ret1, _ := ret[1].([20]int)
	return ret0, ret1
}

// GetTaskStats indicates an expected call of GetTaskStats.
func (mr *MockTaskRunnerMockRecorder) GetTaskStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskStats", reflect.TypeOf((*MockTaskRunner)(nil).GetTaskStats))
}

// Run mocks base method.
func (m *MockTaskRunner) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockTaskRunnerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTaskRunner)(nil).Run))
}

// MockVolumeCache is a mock of IVolumeCache interface.
type MockVolumeCache struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeCacheMockRecorder
}

// MockVolumeCacheMockRecorder is the mock recorder for MockVolumeCache.
type MockVolumeCacheMockRecorder struct {
	mock *MockVolumeCache
}

// NewMockVolumeCache creates a new mock instance.
func NewMockVolumeCache(ctrl *gomock.Controller) *MockVolumeCache {
	mock := &MockVolumeCache{ctrl: ctrl}
	mock.recorder = &MockVolumeCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeCache) EXPECT() *MockVolumeCacheMockRecorder {
	return m.recorder
}

// GetVolume mocks base method.
func (m *MockVolumeCache) GetVolume(arg0 proto.Vid) (*client.VolumeInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0)
	ret0, _ := ret[0].(*client.VolumeInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume.
func (mr *MockVolumeCacheMockRecorder) GetVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockVolumeCache)(nil).GetVolume), arg0)
}

// LoadVolumes mocks base method.
func (m *MockVolumeCache) LoadVolumes() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadVolumes")
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadVolumes indicates an expected call of LoadVolumes.
func (mr *MockVolumeCacheMockRecorder) LoadVolumes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadVolumes", reflect.TypeOf((*MockVolumeCache)(nil).LoadVolumes))
}

// UpdateVolume mocks base method.
func (m *MockVolumeCache) UpdateVolume(arg0 proto.Vid) (*client.VolumeInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVolume", arg0)
	ret0, _ := ret[0].(*client.VolumeInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVolume indicates an expected call of UpdateVolume.
func (mr *MockVolumeCacheMockRecorder) UpdateVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVolume", reflect.TypeOf((*MockVolumeCache)(nil).UpdateVolume), arg0)
}

// MockMigrater is a mock of MMigrator interface.
type MockMigrater struct {
	ctrl     *gomock.Controller
	recorder *MockMigraterMockRecorder
}

// MockMigraterMockRecorder is the mock recorder for MockMigrater.
type MockMigraterMockRecorder struct {
	mock *MockMigrater
}

// NewMockMigrater creates a new mock instance.
func NewMockMigrater(ctrl *gomock.Controller) *MockMigrater {
	mock := &MockMigrater{ctrl: ctrl}
	mock.recorder = &MockMigraterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrater) EXPECT() *MockMigraterMockRecorder {
	return m.recorder
}

// AcquireTask mocks base method.
func (m *MockMigrater) AcquireTask(arg0 context.Context, arg1 string) (proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireTask", arg0, arg1)
	ret0, _ := ret[0].(proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireTask indicates an expected call of AcquireTask.
func (mr *MockMigraterMockRecorder) AcquireTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireTask", reflect.TypeOf((*MockMigrater)(nil).AcquireTask), arg0, arg1)
}

// AddManualTask mocks base method.
func (m *MockMigrater) AddManualTask(arg0 context.Context, arg1 proto.Vuid, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddManualTask", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddManualTask indicates an expected call of AddManualTask.
func (mr *MockMigraterMockRecorder) AddManualTask(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddManualTask", reflect.TypeOf((*MockMigrater)(nil).AddManualTask), arg0, arg1, arg2)
}

// AddTask mocks base method.
func (m *MockMigrater) AddTask(arg0 context.Context, arg1 *proto.MigrateTask) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTask", arg0, arg1)
}

// AddTask indicates an expected call of AddTask.
func (mr *MockMigraterMockRecorder) AddTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockMigrater)(nil).AddTask), arg0, arg1)
}

// CancelTask mocks base method.
func (m *MockMigrater) CancelTask(arg0 context.Context, arg1 *scheduler.OperateTaskArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTask indicates an expected call of CancelTask.
func (mr *MockMigraterMockRecorder) CancelTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTask", reflect.TypeOf((*MockMigrater)(nil).CancelTask), arg0, arg1)
}

// ClearDeletedTaskByID mocks base method.
func (m *MockMigrater) ClearDeletedTaskByID(arg0 proto.DiskID, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearDeletedTaskByID", arg0, arg1)
}

// ClearDeletedTaskByID indicates an expected call of ClearDeletedTaskByID.
func (mr *MockMigraterMockRecorder) ClearDeletedTaskByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDeletedTaskByID", reflect.TypeOf((*MockMigrater)(nil).ClearDeletedTaskByID), arg0, arg1)
}

// ClearDeletedTasks mocks base method.
func (m *MockMigrater) ClearDeletedTasks(arg0 proto.DiskID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearDeletedTasks", arg0)
}

// ClearDeletedTasks indicates an expected call of ClearDeletedTasks.
func (mr *MockMigraterMockRecorder) ClearDeletedTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearDeletedTasks", reflect.TypeOf((*MockMigrater)(nil).ClearDeletedTasks), arg0)
}

// Close mocks base method.
func (m *MockMigrater) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMigraterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMigrater)(nil).Close))
}

// CompleteTask mocks base method.
func (m *MockMigrater) CompleteTask(arg0 context.Context, arg1 *scheduler.OperateTaskArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteTask indicates an expected call of CompleteTask.
func (mr *MockMigraterMockRecorder) CompleteTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTask", reflect.TypeOf((*MockMigrater)(nil).CompleteTask), arg0, arg1)
}

// DeletedTasks mocks base method.
func (m *MockMigrater) DeletedTasks() []DeletedTask {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletedTasks")
	ret0, _ := ret[0].([]DeletedTask)
	return ret0
}

// DeletedTasks indicates an expected call of DeletedTasks.
func (mr *MockMigraterMockRecorder) DeletedTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletedTasks", reflect.TypeOf((*MockMigrater)(nil).DeletedTasks))
}

// DiskProgress mocks base method.
func (m *MockMigrater) DiskProgress(arg0 context.Context, arg1 proto.DiskID) (*scheduler.DiskMigratingStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskProgress", arg0, arg1)
	ret0, _ := ret[0].(*scheduler.DiskMigratingStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiskProgress indicates an expected call of DiskProgress.
func (mr *MockMigraterMockRecorder) DiskProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskProgress", reflect.TypeOf((*MockMigrater)(nil).DiskProgress), arg0, arg1)
}

// Done mocks base method.
func (m *MockMigrater) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockMigraterMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockMigrater)(nil).Done))
}

// Enabled mocks base method.
func (m *MockMigrater) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockMigraterMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockMigrater)(nil).Enabled))
}

// FinishTaskInAdvanceWhenLockFail mocks base method.
func (m *MockMigrater) FinishTaskInAdvanceWhenLockFail(arg0 context.Context, arg1 *proto.MigrateTask) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FinishTaskInAdvanceWhenLockFail", arg0, arg1)
}

// FinishTaskInAdvanceWhenLockFail indicates an expected call of FinishTaskInAdvanceWhenLockFail.
func (mr *MockMigraterMockRecorder) FinishTaskInAdvanceWhenLockFail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishTaskInAdvanceWhenLockFail", reflect.TypeOf((*MockMigrater)(nil).FinishTaskInAdvanceWhenLockFail), arg0, arg1)
}

// GetMigratingDiskNum mocks base method.
func (m *MockMigrater) GetMigratingDiskNum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigratingDiskNum")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMigratingDiskNum indicates an expected call of GetMigratingDiskNum.
func (mr *MockMigraterMockRecorder) GetMigratingDiskNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigratingDiskNum", reflect.TypeOf((*MockMigrater)(nil).GetMigratingDiskNum))
}

// GetTask mocks base method.
func (m *MockMigrater) GetTask(arg0 context.Context, arg1 string) (*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0, arg1)
	ret0, _ := ret[0].(*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockMigraterMockRecorder) GetTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockMigrater)(nil).GetTask), arg0, arg1)
}

// IsDeletedTask mocks base method.
func (m *MockMigrater) IsDeletedTask(arg0 *proto.MigrateTask) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDeletedTask", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDeletedTask indicates an expected call of IsDeletedTask.
func (mr *MockMigraterMockRecorder) IsDeletedTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDeletedTask", reflect.TypeOf((*MockMigrater)(nil).IsDeletedTask), arg0)
}

// IsMigratingDisk mocks base method.
func (m *MockMigrater) IsMigratingDisk(arg0 proto.DiskID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMigratingDisk", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMigratingDisk indicates an expected call of IsMigratingDisk.
func (mr *MockMigraterMockRecorder) IsMigratingDisk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMigratingDisk", reflect.TypeOf((*MockMigrater)(nil).IsMigratingDisk), arg0)
}

// ListAllTask mocks base method.
func (m *MockMigrater) ListAllTask(arg0 context.Context) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllTask", arg0)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllTask indicates an expected call of ListAllTask.
func (mr *MockMigraterMockRecorder) ListAllTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllTask", reflect.TypeOf((*MockMigrater)(nil).ListAllTask), arg0)
}

// ListAllTaskByDiskID mocks base method.
func (m *MockMigrater) ListAllTaskByDiskID(arg0 context.Context, arg1 proto.DiskID) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllTaskByDiskID", arg0, arg1)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllTaskByDiskID indicates an expected call of ListAllTaskByDiskID.
func (mr *MockMigraterMockRecorder) ListAllTaskByDiskID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllTaskByDiskID", reflect.TypeOf((*MockMigrater)(nil).ListAllTaskByDiskID), arg0, arg1)
}

// Load mocks base method.
func (m *MockMigrater) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockMigraterMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockMigrater)(nil).Load))
}

// Progress mocks base method.
func (m *MockMigrater) Progress(arg0 context.Context) ([]proto.DiskID, int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Progress", arg0)
	ret0, _ := ret[0].([]proto.DiskID)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// Progress indicates an expected call of Progress.
func (mr *MockMigraterMockRecorder) Progress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Progress", reflect.TypeOf((*MockMigrater)(nil).Progress), arg0)
}

// QueryTask mocks base method.
func (m *MockMigrater) QueryTask(arg0 context.Context, arg1 string) (*scheduler.MigrateTaskDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTask", arg0, arg1)
	ret0, _ := ret[0].(*scheduler.MigrateTaskDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTask indicates an expected call of QueryTask.
func (mr *MockMigraterMockRecorder) QueryTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTask", reflect.TypeOf((*MockMigrater)(nil).QueryTask), arg0, arg1)
}

// ReclaimTask mocks base method.
func (m *MockMigrater) ReclaimTask(arg0 context.Context, arg1, arg2 string, arg3 []proto.VunitLocation, arg4 proto.VunitLocation, arg5 *client.AllocVunitInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReclaimTask", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReclaimTask indicates an expected call of ReclaimTask.
func (mr *MockMigraterMockRecorder) ReclaimTask(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReclaimTask", reflect.TypeOf((*MockMigrater)(nil).ReclaimTask), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RegisteDiskDropTaskLimitFunc mocks base method.
func (m *MockMigrater) RegisteDiskDropTaskLimitFunc(arg0, arg1 diskTaskLimitFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisteDiskDropTaskLimitFunc", arg0, arg1)
}

// RegisteDiskDropTaskLimitFunc indicates an expected call of RegisteDiskDropTaskLimitFunc.
func (mr *MockMigraterMockRecorder) RegisteDiskDropTaskLimitFunc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisteDiskDropTaskLimitFunc", reflect.TypeOf((*MockMigrater)(nil).RegisteDiskDropTaskLimitFunc), arg0, arg1)
}

// RenewalTask mocks base method.
func (m *MockMigrater) RenewalTask(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewalTask", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenewalTask indicates an expected call of RenewalTask.
func (mr *MockMigraterMockRecorder) RenewalTask(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewalTask", reflect.TypeOf((*MockMigrater)(nil).RenewalTask), arg0, arg1, arg2)
}

// ReportWorkerTaskStats mocks base method.
func (m *MockMigrater) ReportWorkerTaskStats(arg0 *scheduler.TaskReportArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportWorkerTaskStats", arg0)
}

// ReportWorkerTaskStats indicates an expected call of ReportWorkerTaskStats.
func (mr *MockMigraterMockRecorder) ReportWorkerTaskStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportWorkerTaskStats", reflect.TypeOf((*MockMigrater)(nil).ReportWorkerTaskStats), arg0)
}

// Run mocks base method.
func (m *MockMigrater) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockMigraterMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMigrater)(nil).Run))
}

// SetClearJunkTasksWhenLoadingFunc mocks base method.
func (m *MockMigrater) SetClearJunkTasksWhenLoadingFunc(arg0 clearJunkTasksFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClearJunkTasksWhenLoadingFunc", arg0)
}

// SetClearJunkTasksWhenLoadingFunc indicates an expected call of SetClearJunkTasksWhenLoadingFunc.
func (mr *MockMigraterMockRecorder) SetClearJunkTasksWhenLoadingFunc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClearJunkTasksWhenLoadingFunc", reflect.TypeOf((*MockMigrater)(nil).SetClearJunkTasksWhenLoadingFunc), arg0)
}

// SetLockFailHandleFunc mocks base method.
func (m *MockMigrater) SetLockFailHandleFunc(arg0 lockFailFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLockFailHandleFunc", arg0)
}

// SetLockFailHandleFunc indicates an expected call of SetLockFailHandleFunc.
func (mr *MockMigraterMockRecorder) SetLockFailHandleFunc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLockFailHandleFunc", reflect.TypeOf((*MockMigrater)(nil).SetLockFailHandleFunc), arg0)
}

// StatQueueTaskCnt mocks base method.
func (m *MockMigrater) StatQueueTaskCnt() (int, int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatQueueTaskCnt")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	return ret0, ret1, ret2
}

// StatQueueTaskCnt indicates an expected call of StatQueueTaskCnt.
func (mr *MockMigraterMockRecorder) StatQueueTaskCnt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatQueueTaskCnt", reflect.TypeOf((*MockMigrater)(nil).StatQueueTaskCnt))
}

// Stats mocks base method.
func (m *MockMigrater) Stats() scheduler.MigrateTasksStat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(scheduler.MigrateTasksStat)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockMigraterMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockMigrater)(nil).Stats))
}

// WaitEnable mocks base method.
func (m *MockMigrater) WaitEnable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WaitEnable")
}

// WaitEnable indicates an expected call of WaitEnable.
func (mr *MockMigraterMockRecorder) WaitEnable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitEnable", reflect.TypeOf((*MockMigrater)(nil).WaitEnable))
}

// MockVolumeInspector is a mock of IVolumeInspector interface.
type MockVolumeInspector struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeInspectorMockRecorder
}

// MockVolumeInspectorMockRecorder is the mock recorder for MockVolumeInspector.
type MockVolumeInspectorMockRecorder struct {
	mock *MockVolumeInspector
}

// NewMockVolumeInspector creates a new mock instance.
func NewMockVolumeInspector(ctrl *gomock.Controller) *MockVolumeInspector {
	mock := &MockVolumeInspector{ctrl: ctrl}
	mock.recorder = &MockVolumeInspectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeInspector) EXPECT() *MockVolumeInspectorMockRecorder {
	return m.recorder
}

// AcquireInspect mocks base method.
func (m *MockVolumeInspector) AcquireInspect(arg0 context.Context) (*proto.VolumeInspectTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireInspect", arg0)
	ret0, _ := ret[0].(*proto.VolumeInspectTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireInspect indicates an expected call of AcquireInspect.
func (mr *MockVolumeInspectorMockRecorder) AcquireInspect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireInspect", reflect.TypeOf((*MockVolumeInspector)(nil).AcquireInspect), arg0)
}

// Close mocks base method.
func (m *MockVolumeInspector) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockVolumeInspectorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockVolumeInspector)(nil).Close))
}

// CompleteInspect mocks base method.
func (m *MockVolumeInspector) CompleteInspect(arg0 context.Context, arg1 *proto.VolumeInspectRet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompleteInspect", arg0, arg1)
}

// CompleteInspect indicates an expected call of CompleteInspect.
func (mr *MockVolumeInspectorMockRecorder) CompleteInspect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteInspect", reflect.TypeOf((*MockVolumeInspector)(nil).CompleteInspect), arg0, arg1)
}

// Done mocks base method.
func (m *MockVolumeInspector) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockVolumeInspectorMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockVolumeInspector)(nil).Done))
}

// Enabled mocks base method.
func (m *MockVolumeInspector) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockVolumeInspectorMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockVolumeInspector)(nil).Enabled))
}

// GetTaskStats mocks base method.
func (m *MockVolumeInspector) GetTaskStats() ([20]int, [20]int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskStats")
	ret0, _ := ret[0].([20]int)
	ret1, _ := ret[1].([20]int)
	return ret0, ret1
}

// GetTaskStats indicates an expected call of GetTaskStats.
func (mr *MockVolumeInspectorMockRecorder) GetTaskStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskStats", reflect.TypeOf((*MockVolumeInspector)(nil).GetTaskStats))
}

// Run mocks base method.
func (m *MockVolumeInspector) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockVolumeInspectorMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockVolumeInspector)(nil).Run))
}

// MockClusterTopology is a mock of IClusterTopology interface.
type MockClusterTopology struct {
	ctrl     *gomock.Controller
	recorder *MockClusterTopologyMockRecorder
}

// MockClusterTopologyMockRecorder is the mock recorder for MockClusterTopology.
type MockClusterTopologyMockRecorder struct {
	mock *MockClusterTopology
}

// NewMockClusterTopology creates a new mock instance.
func NewMockClusterTopology(ctrl *gomock.Controller) *MockClusterTopology {
	mock := &MockClusterTopology{ctrl: ctrl}
	mock.recorder = &MockClusterTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterTopology) EXPECT() *MockClusterTopologyMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClusterTopology) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClusterTopologyMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClusterTopology)(nil).Close))
}

// Done mocks base method.
func (m *MockClusterTopology) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockClusterTopologyMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockClusterTopology)(nil).Done))
}

// GetIDCDisks mocks base method.
func (m *MockClusterTopology) GetIDCDisks(arg0 string) []*client.DiskInfoSimple {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDCDisks", arg0)
	ret0, _ := ret[0].([]*client.DiskInfoSimple)
	return ret0
}

// GetIDCDisks indicates an expected call of GetIDCDisks.
func (mr *MockClusterTopologyMockRecorder) GetIDCDisks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDCDisks", reflect.TypeOf((*MockClusterTopology)(nil).GetIDCDisks), arg0)
}

// GetIDCs mocks base method.
func (m *MockClusterTopology) GetIDCs() map[string]*IDC {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDCs")
	ret0, _ := ret[0].(map[string]*IDC)
	return ret0
}

// GetIDCs indicates an expected call of GetIDCs.
func (mr *MockClusterTopologyMockRecorder) GetIDCs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDCs", reflect.TypeOf((*MockClusterTopology)(nil).GetIDCs))
}

// GetVolume mocks base method.
func (m *MockClusterTopology) GetVolume(arg0 proto.Vid) (*client.VolumeInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0)
	ret0, _ := ret[0].(*client.VolumeInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume.
func (mr *MockClusterTopologyMockRecorder) GetVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockClusterTopology)(nil).GetVolume), arg0)
}

// IsBrokenDisk mocks base method.
func (m *MockClusterTopology) IsBrokenDisk(arg0 proto.DiskID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBrokenDisk", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBrokenDisk indicates an expected call of IsBrokenDisk.
func (mr *MockClusterTopologyMockRecorder) IsBrokenDisk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBrokenDisk", reflect.TypeOf((*MockClusterTopology)(nil).IsBrokenDisk), arg0)
}

// LoadVolumes mocks base method.
func (m *MockClusterTopology) LoadVolumes() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadVolumes")
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadVolumes indicates an expected call of LoadVolumes.
func (mr *MockClusterTopologyMockRecorder) LoadVolumes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadVolumes", reflect.TypeOf((*MockClusterTopology)(nil).LoadVolumes))
}

// MaxFreeChunksDisk mocks base method.
func (m *MockClusterTopology) MaxFreeChunksDisk(arg0 string) *client.DiskInfoSimple {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxFreeChunksDisk", arg0)
	ret0, _ := ret[0].(*client.DiskInfoSimple)
	return ret0
}

// MaxFreeChunksDisk indicates an expected call of MaxFreeChunksDisk.
func (mr *MockClusterTopologyMockRecorder) MaxFreeChunksDisk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxFreeChunksDisk", reflect.TypeOf((*MockClusterTopology)(nil).MaxFreeChunksDisk), arg0)
}

// UpdateVolume mocks base method.
func (m *MockClusterTopology) UpdateVolume(arg0 proto.Vid) (*client.VolumeInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVolume", arg0)
	ret0, _ := ret[0].(*client.VolumeInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVolume indicates an expected call of UpdateVolume.
func (mr *MockClusterTopologyMockRecorder) UpdateVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVolume", reflect.TypeOf((*MockClusterTopology)(nil).UpdateVolume), arg0)
}
