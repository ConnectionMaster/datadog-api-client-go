/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// NotebookUpdateRequest The description of a notebook update request.
type NotebookUpdateRequest struct {
	Data NotebookUpdateData `json:"data"`
}

// NewNotebookUpdateRequest instantiates a new NotebookUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookUpdateRequest(data NotebookUpdateData) *NotebookUpdateRequest {
	this := NotebookUpdateRequest{}
	this.Data = data
	return &this
}

// NewNotebookUpdateRequestWithDefaults instantiates a new NotebookUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookUpdateRequestWithDefaults() *NotebookUpdateRequest {
	this := NotebookUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *NotebookUpdateRequest) GetData() NotebookUpdateData {
	if o == nil {
		var ret NotebookUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NotebookUpdateRequest) GetDataOk() (*NotebookUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NotebookUpdateRequest) SetData(v NotebookUpdateData) {
	o.Data = v
}

func (o NotebookUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

func (o *NotebookUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Data *NotebookUpdateData `json:"data"`
	}{}
	all := struct {
		Data NotebookUpdateData `json:"data"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Data == nil {
		return fmt.Errorf("Required field data missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Data = all.Data
	return nil
}

type NullableNotebookUpdateRequest struct {
	value *NotebookUpdateRequest
	isSet bool
}

func (v NullableNotebookUpdateRequest) Get() *NotebookUpdateRequest {
	return v.value
}

func (v *NullableNotebookUpdateRequest) Set(val *NotebookUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookUpdateRequest(val *NotebookUpdateRequest) *NullableNotebookUpdateRequest {
	return &NullableNotebookUpdateRequest{value: val, isSet: true}
}

func (v NullableNotebookUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
