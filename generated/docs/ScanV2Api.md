# \ScanV2Api

All URIs are relative to *https://api.nightfall.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScanPayloadV2**](ScanV2Api.md#ScanPayloadV2) | **Post** /v2/scan | Scan for sensitive information in your data



## ScanPayloadV2

> [][]ScanResponseV2 ScanPayloadV2(ctx, scanReqV2)

Scan for sensitive information in your data

Provide a list of arbitrary string data, and scan each item with the provided detectors to uncover sensitive information. Returns a list equal in size to the number of provided string payloads. The item at each list index will be a list of all matches for the provided detectors, or an empty list if no occurrences are found.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanReqV2** | [**ScanRequestV2**](ScanRequestV2.md)|  | 

### Return type

[**[][]ScanResponseV2**](array.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

