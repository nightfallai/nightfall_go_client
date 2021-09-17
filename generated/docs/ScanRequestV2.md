# ScanRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ScanRequestV2Config**](ScanRequestV2Config.md) |  | [optional] 
**Payload** | Pointer to **[]string** | The text sample(s) you wish to scan. This data is passed as a string list, so you may choose to segment your text into multiple items for better granularity. The aggregate size of your text (summed across all items in the list) must not exceed 500 KB for any individual request, and the number of items in that list may not exceed 50,000. | [optional] 

## Methods

### NewScanRequestV2

`func NewScanRequestV2() *ScanRequestV2`

NewScanRequestV2 instantiates a new ScanRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanRequestV2WithDefaults

`func NewScanRequestV2WithDefaults() *ScanRequestV2`

NewScanRequestV2WithDefaults instantiates a new ScanRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ScanRequestV2) GetConfig() ScanRequestV2Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ScanRequestV2) GetConfigOk() (*ScanRequestV2Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ScanRequestV2) SetConfig(v ScanRequestV2Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ScanRequestV2) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPayload

`func (o *ScanRequestV2) GetPayload() []string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ScanRequestV2) GetPayloadOk() (*[]string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ScanRequestV2) SetPayload(v []string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ScanRequestV2) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


