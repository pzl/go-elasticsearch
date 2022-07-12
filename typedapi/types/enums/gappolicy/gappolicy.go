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

// Package gappolicy
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/pipeline.ts#L52-L55
package gappolicy

import "strings"

type GapPolicy struct {
	name string
}

var (
	Skip = GapPolicy{"skip"}

	Insertzeros = GapPolicy{"insert_zeros"}
)

func (g GapPolicy) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GapPolicy) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "skip":
		*g = Skip
	case "insert_zeros":
		*g = Insertzeros
	default:
		*g = GapPolicy{string(text)}
	}

	return nil
}

func (g GapPolicy) String() string {
	return g.name
}
