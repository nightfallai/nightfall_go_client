/*
 * api_platform
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Condition An inclusion condition within a condition set.
type Condition struct {
	// The minimum number of findings required in order for this detector to be reported.
	MinNumFindings int32 `json:"minNumFindings,omitempty"`
	MinConfidence Confidence `json:"minConfidence,omitempty"`
	Detector Detector `json:"detector,omitempty"`
}
