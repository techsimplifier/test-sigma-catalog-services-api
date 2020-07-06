# \ValidateApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCommercialValidationForCustomerPortfolioPost**](ValidateApi.md#ApiCommercialValidationForCustomerPortfolioPost) | **Post** /api/commercialValidation/ForCustomerPortfolio | Fully validate the end state portfolio to determine if it is eligible for purchase by the customer
[**ApiCommercialvalidationForItemPost**](ValidateApi.md#ApiCommercialvalidationForItemPost) | **Post** /api/commercialvalidation/ForItem | Fully validate a ProductCandidate after carrying out inference  with respect to the commercial/technical boundary
[**ApiOrdercandidateValidateCommercialPost**](ValidateApi.md#ApiOrdercandidateValidateCommercialPost) | **Post** /api/ordercandidate/validate/commercial | Fully validate an OrderCandidate after carrying out inference  with respect to the commercial/technical boundary
[**ApiOrdercandidateValidateLockedCommercialPost**](ValidateApi.md#ApiOrdercandidateValidateLockedCommercialPost) | **Post** /api/ordercandidate/validate/locked/commercial | Fully validate an OrderCandidate with respect to the Commercial/Technical boundary  however does not carry out inference
[**ApiOrdercandidateValidateLockedPost**](ValidateApi.md#ApiOrdercandidateValidateLockedPost) | **Post** /api/ordercandidate/validate/locked | Fully validate an OrderCandidate however does not carry out inference
[**ApiOrdercandidateValidatePost**](ValidateApi.md#ApiOrdercandidateValidatePost) | **Post** /api/ordercandidate/validate | Fully validate an OrderCandidate after carrying out inference
[**ApiPortfolioValidatePost**](ValidateApi.md#ApiPortfolioValidatePost) | **Post** /api/portfolio/validate | Fully validate the end state portfolio to determine if it is eligible for purchase by the customer
[**ApiProductcandidateValidateCommercialPost**](ValidateApi.md#ApiProductcandidateValidateCommercialPost) | **Post** /api/productcandidate/validate/commercial | Fully validate a ProductCandidate after carrying out inference  with respect to the commercial/technical boundary
[**ApiProductcandidateValidateLockedCommercialPost**](ValidateApi.md#ApiProductcandidateValidateLockedCommercialPost) | **Post** /api/productcandidate/validate/Locked/commercial | Fully validate an ProductCandidate without performing inference and with respect to the Commercial/Technical boundary
[**ApiProductcandidateValidateLockedPost**](ValidateApi.md#ApiProductcandidateValidateLockedPost) | **Post** /api/productcandidate/validate/Locked | Fully validate a ProductCandidate without performing inference
[**ApiProductcandidateValidatePost**](ValidateApi.md#ApiProductcandidateValidatePost) | **Post** /api/productcandidate/validate | Fully validate a ProductCandidate after carrying out inference
[**ApiValidationForItemPost**](ValidateApi.md#ApiValidationForItemPost) | **Post** /api/validation/ForItem | Fully validate a ProductCandidate after carrying out inference


# **ApiCommercialValidationForCustomerPortfolioPost**
> ApiCommercialValidationForCustomerPortfolioPost(ctx, optional)
Fully validate the end state portfolio to determine if it is eligible for purchase by the customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiCommercialValidationForCustomerPortfolioPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiCommercialValidationForCustomerPortfolioPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portfolio** | **optional.String**| a Json/Xml object containing a CustomerPortfolio | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCommercialvalidationForItemPost**
> ApiCommercialvalidationForItemPost(ctx, optional)
Fully validate a ProductCandidate after carrying out inference  with respect to the commercial/technical boundary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiCommercialvalidationForItemPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiCommercialvalidationForItemPostOpts struct

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

# **ApiOrdercandidateValidateCommercialPost**
> ApiOrdercandidateValidateCommercialPost(ctx, optional)
Fully validate an OrderCandidate after carrying out inference  with respect to the commercial/technical boundary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiOrdercandidateValidateCommercialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiOrdercandidateValidateCommercialPostOpts struct

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

# **ApiOrdercandidateValidateLockedCommercialPost**
> ApiOrdercandidateValidateLockedCommercialPost(ctx, optional)
Fully validate an OrderCandidate with respect to the Commercial/Technical boundary  however does not carry out inference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiOrdercandidateValidateLockedCommercialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiOrdercandidateValidateLockedCommercialPostOpts struct

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

# **ApiOrdercandidateValidateLockedPost**
> ApiOrdercandidateValidateLockedPost(ctx, optional)
Fully validate an OrderCandidate however does not carry out inference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiOrdercandidateValidateLockedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiOrdercandidateValidateLockedPostOpts struct

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

# **ApiOrdercandidateValidatePost**
> ApiOrdercandidateValidatePost(ctx, )
Fully validate an OrderCandidate after carrying out inference

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

# **ApiPortfolioValidatePost**
> ApiPortfolioValidatePost(ctx, optional)
Fully validate the end state portfolio to determine if it is eligible for purchase by the customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiPortfolioValidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiPortfolioValidatePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portfolio** | **optional.String**| a Json/Xml object containing a CustomerPortfolio | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiProductcandidateValidateCommercialPost**
> ApiProductcandidateValidateCommercialPost(ctx, optional)
Fully validate a ProductCandidate after carrying out inference  with respect to the commercial/technical boundary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiProductcandidateValidateCommercialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiProductcandidateValidateCommercialPostOpts struct

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

# **ApiProductcandidateValidateLockedCommercialPost**
> ApiProductcandidateValidateLockedCommercialPost(ctx, optional)
Fully validate an ProductCandidate without performing inference and with respect to the Commercial/Technical boundary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiProductcandidateValidateLockedCommercialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiProductcandidateValidateLockedCommercialPostOpts struct

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

# **ApiProductcandidateValidateLockedPost**
> ApiProductcandidateValidateLockedPost(ctx, optional)
Fully validate a ProductCandidate without performing inference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiProductcandidateValidateLockedPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiProductcandidateValidateLockedPostOpts struct

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

# **ApiProductcandidateValidatePost**
> ApiProductcandidateValidatePost(ctx, optional)
Fully validate a ProductCandidate after carrying out inference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiProductcandidateValidatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiProductcandidateValidatePostOpts struct

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

# **ApiValidationForItemPost**
> ApiValidationForItemPost(ctx, optional)
Fully validate a ProductCandidate after carrying out inference

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidateApiApiValidationForItemPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiApiValidationForItemPostOpts struct

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

