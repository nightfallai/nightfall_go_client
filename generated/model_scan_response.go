/*
Methods

This API exposes detectors for sensitive data in arbitrary string payloads.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ScanResponse struct for ScanResponse
type ScanResponse struct {
	// The text sample that was flagged.
	Fragment *string `json:"fragment,omitempty"`
	// The data type flagged in the text fragment.
	Detector *string `json:"detector,omitempty"`
	Confidence *ScanResponseConfidence `json:"confidence,omitempty"`
	Location *ScanResponseLocation `json:"location,omitempty"`
}

// NewScanResponse instantiates a new ScanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanResponse() *ScanResponse {
	this := ScanResponse{}
	return &this
}

// NewScanResponseWithDefaults instantiates a new ScanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanResponseWithDefaults() *ScanResponse {
	this := ScanResponse{}
	return &this
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *ScanResponse) GetFragment() string {
	if o == nil || o.Fragment == nil {
		var ret string
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanResponse) GetFragmentOk() (*string, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *ScanResponse) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given string and assigns it to the Fragment field.
func (o *ScanResponse) SetFragment(v string) {
	o.Fragment = &v
}

// GetDetector returns the Detector field value if set, zero value otherwise.
func (o *ScanResponse) GetDetector() string {
	if o == nil || o.Detector == nil {
		var ret string
		return ret
	}
	return *o.Detector
}

// GetDetectorOk returns a tuple with the Detector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanResponse) GetDetectorOk() (*string, bool) {
	if o == nil || o.Detector == nil {
		return nil, false
	}
	return o.Detector, true
}

// HasDetector returns a boolean if a field has been set.
func (o *ScanResponse) HasDetector() bool {
	if o != nil && o.Detector != nil {
		return true
	}

	return false
}

// SetDetector gets a reference to the given string and assigns it to the Detector field.
func (o *ScanResponse) SetDetector(v string) {
	o.Detector = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *ScanResponse) GetConfidence() ScanResponseConfidence {
	if o == nil || o.Confidence == nil {
		var ret ScanResponseConfidence
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanResponse) GetConfidenceOk() (*ScanResponseConfidence, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *ScanResponse) HasConfidence() bool {
	if o != nil && o.Confidence != nil {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given ScanResponseConfidence and assigns it to the Confidence field.
func (o *ScanResponse) SetConfidence(v ScanResponseConfidence) {
	o.Confidence = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ScanResponse) GetLocation() ScanResponseLocation {
	if o == nil || o.Location == nil {
		var ret ScanResponseLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanResponse) GetLocationOk() (*ScanResponseLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ScanResponse) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ScanResponseLocation and assigns it to the Location field.
func (o *ScanResponse) SetLocation(v ScanResponseLocation) {
	o.Location = &v
}

func (o ScanResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fragment != nil {
		toSerialize["fragment"] = o.Fragment
	}
	if o.Detector != nil {
		toSerialize["detector"] = o.Detector
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableScanResponse struct {
	value *ScanResponse
	isSet bool
}

func (v NullableScanResponse) Get() *ScanResponse {
	return v.value
}

func (v *NullableScanResponse) Set(val *ScanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanResponse(val *ScanResponse) *NullableScanResponse {
	return &NullableScanResponse{value: val, isSet: true}
}

func (v NullableScanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


