/*
 * Methods
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ContextRuleConfidenceAdjustment The confidence to adjust finding to on found context.
type ContextRuleConfidenceAdjustment struct {
	FixedConfidence Confidence `json:"fixedConfidence,omitempty"`
}
