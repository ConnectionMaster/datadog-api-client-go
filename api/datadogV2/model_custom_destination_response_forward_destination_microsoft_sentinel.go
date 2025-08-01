// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationMicrosoftSentinel The Microsoft Sentinel destination.
type CustomDestinationResponseForwardDestinationMicrosoftSentinel struct {
	// Client ID from the Datadog Azure integration.
	ClientId string `json:"client_id"`
	// Azure data collection endpoint.
	DataCollectionEndpoint string `json:"data_collection_endpoint"`
	// Azure data collection rule ID.
	DataCollectionRuleId string `json:"data_collection_rule_id"`
	// Azure stream name.
	StreamName string `json:"stream_name"`
	// Tenant ID from the Datadog Azure integration.
	TenantId string `json:"tenant_id"`
	// Type of the Microsoft Sentinel destination.
	Type CustomDestinationResponseForwardDestinationMicrosoftSentinelType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseForwardDestinationMicrosoftSentinel instantiates a new CustomDestinationResponseForwardDestinationMicrosoftSentinel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseForwardDestinationMicrosoftSentinel(clientId string, dataCollectionEndpoint string, dataCollectionRuleId string, streamName string, tenantId string, typeVar CustomDestinationResponseForwardDestinationMicrosoftSentinelType) *CustomDestinationResponseForwardDestinationMicrosoftSentinel {
	this := CustomDestinationResponseForwardDestinationMicrosoftSentinel{}
	this.ClientId = clientId
	this.DataCollectionEndpoint = dataCollectionEndpoint
	this.DataCollectionRuleId = dataCollectionRuleId
	this.StreamName = streamName
	this.TenantId = tenantId
	this.Type = typeVar
	return &this
}

// NewCustomDestinationResponseForwardDestinationMicrosoftSentinelWithDefaults instantiates a new CustomDestinationResponseForwardDestinationMicrosoftSentinel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseForwardDestinationMicrosoftSentinelWithDefaults() *CustomDestinationResponseForwardDestinationMicrosoftSentinel {
	this := CustomDestinationResponseForwardDestinationMicrosoftSentinel{}
	var typeVar CustomDestinationResponseForwardDestinationMicrosoftSentinelType = CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONMICROSOFTSENTINELTYPE_MICROSOFT_SENTINEL
	this.Type = typeVar
	return &this
}

// GetClientId returns the ClientId field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) SetClientId(v string) {
	o.ClientId = v
}

// GetDataCollectionEndpoint returns the DataCollectionEndpoint field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetDataCollectionEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataCollectionEndpoint
}

// GetDataCollectionEndpointOk returns a tuple with the DataCollectionEndpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetDataCollectionEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataCollectionEndpoint, true
}

// SetDataCollectionEndpoint sets field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) SetDataCollectionEndpoint(v string) {
	o.DataCollectionEndpoint = v
}

// GetDataCollectionRuleId returns the DataCollectionRuleId field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetDataCollectionRuleId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataCollectionRuleId
}

// GetDataCollectionRuleIdOk returns a tuple with the DataCollectionRuleId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetDataCollectionRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataCollectionRuleId, true
}

// SetDataCollectionRuleId sets field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) SetDataCollectionRuleId(v string) {
	o.DataCollectionRuleId = v
}

// GetStreamName returns the StreamName field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) SetStreamName(v string) {
	o.StreamName = v
}

// GetTenantId returns the TenantId field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) SetTenantId(v string) {
	o.TenantId = v
}

// GetType returns the Type field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetType() CustomDestinationResponseForwardDestinationMicrosoftSentinelType {
	if o == nil {
		var ret CustomDestinationResponseForwardDestinationMicrosoftSentinelType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) GetTypeOk() (*CustomDestinationResponseForwardDestinationMicrosoftSentinelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) SetType(v CustomDestinationResponseForwardDestinationMicrosoftSentinelType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseForwardDestinationMicrosoftSentinel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["data_collection_endpoint"] = o.DataCollectionEndpoint
	toSerialize["data_collection_rule_id"] = o.DataCollectionRuleId
	toSerialize["stream_name"] = o.StreamName
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationResponseForwardDestinationMicrosoftSentinel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId               *string                                                           `json:"client_id"`
		DataCollectionEndpoint *string                                                           `json:"data_collection_endpoint"`
		DataCollectionRuleId   *string                                                           `json:"data_collection_rule_id"`
		StreamName             *string                                                           `json:"stream_name"`
		TenantId               *string                                                           `json:"tenant_id"`
		Type                   *CustomDestinationResponseForwardDestinationMicrosoftSentinelType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.DataCollectionEndpoint == nil {
		return fmt.Errorf("required field data_collection_endpoint missing")
	}
	if all.DataCollectionRuleId == nil {
		return fmt.Errorf("required field data_collection_rule_id missing")
	}
	if all.StreamName == nil {
		return fmt.Errorf("required field stream_name missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "data_collection_endpoint", "data_collection_rule_id", "stream_name", "tenant_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ClientId = *all.ClientId
	o.DataCollectionEndpoint = *all.DataCollectionEndpoint
	o.DataCollectionRuleId = *all.DataCollectionRuleId
	o.StreamName = *all.StreamName
	o.TenantId = *all.TenantId
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
