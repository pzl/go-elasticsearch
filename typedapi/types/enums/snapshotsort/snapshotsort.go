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

// Package snapshotsort
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/snapshot/_types/SnapshotInfo.ts#L67-L78
package snapshotsort

import "strings"

type SnapshotSort struct {
	name string
}

var (
	Starttime = SnapshotSort{"start_time"}

	Duration = SnapshotSort{"duration"}

	Name = SnapshotSort{"name"}

	Indexcount = SnapshotSort{"index_count"}

	Repository = SnapshotSort{"repository"}

	Shardcount = SnapshotSort{"shard_count"}

	Failedshardcount = SnapshotSort{"failed_shard_count"}
)

func (s SnapshotSort) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SnapshotSort) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "start_time":
		*s = Starttime
	case "duration":
		*s = Duration
	case "name":
		*s = Name
	case "index_count":
		*s = Indexcount
	case "repository":
		*s = Repository
	case "shard_count":
		*s = Shardcount
	case "failed_shard_count":
		*s = Failedshardcount
	default:
		*s = SnapshotSort{string(text)}
	}

	return nil
}

func (s SnapshotSort) String() string {
	return s.name
}
