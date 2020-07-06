# \DecomposeApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOrdercandidateDecomposeCommercialPost**](DecomposeApi.md#ApiOrdercandidateDecomposeCommercialPost) | **Post** /api/ordercandidate/decompose/commercial | Fully validates and articulates an OrderCandidate request with respect to the commercial/technical boundary
[**ApiOrdercandidateDecomposePost**](DecomposeApi.md#ApiOrdercandidateDecomposePost) | **Post** /api/ordercandidate/decompose | Fully validates and articulates an OrderCandidate request
[**ApiProductcandidateDecomposeCommercialPost**](DecomposeApi.md#ApiProductcandidateDecomposeCommercialPost) | **Post** /api/productcandidate/decompose/commercial | Fully validates and articulates an ProductCandidate request with respect to the commercial/technical boundary
[**ApiProductcandidateDecomposePost**](DecomposeApi.md#ApiProductcandidateDecomposePost) | **Post** /api/productcandidate/decompose | Fully validates and articulates an ProductCandidate request


# **ApiOrdercandidateDecomposeCommercialPost**
> ApiOrdercandidateDecomposeCommercialPost(ctx, optional)
Fully validates and articulates an OrderCandidate request with respect to the commercial/technical boundary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DecomposeApiApiOrdercandidateDecomposeCommercialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DecomposeApiApiOrdercandidateDecomposeCommercialPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCandidate** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrdercandidateDecomposePost**
> ApiOrdercandidateDecomposePost(ctx, optional)
Fully validates and articulates an OrderCandidate request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DecomposeApiApiOrdercandidateDecomposePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DecomposeApiApiOrdercandidateDecomposePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCandidate** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateDecomposeCommercialPost**
> ApiProductcandidateDecomposeCommercialPost(ctx, optional)
Fully validates and articulates an ProductCandidate request with respect to the commercial/technical boundary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DecomposeApiApiProductcandidateDecomposeCommercialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DecomposeApiApiProductcandidateDecomposeCommercialPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateDecomposePost**
> ApiProductcandidateDecomposePost(ctx, optional)
Fully validates and articulates an ProductCandidate request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DecomposeApiApiProductcandidateDecomposePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DecomposeApiApiProductcandidateDecomposePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

