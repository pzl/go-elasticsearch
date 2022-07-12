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

// DoubleTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L405-L408
type DoubleTermsBucket struct {
	Aggregations  map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount      int64                       `json:"doc_count"`
	DocCountError *int64                      `json:"doc_count_error,omitempty"`
	Key           float64                     `json:"key"`
	KeyAsString   *string                     `json:"key_as_string,omitempty"`
}

// DoubleTermsBucketBuilder holds DoubleTermsBucket struct and provides a builder API.
type DoubleTermsBucketBuilder struct {
	v *DoubleTermsBucket
}

// NewDoubleTermsBucket provides a builder for the DoubleTermsBucket struct.
func NewDoubleTermsBucketBuilder() *DoubleTermsBucketBuilder {
	r := DoubleTermsBucketBuilder{
		&DoubleTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DoubleTermsBucket struct
func (rb *DoubleTermsBucketBuilder) Build() DoubleTermsBucket {
	return *rb.v
}

func (rb *DoubleTermsBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *DoubleTermsBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *DoubleTermsBucketBuilder) DocCount(doccount int64) *DoubleTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *DoubleTermsBucketBuilder) DocCountError(doccounterror int64) *DoubleTermsBucketBuilder {
	rb.v.DocCountError = &doccounterror
	return rb
}

func (rb *DoubleTermsBucketBuilder) Key(key float64) *DoubleTermsBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *DoubleTermsBucketBuilder) KeyAsString(keyasstring string) *DoubleTermsBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}