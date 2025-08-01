// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VirusTotalIntegrationUpdate The definition of the `VirusTotalIntegrationUpdate` object.
type VirusTotalIntegrationUpdate struct {
	// The definition of the `VirusTotalCredentialsUpdate` object.
	Credentials *VirusTotalCredentialsUpdate `json:"credentials,omitempty"`
	// The definition of the `VirusTotalIntegrationType` object.
	Type VirusTotalIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVirusTotalIntegrationUpdate instantiates a new VirusTotalIntegrationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVirusTotalIntegrationUpdate(typeVar VirusTotalIntegrationType) *VirusTotalIntegrationUpdate {
	this := VirusTotalIntegrationUpdate{}
	this.Type = typeVar
	return &this
}

// NewVirusTotalIntegrationUpdateWithDefaults instantiates a new VirusTotalIntegrationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVirusTotalIntegrationUpdateWithDefaults() *VirusTotalIntegrationUpdate {
	this := VirusTotalIntegrationUpdate{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *VirusTotalIntegrationUpdate) GetCredentials() VirusTotalCredentialsUpdate {
	if o == nil || o.Credentials == nil {
		var ret VirusTotalCredentialsUpdate
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirusTotalIntegrationUpdate) GetCredentialsOk() (*VirusTotalCredentialsUpdate, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *VirusTotalIntegrationUpdate) HasCredentials() bool {
	return o != nil && o.Credentials != nil
}

// SetCredentials gets a reference to the given VirusTotalCredentialsUpdate and assigns it to the Credentials field.
func (o *VirusTotalIntegrationUpdate) SetCredentials(v VirusTotalCredentialsUpdate) {
	o.Credentials = &v
}

// GetType returns the Type field value.
func (o *VirusTotalIntegrationUpdate) GetType() VirusTotalIntegrationType {
	if o == nil {
		var ret VirusTotalIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VirusTotalIntegrationUpdate) GetTypeOk() (*VirusTotalIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *VirusTotalIntegrationUpdate) SetType(v VirusTotalIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VirusTotalIntegrationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VirusTotalIntegrationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Credentials *VirusTotalCredentialsUpdate `json:"credentials,omitempty"`
		Type        *VirusTotalIntegrationType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"credentials", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Credentials = all.Credentials
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
