// Copyright 2023 The Cockroach Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
// CockroachDB Cloud API
// API version: 2023-04-10

package client

import (
	"encoding/json"
)

// CMEKClusterSpecification struct for CMEKClusterSpecification.
type CMEKClusterSpecification struct {
	RegionSpecs []CMEKRegionSpecification `json:"region_specs"`
}

// NewCMEKClusterSpecification instantiates a new CMEKClusterSpecification object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKClusterSpecification(regionSpecs []CMEKRegionSpecification) *CMEKClusterSpecification {
	p := CMEKClusterSpecification{}
	p.RegionSpecs = regionSpecs
	return &p
}

// NewCMEKClusterSpecificationWithDefaults instantiates a new CMEKClusterSpecification object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCMEKClusterSpecificationWithDefaults() *CMEKClusterSpecification {
	p := CMEKClusterSpecification{}
	return &p
}

// GetRegionSpecs returns the RegionSpecs field value.
func (o *CMEKClusterSpecification) GetRegionSpecs() []CMEKRegionSpecification {
	if o == nil {
		var ret []CMEKRegionSpecification
		return ret
	}

	return o.RegionSpecs
}

// SetRegionSpecs sets field value.
func (o *CMEKClusterSpecification) SetRegionSpecs(v []CMEKRegionSpecification) {
	o.RegionSpecs = v
}

func (o CMEKClusterSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["region_specs"] = o.RegionSpecs
	}
	return json.Marshal(toSerialize)
}
