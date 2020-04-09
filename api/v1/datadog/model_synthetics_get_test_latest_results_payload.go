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

// SyntheticsGetTestLatestResultsPayload TODO.
type SyntheticsGetTestLatestResultsPayload struct {
	// TODO.
	FromTs float64 `json:"from_ts"`
	// TODO.
	ProbeDc *[]string `json:"probe_dc,omitempty"`
	// TODO.
	ToTs float64 `json:"to_ts"`
}

// NewSyntheticsGetTestLatestResultsPayload instantiates a new SyntheticsGetTestLatestResultsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsGetTestLatestResultsPayload(fromTs float64, toTs float64) *SyntheticsGetTestLatestResultsPayload {
	this := SyntheticsGetTestLatestResultsPayload{}
	this.FromTs = fromTs
	this.ToTs = toTs
	return &this
}

// NewSyntheticsGetTestLatestResultsPayloadWithDefaults instantiates a new SyntheticsGetTestLatestResultsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsGetTestLatestResultsPayloadWithDefaults() *SyntheticsGetTestLatestResultsPayload {
	this := SyntheticsGetTestLatestResultsPayload{}
	return &this
}

// GetFromTs returns the FromTs field value
func (o *SyntheticsGetTestLatestResultsPayload) GetFromTs() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGetTestLatestResultsPayload) GetFromTsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromTs, true
}

// SetFromTs sets field value
func (o *SyntheticsGetTestLatestResultsPayload) SetFromTs(v float64) {
	o.FromTs = v
}

// GetProbeDc returns the ProbeDc field value if set, zero value otherwise.
func (o *SyntheticsGetTestLatestResultsPayload) GetProbeDc() []string {
	if o == nil || o.ProbeDc == nil {
		var ret []string
		return ret
	}
	return *o.ProbeDc
}

// GetProbeDcOk returns a tuple with the ProbeDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsGetTestLatestResultsPayload) GetProbeDcOk() (*[]string, bool) {
	if o == nil || o.ProbeDc == nil {
		return nil, false
	}
	return o.ProbeDc, true
}

// HasProbeDc returns a boolean if a field has been set.
func (o *SyntheticsGetTestLatestResultsPayload) HasProbeDc() bool {
	if o != nil && o.ProbeDc != nil {
		return true
	}

	return false
}

// SetProbeDc gets a reference to the given []string and assigns it to the ProbeDc field.
func (o *SyntheticsGetTestLatestResultsPayload) SetProbeDc(v []string) {
	o.ProbeDc = &v
}

// GetToTs returns the ToTs field value
func (o *SyntheticsGetTestLatestResultsPayload) GetToTs() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value
// and a boolean to check if the value has been set.
func (o *SyntheticsGetTestLatestResultsPayload) GetToTsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToTs, true
}

// SetToTs sets field value
func (o *SyntheticsGetTestLatestResultsPayload) SetToTs(v float64) {
	o.ToTs = v
}

func (o SyntheticsGetTestLatestResultsPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from_ts"] = o.FromTs
	}
	if o.ProbeDc != nil {
		toSerialize["probe_dc"] = o.ProbeDc
	}
	if true {
		toSerialize["to_ts"] = o.ToTs
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsGetTestLatestResultsPayload struct {
	value *SyntheticsGetTestLatestResultsPayload
	isSet bool
}

func (v NullableSyntheticsGetTestLatestResultsPayload) Get() *SyntheticsGetTestLatestResultsPayload {
	return v.value
}

func (v *NullableSyntheticsGetTestLatestResultsPayload) Set(val *SyntheticsGetTestLatestResultsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsGetTestLatestResultsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsGetTestLatestResultsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsGetTestLatestResultsPayload(val *SyntheticsGetTestLatestResultsPayload) *NullableSyntheticsGetTestLatestResultsPayload {
	return &NullableSyntheticsGetTestLatestResultsPayload{value: val, isSet: true}
}

func (v NullableSyntheticsGetTestLatestResultsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsGetTestLatestResultsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
