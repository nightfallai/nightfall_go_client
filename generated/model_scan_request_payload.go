/*
 * Methods
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScanRequestPayload The text sample(s) you wish to scan. This data is passed as a string list, so you may choose to segment your text into multiple items for better granularity. The aggregate size of your text (summed across all items in the list) must not exceed 500 KB for any individual request, and the number of items in that list may not exceed 50,000.
type ScanRequestPayload struct {
	// A collection of strings to scan
	Items []string `json:"items,omitempty"`
}
