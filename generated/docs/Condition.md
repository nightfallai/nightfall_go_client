# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinNumFindings** | Pointer to **int32** | The minimum number of findings required in order for this detector to be reported. | [optional] 
**MinConfidence** | Pointer to [**Confidence**](Confidence.md) |  | [optional] 
**Detector** | Pointer to [**Detector**](Detector.md) |  | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinNumFindings

`func (o *Condition) GetMinNumFindings() int32`

GetMinNumFindings returns the MinNumFindings field if non-nil, zero value otherwise.

### GetMinNumFindingsOk

`func (o *Condition) GetMinNumFindingsOk() (*int32, bool)`

GetMinNumFindingsOk returns a tuple with the MinNumFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumFindings

`func (o *Condition) SetMinNumFindings(v int32)`

SetMinNumFindings sets MinNumFindings field to given value.

### HasMinNumFindings

`func (o *Condition) HasMinNumFindings() bool`

HasMinNumFindings returns a boolean if a field has been set.

### GetMinConfidence

`func (o *Condition) GetMinConfidence() Confidence`

GetMinConfidence returns the MinConfidence field if non-nil, zero value otherwise.

### GetMinConfidenceOk

`func (o *Condition) GetMinConfidenceOk() (*Confidence, bool)`

GetMinConfidenceOk returns a tuple with the MinConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConfidence

`func (o *Condition) SetMinConfidence(v Confidence)`

SetMinConfidence sets MinConfidence field to given value.

### HasMinConfidence

`func (o *Condition) HasMinConfidence() bool`

HasMinConfidence returns a boolean if a field has been set.

### GetDetector

`func (o *Condition) GetDetector() Detector`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *Condition) GetDetectorOk() (*Detector, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *Condition) SetDetector(v Detector)`

SetDetector sets Detector field to given value.

### HasDetector

`func (o *Condition) HasDetector() bool`

HasDetector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


