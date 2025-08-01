// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventAttributesAuthor The entity that made the change.
type ChangeEventAttributesAuthor struct {
	// The name of the user or system that made the change.
	Name *string `json:"name,omitempty"`
	// The type of the author.
	Type *ChangeEventAttributesAuthorType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeEventAttributesAuthor instantiates a new ChangeEventAttributesAuthor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventAttributesAuthor() *ChangeEventAttributesAuthor {
	this := ChangeEventAttributesAuthor{}
	return &this
}

// NewChangeEventAttributesAuthorWithDefaults instantiates a new ChangeEventAttributesAuthor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventAttributesAuthorWithDefaults() *ChangeEventAttributesAuthor {
	this := ChangeEventAttributesAuthor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChangeEventAttributesAuthor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributesAuthor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChangeEventAttributesAuthor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChangeEventAttributesAuthor) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ChangeEventAttributesAuthor) GetType() ChangeEventAttributesAuthorType {
	if o == nil || o.Type == nil {
		var ret ChangeEventAttributesAuthorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventAttributesAuthor) GetTypeOk() (*ChangeEventAttributesAuthorType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ChangeEventAttributesAuthor) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ChangeEventAttributesAuthorType and assigns it to the Type field.
func (o *ChangeEventAttributesAuthor) SetType(v ChangeEventAttributesAuthorType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventAttributesAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEventAttributesAuthor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string                          `json:"name,omitempty"`
		Type *ChangeEventAttributesAuthorType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
