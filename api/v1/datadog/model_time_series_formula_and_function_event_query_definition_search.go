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

// TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch Search options.
type TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch struct {
	// Events search string.
	Query string `json:"query"`
}

// NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch(query string) *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch {
	this := TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch{}
	this.Query = query
	return &this
}

// NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSearchWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSearchWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch {
	this := TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch{}
	return &this
}

// GetQuery returns the Query field value
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) SetQuery(v string) {
	o.Query = v
}

func (o TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch struct {
	value *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch
	isSet bool
}

func (v NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) Get() *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch {
	return v.value
}

func (v *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) Set(val *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch(val *TimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch {
	return &NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch{value: val, isSet: true}
}

func (v NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesFormulaAndFunctionEventQueryDefinitionSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
