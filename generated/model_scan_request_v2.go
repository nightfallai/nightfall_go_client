/*
 * Methods
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScanRequestV2 The request body of the v2/scan endpoint
type ScanRequestV2 struct {
	Config ScanRequestV2Config `json:"config,omitempty"`
	Payload ScanRequestPayload `json:"payload,omitempty"`
}
