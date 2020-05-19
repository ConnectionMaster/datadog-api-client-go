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

// DistributionWidgetDefinition The Distribution visualization is another way of showing metrics aggregated across one or several tags, such as hosts. Unlike the heat map, a distribution graph’s x-axis is quantity rather than time.
type DistributionWidgetDefinition struct {
	LegendSize *WidgetLegendSize `json:"legend_size,omitempty"`
	// Array of one request object to display in the widget.  See the dedicated [Request JSON schema documentation](https://docs.datadoghq.com/dashboards/graphing_json/request_json)  to learn how to build the `REQUEST_SCHEMA`.
	Requests []DistributionWidgetRequest `json:"requests"`
	// Whether or not to display the legend on this widget.
	ShowLegend *bool       `json:"show_legend,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
	// Title of the widget.
	Title      *string          `json:"title,omitempty"`
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string                          `json:"title_size,omitempty"`
	Type      DistributionWidgetDefinitionType `json:"type"`
}

// NewDistributionWidgetDefinition instantiates a new DistributionWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionWidgetDefinition(requests []DistributionWidgetRequest, type_ DistributionWidgetDefinitionType) *DistributionWidgetDefinition {
	this := DistributionWidgetDefinition{}
	this.Requests = requests
	this.Type = type_
	return &this
}

// NewDistributionWidgetDefinitionWithDefaults instantiates a new DistributionWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionWidgetDefinitionWithDefaults() *DistributionWidgetDefinition {
	this := DistributionWidgetDefinition{}
	var type_ DistributionWidgetDefinitionType = "distribution"
	this.Type = type_
	return &this
}

// GetLegendSize returns the LegendSize field value if set, zero value otherwise.
func (o *DistributionWidgetDefinition) GetLegendSize() WidgetLegendSize {
	if o == nil || o.LegendSize == nil {
		var ret WidgetLegendSize
		return ret
	}
	return *o.LegendSize
}

// GetLegendSizeOk returns a tuple with the LegendSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetLegendSizeOk() (*WidgetLegendSize, bool) {
	if o == nil || o.LegendSize == nil {
		return nil, false
	}
	return o.LegendSize, true
}

// HasLegendSize returns a boolean if a field has been set.
func (o *DistributionWidgetDefinition) HasLegendSize() bool {
	if o != nil && o.LegendSize != nil {
		return true
	}

	return false
}

// SetLegendSize gets a reference to the given WidgetLegendSize and assigns it to the LegendSize field.
func (o *DistributionWidgetDefinition) SetLegendSize(v WidgetLegendSize) {
	o.LegendSize = &v
}

// GetRequests returns the Requests field value
func (o *DistributionWidgetDefinition) GetRequests() []DistributionWidgetRequest {
	if o == nil {
		var ret []DistributionWidgetRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetRequestsOk() (*[]DistributionWidgetRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value
func (o *DistributionWidgetDefinition) SetRequests(v []DistributionWidgetRequest) {
	o.Requests = v
}

// GetShowLegend returns the ShowLegend field value if set, zero value otherwise.
func (o *DistributionWidgetDefinition) GetShowLegend() bool {
	if o == nil || o.ShowLegend == nil {
		var ret bool
		return ret
	}
	return *o.ShowLegend
}

// GetShowLegendOk returns a tuple with the ShowLegend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetShowLegendOk() (*bool, bool) {
	if o == nil || o.ShowLegend == nil {
		return nil, false
	}
	return o.ShowLegend, true
}

// HasShowLegend returns a boolean if a field has been set.
func (o *DistributionWidgetDefinition) HasShowLegend() bool {
	if o != nil && o.ShowLegend != nil {
		return true
	}

	return false
}

// SetShowLegend gets a reference to the given bool and assigns it to the ShowLegend field.
func (o *DistributionWidgetDefinition) SetShowLegend(v bool) {
	o.ShowLegend = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DistributionWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DistributionWidgetDefinition) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *DistributionWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DistributionWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DistributionWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DistributionWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *DistributionWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *DistributionWidgetDefinition) HasTitleAlign() bool {
	if o != nil && o.TitleAlign != nil {
		return true
	}

	return false
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *DistributionWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *DistributionWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *DistributionWidgetDefinition) HasTitleSize() bool {
	if o != nil && o.TitleSize != nil {
		return true
	}

	return false
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *DistributionWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value
func (o *DistributionWidgetDefinition) GetType() DistributionWidgetDefinitionType {
	if o == nil {
		var ret DistributionWidgetDefinitionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DistributionWidgetDefinition) GetTypeOk() (*DistributionWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DistributionWidgetDefinition) SetType(v DistributionWidgetDefinitionType) {
	o.Type = v
}

func (o DistributionWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LegendSize != nil {
		toSerialize["legend_size"] = o.LegendSize
	}
	if true {
		toSerialize["requests"] = o.Requests
	}
	if o.ShowLegend != nil {
		toSerialize["show_legend"] = o.ShowLegend
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.TitleAlign != nil {
		toSerialize["title_align"] = o.TitleAlign
	}
	if o.TitleSize != nil {
		toSerialize["title_size"] = o.TitleSize
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDistributionWidgetDefinition struct {
	value *DistributionWidgetDefinition
	isSet bool
}

func (v NullableDistributionWidgetDefinition) Get() *DistributionWidgetDefinition {
	return v.value
}

func (v *NullableDistributionWidgetDefinition) Set(val *DistributionWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionWidgetDefinition(val *DistributionWidgetDefinition) *NullableDistributionWidgetDefinition {
	return &NullableDistributionWidgetDefinition{value: val, isSet: true}
}

func (v NullableDistributionWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
