# WordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **[]string** | A list of words for wordList. | [optional] 
**IsCaseSensitive** | Pointer to **bool** | The case sensitivity for words in the wordList. If false, ignore the case of findings. | [optional] 

## Methods

### NewWordList

`func NewWordList() *WordList`

NewWordList instantiates a new WordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWordListWithDefaults

`func NewWordListWithDefaults() *WordList`

NewWordListWithDefaults instantiates a new WordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *WordList) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WordList) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WordList) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *WordList) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetIsCaseSensitive

`func (o *WordList) GetIsCaseSensitive() bool`

GetIsCaseSensitive returns the IsCaseSensitive field if non-nil, zero value otherwise.

### GetIsCaseSensitiveOk

`func (o *WordList) GetIsCaseSensitiveOk() (*bool, bool)`

GetIsCaseSensitiveOk returns a tuple with the IsCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaseSensitive

`func (o *WordList) SetIsCaseSensitive(v bool)`

SetIsCaseSensitive sets IsCaseSensitive field to given value.

### HasIsCaseSensitive

`func (o *WordList) HasIsCaseSensitive() bool`

HasIsCaseSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


