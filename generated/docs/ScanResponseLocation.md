# ScanResponseLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByteRange** | Pointer to [**ScanResponseLocationByteRange**](ScanResponseLocationByteRange.md) |  | [optional] 
**UnicodeRange** | Pointer to [**ScanResponseLocationUnicodeRange**](ScanResponseLocationUnicodeRange.md) |  | [optional] 

## Methods

### NewScanResponseLocation

`func NewScanResponseLocation() *ScanResponseLocation`

NewScanResponseLocation instantiates a new ScanResponseLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanResponseLocationWithDefaults

`func NewScanResponseLocationWithDefaults() *ScanResponseLocation`

NewScanResponseLocationWithDefaults instantiates a new ScanResponseLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByteRange

`func (o *ScanResponseLocation) GetByteRange() ScanResponseLocationByteRange`

GetByteRange returns the ByteRange field if non-nil, zero value otherwise.

### GetByteRangeOk

`func (o *ScanResponseLocation) GetByteRangeOk() (*ScanResponseLocationByteRange, bool)`

GetByteRangeOk returns a tuple with the ByteRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteRange

`func (o *ScanResponseLocation) SetByteRange(v ScanResponseLocationByteRange)`

SetByteRange sets ByteRange field to given value.

### HasByteRange

`func (o *ScanResponseLocation) HasByteRange() bool`

HasByteRange returns a boolean if a field has been set.

### GetUnicodeRange

`func (o *ScanResponseLocation) GetUnicodeRange() ScanResponseLocationUnicodeRange`

GetUnicodeRange returns the UnicodeRange field if non-nil, zero value otherwise.

### GetUnicodeRangeOk

`func (o *ScanResponseLocation) GetUnicodeRangeOk() (*ScanResponseLocationUnicodeRange, bool)`

GetUnicodeRangeOk returns a tuple with the UnicodeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicodeRange

`func (o *ScanResponseLocation) SetUnicodeRange(v ScanResponseLocationUnicodeRange)`

SetUnicodeRange sets UnicodeRange field to given value.

### HasUnicodeRange

`func (o *ScanResponseLocation) HasUnicodeRange() bool`

HasUnicodeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


