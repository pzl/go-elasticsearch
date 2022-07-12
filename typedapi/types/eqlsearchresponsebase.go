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

// EqlSearchResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/eql/_types/EqlSearchResponseBase.ts#L25-L50
type EqlSearchResponseBase struct {
	// Hits Contains matching events and sequences. Also contains related metadata.
	Hits EqlHits `json:"hits"`
	// Id Identifier for the search.
	Id *Id `json:"id,omitempty"`
	// IsPartial If true, the response does not contain complete search results.
	IsPartial *bool `json:"is_partial,omitempty"`
	// IsRunning If true, the search request is still executing.
	IsRunning *bool `json:"is_running,omitempty"`
	// TimedOut If true, the request timed out before completion.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took Milliseconds it took Elasticsearch to execute the request.
	Took *DurationValueUnitMillis `json:"took,omitempty"`
}

// EqlSearchResponseBaseBuilder holds EqlSearchResponseBase struct and provides a builder API.
type EqlSearchResponseBaseBuilder struct {
	v *EqlSearchResponseBase
}

// NewEqlSearchResponseBase provides a builder for the EqlSearchResponseBase struct.
func NewEqlSearchResponseBaseBuilder() *EqlSearchResponseBaseBuilder {
	r := EqlSearchResponseBaseBuilder{
		&EqlSearchResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the EqlSearchResponseBase struct
func (rb *EqlSearchResponseBaseBuilder) Build() EqlSearchResponseBase {
	return *rb.v
}

// Hits Contains matching events and sequences. Also contains related metadata.

func (rb *EqlSearchResponseBaseBuilder) Hits(hits *EqlHitsBuilder) *EqlSearchResponseBaseBuilder {
	v := hits.Build()
	rb.v.Hits = v
	return rb
}

// Id Identifier for the search.

func (rb *EqlSearchResponseBaseBuilder) Id(id Id) *EqlSearchResponseBaseBuilder {
	rb.v.Id = &id
	return rb
}

// IsPartial If true, the response does not contain complete search results.

func (rb *EqlSearchResponseBaseBuilder) IsPartial(ispartial bool) *EqlSearchResponseBaseBuilder {
	rb.v.IsPartial = &ispartial
	return rb
}

// IsRunning If true, the search request is still executing.

func (rb *EqlSearchResponseBaseBuilder) IsRunning(isrunning bool) *EqlSearchResponseBaseBuilder {
	rb.v.IsRunning = &isrunning
	return rb
}

// TimedOut If true, the request timed out before completion.

func (rb *EqlSearchResponseBaseBuilder) TimedOut(timedout bool) *EqlSearchResponseBaseBuilder {
	rb.v.TimedOut = &timedout
	return rb
}

// Took Milliseconds it took Elasticsearch to execute the request.

func (rb *EqlSearchResponseBaseBuilder) Took(took *DurationValueUnitMillisBuilder) *EqlSearchResponseBaseBuilder {
	v := took.Build()
	rb.v.Took = &v
	return rb
}
