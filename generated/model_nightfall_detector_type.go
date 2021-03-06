/*
 * Methods
 *
 * This API exposes detectors for sensitive data in arbitrary string payloads.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NightfallDetectorType The name for a Nightfall detector.
type NightfallDetectorType string

// List of NightfallDetectorType
const (
	NIGHTFALLDETECTORTYPE_API_KEY NightfallDetectorType = "API_KEY"
	NIGHTFALLDETECTORTYPE_RANDOMLY_GENERATED_TOKEN NightfallDetectorType = "RANDOMLY_GENERATED_TOKEN"
	NIGHTFALLDETECTORTYPE_CRYPTOGRAPHIC_KEY NightfallDetectorType = "CRYPTOGRAPHIC_KEY"
	NIGHTFALLDETECTORTYPE_CREDIT_CARD_NUMBER NightfallDetectorType = "CREDIT_CARD_NUMBER"
	NIGHTFALLDETECTORTYPE_US_SOCIAL_SECURITY_NUMBER NightfallDetectorType = "US_SOCIAL_SECURITY_NUMBER"
	NIGHTFALLDETECTORTYPE_AMERICAN_BANKERS_CUSIP_ID NightfallDetectorType = "AMERICAN_BANKERS_CUSIP_ID"
	NIGHTFALLDETECTORTYPE_US_BANK_ROUTING_MICR NightfallDetectorType = "US_BANK_ROUTING_MICR"
	NIGHTFALLDETECTORTYPE_ICD9_CODE NightfallDetectorType = "ICD9_CODE"
	NIGHTFALLDETECTORTYPE_ICD10_CODE NightfallDetectorType = "ICD10_CODE"
	NIGHTFALLDETECTORTYPE_US_DRIVERS_LICENSE_NUMBER NightfallDetectorType = "US_DRIVERS_LICENSE_NUMBER"
	NIGHTFALLDETECTORTYPE_US_PASSPORT NightfallDetectorType = "US_PASSPORT"
	NIGHTFALLDETECTORTYPE_PHONE_NUMBER NightfallDetectorType = "PHONE_NUMBER"
	NIGHTFALLDETECTORTYPE_IP_ADDRESS NightfallDetectorType = "IP_ADDRESS"
	NIGHTFALLDETECTORTYPE_EMAIL_ADDRESS NightfallDetectorType = "EMAIL_ADDRESS"
)
