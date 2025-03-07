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

// ServerlessClusterConfig struct for ServerlessClusterConfig.
type ServerlessClusterConfig struct {
	// routing_id is used to identify the cluster in a connection string.
	RoutingId string `json:"routing_id"`
	// spend_limit is the maximum monthly charge for a cluster, in US cents. We recommend using usage_limits instead, since spend_limit will be deprecated in the future.
	SpendLimit  *int32       `json:"spend_limit,omitempty"`
	UsageLimits *UsageLimits `json:"usage_limits,omitempty"`
}

// NewServerlessClusterConfig instantiates a new ServerlessClusterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessClusterConfig(routingId string) *ServerlessClusterConfig {
	p := ServerlessClusterConfig{}
	p.RoutingId = routingId
	return &p
}

// NewServerlessClusterConfigWithDefaults instantiates a new ServerlessClusterConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessClusterConfigWithDefaults() *ServerlessClusterConfig {
	p := ServerlessClusterConfig{}
	return &p
}

// GetRoutingId returns the RoutingId field value.
func (o *ServerlessClusterConfig) GetRoutingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingId
}

// SetRoutingId sets field value.
func (o *ServerlessClusterConfig) SetRoutingId(v string) {
	o.RoutingId = v
}

// GetSpendLimit returns the SpendLimit field value if set, zero value otherwise.
func (o *ServerlessClusterConfig) GetSpendLimit() int32 {
	if o == nil || o.SpendLimit == nil {
		var ret int32
		return ret
	}
	return *o.SpendLimit
}

// SetSpendLimit gets a reference to the given int32 and assigns it to the SpendLimit field.
func (o *ServerlessClusterConfig) SetSpendLimit(v int32) {
	o.SpendLimit = &v
}

// GetUsageLimits returns the UsageLimits field value if set, zero value otherwise.
func (o *ServerlessClusterConfig) GetUsageLimits() UsageLimits {
	if o == nil || o.UsageLimits == nil {
		var ret UsageLimits
		return ret
	}
	return *o.UsageLimits
}

// SetUsageLimits gets a reference to the given UsageLimits and assigns it to the UsageLimits field.
func (o *ServerlessClusterConfig) SetUsageLimits(v UsageLimits) {
	o.UsageLimits = &v
}

func (o ServerlessClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["routing_id"] = o.RoutingId
	}
	if o.SpendLimit != nil {
		toSerialize["spend_limit"] = o.SpendLimit
	}
	if o.UsageLimits != nil {
		toSerialize["usage_limits"] = o.UsageLimits
	}
	return json.Marshal(toSerialize)
}
