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

// ApplicationKeyCreateData Object used to create an application key.
type ApplicationKeyCreateData struct {
	Attributes ApplicationKeyCreateAttributes `json:"attributes"`
	Type       ApplicationKeysType            `json:"type"`
}

// NewApplicationKeyCreateData instantiates a new ApplicationKeyCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyCreateData(attributes ApplicationKeyCreateAttributes, type_ ApplicationKeysType) *ApplicationKeyCreateData {
	this := ApplicationKeyCreateData{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewApplicationKeyCreateDataWithDefaults instantiates a new ApplicationKeyCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyCreateDataWithDefaults() *ApplicationKeyCreateData {
	this := ApplicationKeyCreateData{}
	var type_ ApplicationKeysType = APPLICATIONKEYSTYPE_APPLICATION_KEYS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *ApplicationKeyCreateData) GetAttributes() ApplicationKeyCreateAttributes {
	if o == nil {
		var ret ApplicationKeyCreateAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyCreateData) GetAttributesOk() (*ApplicationKeyCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ApplicationKeyCreateData) SetAttributes(v ApplicationKeyCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *ApplicationKeyCreateData) GetType() ApplicationKeysType {
	if o == nil {
		var ret ApplicationKeysType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationKeyCreateData) GetTypeOk() (*ApplicationKeysType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationKeyCreateData) SetType(v ApplicationKeysType) {
	o.Type = v
}

func (o ApplicationKeyCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationKeyCreateData struct {
	value *ApplicationKeyCreateData
	isSet bool
}

func (v NullableApplicationKeyCreateData) Get() *ApplicationKeyCreateData {
	return v.value
}

func (v *NullableApplicationKeyCreateData) Set(val *ApplicationKeyCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeyCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyCreateData(val *ApplicationKeyCreateData) *NullableApplicationKeyCreateData {
	return &NullableApplicationKeyCreateData{value: val, isSet: true}
}

func (v NullableApplicationKeyCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
