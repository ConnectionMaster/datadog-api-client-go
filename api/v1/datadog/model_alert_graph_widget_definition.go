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

// AlertGraphWidgetDefinition Alert graphs are timeseries graphs showing the current status of any monitor defined on your system.
type AlertGraphWidgetDefinition struct {
	// ID of the alert to use in the widget.
	AlertId string      `json:"alert_id"`
	Time    *WidgetTime `json:"time,omitempty"`
	// The title of the widget.
	Title      *string          `json:"title,omitempty"`
	TitleAlign *WidgetTextAlign `json:"title_align,omitempty"`
	// Size of the title.
	TitleSize *string                        `json:"title_size,omitempty"`
	Type      AlertGraphWidgetDefinitionType `json:"type"`
	VizType   WidgetVizType                  `json:"viz_type"`
}

// NewAlertGraphWidgetDefinition instantiates a new AlertGraphWidgetDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertGraphWidgetDefinition(alertId string, type_ AlertGraphWidgetDefinitionType, vizType WidgetVizType) *AlertGraphWidgetDefinition {
	this := AlertGraphWidgetDefinition{}
	this.AlertId = alertId
	this.Type = type_
	this.VizType = vizType
	return &this
}

// NewAlertGraphWidgetDefinitionWithDefaults instantiates a new AlertGraphWidgetDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertGraphWidgetDefinitionWithDefaults() *AlertGraphWidgetDefinition {
	this := AlertGraphWidgetDefinition{}
	var type_ AlertGraphWidgetDefinitionType = ALERTGRAPHWIDGETDEFINITIONTYPE_ALERT_GRAPH
	this.Type = type_
	return &this
}

// GetAlertId returns the AlertId field value
func (o *AlertGraphWidgetDefinition) GetAlertId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetAlertIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertId, true
}

// SetAlertId sets field value
func (o *AlertGraphWidgetDefinition) SetAlertId(v string) {
	o.AlertId = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTime() WidgetTime {
	if o == nil || o.Time == nil {
		var ret WidgetTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTimeOk() (*WidgetTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given WidgetTime and assigns it to the Time field.
func (o *AlertGraphWidgetDefinition) SetTime(v WidgetTime) {
	o.Time = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AlertGraphWidgetDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetTitleAlign returns the TitleAlign field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTitleAlign() WidgetTextAlign {
	if o == nil || o.TitleAlign == nil {
		var ret WidgetTextAlign
		return ret
	}
	return *o.TitleAlign
}

// GetTitleAlignOk returns a tuple with the TitleAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool) {
	if o == nil || o.TitleAlign == nil {
		return nil, false
	}
	return o.TitleAlign, true
}

// HasTitleAlign returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTitleAlign() bool {
	if o != nil && o.TitleAlign != nil {
		return true
	}

	return false
}

// SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.
func (o *AlertGraphWidgetDefinition) SetTitleAlign(v WidgetTextAlign) {
	o.TitleAlign = &v
}

// GetTitleSize returns the TitleSize field value if set, zero value otherwise.
func (o *AlertGraphWidgetDefinition) GetTitleSize() string {
	if o == nil || o.TitleSize == nil {
		var ret string
		return ret
	}
	return *o.TitleSize
}

// GetTitleSizeOk returns a tuple with the TitleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTitleSizeOk() (*string, bool) {
	if o == nil || o.TitleSize == nil {
		return nil, false
	}
	return o.TitleSize, true
}

// HasTitleSize returns a boolean if a field has been set.
func (o *AlertGraphWidgetDefinition) HasTitleSize() bool {
	if o != nil && o.TitleSize != nil {
		return true
	}

	return false
}

// SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.
func (o *AlertGraphWidgetDefinition) SetTitleSize(v string) {
	o.TitleSize = &v
}

// GetType returns the Type field value
func (o *AlertGraphWidgetDefinition) GetType() AlertGraphWidgetDefinitionType {
	if o == nil {
		var ret AlertGraphWidgetDefinitionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetTypeOk() (*AlertGraphWidgetDefinitionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AlertGraphWidgetDefinition) SetType(v AlertGraphWidgetDefinitionType) {
	o.Type = v
}

// GetVizType returns the VizType field value
func (o *AlertGraphWidgetDefinition) GetVizType() WidgetVizType {
	if o == nil {
		var ret WidgetVizType
		return ret
	}

	return o.VizType
}

// GetVizTypeOk returns a tuple with the VizType field value
// and a boolean to check if the value has been set.
func (o *AlertGraphWidgetDefinition) GetVizTypeOk() (*WidgetVizType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VizType, true
}

// SetVizType sets field value
func (o *AlertGraphWidgetDefinition) SetVizType(v WidgetVizType) {
	o.VizType = v
}

func (o AlertGraphWidgetDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alert_id"] = o.AlertId
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
	if true {
		toSerialize["viz_type"] = o.VizType
	}
	return json.Marshal(toSerialize)
}

type NullableAlertGraphWidgetDefinition struct {
	value *AlertGraphWidgetDefinition
	isSet bool
}

func (v NullableAlertGraphWidgetDefinition) Get() *AlertGraphWidgetDefinition {
	return v.value
}

func (v *NullableAlertGraphWidgetDefinition) Set(val *AlertGraphWidgetDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertGraphWidgetDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertGraphWidgetDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertGraphWidgetDefinition(val *AlertGraphWidgetDefinition) *NullableAlertGraphWidgetDefinition {
	return &NullableAlertGraphWidgetDefinition{value: val, isSet: true}
}

func (v NullableAlertGraphWidgetDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertGraphWidgetDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
