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
	"fmt"
)

// NetworkVisibilityType the model 'NetworkVisibilityType'.
type NetworkVisibilityType string

// List of NetworkVisibility.Type.
const (
	NETWORKVISIBILITYTYPE_PUBLIC  NetworkVisibilityType = "PUBLIC"
	NETWORKVISIBILITYTYPE_PRIVATE NetworkVisibilityType = "PRIVATE"
)

// All allowed values of NetworkVisibilityType enum.
var AllowedNetworkVisibilityTypeEnumValues = []NetworkVisibilityType{
	"PUBLIC",
	"PRIVATE",
}

func (v *NetworkVisibilityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkVisibilityType(value)
	for _, existing := range AllowedNetworkVisibilityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkVisibilityType", value)
}

// NewNetworkVisibilityTypeFromValue returns a pointer to a valid NetworkVisibilityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNetworkVisibilityTypeFromValue(v string) (*NetworkVisibilityType, error) {
	ev := NetworkVisibilityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkVisibilityType: valid values are %v", v, AllowedNetworkVisibilityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NetworkVisibilityType) IsValid() bool {
	for _, existing := range AllowedNetworkVisibilityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkVisibility.Type value.
func (v NetworkVisibilityType) Ptr() *NetworkVisibilityType {
	return &v
}
