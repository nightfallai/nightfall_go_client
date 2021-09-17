# ScanRequestV2Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionSetUUID** | Pointer to **string** | The UUID corresponding to the condition set with which you wish to scan the request payload. | [optional] 
**ConditionSet** | Pointer to [**ScanRequestV2ConfigConditionSet**](ScanRequestV2ConfigConditionSet.md) |  | [optional] 

## Methods

### NewScanRequestV2Config

`func NewScanRequestV2Config() *ScanRequestV2Config`

NewScanRequestV2Config instantiates a new ScanRequestV2Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanRequestV2ConfigWithDefaults

`func NewScanRequestV2ConfigWithDefaults() *ScanRequestV2Config`

NewScanRequestV2ConfigWithDefaults instantiates a new ScanRequestV2Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionSetUUID

`func (o *ScanRequestV2Config) GetConditionSetUUID() string`

GetConditionSetUUID returns the ConditionSetUUID field if non-nil, zero value otherwise.

### GetConditionSetUUIDOk

`func (o *ScanRequestV2Config) GetConditionSetUUIDOk() (*string, bool)`

GetConditionSetUUIDOk returns a tuple with the ConditionSetUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSetUUID

`func (o *ScanRequestV2Config) SetConditionSetUUID(v string)`

SetConditionSetUUID sets ConditionSetUUID field to given value.

### HasConditionSetUUID

`func (o *ScanRequestV2Config) HasConditionSetUUID() bool`

HasConditionSetUUID returns a boolean if a field has been set.

### GetConditionSet

`func (o *ScanRequestV2Config) GetConditionSet() ScanRequestV2ConfigConditionSet`

GetConditionSet returns the ConditionSet field if non-nil, zero value otherwise.

### GetConditionSetOk

`func (o *ScanRequestV2Config) GetConditionSetOk() (*ScanRequestV2ConfigConditionSet, bool)`

GetConditionSetOk returns a tuple with the ConditionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSet

`func (o *ScanRequestV2Config) SetConditionSet(v ScanRequestV2ConfigConditionSet)`

SetConditionSet sets ConditionSet field to given value.

### HasConditionSet

`func (o *ScanRequestV2Config) HasConditionSet() bool`

HasConditionSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


