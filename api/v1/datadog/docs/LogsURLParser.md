# LogsURLParser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to **[]string** | Array of source attributes. | [default to ["http.url"]]
**Target** | Pointer to **string** | Name of the parent attribute that contains all the extracted details from the &#x60;sources&#x60;. | [default to "http.url_details"]
**Type** | Pointer to **string** | Type of processor. | [optional] [readonly] [default to "url-parser"]
**IsEnabled** | Pointer to **bool** | Whether or not the processor is enabled. | [optional] [default to false]
**Name** | Pointer to **string** | Name of the processor. | [optional] 

## Methods

### NewLogsURLParser

`func NewLogsURLParser(sources []string, target string, ) *LogsURLParser`

NewLogsURLParser instantiates a new LogsURLParser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsURLParserWithDefaults

`func NewLogsURLParserWithDefaults() *LogsURLParser`

NewLogsURLParserWithDefaults instantiates a new LogsURLParser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *LogsURLParser) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *LogsURLParser) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *LogsURLParser) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *LogsURLParser) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetTarget

`func (o *LogsURLParser) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LogsURLParser) GetTargetOk() (string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *LogsURLParser) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *LogsURLParser) SetTarget(v string)`

SetTarget gets a reference to the given string and assigns it to the Target field.

### GetType

`func (o *LogsURLParser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsURLParser) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *LogsURLParser) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *LogsURLParser) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetIsEnabled

`func (o *LogsURLParser) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LogsURLParser) GetIsEnabledOk() (bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabled

`func (o *LogsURLParser) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabled

`func (o *LogsURLParser) SetIsEnabled(v bool)`

SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.

### GetName

`func (o *LogsURLParser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogsURLParser) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LogsURLParser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LogsURLParser) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


### AsLogsProcessor

`func (s *LogsURLParser) AsLogsProcessor() LogsProcessor`

Convenience method to wrap this instance of LogsURLParser in LogsProcessor

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

