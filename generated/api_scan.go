/*
Methods

This API exposes detectors for sensitive data in arbitrary string payloads.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// ScanApiService ScanApi service
type ScanApiService service

type ApiScanPayloadV2Request struct {
	ctx _context.Context
	ApiService *ScanApiService
	scanReqV2 *ScanRequestV2
}

func (r ApiScanPayloadV2Request) ScanReqV2(scanReqV2 ScanRequestV2) ApiScanPayloadV2Request {
	r.scanReqV2 = &scanReqV2
	return r
}

func (r ApiScanPayloadV2Request) Execute() ([][]ScanResponseV2, *_nethttp.Response, error) {
	return r.ApiService.ScanPayloadV2Execute(r)
}

/*
ScanPayloadV2 Scan for sensitive information in your data

Provide a list of arbitrary string data, and scan each item with the provided detectors to uncover sensitive information. Returns a list equal in size to the number of provided string payloads. The item at each list index will be a list of all matches for the provided detectors, or an empty list if no occurrences are found.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScanPayloadV2Request
*/
func (a *ScanApiService) ScanPayloadV2(ctx _context.Context) ApiScanPayloadV2Request {
	return ApiScanPayloadV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return [][]ScanResponseV2
func (a *ScanApiService) ScanPayloadV2Execute(r ApiScanPayloadV2Request) ([][]ScanResponseV2, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  [][]ScanResponseV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScanApiService.ScanPayloadV2")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/scan"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.scanReqV2 == nil {
		return localVarReturnValue, nil, reportError("scanReqV2 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.scanReqV2
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}