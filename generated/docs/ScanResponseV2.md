# ScanResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fragment** | Pointer to **string** | The text sample that was flagged. | [optional] 
**DetectorName** | Pointer to **string** | The data type flagged in the text fragment. | [optional] 
**Confidence** | Pointer to [**Confidence**](Confidence.md) |  | [optional] 
**Location** | Pointer to [**ScanResponseLocation**](ScanResponseLocation.md) |  | [optional] 

## Methods

### NewScanResponseV2

`func NewScanResponseV2() *ScanResponseV2`

NewScanResponseV2 instantiates a new ScanResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanResponseV2WithDefaults

`func NewScanResponseV2WithDefaults() *ScanResponseV2`

NewScanResponseV2WithDefaults instantiates a new ScanResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFragment

`func (o *ScanResponseV2) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *ScanResponseV2) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *ScanResponseV2) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *ScanResponseV2) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetDetectorName

`func (o *ScanResponseV2) GetDetectorName() string`

GetDetectorName returns the DetectorName field if non-nil, zero value otherwise.

### GetDetectorNameOk

`func (o *ScanResponseV2) GetDetectorNameOk() (*string, bool)`

GetDetectorNameOk returns a tuple with the DetectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectorName

`func (o *ScanResponseV2) SetDetectorName(v string)`

SetDetectorName sets DetectorName field to given value.

### HasDetectorName

`func (o *ScanResponseV2) HasDetectorName() bool`

HasDetectorName returns a boolean if a field has been set.

### GetConfidence

`func (o *ScanResponseV2) GetConfidence() Confidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ScanResponseV2) GetConfidenceOk() (*Confidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ScanResponseV2) SetConfidence(v Confidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ScanResponseV2) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLocation

`func (o *ScanResponseV2) GetLocation() ScanResponseLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ScanResponseV2) GetLocationOk() (*ScanResponseLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ScanResponseV2) SetLocation(v ScanResponseLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ScanResponseV2) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


