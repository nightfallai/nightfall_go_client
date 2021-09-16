# ExclusionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchType** | Pointer to [**MatchType**](MatchType.md) |  | [optional] 
**ExclusionType** | Pointer to [**ExclusionType**](ExclusionType.md) |  | [optional] 
**Regex** | Pointer to [**Regex**](Regex.md) |  | [optional] 
**WordList** | Pointer to [**WordList**](WordList.md) |  | [optional] 

## Methods

### NewExclusionRule

`func NewExclusionRule() *ExclusionRule`

NewExclusionRule instantiates a new ExclusionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExclusionRuleWithDefaults

`func NewExclusionRuleWithDefaults() *ExclusionRule`

NewExclusionRuleWithDefaults instantiates a new ExclusionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchType

`func (o *ExclusionRule) GetMatchType() MatchType`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *ExclusionRule) GetMatchTypeOk() (*MatchType, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *ExclusionRule) SetMatchType(v MatchType)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *ExclusionRule) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetExclusionType

`func (o *ExclusionRule) GetExclusionType() ExclusionType`

GetExclusionType returns the ExclusionType field if non-nil, zero value otherwise.

### GetExclusionTypeOk

`func (o *ExclusionRule) GetExclusionTypeOk() (*ExclusionType, bool)`

GetExclusionTypeOk returns a tuple with the ExclusionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionType

`func (o *ExclusionRule) SetExclusionType(v ExclusionType)`

SetExclusionType sets ExclusionType field to given value.

### HasExclusionType

`func (o *ExclusionRule) HasExclusionType() bool`

HasExclusionType returns a boolean if a field has been set.

### GetRegex

`func (o *ExclusionRule) GetRegex() Regex`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ExclusionRule) GetRegexOk() (*Regex, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ExclusionRule) SetRegex(v Regex)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ExclusionRule) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetWordList

`func (o *ExclusionRule) GetWordList() WordList`

GetWordList returns the WordList field if non-nil, zero value otherwise.

### GetWordListOk

`func (o *ExclusionRule) GetWordListOk() (*WordList, bool)`

GetWordListOk returns a tuple with the WordList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordList

`func (o *ExclusionRule) SetWordList(v WordList)`

SetWordList sets WordList field to given value.

### HasWordList

`func (o *ExclusionRule) HasWordList() bool`

HasWordList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


