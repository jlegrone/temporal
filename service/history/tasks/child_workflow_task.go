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

package tasks

import (
	"time"

	"go.temporal.io/server/common/definition"
)

var _ Task = (*StartChildExecutionTask)(nil)

type (
	StartChildExecutionTask struct {
		definition.WorkflowKey
		VisibilityTimestamp time.Time
		TaskID              int64
		TargetNamespaceID   string
		TargetWorkflowID    string
		InitiatedID         int64
		Version             int64
	}
)

func (u *StartChildExecutionTask) GetKey() Key {
	return Key{
		FireTime: time.Unix(0, 0),
		TaskID:   u.TaskID,
	}
}

func (u *StartChildExecutionTask) GetVersion() int64 {
	return u.Version
}

func (u *StartChildExecutionTask) SetVersion(version int64) {
	u.Version = version
}

func (u *StartChildExecutionTask) GetTaskID() int64 {
	return u.TaskID
}

func (u *StartChildExecutionTask) SetTaskID(id int64) {
	u.TaskID = id
}

func (u *StartChildExecutionTask) GetVisibilityTime() time.Time {
	return u.VisibilityTimestamp
}

func (u *StartChildExecutionTask) SetVisibilityTime(timestamp time.Time) {
	u.VisibilityTimestamp = timestamp
}

func (u *StartChildExecutionTask) GetCategory() Category {
	return CategoryTransfer
}
