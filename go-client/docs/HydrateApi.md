# \HydrateApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProductcandidateGetcommerciallyhydratedproductPost**](HydrateApi.md#ApiProductcandidateGetcommerciallyhydratedproductPost) | **Post** /api/productcandidate/getcommerciallyhydratedproduct | Validates a supplied Product Candidate as per the Validation resource and   return a hydrated product containing only the commercial data and no technical data  Note: Technical items are excluded unless explicitly included in the input candidate
[**ApiProductcandidateGetcommerciallyhydratedproductwithoutvalidatingPost**](HydrateApi.md#ApiProductcandidateGetcommerciallyhydratedproductwithoutvalidatingPost) | **Post** /api/productcandidate/getcommerciallyhydratedproductwithoutvalidating | Validates a supplied Product Candidate as per the Validation resource and   return a hydrated product containing only the commercial data and no technical data  Note: Technical items are excluded unless explicitly included in the input candidate
[**ApiProductcandidateGetfullyhydratedproductBVPost**](HydrateApi.md#ApiProductcandidateGetfullyhydratedproductBVPost) | **Post** /api/productcandidate/getfullyhydratedproductBV | Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated Product XML made up of the user selected options.
[**ApiProductcandidateGetfullyhydratedproductPost**](HydrateApi.md#ApiProductcandidateGetfullyhydratedproductPost) | **Post** /api/productcandidate/getfullyhydratedproduct | Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated Product XML made up of the user selected options.
[**ApiProductcandidateGetfullyhydratedproductwithoutvalidatingPost**](HydrateApi.md#ApiProductcandidateGetfullyhydratedproductwithoutvalidatingPost) | **Post** /api/productcandidate/getfullyhydratedproductwithoutvalidating | Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated Product XML made up of the user selected options.
[**ApiProductcandidateGettechnicallyhydratedproductPost**](HydrateApi.md#ApiProductcandidateGettechnicallyhydratedproductPost) | **Post** /api/productcandidate/gettechnicallyhydratedproduct | Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated commercial Product XML with technical ambiguities included
[**ApiProductcandidateGettechnicallyhydratedproductwithoutvalidatingPost**](HydrateApi.md#ApiProductcandidateGettechnicallyhydratedproductwithoutvalidatingPost) | **Post** /api/productcandidate/gettechnicallyhydratedproductwithoutvalidating | Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated commercial Product XML with technical ambiguities included


# **ApiProductcandidateGetcommerciallyhydratedproductPost**
> ApiProductcandidateGetcommerciallyhydratedproductPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and   return a hydrated product containing only the commercial data and no technical data  Note: Technical items are excluded unless explicitly included in the input candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGetcommerciallyhydratedproductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGetcommerciallyhydratedproductPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateGetcommerciallyhydratedproductwithoutvalidatingPost**
> ApiProductcandidateGetcommerciallyhydratedproductwithoutvalidatingPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and   return a hydrated product containing only the commercial data and no technical data  Note: Technical items are excluded unless explicitly included in the input candidate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGetcommerciallyhydratedproductwithoutvalidatingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGetcommerciallyhydratedproductwithoutvalidatingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateGetfullyhydratedproductBVPost**
> ApiProductcandidateGetfullyhydratedproductBVPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated Product XML made up of the user selected options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGetfullyhydratedproductBVPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGetfullyhydratedproductBVPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)| A xml/json object that contains a product candidate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateGetfullyhydratedproductPost**
> ApiProductcandidateGetfullyhydratedproductPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated Product XML made up of the user selected options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGetfullyhydratedproductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGetfullyhydratedproductPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)| A xml/json object that contains a product candidate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateGetfullyhydratedproductwithoutvalidatingPost**
> ApiProductcandidateGetfullyhydratedproductwithoutvalidatingPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated Product XML made up of the user selected options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGetfullyhydratedproductwithoutvalidatingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGetfullyhydratedproductwithoutvalidatingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)| A xml/json object that contains a product candidate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateGettechnicallyhydratedproductPost**
> ApiProductcandidateGettechnicallyhydratedproductPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated commercial Product XML with technical ambiguities included

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGettechnicallyhydratedproductPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGettechnicallyhydratedproductPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)| A xml/json object that contains a product candidate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateGettechnicallyhydratedproductwithoutvalidatingPost**
> ApiProductcandidateGettechnicallyhydratedproductwithoutvalidatingPost(ctx, optional)
Validates a supplied Product Candidate as per the Validation resource and return a structure representing a fully hydrated commercial Product XML with technical ambiguities included

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HydrateApiApiProductcandidateGettechnicallyhydratedproductwithoutvalidatingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HydrateApiApiProductcandidateGettechnicallyhydratedproductwithoutvalidatingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | [**optional.Interface of ProductCandidateDecomposeRequest**](ProductCandidateDecomposeRequest.md)| A xml/json object that contains a product candidate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

