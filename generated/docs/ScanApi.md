# \ScanApi

All URIs are relative to *https://api.nightfall.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScanPayload**](ScanApi.md#ScanPayload) | **Post** /scan | Scan for sensitive information in your data



## ScanPayload

> [][]ScanResponse ScanPayload(ctx, scanReq)

Scan for sensitive information in your data

Provide a list of arbitrary string data, and scan each item with the provided detectors to uncover sensitive information. Returns a list equal in size to the number of provided string payloads. The item at each list index will be a list of all matches for the provided detectors, or an empty list if no occurrences are found.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanReq** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**[][]ScanResponse**](array.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

