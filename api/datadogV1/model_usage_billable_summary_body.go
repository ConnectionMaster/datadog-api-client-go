// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageBillableSummaryBody Response with properties for each aggregated usage type.
type UsageBillableSummaryBody struct {
	// The total account usage.
	AccountBillableUsage *int64 `json:"account_billable_usage,omitempty"`
	// The total account committed usage.
	AccountCommittedUsage *int64 `json:"account_committed_usage,omitempty"`
	// The total account on-demand usage.
	AccountOnDemandUsage *int64 `json:"account_on_demand_usage,omitempty"`
	// Elapsed usage hours for some billable product.
	ElapsedUsageHours *int64 `json:"elapsed_usage_hours,omitempty"`
	// The first billable hour for the org.
	FirstBillableUsageHour *time.Time `json:"first_billable_usage_hour,omitempty"`
	// The last billable hour for the org.
	LastBillableUsageHour *time.Time `json:"last_billable_usage_hour,omitempty"`
	// The number of units used within the billable timeframe.
	OrgBillableUsage *int64 `json:"org_billable_usage,omitempty"`
	// The percentage of account usage the org represents.
	PercentageInAccount *float64 `json:"percentage_in_account,omitempty"`
	// Units pertaining to the usage.
	UsageUnit *string `json:"usage_unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageBillableSummaryBody instantiates a new UsageBillableSummaryBody object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageBillableSummaryBody() *UsageBillableSummaryBody {
	this := UsageBillableSummaryBody{}
	return &this
}

// NewUsageBillableSummaryBodyWithDefaults instantiates a new UsageBillableSummaryBody object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageBillableSummaryBodyWithDefaults() *UsageBillableSummaryBody {
	this := UsageBillableSummaryBody{}
	return &this
}

// GetAccountBillableUsage returns the AccountBillableUsage field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetAccountBillableUsage() int64 {
	if o == nil || o.AccountBillableUsage == nil {
		var ret int64
		return ret
	}
	return *o.AccountBillableUsage
}

// GetAccountBillableUsageOk returns a tuple with the AccountBillableUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetAccountBillableUsageOk() (*int64, bool) {
	if o == nil || o.AccountBillableUsage == nil {
		return nil, false
	}
	return o.AccountBillableUsage, true
}

// HasAccountBillableUsage returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasAccountBillableUsage() bool {
	return o != nil && o.AccountBillableUsage != nil
}

// SetAccountBillableUsage gets a reference to the given int64 and assigns it to the AccountBillableUsage field.
func (o *UsageBillableSummaryBody) SetAccountBillableUsage(v int64) {
	o.AccountBillableUsage = &v
}

// GetAccountCommittedUsage returns the AccountCommittedUsage field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetAccountCommittedUsage() int64 {
	if o == nil || o.AccountCommittedUsage == nil {
		var ret int64
		return ret
	}
	return *o.AccountCommittedUsage
}

// GetAccountCommittedUsageOk returns a tuple with the AccountCommittedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetAccountCommittedUsageOk() (*int64, bool) {
	if o == nil || o.AccountCommittedUsage == nil {
		return nil, false
	}
	return o.AccountCommittedUsage, true
}

// HasAccountCommittedUsage returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasAccountCommittedUsage() bool {
	return o != nil && o.AccountCommittedUsage != nil
}

// SetAccountCommittedUsage gets a reference to the given int64 and assigns it to the AccountCommittedUsage field.
func (o *UsageBillableSummaryBody) SetAccountCommittedUsage(v int64) {
	o.AccountCommittedUsage = &v
}

// GetAccountOnDemandUsage returns the AccountOnDemandUsage field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetAccountOnDemandUsage() int64 {
	if o == nil || o.AccountOnDemandUsage == nil {
		var ret int64
		return ret
	}
	return *o.AccountOnDemandUsage
}

// GetAccountOnDemandUsageOk returns a tuple with the AccountOnDemandUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetAccountOnDemandUsageOk() (*int64, bool) {
	if o == nil || o.AccountOnDemandUsage == nil {
		return nil, false
	}
	return o.AccountOnDemandUsage, true
}

// HasAccountOnDemandUsage returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasAccountOnDemandUsage() bool {
	return o != nil && o.AccountOnDemandUsage != nil
}

// SetAccountOnDemandUsage gets a reference to the given int64 and assigns it to the AccountOnDemandUsage field.
func (o *UsageBillableSummaryBody) SetAccountOnDemandUsage(v int64) {
	o.AccountOnDemandUsage = &v
}

// GetElapsedUsageHours returns the ElapsedUsageHours field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetElapsedUsageHours() int64 {
	if o == nil || o.ElapsedUsageHours == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedUsageHours
}

// GetElapsedUsageHoursOk returns a tuple with the ElapsedUsageHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetElapsedUsageHoursOk() (*int64, bool) {
	if o == nil || o.ElapsedUsageHours == nil {
		return nil, false
	}
	return o.ElapsedUsageHours, true
}

// HasElapsedUsageHours returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasElapsedUsageHours() bool {
	return o != nil && o.ElapsedUsageHours != nil
}

// SetElapsedUsageHours gets a reference to the given int64 and assigns it to the ElapsedUsageHours field.
func (o *UsageBillableSummaryBody) SetElapsedUsageHours(v int64) {
	o.ElapsedUsageHours = &v
}

// GetFirstBillableUsageHour returns the FirstBillableUsageHour field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetFirstBillableUsageHour() time.Time {
	if o == nil || o.FirstBillableUsageHour == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstBillableUsageHour
}

// GetFirstBillableUsageHourOk returns a tuple with the FirstBillableUsageHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetFirstBillableUsageHourOk() (*time.Time, bool) {
	if o == nil || o.FirstBillableUsageHour == nil {
		return nil, false
	}
	return o.FirstBillableUsageHour, true
}

// HasFirstBillableUsageHour returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasFirstBillableUsageHour() bool {
	return o != nil && o.FirstBillableUsageHour != nil
}

// SetFirstBillableUsageHour gets a reference to the given time.Time and assigns it to the FirstBillableUsageHour field.
func (o *UsageBillableSummaryBody) SetFirstBillableUsageHour(v time.Time) {
	o.FirstBillableUsageHour = &v
}

// GetLastBillableUsageHour returns the LastBillableUsageHour field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetLastBillableUsageHour() time.Time {
	if o == nil || o.LastBillableUsageHour == nil {
		var ret time.Time
		return ret
	}
	return *o.LastBillableUsageHour
}

// GetLastBillableUsageHourOk returns a tuple with the LastBillableUsageHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetLastBillableUsageHourOk() (*time.Time, bool) {
	if o == nil || o.LastBillableUsageHour == nil {
		return nil, false
	}
	return o.LastBillableUsageHour, true
}

// HasLastBillableUsageHour returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasLastBillableUsageHour() bool {
	return o != nil && o.LastBillableUsageHour != nil
}

// SetLastBillableUsageHour gets a reference to the given time.Time and assigns it to the LastBillableUsageHour field.
func (o *UsageBillableSummaryBody) SetLastBillableUsageHour(v time.Time) {
	o.LastBillableUsageHour = &v
}

// GetOrgBillableUsage returns the OrgBillableUsage field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetOrgBillableUsage() int64 {
	if o == nil || o.OrgBillableUsage == nil {
		var ret int64
		return ret
	}
	return *o.OrgBillableUsage
}

// GetOrgBillableUsageOk returns a tuple with the OrgBillableUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetOrgBillableUsageOk() (*int64, bool) {
	if o == nil || o.OrgBillableUsage == nil {
		return nil, false
	}
	return o.OrgBillableUsage, true
}

// HasOrgBillableUsage returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasOrgBillableUsage() bool {
	return o != nil && o.OrgBillableUsage != nil
}

// SetOrgBillableUsage gets a reference to the given int64 and assigns it to the OrgBillableUsage field.
func (o *UsageBillableSummaryBody) SetOrgBillableUsage(v int64) {
	o.OrgBillableUsage = &v
}

// GetPercentageInAccount returns the PercentageInAccount field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetPercentageInAccount() float64 {
	if o == nil || o.PercentageInAccount == nil {
		var ret float64
		return ret
	}
	return *o.PercentageInAccount
}

// GetPercentageInAccountOk returns a tuple with the PercentageInAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetPercentageInAccountOk() (*float64, bool) {
	if o == nil || o.PercentageInAccount == nil {
		return nil, false
	}
	return o.PercentageInAccount, true
}

// HasPercentageInAccount returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasPercentageInAccount() bool {
	return o != nil && o.PercentageInAccount != nil
}

// SetPercentageInAccount gets a reference to the given float64 and assigns it to the PercentageInAccount field.
func (o *UsageBillableSummaryBody) SetPercentageInAccount(v float64) {
	o.PercentageInAccount = &v
}

// GetUsageUnit returns the UsageUnit field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetUsageUnit() string {
	if o == nil || o.UsageUnit == nil {
		var ret string
		return ret
	}
	return *o.UsageUnit
}

// GetUsageUnitOk returns a tuple with the UsageUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetUsageUnitOk() (*string, bool) {
	if o == nil || o.UsageUnit == nil {
		return nil, false
	}
	return o.UsageUnit, true
}

// HasUsageUnit returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasUsageUnit() bool {
	return o != nil && o.UsageUnit != nil
}

// SetUsageUnit gets a reference to the given string and assigns it to the UsageUnit field.
func (o *UsageBillableSummaryBody) SetUsageUnit(v string) {
	o.UsageUnit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageBillableSummaryBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountBillableUsage != nil {
		toSerialize["account_billable_usage"] = o.AccountBillableUsage
	}
	if o.AccountCommittedUsage != nil {
		toSerialize["account_committed_usage"] = o.AccountCommittedUsage
	}
	if o.AccountOnDemandUsage != nil {
		toSerialize["account_on_demand_usage"] = o.AccountOnDemandUsage
	}
	if o.ElapsedUsageHours != nil {
		toSerialize["elapsed_usage_hours"] = o.ElapsedUsageHours
	}
	if o.FirstBillableUsageHour != nil {
		if o.FirstBillableUsageHour.Nanosecond() == 0 {
			toSerialize["first_billable_usage_hour"] = o.FirstBillableUsageHour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["first_billable_usage_hour"] = o.FirstBillableUsageHour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LastBillableUsageHour != nil {
		if o.LastBillableUsageHour.Nanosecond() == 0 {
			toSerialize["last_billable_usage_hour"] = o.LastBillableUsageHour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_billable_usage_hour"] = o.LastBillableUsageHour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgBillableUsage != nil {
		toSerialize["org_billable_usage"] = o.OrgBillableUsage
	}
	if o.PercentageInAccount != nil {
		toSerialize["percentage_in_account"] = o.PercentageInAccount
	}
	if o.UsageUnit != nil {
		toSerialize["usage_unit"] = o.UsageUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageBillableSummaryBody) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountBillableUsage   *int64     `json:"account_billable_usage,omitempty"`
		AccountCommittedUsage  *int64     `json:"account_committed_usage,omitempty"`
		AccountOnDemandUsage   *int64     `json:"account_on_demand_usage,omitempty"`
		ElapsedUsageHours      *int64     `json:"elapsed_usage_hours,omitempty"`
		FirstBillableUsageHour *time.Time `json:"first_billable_usage_hour,omitempty"`
		LastBillableUsageHour  *time.Time `json:"last_billable_usage_hour,omitempty"`
		OrgBillableUsage       *int64     `json:"org_billable_usage,omitempty"`
		PercentageInAccount    *float64   `json:"percentage_in_account,omitempty"`
		UsageUnit              *string    `json:"usage_unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_billable_usage", "account_committed_usage", "account_on_demand_usage", "elapsed_usage_hours", "first_billable_usage_hour", "last_billable_usage_hour", "org_billable_usage", "percentage_in_account", "usage_unit"})
	} else {
		return err
	}
	o.AccountBillableUsage = all.AccountBillableUsage
	o.AccountCommittedUsage = all.AccountCommittedUsage
	o.AccountOnDemandUsage = all.AccountOnDemandUsage
	o.ElapsedUsageHours = all.ElapsedUsageHours
	o.FirstBillableUsageHour = all.FirstBillableUsageHour
	o.LastBillableUsageHour = all.LastBillableUsageHour
	o.OrgBillableUsage = all.OrgBillableUsage
	o.PercentageInAccount = all.PercentageInAccount
	o.UsageUnit = all.UsageUnit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
