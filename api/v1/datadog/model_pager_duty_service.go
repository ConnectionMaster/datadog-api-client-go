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

// PagerDutyService Configure your Datadog-PagerDuty integration directly through the Datadog API. For more informations, see the [PagerDuty integration page](https://docs.datadoghq.com/integrations/pagerduty/).
type PagerDutyService struct {
	// Your Service name associated service key in Pagerduty.
	ServiceKey string `json:"service_key"`
	// Your Service name in PagerDuty.
	ServiceName string `json:"service_name"`
}

// NewPagerDutyService instantiates a new PagerDutyService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagerDutyService(serviceKey string, serviceName string) *PagerDutyService {
	this := PagerDutyService{}
	this.ServiceKey = serviceKey
	this.ServiceName = serviceName
	return &this
}

// NewPagerDutyServiceWithDefaults instantiates a new PagerDutyService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagerDutyServiceWithDefaults() *PagerDutyService {
	this := PagerDutyService{}
	return &this
}

// GetServiceKey returns the ServiceKey field value
func (o *PagerDutyService) GetServiceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceKey
}

// SetServiceKey sets field value
func (o *PagerDutyService) SetServiceKey(v string) {
	o.ServiceKey = v
}

// GetServiceName returns the ServiceName field value
func (o *PagerDutyService) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// SetServiceName sets field value
func (o *PagerDutyService) SetServiceName(v string) {
	o.ServiceName = v
}

func (o PagerDutyService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service_key"] = o.ServiceKey
	}
	if true {
		toSerialize["service_name"] = o.ServiceName
	}
	return json.Marshal(toSerialize)
}

type NullablePagerDutyService struct {
	value *PagerDutyService
	isSet bool
}

func (v NullablePagerDutyService) Get() *PagerDutyService {
	return v.value
}

func (v NullablePagerDutyService) Set(val *PagerDutyService) {
	v.value = val
	v.isSet = true
}

func (v NullablePagerDutyService) IsSet() bool {
	return v.isSet
}

func (v NullablePagerDutyService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagerDutyService(val *PagerDutyService) *NullablePagerDutyService {
	return &NullablePagerDutyService{value: val, isSet: true}
}

func (v NullablePagerDutyService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagerDutyService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}