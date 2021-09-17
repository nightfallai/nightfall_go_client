# ScanRequestV2ConfigConditionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalOp** | Pointer to **string** | Evaluate whether the conditions satisfy ANY or ALL criteria. Empty value defaults to ANY. | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) | A list of conditions. | [optional] 

## Methods

### NewScanRequestV2ConfigConditionSet

`func NewScanRequestV2ConfigConditionSet() *ScanRequestV2ConfigConditionSet`

NewScanRequestV2ConfigConditionSet instantiates a new ScanRequestV2ConfigConditionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanRequestV2ConfigConditionSetWithDefaults

`func NewScanRequestV2ConfigConditionSetWithDefaults() *ScanRequestV2ConfigConditionSet`

NewScanRequestV2ConfigConditionSetWithDefaults instantiates a new ScanRequestV2ConfigConditionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalOp

`func (o *ScanRequestV2ConfigConditionSet) GetLogicalOp() string`

GetLogicalOp returns the LogicalOp field if non-nil, zero value otherwise.

### GetLogicalOpOk

`func (o *ScanRequestV2ConfigConditionSet) GetLogicalOpOk() (*string, bool)`

GetLogicalOpOk returns a tuple with the LogicalOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOp

`func (o *ScanRequestV2ConfigConditionSet) SetLogicalOp(v string)`

SetLogicalOp sets LogicalOp field to given value.

### HasLogicalOp

`func (o *ScanRequestV2ConfigConditionSet) HasLogicalOp() bool`

HasLogicalOp returns a boolean if a field has been set.

### GetConditions

`func (o *ScanRequestV2ConfigConditionSet) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ScanRequestV2ConfigConditionSet) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ScanRequestV2ConfigConditionSet) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ScanRequestV2ConfigConditionSet) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


