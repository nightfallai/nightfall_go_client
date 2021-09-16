# Regex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | The regex pattern to match on. | [optional] 
**IsCaseSensitive** | Pointer to **bool** | The case sensitivity for the regex pattern. | [optional] 

## Methods

### NewRegex

`func NewRegex() *Regex`

NewRegex instantiates a new Regex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegexWithDefaults

`func NewRegexWithDefaults() *Regex`

NewRegexWithDefaults instantiates a new Regex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *Regex) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Regex) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Regex) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *Regex) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetIsCaseSensitive

`func (o *Regex) GetIsCaseSensitive() bool`

GetIsCaseSensitive returns the IsCaseSensitive field if non-nil, zero value otherwise.

### GetIsCaseSensitiveOk

`func (o *Regex) GetIsCaseSensitiveOk() (*bool, bool)`

GetIsCaseSensitiveOk returns a tuple with the IsCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaseSensitive

`func (o *Regex) SetIsCaseSensitive(v bool)`

SetIsCaseSensitive sets IsCaseSensitive field to given value.

### HasIsCaseSensitive

`func (o *Regex) HasIsCaseSensitive() bool`

HasIsCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


