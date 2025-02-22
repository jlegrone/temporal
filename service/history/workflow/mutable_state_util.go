// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"time"

	persistencespb "go.temporal.io/server/api/persistence/v1"
	"go.temporal.io/server/common/definition"
	"go.temporal.io/server/service/history/tasks"
)

// NOTE: do not use make(type, len(input))
// since this will assume initial length being len(inputs)
// always use make(type, 0, len(input))
func convertSyncActivityInfos(
	now time.Time,
	workflowKey definition.WorkflowKey,
	activityInfos map[int64]*persistencespb.ActivityInfo,
	inputs map[int64]struct{},
) []tasks.Task {
	outputs := make([]tasks.Task, 0, len(inputs))
	for item := range inputs {
		activityInfo, ok := activityInfos[item]
		if ok {
			outputs = append(outputs, &tasks.SyncActivityTask{
				WorkflowKey:         workflowKey,
				Version:             activityInfo.Version,
				ScheduledID:         activityInfo.ScheduleId,
				VisibilityTimestamp: now,
			})
		}
	}
	return outputs
}

// TODO: can we deprecate this method and
// let task generator correctly set task version and
// visibility timestamp?
func setTaskInfo(
	version int64,
	timestamp time.Time,
	insertTasks map[tasks.Category][]tasks.Task,
) {
	// set the task version,
	// as well as the Timestamp if not scheduled task
	for category, tasksByCategory := range insertTasks {
		if category == tasks.CategoryReplication {
			continue
		}

		for _, task := range tasksByCategory {
			task.SetVersion(version)
			if category.Type() == tasks.CategoryTypeImmediate {
				task.SetVisibilityTime(timestamp)
			}
		}
	}
}
