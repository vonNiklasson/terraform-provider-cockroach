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

// CMEKClusterInfo CMEKClusterInfo contains the status of CMEK across an entire cluster, including within each one its regions..
type CMEKClusterInfo struct {
	RegionInfos *[]CMEKRegionInfo `json:"region_infos,omitempty"`
	Status      *CMEKStatus       `json:"status,omitempty"`
}

// NewCMEKClusterInfo instantiates a new CMEKClusterInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMEKClusterInfo() *CMEKClusterInfo {
	p := CMEKClusterInfo{}
	return &p
}

// GetRegionInfos returns the RegionInfos field value if set, zero value otherwise.
func (o *CMEKClusterInfo) GetRegionInfos() []CMEKRegionInfo {
	if o == nil || o.RegionInfos == nil {
		var ret []CMEKRegionInfo
		return ret
	}
	return *o.RegionInfos
}

// SetRegionInfos gets a reference to the given []CMEKRegionInfo and assigns it to the RegionInfos field.
func (o *CMEKClusterInfo) SetRegionInfos(v []CMEKRegionInfo) {
	o.RegionInfos = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CMEKClusterInfo) GetStatus() CMEKStatus {
	if o == nil || o.Status == nil {
		var ret CMEKStatus
		return ret
	}
	return *o.Status
}

// SetStatus gets a reference to the given CMEKStatus and assigns it to the Status field.
func (o *CMEKClusterInfo) SetStatus(v CMEKStatus) {
	o.Status = &v
}

func (o CMEKClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RegionInfos != nil {
		toSerialize["region_infos"] = o.RegionInfos
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}
