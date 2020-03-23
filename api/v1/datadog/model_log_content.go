/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// LogContent struct for LogContent
type LogContent struct {
	// TODO.
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Name of the machine from where the logs are being sent.
	Host *string `json:"host,omitempty"`
	// TODO.
	Message *string `json:"message,omitempty"`
	// TODO.
	Service *string `json:"service,omitempty"`
	// TODO.
	Tags *[]interface{} `json:"tags,omitempty"`
	// TODO.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewLogContent instantiates a new LogContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogContent() *LogContent {
	this := LogContent{}
	return &this
}

// NewLogContentWithDefaults instantiates a new LogContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogContentWithDefaults() *LogContent {
	this := LogContent{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *LogContent) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogContent) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *LogContent) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *LogContent) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *LogContent) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogContent) GetHostOk() (string, bool) {
	if o == nil || o.Host == nil {
		var ret string
		return ret, false
	}
	return *o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *LogContent) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *LogContent) SetHost(v string) {
	o.Host = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogContent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogContent) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogContent) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogContent) SetMessage(v string) {
	o.Message = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *LogContent) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogContent) GetServiceOk() (string, bool) {
	if o == nil || o.Service == nil {
		var ret string
		return ret, false
	}
	return *o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *LogContent) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *LogContent) SetService(v string) {
	o.Service = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LogContent) GetTags() []interface{} {
	if o == nil || o.Tags == nil {
		var ret []interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogContent) GetTagsOk() ([]interface{}, bool) {
	if o == nil || o.Tags == nil {
		var ret []interface{}
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LogContent) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []interface{} and assigns it to the Tags field.
func (o *LogContent) SetTags(v []interface{}) {
	o.Tags = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogContent) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LogContent) GetTimestampOk() (time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogContent) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *LogContent) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o LogContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableLogContent struct {
	value *LogContent
	isSet bool
}

func (v NullableLogContent) Get() *LogContent {
	return v.value
}

func (v NullableLogContent) Set(val *LogContent) {
	v.value = val
	v.isSet = true
}

func (v NullableLogContent) IsSet() bool {
	return v.isSet
}

func (v NullableLogContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogContent(val *LogContent) *NullableLogContent {
	return &NullableLogContent{value: val, isSet: true}
}

func (v NullableLogContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}