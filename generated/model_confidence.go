/*
Methods

This API exposes detectors for sensitive data in arbitrary string payloads.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Confidence The confidence level of a finding.
type Confidence string

// List of Confidence
const (
	CONFIDENCE_VERY_UNLIKELY Confidence = "VERY_UNLIKELY"
	CONFIDENCE_UNLIKELY Confidence = "UNLIKELY"
	CONFIDENCE_POSSIBLE Confidence = "POSSIBLE"
	CONFIDENCE_LIKELY Confidence = "LIKELY"
	CONFIDENCE_VERY_LIKELY Confidence = "VERY_LIKELY"
)

var allowedConfidenceEnumValues = []Confidence{
	"VERY_UNLIKELY",
	"UNLIKELY",
	"POSSIBLE",
	"LIKELY",
	"VERY_LIKELY",
}

func (v *Confidence) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Confidence(value)
	for _, existing := range allowedConfidenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Confidence", value)
}

// NewConfidenceFromValue returns a pointer to a valid Confidence
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfidenceFromValue(v string) (*Confidence, error) {
	ev := Confidence(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Confidence: valid values are %v", v, allowedConfidenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Confidence) IsValid() bool {
	for _, existing := range allowedConfidenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Confidence value
func (v Confidence) Ptr() *Confidence {
	return &v
}

type NullableConfidence struct {
	value *Confidence
	isSet bool
}

func (v NullableConfidence) Get() *Confidence {
	return v.value
}

func (v *NullableConfidence) Set(val *Confidence) {
	v.value = val
	v.isSet = true
}

func (v NullableConfidence) IsSet() bool {
	return v.isSet
}

func (v *NullableConfidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfidence(val *Confidence) *NullableConfidence {
	return &NullableConfidence{value: val, isSet: true}
}

func (v NullableConfidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

