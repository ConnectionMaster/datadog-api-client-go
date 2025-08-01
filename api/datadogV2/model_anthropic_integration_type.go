// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnthropicIntegrationType The definition of the `AnthropicIntegrationType` object.
type AnthropicIntegrationType string

// List of AnthropicIntegrationType.
const (
	ANTHROPICINTEGRATIONTYPE_ANTHROPIC AnthropicIntegrationType = "Anthropic"
)

var allowedAnthropicIntegrationTypeEnumValues = []AnthropicIntegrationType{
	ANTHROPICINTEGRATIONTYPE_ANTHROPIC,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AnthropicIntegrationType) GetAllowedValues() []AnthropicIntegrationType {
	return allowedAnthropicIntegrationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AnthropicIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AnthropicIntegrationType(value)
	return nil
}

// NewAnthropicIntegrationTypeFromValue returns a pointer to a valid AnthropicIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAnthropicIntegrationTypeFromValue(v string) (*AnthropicIntegrationType, error) {
	ev := AnthropicIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AnthropicIntegrationType: valid values are %v", v, allowedAnthropicIntegrationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AnthropicIntegrationType) IsValid() bool {
	for _, existing := range allowedAnthropicIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnthropicIntegrationType value.
func (v AnthropicIntegrationType) Ptr() *AnthropicIntegrationType {
	return &v
}
