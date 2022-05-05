// Copyright (C) 2019-2021 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"sync"

	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

// metaCache store the meta information for milvus collectipon/partition
// if performs "lazy" load and invalidate related cache if any error returns
type metaCache struct {
	// collections name=>collection meta
	collections sync.Map
	// partitions name=>partition meta
	partitions sync.Map
	// client instance to call describe collection methods, etc.
	client Client

	// workers performs fetch meta tasks
	workers sync.Map
}

// hasCollection returns collection
func (c *metaCache) hasCollection(ctx context.Context, name string) (bool, error) {
	raw, ok := c.collections.Load(name)
	coll := raw.(*entity.Collection)
	if !ok {
		var err error
		// meta not fetched, reload
		coll, err = c.fetchCollectionMeta(ctx, name)
		if err != nil {
			return false, err
		}
	}
	return coll.ID > 0, nil
}

// invalidateCollection delete the cached meta for collection if any
func (c *metaCache) invalidateCollection(name string) {
	c.collections.Delete(name)
}

// fetchCollectionMeta uses the client to call DescribeCollection
func (c *metaCache) fetchCollectionMeta(ctx context.Context, name string) (coll *entity.Collection, err error) {
	w, _ := c.workers.LoadOrStore(name, &metaWorker{})
	w.(*metaWorker).Do(func() {
		coll, err = c.client.DescribeCollection(ctx, name)
		if err == nil {
			c.collections.Store(name, coll)
		}
	})
	return
}

type metaWorker struct {
	sync.Once
}
