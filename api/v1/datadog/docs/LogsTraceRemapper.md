# LogsTraceRemapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Array of source attributes. | [optional] [default to ["dd.trace_id"]]
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "trace-id-remapper"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsTraceRemapper

`func NewLogsTraceRemapper() *LogsTraceRemapper`

NewLogsTraceRemapper instantiates a new LogsTraceRemapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsTraceRemapperWithDefaults

`func NewLogsTraceRemapperWithDefaults() *LogsTraceRemapper`

NewLogsTraceRemapperWithDefaults instantiates a new LogsTraceRemapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *LogsTraceRemapper) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsTraceRemapper) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *LogsTraceRemapper) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *LogsTraceRemapper) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetType

`func (o *LogsTraceRemapper) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsTraceRemapper) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsTraceRemapper) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsTraceRemapper) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsTraceRemapper) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsTraceRemapper) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsTraceRemapper) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsTraceRemapper) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsTraceRemapper) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsTraceRemapper) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsTraceRemapper) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsTraceRemapper) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsTraceRemapper) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsTraceRemapper in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

