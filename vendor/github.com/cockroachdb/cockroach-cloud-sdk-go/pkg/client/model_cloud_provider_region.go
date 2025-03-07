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

// CloudProviderRegion struct for CloudProviderRegion.
type CloudProviderRegion struct {
	// Distance in miles, based on client IP address.
	Distance   float32           `json:"distance"`
	Location   string            `json:"location"`
	Name       string            `json:"name"`
	Provider   CloudProviderType `json:"provider"`
	Serverless bool              `json:"serverless"`
}

// NewCloudProviderRegion instantiates a new CloudProviderRegion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderRegion(distance float32, location string, name string, provider CloudProviderType, serverless bool) *CloudProviderRegion {
	p := CloudProviderRegion{}
	p.Distance = distance
	p.Location = location
	p.Name = name
	p.Provider = provider
	p.Serverless = serverless
	return &p
}

// NewCloudProviderRegionWithDefaults instantiates a new CloudProviderRegion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderRegionWithDefaults() *CloudProviderRegion {
	p := CloudProviderRegion{}
	return &p
}

// GetDistance returns the Distance field value.
func (o *CloudProviderRegion) GetDistance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Distance
}

// SetDistance sets field value.
func (o *CloudProviderRegion) SetDistance(v float32) {
	o.Distance = v
}

// GetLocation returns the Location field value.
func (o *CloudProviderRegion) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// SetLocation sets field value.
func (o *CloudProviderRegion) SetLocation(v string) {
	o.Location = v
}

// GetName returns the Name field value.
func (o *CloudProviderRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *CloudProviderRegion) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value.
func (o *CloudProviderRegion) GetProvider() CloudProviderType {
	if o == nil {
		var ret CloudProviderType
		return ret
	}

	return o.Provider
}

// SetProvider sets field value.
func (o *CloudProviderRegion) SetProvider(v CloudProviderType) {
	o.Provider = v
}

// GetServerless returns the Serverless field value.
func (o *CloudProviderRegion) GetServerless() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Serverless
}

// SetServerless sets field value.
func (o *CloudProviderRegion) SetServerless(v bool) {
	o.Serverless = v
}

func (o CloudProviderRegion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distance"] = o.Distance
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["serverless"] = o.Serverless
	}
	return json.Marshal(toSerialize)
}
