/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// UsageHostsResponse struct for UsageHostsResponse
type UsageHostsResponse struct {
	Usage *[]UsageHostHour `json:"usage,omitempty"`
}

// NewUsageHostsResponse instantiates a new UsageHostsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageHostsResponse() *UsageHostsResponse {
	this := UsageHostsResponse{}
	return &this
}

// NewUsageHostsResponseWithDefaults instantiates a new UsageHostsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageHostsResponseWithDefaults() *UsageHostsResponse {
	this := UsageHostsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageHostsResponse) GetUsage() []UsageHostHour {
	if o == nil || o.Usage == nil {
		var ret []UsageHostHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsageHostsResponse) GetUsageOk() ([]UsageHostHour, bool) {
	if o == nil || o.Usage == nil {
		var ret []UsageHostHour
		return ret, false
	}
	return *o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageHostsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageHostHour and assigns it to the Usage field.
func (o *UsageHostsResponse) SetUsage(v []UsageHostHour) {
	o.Usage = &v
}

func (o UsageHostsResponse) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageHostsResponse struct {
	value *UsageHostsResponse
	isSet bool
}

func (v NullableUsageHostsResponse) Get() *UsageHostsResponse {
	return v.value
}

func (v NullableUsageHostsResponse) Set(val *UsageHostsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageHostsResponse) IsSet() bool {
	return v.isSet
}

func (v NullableUsageHostsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageHostsResponse(val *UsageHostsResponse) *NullableUsageHostsResponse {
	return &NullableUsageHostsResponse{value: val, isSet: true}
}

func (v NullableUsageHostsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageHostsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
