/*
 * api_platform
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScanResponse struct for ScanResponse
type ScanResponse struct {
	// The text sample that was flagged.
	Fragment string `json:"fragment,omitempty"`
	// The data type flagged in the text fragment.
	Detector string `json:"detector,omitempty"`
	Confidence ScanResponseConfidence `json:"confidence,omitempty"`
	Location ScanResponseLocation `json:"location,omitempty"`
}
