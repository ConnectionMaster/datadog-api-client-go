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

// SyntheticsBrowserTestResultShortResult TODO.
type SyntheticsBrowserTestResultShortResult struct {
	Device *SyntheticsDevice `json:"device,omitempty"`
	// TODO.
	Duration *float64 `json:"duration,omitempty"`
	// TODO.
	ErrorCount *int64 `json:"errorCount,omitempty"`
	// TODO.
	StepCountCompleted *int64 `json:"stepCountCompleted,omitempty"`
	// TODO.
	StepCountTotal *int64 `json:"stepCountTotal,omitempty"`
}

// NewSyntheticsBrowserTestResultShortResult instantiates a new SyntheticsBrowserTestResultShortResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBrowserTestResultShortResult() *SyntheticsBrowserTestResultShortResult {
	this := SyntheticsBrowserTestResultShortResult{}
	return &this
}

// NewSyntheticsBrowserTestResultShortResultWithDefaults instantiates a new SyntheticsBrowserTestResultShortResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBrowserTestResultShortResultWithDefaults() *SyntheticsBrowserTestResultShortResult {
	this := SyntheticsBrowserTestResultShortResult{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultShortResult) GetDevice() SyntheticsDevice {
	if o == nil || o.Device == nil {
		var ret SyntheticsDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultShortResult) GetDeviceOk() (*SyntheticsDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultShortResult) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given SyntheticsDevice and assigns it to the Device field.
func (o *SyntheticsBrowserTestResultShortResult) SetDevice(v SyntheticsDevice) {
	o.Device = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultShortResult) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultShortResult) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultShortResult) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *SyntheticsBrowserTestResultShortResult) SetDuration(v float64) {
	o.Duration = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultShortResult) GetErrorCount() int64 {
	if o == nil || o.ErrorCount == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultShortResult) GetErrorCountOk() (*int64, bool) {
	if o == nil || o.ErrorCount == nil {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultShortResult) HasErrorCount() bool {
	if o != nil && o.ErrorCount != nil {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int64 and assigns it to the ErrorCount field.
func (o *SyntheticsBrowserTestResultShortResult) SetErrorCount(v int64) {
	o.ErrorCount = &v
}

// GetStepCountCompleted returns the StepCountCompleted field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultShortResult) GetStepCountCompleted() int64 {
	if o == nil || o.StepCountCompleted == nil {
		var ret int64
		return ret
	}
	return *o.StepCountCompleted
}

// GetStepCountCompletedOk returns a tuple with the StepCountCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultShortResult) GetStepCountCompletedOk() (*int64, bool) {
	if o == nil || o.StepCountCompleted == nil {
		return nil, false
	}
	return o.StepCountCompleted, true
}

// HasStepCountCompleted returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultShortResult) HasStepCountCompleted() bool {
	if o != nil && o.StepCountCompleted != nil {
		return true
	}

	return false
}

// SetStepCountCompleted gets a reference to the given int64 and assigns it to the StepCountCompleted field.
func (o *SyntheticsBrowserTestResultShortResult) SetStepCountCompleted(v int64) {
	o.StepCountCompleted = &v
}

// GetStepCountTotal returns the StepCountTotal field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestResultShortResult) GetStepCountTotal() int64 {
	if o == nil || o.StepCountTotal == nil {
		var ret int64
		return ret
	}
	return *o.StepCountTotal
}

// GetStepCountTotalOk returns a tuple with the StepCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestResultShortResult) GetStepCountTotalOk() (*int64, bool) {
	if o == nil || o.StepCountTotal == nil {
		return nil, false
	}
	return o.StepCountTotal, true
}

// HasStepCountTotal returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestResultShortResult) HasStepCountTotal() bool {
	if o != nil && o.StepCountTotal != nil {
		return true
	}

	return false
}

// SetStepCountTotal gets a reference to the given int64 and assigns it to the StepCountTotal field.
func (o *SyntheticsBrowserTestResultShortResult) SetStepCountTotal(v int64) {
	o.StepCountTotal = &v
}

func (o SyntheticsBrowserTestResultShortResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ErrorCount != nil {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if o.StepCountCompleted != nil {
		toSerialize["stepCountCompleted"] = o.StepCountCompleted
	}
	if o.StepCountTotal != nil {
		toSerialize["stepCountTotal"] = o.StepCountTotal
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsBrowserTestResultShortResult struct {
	value *SyntheticsBrowserTestResultShortResult
	isSet bool
}

func (v NullableSyntheticsBrowserTestResultShortResult) Get() *SyntheticsBrowserTestResultShortResult {
	return v.value
}

func (v *NullableSyntheticsBrowserTestResultShortResult) Set(val *SyntheticsBrowserTestResultShortResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserTestResultShortResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserTestResultShortResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserTestResultShortResult(val *SyntheticsBrowserTestResultShortResult) *NullableSyntheticsBrowserTestResultShortResult {
	return &NullableSyntheticsBrowserTestResultShortResult{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserTestResultShortResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserTestResultShortResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
