/*
 * Methods
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContextRuleProximity The object containing the length of characters before and after finding to evaluate context.
type ContextRuleProximity struct {
	// The number of leading characters to include as context before the finding itself.
	WindowBefore int32 `json:"windowBefore,omitempty"`
	// The number of trailing characters to include as context after the finding itself.
	WindowAfter int32 `json:"windowAfter,omitempty"`
}