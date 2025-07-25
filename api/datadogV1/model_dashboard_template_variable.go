// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardTemplateVariable Template variable.
type DashboardTemplateVariable struct {
	// The list of values that the template variable drop-down is limited to.
	AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
	// (deprecated) The default value for the template variable on dashboard load. Cannot be used in conjunction with `defaults`.
	// Deprecated
	Default datadog.NullableString `json:"default,omitempty"`
	// One or many default values for template variables on load. If more than one default is specified, they will be unioned together with `OR`. Cannot be used in conjunction with `default`.
	Defaults []string `json:"defaults,omitempty"`
	// The name of the variable.
	Name string `json:"name"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable drop-down.
	Prefix datadog.NullableString `json:"prefix,omitempty"`
	// The type of variable. This is to differentiate between filter variables (interpolated in query) and group by variables (interpolated into group by).
	Type datadog.NullableString `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardTemplateVariable instantiates a new DashboardTemplateVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardTemplateVariable(name string) *DashboardTemplateVariable {
	this := DashboardTemplateVariable{}
	this.Name = name
	return &this
}

// NewDashboardTemplateVariableWithDefaults instantiates a new DashboardTemplateVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardTemplateVariableWithDefaults() *DashboardTemplateVariable {
	this := DashboardTemplateVariable{}
	return &this
}

// GetAvailableValues returns the AvailableValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardTemplateVariable) GetAvailableValues() []string {
	if o == nil || o.AvailableValues.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AvailableValues.Get()
}

// GetAvailableValuesOk returns a tuple with the AvailableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardTemplateVariable) GetAvailableValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableValues.Get(), o.AvailableValues.IsSet()
}

// HasAvailableValues returns a boolean if a field has been set.
func (o *DashboardTemplateVariable) HasAvailableValues() bool {
	return o != nil && o.AvailableValues.IsSet()
}

// SetAvailableValues gets a reference to the given datadog.NullableList[string] and assigns it to the AvailableValues field.
func (o *DashboardTemplateVariable) SetAvailableValues(v []string) {
	o.AvailableValues.Set(&v)
}

// SetAvailableValuesNil sets the value for AvailableValues to be an explicit nil.
func (o *DashboardTemplateVariable) SetAvailableValuesNil() {
	o.AvailableValues.Set(nil)
}

// UnsetAvailableValues ensures that no value is present for AvailableValues, not even an explicit nil.
func (o *DashboardTemplateVariable) UnsetAvailableValues() {
	o.AvailableValues.Unset()
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *DashboardTemplateVariable) GetDefault() string {
	if o == nil || o.Default.Get() == nil {
		var ret string
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
// Deprecated
func (o *DashboardTemplateVariable) GetDefaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *DashboardTemplateVariable) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given datadog.NullableString and assigns it to the Default field.
// Deprecated
func (o *DashboardTemplateVariable) SetDefault(v string) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *DashboardTemplateVariable) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *DashboardTemplateVariable) UnsetDefault() {
	o.Default.Unset()
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *DashboardTemplateVariable) GetDefaults() []string {
	if o == nil || o.Defaults == nil {
		var ret []string
		return ret
	}
	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardTemplateVariable) GetDefaultsOk() (*[]string, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return &o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *DashboardTemplateVariable) HasDefaults() bool {
	return o != nil && o.Defaults != nil
}

// SetDefaults gets a reference to the given []string and assigns it to the Defaults field.
func (o *DashboardTemplateVariable) SetDefaults(v []string) {
	o.Defaults = v
}

// GetName returns the Name field value.
func (o *DashboardTemplateVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DashboardTemplateVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DashboardTemplateVariable) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardTemplateVariable) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardTemplateVariable) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *DashboardTemplateVariable) HasPrefix() bool {
	return o != nil && o.Prefix.IsSet()
}

// SetPrefix gets a reference to the given datadog.NullableString and assigns it to the Prefix field.
func (o *DashboardTemplateVariable) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil.
func (o *DashboardTemplateVariable) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil.
func (o *DashboardTemplateVariable) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardTemplateVariable) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardTemplateVariable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *DashboardTemplateVariable) HasType() bool {
	return o != nil && o.Type.IsSet()
}

// SetType gets a reference to the given datadog.NullableString and assigns it to the Type field.
func (o *DashboardTemplateVariable) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil.
func (o *DashboardTemplateVariable) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil.
func (o *DashboardTemplateVariable) UnsetType() {
	o.Type.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardTemplateVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AvailableValues.IsSet() {
		toSerialize["available_values"] = o.AvailableValues.Get()
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	toSerialize["name"] = o.Name
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardTemplateVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvailableValues datadog.NullableList[string] `json:"available_values,omitempty"`
		Default         datadog.NullableString       `json:"default,omitempty"`
		Defaults        []string                     `json:"defaults,omitempty"`
		Name            *string                      `json:"name"`
		Prefix          datadog.NullableString       `json:"prefix,omitempty"`
		Type            datadog.NullableString       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"available_values", "default", "defaults", "name", "prefix", "type"})
	} else {
		return err
	}
	o.AvailableValues = all.AvailableValues
	o.Default = all.Default
	o.Defaults = all.Defaults
	o.Name = *all.Name
	o.Prefix = all.Prefix
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
