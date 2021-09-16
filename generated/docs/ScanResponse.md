# ScanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **string** | The text sample that was flagged. | [optional] 
**Detector** | Pointer to **string** | The data type flagged in the text fragment. | [optional] 
**Confidence** | Pointer to [**ScanResponseConfidence**](ScanResponseConfidence.md) |  | [optional] 
**Location** | Pointer to [**ScanResponseLocation**](ScanResponseLocation.md) |  | [optional] 

## Methods

### NewScanResponse

`func NewScanResponse() *ScanResponse`

NewScanResponse instantiates a new ScanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanResponseWithDefaults

`func NewScanResponseWithDefaults() *ScanResponse`

NewScanResponseWithDefaults instantiates a new ScanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *ScanResponse) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *ScanResponse) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *ScanResponse) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *ScanResponse) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetDetector

`func (o *ScanResponse) GetDetector() string`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *ScanResponse) GetDetectorOk() (*string, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *ScanResponse) SetDetector(v string)`

SetDetector sets Detector field to given value.

### HasDetector

`func (o *ScanResponse) HasDetector() bool`

HasDetector returns a boolean if a field has been set.

### GetConfidence

`func (o *ScanResponse) GetConfidence() ScanResponseConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ScanResponse) GetConfidenceOk() (*ScanResponseConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ScanResponse) SetConfidence(v ScanResponseConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ScanResponse) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLocation

`func (o *ScanResponse) GetLocation() ScanResponseLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ScanResponse) GetLocationOk() (*ScanResponseLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ScanResponse) SetLocation(v ScanResponseLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ScanResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


