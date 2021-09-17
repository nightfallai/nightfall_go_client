# ContextRuleProximity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowBefore** | Pointer to **int32** | The number of leading characters to include as context before the finding itself. | [optional] 
**WindowAfter** | Pointer to **int32** | The number of trailing characters to include as context after the finding itself. | [optional] 

## Methods

### NewContextRuleProximity

`func NewContextRuleProximity() *ContextRuleProximity`

NewContextRuleProximity instantiates a new ContextRuleProximity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRuleProximityWithDefaults

`func NewContextRuleProximityWithDefaults() *ContextRuleProximity`

NewContextRuleProximityWithDefaults instantiates a new ContextRuleProximity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowBefore

`func (o *ContextRuleProximity) GetWindowBefore() int32`

GetWindowBefore returns the WindowBefore field if non-nil, zero value otherwise.

### GetWindowBeforeOk

`func (o *ContextRuleProximity) GetWindowBeforeOk() (*int32, bool)`

GetWindowBeforeOk returns a tuple with the WindowBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowBefore

`func (o *ContextRuleProximity) SetWindowBefore(v int32)`

SetWindowBefore sets WindowBefore field to given value.

### HasWindowBefore

`func (o *ContextRuleProximity) HasWindowBefore() bool`

HasWindowBefore returns a boolean if a field has been set.

### GetWindowAfter

`func (o *ContextRuleProximity) GetWindowAfter() int32`

GetWindowAfter returns the WindowAfter field if non-nil, zero value otherwise.

### GetWindowAfterOk

`func (o *ContextRuleProximity) GetWindowAfterOk() (*int32, bool)`

GetWindowAfterOk returns a tuple with the WindowAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowAfter

`func (o *ContextRuleProximity) SetWindowAfter(v int32)`

SetWindowAfter sets WindowAfter field to given value.

### HasWindowAfter

`func (o *ContextRuleProximity) HasWindowAfter() bool`

HasWindowAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


