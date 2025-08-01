// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSIntegrationIamPermissionsResponseData AWS Integration IAM Permissions response data.
type AWSIntegrationIamPermissionsResponseData struct {
	// AWS Integration IAM Permissions response attributes.
	Attributes *AWSIntegrationIamPermissionsResponseAttributes `json:"attributes,omitempty"`
	// The `AWSIntegrationIamPermissionsResponseData` `id`.
	Id *string `json:"id,omitempty"`
	// The `AWSIntegrationIamPermissionsResponseData` `type`.
	Type *AWSIntegrationIamPermissionsResponseDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSIntegrationIamPermissionsResponseData instantiates a new AWSIntegrationIamPermissionsResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSIntegrationIamPermissionsResponseData() *AWSIntegrationIamPermissionsResponseData {
	this := AWSIntegrationIamPermissionsResponseData{}
	var id string = "permissions"
	this.Id = &id
	var typeVar AWSIntegrationIamPermissionsResponseDataType = AWSINTEGRATIONIAMPERMISSIONSRESPONSEDATATYPE_PERMISSIONS
	this.Type = &typeVar
	return &this
}

// NewAWSIntegrationIamPermissionsResponseDataWithDefaults instantiates a new AWSIntegrationIamPermissionsResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSIntegrationIamPermissionsResponseDataWithDefaults() *AWSIntegrationIamPermissionsResponseData {
	this := AWSIntegrationIamPermissionsResponseData{}
	var id string = "permissions"
	this.Id = &id
	var typeVar AWSIntegrationIamPermissionsResponseDataType = AWSINTEGRATIONIAMPERMISSIONSRESPONSEDATATYPE_PERMISSIONS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AWSIntegrationIamPermissionsResponseData) GetAttributes() AWSIntegrationIamPermissionsResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret AWSIntegrationIamPermissionsResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSIntegrationIamPermissionsResponseData) GetAttributesOk() (*AWSIntegrationIamPermissionsResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AWSIntegrationIamPermissionsResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given AWSIntegrationIamPermissionsResponseAttributes and assigns it to the Attributes field.
func (o *AWSIntegrationIamPermissionsResponseData) SetAttributes(v AWSIntegrationIamPermissionsResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSIntegrationIamPermissionsResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSIntegrationIamPermissionsResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSIntegrationIamPermissionsResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSIntegrationIamPermissionsResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AWSIntegrationIamPermissionsResponseData) GetType() AWSIntegrationIamPermissionsResponseDataType {
	if o == nil || o.Type == nil {
		var ret AWSIntegrationIamPermissionsResponseDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSIntegrationIamPermissionsResponseData) GetTypeOk() (*AWSIntegrationIamPermissionsResponseDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AWSIntegrationIamPermissionsResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AWSIntegrationIamPermissionsResponseDataType and assigns it to the Type field.
func (o *AWSIntegrationIamPermissionsResponseData) SetType(v AWSIntegrationIamPermissionsResponseDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSIntegrationIamPermissionsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *AWSIntegrationIamPermissionsResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *AWSIntegrationIamPermissionsResponseAttributes `json:"attributes,omitempty"`
		Id         *string                                         `json:"id,omitempty"`
		Type       *AWSIntegrationIamPermissionsResponseDataType   `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
