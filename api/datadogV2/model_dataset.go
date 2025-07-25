// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Dataset Dataset object.
//
// ### Datasets Constraints
// - **Tag Limit per Dataset**:
//   - Each restricted dataset supports a maximum of 10 key:value pairs per product.
//
// - **Tag Key Rules per Telemetry Type**:
//   - Only one tag key or attribute may be used to define access within a single telemetry type.
//   - The same or different tag key may be used across different telemetry types.
//
// - **Tag Value Uniqueness**:
//   - Tag values must be unique within a single dataset.
//   - A tag value used in one dataset cannot be reused in another dataset of the same telemetry type.
type Dataset struct {
	// Dataset metadata and configuration(s).
	Attributes DatasetAttributes `json:"attributes"`
	// Unique identifier for the dataset.
	Id *string `json:"id,omitempty"`
	// Resource type, always "dataset".
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataset instantiates a new Dataset object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataset(attributes DatasetAttributes, typeVar string) *Dataset {
	this := Dataset{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewDatasetWithDefaults instantiates a new Dataset object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetWithDefaults() *Dataset {
	this := Dataset{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *Dataset) GetAttributes() DatasetAttributes {
	if o == nil {
		var ret DatasetAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Dataset) GetAttributesOk() (*DatasetAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *Dataset) SetAttributes(v DatasetAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dataset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dataset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dataset) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dataset) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *Dataset) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Dataset) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Dataset) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Dataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Dataset) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DatasetAttributes `json:"attributes"`
		Id         *string            `json:"id,omitempty"`
		Type       *string            `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = all.Id
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
