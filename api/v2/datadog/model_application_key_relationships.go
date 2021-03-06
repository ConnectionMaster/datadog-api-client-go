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

// ApplicationKeyRelationships Resources related to the application key.
type ApplicationKeyRelationships struct {
	CreatedBy *RelationshipToUser `json:"created_by,omitempty"`
}

// NewApplicationKeyRelationships instantiates a new ApplicationKeyRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationKeyRelationships() *ApplicationKeyRelationships {
	this := ApplicationKeyRelationships{}
	return &this
}

// NewApplicationKeyRelationshipsWithDefaults instantiates a new ApplicationKeyRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationKeyRelationshipsWithDefaults() *ApplicationKeyRelationships {
	this := ApplicationKeyRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ApplicationKeyRelationships) GetCreatedBy() RelationshipToUser {
	if o == nil || o.CreatedBy == nil {
		var ret RelationshipToUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationKeyRelationships) GetCreatedByOk() (*RelationshipToUser, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ApplicationKeyRelationships) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given RelationshipToUser and assigns it to the CreatedBy field.
func (o *ApplicationKeyRelationships) SetCreatedBy(v RelationshipToUser) {
	o.CreatedBy = &v
}

func (o ApplicationKeyRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationKeyRelationships struct {
	value *ApplicationKeyRelationships
	isSet bool
}

func (v NullableApplicationKeyRelationships) Get() *ApplicationKeyRelationships {
	return v.value
}

func (v *NullableApplicationKeyRelationships) Set(val *ApplicationKeyRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeyRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeyRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeyRelationships(val *ApplicationKeyRelationships) *NullableApplicationKeyRelationships {
	return &NullableApplicationKeyRelationships{value: val, isSet: true}
}

func (v NullableApplicationKeyRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeyRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
