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

// ListInvoicesResponse struct for ListInvoicesResponse.
type ListInvoicesResponse struct {
	// invoices are sorted by period_start time.
	Invoices []Invoice `json:"invoices"`
}

// NewListInvoicesResponse instantiates a new ListInvoicesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvoicesResponse(invoices []Invoice) *ListInvoicesResponse {
	p := ListInvoicesResponse{}
	p.Invoices = invoices
	return &p
}

// NewListInvoicesResponseWithDefaults instantiates a new ListInvoicesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvoicesResponseWithDefaults() *ListInvoicesResponse {
	p := ListInvoicesResponse{}
	return &p
}

// GetInvoices returns the Invoices field value.
func (o *ListInvoicesResponse) GetInvoices() []Invoice {
	if o == nil {
		var ret []Invoice
		return ret
	}

	return o.Invoices
}

// SetInvoices sets field value.
func (o *ListInvoicesResponse) SetInvoices(v []Invoice) {
	o.Invoices = v
}

func (o ListInvoicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["invoices"] = o.Invoices
	}
	return json.Marshal(toSerialize)
}
