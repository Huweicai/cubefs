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

package null

import (
	"context"

	"github.com/cubefs/cubefs/blobstore/util/limit"
)

type nullLimit struct{}

// New returns limiter with nothing
func New() limit.Limiter {
	return nullLimit{}
}

func (l nullLimit) Running() int                                                      { return -1 }
func (l nullLimit) Acquire(keys ...interface{}) error                                 { return nil }
func (l nullLimit) Release(keys ...interface{})                                       { _ = struct{}{} }
func (l nullLimit) AcquireWithContext(ctx context.Context, keys ...interface{}) error { return nil }
