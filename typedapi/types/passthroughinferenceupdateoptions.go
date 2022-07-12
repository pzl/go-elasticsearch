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

// PassThroughInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/ml/_types/inference.ts#L344-L349
type PassThroughInferenceUpdateOptions struct {
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// PassThroughInferenceUpdateOptionsBuilder holds PassThroughInferenceUpdateOptions struct and provides a builder API.
type PassThroughInferenceUpdateOptionsBuilder struct {
	v *PassThroughInferenceUpdateOptions
}

// NewPassThroughInferenceUpdateOptions provides a builder for the PassThroughInferenceUpdateOptions struct.
func NewPassThroughInferenceUpdateOptionsBuilder() *PassThroughInferenceUpdateOptionsBuilder {
	r := PassThroughInferenceUpdateOptionsBuilder{
		&PassThroughInferenceUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the PassThroughInferenceUpdateOptions struct
func (rb *PassThroughInferenceUpdateOptionsBuilder) Build() PassThroughInferenceUpdateOptions {
	return *rb.v
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *PassThroughInferenceUpdateOptionsBuilder) ResultsField(resultsfield string) *PassThroughInferenceUpdateOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *PassThroughInferenceUpdateOptionsBuilder) Tokenization(tokenization *NlpTokenizationUpdateOptionsBuilder) *PassThroughInferenceUpdateOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
