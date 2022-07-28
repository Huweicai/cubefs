// Copyright 2022 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package blobnode

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	api "github.com/cubefs/cubefs/blobstore/api/scheduler"
	"github.com/cubefs/cubefs/blobstore/common/proto"
)

func mockGenTasklet(bids []proto.BlobID) (ret []*ShardInfoSimple) {
	for _, bid := range bids {
		info := ShardInfoSimple{
			Bid:  bid,
			Size: 0,
		}
		ret = append(ret, &info)
	}
	return
}

var mocktasklets = []Tasklet{
	{bids: mockGenTasklet([]proto.BlobID{1})},
	{bids: mockGenTasklet([]proto.BlobID{2})},
	{bids: mockGenTasklet([]proto.BlobID{3})},
	{bids: mockGenTasklet([]proto.BlobID{4})},
	{bids: mockGenTasklet([]proto.BlobID{5})},
	{bids: mockGenTasklet([]proto.BlobID{6})},
	{bids: mockGenTasklet([]proto.BlobID{7})},
	{bids: mockGenTasklet([]proto.BlobID{8})},
	{bids: mockGenTasklet([]proto.BlobID{9})},
	{bids: mockGenTasklet([]proto.BlobID{10})},
	{bids: mockGenTasklet([]proto.BlobID{11})},
	{bids: mockGenTasklet([]proto.BlobID{12})},
}

type mockMigrateWorker struct {
	tasklet       []Tasklet
	taskletRetErr error
}

func NewmockMigrateWorker(task MigrateTaskEx) ITaskWorker {
	return &mockMigrateWorker{
		tasklet:       mocktasklets,
		taskletRetErr: nil,
	}
}

func (w *mockMigrateWorker) GenTasklets(ctx context.Context) ([]Tasklet, *WorkError) {
	time.Sleep(3600 * time.Second)
	return w.tasklet, nil
}

func (w *mockMigrateWorker) ExecTasklet(ctx context.Context, t Tasklet) *WorkError {
	return nil
}

func (w *mockMigrateWorker) Check(ctx context.Context) *WorkError {
	return nil
}

func (w *mockMigrateWorker) CancelArgs() (taskID string, taskType proto.TaskType, src []proto.VunitLocation, dest proto.VunitLocation) {
	return "test_mock_task", w.TaskType(), []proto.VunitLocation{}, proto.VunitLocation{}
}

func (w *mockMigrateWorker) CompleteArgs() (taskID string, taskType proto.TaskType, src []proto.VunitLocation, dest proto.VunitLocation) {
	return "test_mock_task", w.TaskType(), []proto.VunitLocation{}, proto.VunitLocation{}
}

func (w *mockMigrateWorker) ReclaimArgs() (taskID string, taskType proto.TaskType, src []proto.VunitLocation, dest proto.VunitLocation) {
	return "test_mock_task", w.TaskType(), []proto.VunitLocation{}, proto.VunitLocation{}
}

func (w *mockMigrateWorker) TaskType() proto.TaskType {
	return proto.TaskTypeBalance
}

func (w *mockMigrateWorker) GetBenchmarkBids() []*ShardInfoSimple {
	return nil
}

type mockWorkerFactory struct {
	newMigWorkerFn func(task MigrateTaskEx) ITaskWorker
}

func (mwf *mockWorkerFactory) NewMigrateWorker(task MigrateTaskEx) ITaskWorker {
	return mwf.newMigWorkerFn(task)
}

type mockScheCli struct {
	cancelRet   error
	completeRet error
	reclaimRet  error
	step        string
}

func (mock *mockScheCli) CancelTask(ctx context.Context, args *api.CancelTaskArgs) error {
	mock.step = "CancelOrReclaim"
	return mock.cancelRet
}

func (mock *mockScheCli) CompleteTask(ctx context.Context, args *api.CompleteTaskArgs) error {
	mock.step = "Complete"
	return mock.completeRet
}

func (mock *mockScheCli) ReclaimTask(ctx context.Context, args *api.ReclaimTaskArgs) error {
	mock.step = "CancelOrReclaim"
	return mock.reclaimRet
}

func (mock *mockScheCli) ReportTask(ctx context.Context, args *api.TaskReportArgs) (err error) {
	return nil
}

func initTestTaskRunnerMgr(t *testing.T, taskCnt int) *TaskRunnerMgr {
	cli := mockScheCli{
		cancelRet:   nil,
		completeRet: nil,
		reclaimRet:  nil,
	}
	wf := mockWorkerFactory{newMigWorkerFn: NewmockMigrateWorker}
	tm := NewTaskRunnerMgr(0, 2, 2, 2, 2, &cli, &wf)
	ctx := context.Background()
	for i := 0; i < taskCnt; i++ {
		taskID := fmt.Sprintf("repair_%d", i+1)
		task := MigrateTaskEx{
			taskInfo: &proto.MigrateTask{TaskID: taskID, TaskType: proto.TaskTypeDiskRepair},
		}
		err := tm.AddTask(ctx, task)
		require.NoError(t, err)
	}

	for i := 0; i < taskCnt; i++ {
		taskID := fmt.Sprintf("balance_%d", i+1)
		task := MigrateTaskEx{
			taskInfo: &proto.MigrateTask{TaskID: taskID, TaskType: proto.TaskTypeBalance},
		}
		err := tm.AddTask(ctx, task)
		require.NoError(t, err)
	}

	for i := 0; i < taskCnt; i++ {
		taskID := fmt.Sprintf("diskDrop_%d", i+1)
		task := MigrateTaskEx{
			taskInfo: &proto.MigrateTask{TaskID: taskID, TaskType: proto.TaskTypeDiskDrop},
		}
		err := tm.AddTask(ctx, task)
		require.NoError(t, err)
	}
	return tm
}

func TestTaskRunnerMgr(t *testing.T) {
	tm := initTestTaskRunnerMgr(t, 10)
	time.Sleep(200 * time.Millisecond)
	require.Equal(t, 10, len(tm.GetAliveTask(proto.TaskTypeDiskRepair)))
	require.Equal(t, 10, len(tm.GetAliveTask(proto.TaskTypeBalance)))
	require.Equal(t, 10, len(tm.GetAliveTask(proto.TaskTypeDiskDrop)))

	tm.StopAllAliveRunner()
	require.Equal(t, 0, len(tm.GetAliveTask(proto.TaskTypeDiskRepair)))
	require.Equal(t, 0, len(tm.GetAliveTask(proto.TaskTypeBalance)))
	require.Equal(t, 0, len(tm.GetAliveTask(proto.TaskTypeDiskDrop)))
}
