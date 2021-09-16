# \ScanApi

All URIs are relative to *https://api.nightfall.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScanPayloadV2**](ScanApi.md#ScanPayloadV2) | **Post** /v2/scan | Scan for sensitive information in your data



## ScanPayloadV2

> [][]ScanResponseV2 ScanPayloadV2(ctx).ScanReqV2(scanReqV2).Execute()

Scan for sensitive information in your data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    scanReqV2 := *openapiclient.NewScanRequestV2() // ScanRequestV2 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScanApi.ScanPayloadV2(context.Background()).ScanReqV2(scanReqV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanApi.ScanPayloadV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScanPayloadV2`: [][]ScanResponseV2
    fmt.Fprintf(os.Stdout, "Response from `ScanApi.ScanPayloadV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScanPayloadV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanReqV2** | [**ScanRequestV2**](ScanRequestV2.md) |  | 

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

