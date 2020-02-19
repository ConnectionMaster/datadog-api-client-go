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

// ServiceLevelObjectivesBulkDeleted The bulk partial delete service level objective object endpoint response. This endpoint operates on multiple service level objective objects, so it may be partially successful. In such cases, the \"data\" and \"error\" fields in this response indicate which deletions succeeded and failed.
type ServiceLevelObjectivesBulkDeleted struct {
	Data   ServiceLevelObjectivesBulkDeletedData     `json:"data"`
	Errors []ServiceLevelObjectivesBulkDeletedErrors `json:"errors"`
}

// NewServiceLevelObjectivesBulkDeleted instantiates a new ServiceLevelObjectivesBulkDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectivesBulkDeleted(data ServiceLevelObjectivesBulkDeletedData, errors []ServiceLevelObjectivesBulkDeletedErrors) *ServiceLevelObjectivesBulkDeleted {
	this := ServiceLevelObjectivesBulkDeleted{}
	this.Data = data
	this.Errors = errors
	return &this
}

// NewServiceLevelObjectivesBulkDeletedWithDefaults instantiates a new ServiceLevelObjectivesBulkDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectivesBulkDeletedWithDefaults() *ServiceLevelObjectivesBulkDeleted {
	this := ServiceLevelObjectivesBulkDeleted{}
	return &this
}

// GetData returns the Data field value
func (o *ServiceLevelObjectivesBulkDeleted) GetData() ServiceLevelObjectivesBulkDeletedData {
	if o == nil {
		var ret ServiceLevelObjectivesBulkDeletedData
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *ServiceLevelObjectivesBulkDeleted) SetData(v ServiceLevelObjectivesBulkDeletedData) {
	o.Data = v
}

// GetErrors returns the Errors field value
func (o *ServiceLevelObjectivesBulkDeleted) GetErrors() []ServiceLevelObjectivesBulkDeletedErrors {
	if o == nil {
		var ret []ServiceLevelObjectivesBulkDeletedErrors
		return ret
	}

	return o.Errors
}

// SetErrors sets field value
func (o *ServiceLevelObjectivesBulkDeleted) SetErrors(v []ServiceLevelObjectivesBulkDeletedErrors) {
	o.Errors = v
}

func (o ServiceLevelObjectivesBulkDeleted) MarshalJSON() ([]byte, error) {
	//TODO: serialize parents?
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableServiceLevelObjectivesBulkDeleted struct {
	value *ServiceLevelObjectivesBulkDeleted
	isSet bool
}

func (v NullableServiceLevelObjectivesBulkDeleted) Get() *ServiceLevelObjectivesBulkDeleted {
	return v.value
}

func (v NullableServiceLevelObjectivesBulkDeleted) Set(val *ServiceLevelObjectivesBulkDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelObjectivesBulkDeleted) IsSet() bool {
	return v.isSet
}

func (v NullableServiceLevelObjectivesBulkDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelObjectivesBulkDeleted(val *ServiceLevelObjectivesBulkDeleted) *NullableServiceLevelObjectivesBulkDeleted {
	return &NullableServiceLevelObjectivesBulkDeleted{value: val, isSet: true}
}

func (v NullableServiceLevelObjectivesBulkDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelObjectivesBulkDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
