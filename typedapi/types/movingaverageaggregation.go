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

// MovingAverageAggregation holds the union for the following types:
//     EwmaMovingAverageAggregation
//     HoltMovingAverageAggregation
//     HoltWintersMovingAverageAggregation
//     LinearMovingAverageAggregation
//     SimpleMovingAverageAggregation
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/pipeline.ts#L176-L182
type MovingAverageAggregation interface{}

// MovingAverageAggregationBuilder holds MovingAverageAggregation struct and provides a builder API.
type MovingAverageAggregationBuilder struct {
	v MovingAverageAggregation
}

// NewMovingAverageAggregation provides a builder for the MovingAverageAggregation struct.
func NewMovingAverageAggregationBuilder() *MovingAverageAggregationBuilder {
	return &MovingAverageAggregationBuilder{}
}

// Build finalize the chain and returns the MovingAverageAggregation struct
func (u *MovingAverageAggregationBuilder) Build() MovingAverageAggregation {
	return u.v
}

func (u *MovingAverageAggregationBuilder) EwmaMovingAverageAggregation(ewmamovingaverageaggregation *EwmaMovingAverageAggregationBuilder) *MovingAverageAggregationBuilder {
	v := ewmamovingaverageaggregation.Build()
	u.v = &v
	return u
}

func (u *MovingAverageAggregationBuilder) HoltMovingAverageAggregation(holtmovingaverageaggregation *HoltMovingAverageAggregationBuilder) *MovingAverageAggregationBuilder {
	v := holtmovingaverageaggregation.Build()
	u.v = &v
	return u
}

func (u *MovingAverageAggregationBuilder) HoltWintersMovingAverageAggregation(holtwintersmovingaverageaggregation *HoltWintersMovingAverageAggregationBuilder) *MovingAverageAggregationBuilder {
	v := holtwintersmovingaverageaggregation.Build()
	u.v = &v
	return u
}

func (u *MovingAverageAggregationBuilder) LinearMovingAverageAggregation(linearmovingaverageaggregation *LinearMovingAverageAggregationBuilder) *MovingAverageAggregationBuilder {
	v := linearmovingaverageaggregation.Build()
	u.v = &v
	return u
}

func (u *MovingAverageAggregationBuilder) SimpleMovingAverageAggregation(simplemovingaverageaggregation *SimpleMovingAverageAggregationBuilder) *MovingAverageAggregationBuilder {
	v := simplemovingaverageaggregation.Build()
	u.v = &v
	return u
}