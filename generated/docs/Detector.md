# Detector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name for this detector&#39;s findings in the response. | [optional] 
**DetectorType** | Pointer to [**DetectorType**](DetectorType.md) |  | [optional] 
**NightfallDetector** | Pointer to [**NightfallDetectorType**](NightfallDetectorType.md) |  | [optional] 
**Regex** | Pointer to [**Regex**](Regex.md) |  | [optional] 
**WordList** | Pointer to [**WordList**](WordList.md) |  | [optional] 
**ContextRules** | Pointer to [**[]ContextRule**](ContextRule.md) | A list of context rules. | [optional] 
**ExclusionRules** | Pointer to [**[]ExclusionRule**](ExclusionRule.md) | A list of exclusion rules. | [optional] 

## Methods

### NewDetector

`func NewDetector() *Detector`

NewDetector instantiates a new Detector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetectorWithDefaults

`func NewDetectorWithDefaults() *Detector`

NewDetectorWithDefaults instantiates a new Detector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *Detector) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Detector) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Detector) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Detector) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDetectorType

`func (o *Detector) GetDetectorType() DetectorType`

GetDetectorType returns the DetectorType field if non-nil, zero value otherwise.

### GetDetectorTypeOk

`func (o *Detector) GetDetectorTypeOk() (*DetectorType, bool)`

GetDetectorTypeOk returns a tuple with the DetectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectorType

`func (o *Detector) SetDetectorType(v DetectorType)`

SetDetectorType sets DetectorType field to given value.

### HasDetectorType

`func (o *Detector) HasDetectorType() bool`

HasDetectorType returns a boolean if a field has been set.

### GetNightfallDetector

`func (o *Detector) GetNightfallDetector() NightfallDetectorType`

GetNightfallDetector returns the NightfallDetector field if non-nil, zero value otherwise.

### GetNightfallDetectorOk

`func (o *Detector) GetNightfallDetectorOk() (*NightfallDetectorType, bool)`

GetNightfallDetectorOk returns a tuple with the NightfallDetector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightfallDetector

`func (o *Detector) SetNightfallDetector(v NightfallDetectorType)`

SetNightfallDetector sets NightfallDetector field to given value.

### HasNightfallDetector

`func (o *Detector) HasNightfallDetector() bool`

HasNightfallDetector returns a boolean if a field has been set.

### GetRegex

`func (o *Detector) GetRegex() Regex`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *Detector) GetRegexOk() (*Regex, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *Detector) SetRegex(v Regex)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *Detector) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetWordList

`func (o *Detector) GetWordList() WordList`

GetWordList returns the WordList field if non-nil, zero value otherwise.

### GetWordListOk

`func (o *Detector) GetWordListOk() (*WordList, bool)`

GetWordListOk returns a tuple with the WordList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordList

`func (o *Detector) SetWordList(v WordList)`

SetWordList sets WordList field to given value.

### HasWordList

`func (o *Detector) HasWordList() bool`

HasWordList returns a boolean if a field has been set.

### GetContextRules

`func (o *Detector) GetContextRules() []ContextRule`

GetContextRules returns the ContextRules field if non-nil, zero value otherwise.

### GetContextRulesOk

`func (o *Detector) GetContextRulesOk() (*[]ContextRule, bool)`

GetContextRulesOk returns a tuple with the ContextRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRules

`func (o *Detector) SetContextRules(v []ContextRule)`

SetContextRules sets ContextRules field to given value.

### HasContextRules

`func (o *Detector) HasContextRules() bool`

HasContextRules returns a boolean if a field has been set.

### GetExclusionRules

`func (o *Detector) GetExclusionRules() []ExclusionRule`

GetExclusionRules returns the ExclusionRules field if non-nil, zero value otherwise.

### GetExclusionRulesOk

`func (o *Detector) GetExclusionRulesOk() (*[]ExclusionRule, bool)`

GetExclusionRulesOk returns a tuple with the ExclusionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionRules

`func (o *Detector) SetExclusionRules(v []ExclusionRule)`

SetExclusionRules sets ExclusionRules field to given value.

### HasExclusionRules

`func (o *Detector) HasExclusionRules() bool`

HasExclusionRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


