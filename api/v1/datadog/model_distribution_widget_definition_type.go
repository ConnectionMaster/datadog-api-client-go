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

import (
	"fmt"
)

// DistributionWidgetDefinitionType Type of the distribution widget.
type DistributionWidgetDefinitionType string

// List of DistributionWidgetDefinitionType
const (
	DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION DistributionWidgetDefinitionType = "distribution"
)

func (v *DistributionWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DistributionWidgetDefinitionType(value)
	for _, existing := range []DistributionWidgetDefinitionType{"distribution"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DistributionWidgetDefinitionType", *v)
}

// Ptr returns reference to DistributionWidgetDefinitionType value
func (v DistributionWidgetDefinitionType) Ptr() *DistributionWidgetDefinitionType {
	return &v
}

type NullableDistributionWidgetDefinitionType struct {
	value *DistributionWidgetDefinitionType
	isSet bool
}

func (v NullableDistributionWidgetDefinitionType) Get() *DistributionWidgetDefinitionType {
	return v.value
}

func (v *NullableDistributionWidgetDefinitionType) Set(val *DistributionWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionWidgetDefinitionType(val *DistributionWidgetDefinitionType) *NullableDistributionWidgetDefinitionType {
	return &NullableDistributionWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableDistributionWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
