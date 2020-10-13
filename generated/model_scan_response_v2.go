/*
 * api_platform
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScanResponseV2 struct for ScanResponseV2
type ScanResponseV2 struct {
	// The text sample that was flagged.
	Fragment string `json:"fragment,omitempty"`
	// The data type flagged in the text fragment.
	DetectorName string `json:"detectorName,omitempty"`
	Confidence Confidence `json:"confidence,omitempty"`
	Location ScanResponseLocation `json:"location,omitempty"`
}