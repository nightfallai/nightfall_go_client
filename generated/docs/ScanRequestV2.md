# ScanRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**ScanRequestV2Config**](ScanRequestV2_config.md) |  | [optional] 
**Payload** | **[]string** | The text sample(s) you wish to scan. This data is passed as a string list, so you may choose to segment your text into multiple items for better granularity. The aggregate size of your text (summed across all items in the list) must not exceed 500 KB for any individual request, and the number of items in that list may not exceed 50,000. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


