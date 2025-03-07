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

// UpdateDatabaseRequest struct for UpdateDatabaseRequest.
type UpdateDatabaseRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

// NewUpdateDatabaseRequest instantiates a new UpdateDatabaseRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDatabaseRequest(name string, newName string) *UpdateDatabaseRequest {
	p := UpdateDatabaseRequest{}
	p.Name = name
	p.NewName = newName
	return &p
}

// NewUpdateDatabaseRequestWithDefaults instantiates a new UpdateDatabaseRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDatabaseRequestWithDefaults() *UpdateDatabaseRequest {
	p := UpdateDatabaseRequest{}
	return &p
}

// GetName returns the Name field value.
func (o *UpdateDatabaseRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value.
func (o *UpdateDatabaseRequest) SetName(v string) {
	o.Name = v
}

// GetNewName returns the NewName field value.
func (o *UpdateDatabaseRequest) GetNewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewName
}

// SetNewName sets field value.
func (o *UpdateDatabaseRequest) SetNewName(v string) {
	o.NewName = v
}

func (o UpdateDatabaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["new_name"] = o.NewName
	}
	return json.Marshal(toSerialize)
}
