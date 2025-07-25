// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventCreateResponseAttributesAttributesEvt JSON object of event system attributes.
type EventCreateResponseAttributesAttributesEvt struct {
	// Event identifier. This field is deprecated and will be removed in a future version. Use the `uid` field instead.
	// Deprecated
	Id *string `json:"id,omitempty"`
	// A unique identifier for the event. You can use this identifier to query or reference the event.
	Uid *string `json:"uid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventCreateResponseAttributesAttributesEvt instantiates a new EventCreateResponseAttributesAttributesEvt object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventCreateResponseAttributesAttributesEvt() *EventCreateResponseAttributesAttributesEvt {
	this := EventCreateResponseAttributesAttributesEvt{}
	return &this
}

// NewEventCreateResponseAttributesAttributesEvtWithDefaults instantiates a new EventCreateResponseAttributesAttributesEvt object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventCreateResponseAttributesAttributesEvtWithDefaults() *EventCreateResponseAttributesAttributesEvt {
	this := EventCreateResponseAttributesAttributesEvt{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
// Deprecated
func (o *EventCreateResponseAttributesAttributesEvt) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EventCreateResponseAttributesAttributesEvt) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventCreateResponseAttributesAttributesEvt) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
// Deprecated
func (o *EventCreateResponseAttributesAttributesEvt) SetId(v string) {
	o.Id = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *EventCreateResponseAttributesAttributesEvt) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponseAttributesAttributesEvt) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *EventCreateResponseAttributesAttributesEvt) HasUid() bool {
	return o != nil && o.Uid != nil
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *EventCreateResponseAttributesAttributesEvt) SetUid(v string) {
	o.Uid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventCreateResponseAttributesAttributesEvt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventCreateResponseAttributesAttributesEvt) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id  *string `json:"id,omitempty"`
		Uid *string `json:"uid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "uid"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Uid = all.Uid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
