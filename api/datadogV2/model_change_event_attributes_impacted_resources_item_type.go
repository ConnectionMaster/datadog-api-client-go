// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventAttributesImpactedResourcesItemType The type of the impacted resource.
type ChangeEventAttributesImpactedResourcesItemType string

// List of ChangeEventAttributesImpactedResourcesItemType.
const (
	CHANGEEVENTATTRIBUTESIMPACTEDRESOURCESITEMTYPE_SERVICE ChangeEventAttributesImpactedResourcesItemType = "service"
)

var allowedChangeEventAttributesImpactedResourcesItemTypeEnumValues = []ChangeEventAttributesImpactedResourcesItemType{
	CHANGEEVENTATTRIBUTESIMPACTEDRESOURCESITEMTYPE_SERVICE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeEventAttributesImpactedResourcesItemType) GetAllowedValues() []ChangeEventAttributesImpactedResourcesItemType {
	return allowedChangeEventAttributesImpactedResourcesItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeEventAttributesImpactedResourcesItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeEventAttributesImpactedResourcesItemType(value)
	return nil
}

// NewChangeEventAttributesImpactedResourcesItemTypeFromValue returns a pointer to a valid ChangeEventAttributesImpactedResourcesItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeEventAttributesImpactedResourcesItemTypeFromValue(v string) (*ChangeEventAttributesImpactedResourcesItemType, error) {
	ev := ChangeEventAttributesImpactedResourcesItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeEventAttributesImpactedResourcesItemType: valid values are %v", v, allowedChangeEventAttributesImpactedResourcesItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeEventAttributesImpactedResourcesItemType) IsValid() bool {
	for _, existing := range allowedChangeEventAttributesImpactedResourcesItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeEventAttributesImpactedResourcesItemType value.
func (v ChangeEventAttributesImpactedResourcesItemType) Ptr() *ChangeEventAttributesImpactedResourcesItemType {
	return &v
}
