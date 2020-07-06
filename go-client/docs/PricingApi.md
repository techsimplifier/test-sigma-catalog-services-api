# \PricingApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOrdercandidatePriceFixedPost**](PricingApi.md#ApiOrdercandidatePriceFixedPost) | **Post** /api/ordercandidate/price/fixed | Returns a Fixed price summary for the given order candidate
[**ApiPriceForOrderCandidateFixedPost**](PricingApi.md#ApiPriceForOrderCandidateFixedPost) | **Post** /api/price/forOrderCandidate/fixed | Returns a Fixed price summary for the given order candidate
[**ApiPriceForProductCandidateFixedPost**](PricingApi.md#ApiPriceForProductCandidateFixedPost) | **Post** /api/price/forProductCandidate/fixed | Returns a Fixed price summary for the given Product candidate
[**ApiPricingGetPriceForProductCandidatePost**](PricingApi.md#ApiPricingGetPriceForProductCandidatePost) | **Post** /api/pricing/GetPriceForProductCandidate | Returns a Fixed price summary for the given Product candidate
[**ApiProductcandidatePriceFixedPost**](PricingApi.md#ApiProductcandidatePriceFixedPost) | **Post** /api/productcandidate/price/fixed | Returns a Fixed price summary for the given Product candidate


# **ApiOrdercandidatePriceFixedPost**
> ApiOrdercandidatePriceFixedPost(ctx, optional)
Returns a Fixed price summary for the given order candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PricingApiApiOrdercandidatePriceFixedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiOrdercandidatePriceFixedPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCandidate** | **optional.String**| a Json/Xml object containing an OrderCandidate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPriceForOrderCandidateFixedPost**
> ApiPriceForOrderCandidateFixedPost(ctx, optional)
Returns a Fixed price summary for the given order candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PricingApiApiPriceForOrderCandidateFixedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiPriceForOrderCandidateFixedPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCandidate** | **optional.String**| a Json/Xml object containing an OrderCandidate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPriceForProductCandidateFixedPost**
> ApiPriceForProductCandidateFixedPost(ctx, optional)
Returns a Fixed price summary for the given Product candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PricingApiApiPriceForProductCandidateFixedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiPriceForProductCandidateFixedPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | **optional.String**| a Json/Xml object containing a ProductCandidate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPricingGetPriceForProductCandidatePost**
> ApiPricingGetPriceForProductCandidatePost(ctx, optional)
Returns a Fixed price summary for the given Product candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PricingApiApiPricingGetPriceForProductCandidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiPricingGetPriceForProductCandidatePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | **optional.String**| a Json/Xml object containing a ProductCandidate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidatePriceFixedPost**
> ApiProductcandidatePriceFixedPost(ctx, optional)
Returns a Fixed price summary for the given Product candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PricingApiApiProductcandidatePriceFixedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PricingApiApiProductcandidatePriceFixedPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | **optional.String**| a Json/Xml object containing a ProductCandidate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

