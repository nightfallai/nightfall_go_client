# ContextRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regex** | Pointer to [**Regex**](Regex.md) |  | [optional] 
**Proximity** | Pointer to [**ContextRuleProximity**](ContextRuleProximity.md) |  | [optional] 
**ConfidenceAdjustment** | Pointer to [**ContextRuleConfidenceAdjustment**](ContextRuleConfidenceAdjustment.md) |  | [optional] 

## Methods

### NewContextRule

`func NewContextRule() *ContextRule`

NewContextRule instantiates a new ContextRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRuleWithDefaults

`func NewContextRuleWithDefaults() *ContextRule`

NewContextRuleWithDefaults instantiates a new ContextRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegex

`func (o *ContextRule) GetRegex() Regex`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ContextRule) GetRegexOk() (*Regex, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ContextRule) SetRegex(v Regex)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ContextRule) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetProximity

`func (o *ContextRule) GetProximity() ContextRuleProximity`

GetProximity returns the Proximity field if non-nil, zero value otherwise.

### GetProximityOk

`func (o *ContextRule) GetProximityOk() (*ContextRuleProximity, bool)`

GetProximityOk returns a tuple with the Proximity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximity

`func (o *ContextRule) SetProximity(v ContextRuleProximity)`

SetProximity sets Proximity field to given value.

### HasProximity

`func (o *ContextRule) HasProximity() bool`

HasProximity returns a boolean if a field has been set.

### GetConfidenceAdjustment

`func (o *ContextRule) GetConfidenceAdjustment() ContextRuleConfidenceAdjustment`

GetConfidenceAdjustment returns the ConfidenceAdjustment field if non-nil, zero value otherwise.

### GetConfidenceAdjustmentOk

`func (o *ContextRule) GetConfidenceAdjustmentOk() (*ContextRuleConfidenceAdjustment, bool)`

GetConfidenceAdjustmentOk returns a tuple with the ConfidenceAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceAdjustment

`func (o *ContextRule) SetConfidenceAdjustment(v ContextRuleConfidenceAdjustment)`

SetConfidenceAdjustment sets ConfidenceAdjustment field to given value.

### HasConfidenceAdjustment

`func (o *ContextRule) HasConfidenceAdjustment() bool`

HasConfidenceAdjustment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


