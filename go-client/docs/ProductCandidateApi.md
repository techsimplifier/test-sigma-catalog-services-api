# \ProductCandidateApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProductCandidateGetPriceForPartialProductCandidatePost**](ProductCandidateApi.md#ApiProductCandidateGetPriceForPartialProductCandidatePost) | **Post** /api/productCandidate/GetPriceForPartialProductCandidate | Infers the missing data and resolves the pricing for the specified Product Candidate.
[**ApiProductCandidateInferProductCandidatePost**](ProductCandidateApi.md#ApiProductCandidateInferProductCandidatePost) | **Post** /api/productCandidate/InferProductCandidate | Attempts to augment the specified Product Candidate Structure with inferred data as per the specification.


# **ApiProductCandidateGetPriceForPartialProductCandidatePost**
> ApiProductCandidateGetPriceForPartialProductCandidatePost(ctx, optional)
Infers the missing data and resolves the pricing for the specified Product Candidate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductCandidateApiApiProductCandidateGetPriceForPartialProductCandidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCandidateApiApiProductCandidateGetPriceForPartialProductCandidatePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCandidate** | **optional.String**| A xml/json object that contains a Product Candidate. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductCandidateInferProductCandidatePost**
> ApiProductCandidateInferProductCandidatePost(ctx, )
Attempts to augment the specified Product Candidate Structure with inferred data as per the specification.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

