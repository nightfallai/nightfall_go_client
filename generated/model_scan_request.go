/*
 * Methods
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScanRequest The request body of the /v1/scan endpoint
type ScanRequest struct {
	// A list of detector objects with which you wish to scan the request payload.
	Detectors []ScanRequestDetectors `json:"detectors,omitempty"`
	Payload ScanRequestPayload `json:"payload,omitempty"`
}