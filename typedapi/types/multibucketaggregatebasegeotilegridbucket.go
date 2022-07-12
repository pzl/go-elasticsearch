// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// MultiBucketAggregateBaseGeoTileGridBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L314-L316
type MultiBucketAggregateBaseGeoTileGridBucket struct {
	Buckets BucketsGeoTileGridBucket `json:"buckets"`
	Meta    *Metadata                `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseGeoTileGridBucketBuilder holds MultiBucketAggregateBaseGeoTileGridBucket struct and provides a builder API.
type MultiBucketAggregateBaseGeoTileGridBucketBuilder struct {
	v *MultiBucketAggregateBaseGeoTileGridBucket
}

// NewMultiBucketAggregateBaseGeoTileGridBucket provides a builder for the MultiBucketAggregateBaseGeoTileGridBucket struct.
func NewMultiBucketAggregateBaseGeoTileGridBucketBuilder() *MultiBucketAggregateBaseGeoTileGridBucketBuilder {
	r := MultiBucketAggregateBaseGeoTileGridBucketBuilder{
		&MultiBucketAggregateBaseGeoTileGridBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseGeoTileGridBucket struct
func (rb *MultiBucketAggregateBaseGeoTileGridBucketBuilder) Build() MultiBucketAggregateBaseGeoTileGridBucket {
	return *rb.v
}

func (rb *MultiBucketAggregateBaseGeoTileGridBucketBuilder) Buckets(buckets *BucketsGeoTileGridBucketBuilder) *MultiBucketAggregateBaseGeoTileGridBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiBucketAggregateBaseGeoTileGridBucketBuilder) Meta(meta *MetadataBuilder) *MultiBucketAggregateBaseGeoTileGridBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
