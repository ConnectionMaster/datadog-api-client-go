// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatsigAPIKeyType The definition of the `StatsigAPIKey` object.
type StatsigAPIKeyType string

// List of StatsigAPIKeyType.
const (
	STATSIGAPIKEYTYPE_STATSIGAPIKEY StatsigAPIKeyType = "StatsigAPIKey"
)

var allowedStatsigAPIKeyTypeEnumValues = []StatsigAPIKeyType{
	STATSIGAPIKEYTYPE_STATSIGAPIKEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StatsigAPIKeyType) GetAllowedValues() []StatsigAPIKeyType {
	return allowedStatsigAPIKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StatsigAPIKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StatsigAPIKeyType(value)
	return nil
}

// NewStatsigAPIKeyTypeFromValue returns a pointer to a valid StatsigAPIKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStatsigAPIKeyTypeFromValue(v string) (*StatsigAPIKeyType, error) {
	ev := StatsigAPIKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StatsigAPIKeyType: valid values are %v", v, allowedStatsigAPIKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StatsigAPIKeyType) IsValid() bool {
	for _, existing := range allowedStatsigAPIKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatsigAPIKeyType value.
func (v StatsigAPIKeyType) Ptr() *StatsigAPIKeyType {
	return &v
}
