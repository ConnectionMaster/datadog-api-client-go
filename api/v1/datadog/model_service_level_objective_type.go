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

// ServiceLevelObjectiveType The type of the service level objective.
type ServiceLevelObjectiveType string

// List of ServiceLevelObjectiveType
const (
	SERVICELEVELOBJECTIVETYPE_METRIC  ServiceLevelObjectiveType = "metric"
	SERVICELEVELOBJECTIVETYPE_MONITOR ServiceLevelObjectiveType = "monitor"
)

// Ptr returns reference to ServiceLevelObjectiveType value
func (v ServiceLevelObjectiveType) Ptr() *ServiceLevelObjectiveType {
	return &v
}

type NullableServiceLevelObjectiveType struct {
	value *ServiceLevelObjectiveType
	isSet bool
}

func (v NullableServiceLevelObjectiveType) Get() *ServiceLevelObjectiveType {
	return v.value
}

func (v NullableServiceLevelObjectiveType) Set(val *ServiceLevelObjectiveType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelObjectiveType) IsSet() bool {
	return v.isSet
}

func (v NullableServiceLevelObjectiveType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelObjectiveType(val *ServiceLevelObjectiveType) *NullableServiceLevelObjectiveType {
	return &NullableServiceLevelObjectiveType{value: val, isSet: true}
}

func (v NullableServiceLevelObjectiveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelObjectiveType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}