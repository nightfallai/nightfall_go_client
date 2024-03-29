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

// Detector struct for Detector
type Detector struct {
	// The display name for this detector's findings in the response.
	DisplayName *string `json:"displayName,omitempty"`
	DetectorType *DetectorType `json:"detectorType,omitempty"`
	NightfallDetector *NightfallDetectorType `json:"nightfallDetector,omitempty"`
	Regex *Regex `json:"regex,omitempty"`
	WordList *WordList `json:"wordList,omitempty"`
	// A list of context rules.
	ContextRules *[]ContextRule `json:"contextRules,omitempty"`
	// A list of exclusion rules.
	ExclusionRules *[]ExclusionRule `json:"exclusionRules,omitempty"`
}

// NewDetector instantiates a new Detector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetector() *Detector {
	this := Detector{}
	return &this
}

// NewDetectorWithDefaults instantiates a new Detector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectorWithDefaults() *Detector {
	this := Detector{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Detector) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Detector) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Detector) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDetectorType returns the DetectorType field value if set, zero value otherwise.
func (o *Detector) GetDetectorType() DetectorType {
	if o == nil || o.DetectorType == nil {
		var ret DetectorType
		return ret
	}
	return *o.DetectorType
}

// GetDetectorTypeOk returns a tuple with the DetectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetDetectorTypeOk() (*DetectorType, bool) {
	if o == nil || o.DetectorType == nil {
		return nil, false
	}
	return o.DetectorType, true
}

// HasDetectorType returns a boolean if a field has been set.
func (o *Detector) HasDetectorType() bool {
	if o != nil && o.DetectorType != nil {
		return true
	}

	return false
}

// SetDetectorType gets a reference to the given DetectorType and assigns it to the DetectorType field.
func (o *Detector) SetDetectorType(v DetectorType) {
	o.DetectorType = &v
}

// GetNightfallDetector returns the NightfallDetector field value if set, zero value otherwise.
func (o *Detector) GetNightfallDetector() NightfallDetectorType {
	if o == nil || o.NightfallDetector == nil {
		var ret NightfallDetectorType
		return ret
	}
	return *o.NightfallDetector
}

// GetNightfallDetectorOk returns a tuple with the NightfallDetector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetNightfallDetectorOk() (*NightfallDetectorType, bool) {
	if o == nil || o.NightfallDetector == nil {
		return nil, false
	}
	return o.NightfallDetector, true
}

// HasNightfallDetector returns a boolean if a field has been set.
func (o *Detector) HasNightfallDetector() bool {
	if o != nil && o.NightfallDetector != nil {
		return true
	}

	return false
}

// SetNightfallDetector gets a reference to the given NightfallDetectorType and assigns it to the NightfallDetector field.
func (o *Detector) SetNightfallDetector(v NightfallDetectorType) {
	o.NightfallDetector = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *Detector) GetRegex() Regex {
	if o == nil || o.Regex == nil {
		var ret Regex
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetRegexOk() (*Regex, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *Detector) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// SetRegex gets a reference to the given Regex and assigns it to the Regex field.
func (o *Detector) SetRegex(v Regex) {
	o.Regex = &v
}

// GetWordList returns the WordList field value if set, zero value otherwise.
func (o *Detector) GetWordList() WordList {
	if o == nil || o.WordList == nil {
		var ret WordList
		return ret
	}
	return *o.WordList
}

// GetWordListOk returns a tuple with the WordList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetWordListOk() (*WordList, bool) {
	if o == nil || o.WordList == nil {
		return nil, false
	}
	return o.WordList, true
}

// HasWordList returns a boolean if a field has been set.
func (o *Detector) HasWordList() bool {
	if o != nil && o.WordList != nil {
		return true
	}

	return false
}

// SetWordList gets a reference to the given WordList and assigns it to the WordList field.
func (o *Detector) SetWordList(v WordList) {
	o.WordList = &v
}

// GetContextRules returns the ContextRules field value if set, zero value otherwise.
func (o *Detector) GetContextRules() []ContextRule {
	if o == nil || o.ContextRules == nil {
		var ret []ContextRule
		return ret
	}
	return *o.ContextRules
}

// GetContextRulesOk returns a tuple with the ContextRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetContextRulesOk() (*[]ContextRule, bool) {
	if o == nil || o.ContextRules == nil {
		return nil, false
	}
	return o.ContextRules, true
}

// HasContextRules returns a boolean if a field has been set.
func (o *Detector) HasContextRules() bool {
	if o != nil && o.ContextRules != nil {
		return true
	}

	return false
}

// SetContextRules gets a reference to the given []ContextRule and assigns it to the ContextRules field.
func (o *Detector) SetContextRules(v []ContextRule) {
	o.ContextRules = &v
}

// GetExclusionRules returns the ExclusionRules field value if set, zero value otherwise.
func (o *Detector) GetExclusionRules() []ExclusionRule {
	if o == nil || o.ExclusionRules == nil {
		var ret []ExclusionRule
		return ret
	}
	return *o.ExclusionRules
}

// GetExclusionRulesOk returns a tuple with the ExclusionRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Detector) GetExclusionRulesOk() (*[]ExclusionRule, bool) {
	if o == nil || o.ExclusionRules == nil {
		return nil, false
	}
	return o.ExclusionRules, true
}

// HasExclusionRules returns a boolean if a field has been set.
func (o *Detector) HasExclusionRules() bool {
	if o != nil && o.ExclusionRules != nil {
		return true
	}

	return false
}

// SetExclusionRules gets a reference to the given []ExclusionRule and assigns it to the ExclusionRules field.
func (o *Detector) SetExclusionRules(v []ExclusionRule) {
	o.ExclusionRules = &v
}

func (o Detector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DetectorType != nil {
		toSerialize["detectorType"] = o.DetectorType
	}
	if o.NightfallDetector != nil {
		toSerialize["nightfallDetector"] = o.NightfallDetector
	}
	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}
	if o.WordList != nil {
		toSerialize["wordList"] = o.WordList
	}
	if o.ContextRules != nil {
		toSerialize["contextRules"] = o.ContextRules
	}
	if o.ExclusionRules != nil {
		toSerialize["exclusionRules"] = o.ExclusionRules
	}
	return json.Marshal(toSerialize)
}

type NullableDetector struct {
	value *Detector
	isSet bool
}

func (v NullableDetector) Get() *Detector {
	return v.value
}

func (v *NullableDetector) Set(val *Detector) {
	v.value = val
	v.isSet = true
}

func (v NullableDetector) IsSet() bool {
	return v.isSet
}

func (v *NullableDetector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetector(val *Detector) *NullableDetector {
	return &NullableDetector{value: val, isSet: true}
}

func (v NullableDetector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


