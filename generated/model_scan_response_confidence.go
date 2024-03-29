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

// ScanResponseConfidence struct for ScanResponseConfidence
type ScanResponseConfidence struct {
	Bucket *Confidence `json:"bucket,omitempty"`
}

// NewScanResponseConfidence instantiates a new ScanResponseConfidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanResponseConfidence() *ScanResponseConfidence {
	this := ScanResponseConfidence{}
	return &this
}

// NewScanResponseConfidenceWithDefaults instantiates a new ScanResponseConfidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanResponseConfidenceWithDefaults() *ScanResponseConfidence {
	this := ScanResponseConfidence{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ScanResponseConfidence) GetBucket() Confidence {
	if o == nil || o.Bucket == nil {
		var ret Confidence
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanResponseConfidence) GetBucketOk() (*Confidence, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ScanResponseConfidence) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given Confidence and assigns it to the Bucket field.
func (o *ScanResponseConfidence) SetBucket(v Confidence) {
	o.Bucket = &v
}

func (o ScanResponseConfidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	return json.Marshal(toSerialize)
}

type NullableScanResponseConfidence struct {
	value *ScanResponseConfidence
	isSet bool
}

func (v NullableScanResponseConfidence) Get() *ScanResponseConfidence {
	return v.value
}

func (v *NullableScanResponseConfidence) Set(val *ScanResponseConfidence) {
	v.value = val
	v.isSet = true
}

func (v NullableScanResponseConfidence) IsSet() bool {
	return v.isSet
}

func (v *NullableScanResponseConfidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanResponseConfidence(val *ScanResponseConfidence) *NullableScanResponseConfidence {
	return &NullableScanResponseConfidence{value: val, isSet: true}
}

func (v NullableScanResponseConfidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanResponseConfidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


