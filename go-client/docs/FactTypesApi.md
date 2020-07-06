# \FactTypesApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiFactTypesFactTypeNamesFactsGet**](FactTypesApi.md#ApiFactTypesFactTypeNamesFactsGet) | **Get** /api/factTypes({factTypeNames})/facts | Returns all facts filtered by the fact type
[**ApiFactTypesFactTypeNamesFactsValuesGet**](FactTypesApi.md#ApiFactTypesFactTypeNamesFactsValuesGet) | **Get** /api/factTypes({factTypeNames})/facts/values | Returns all Fact values for specified Fact types.
[**ApiFactsGet**](FactTypesApi.md#ApiFactsGet) | **Get** /api/facts | Returns All Facts


# **ApiFactTypesFactTypeNamesFactsGet**
> ApiFactTypesFactTypeNamesFactsGet(ctx, factTypeNames)
Returns all facts filtered by the fact type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **factTypeNames** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFactTypesFactTypeNamesFactsValuesGet**
> ApiFactTypesFactTypeNamesFactsValuesGet(ctx, factTypeNames, optional)
Returns all Fact values for specified Fact types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **factTypeNames** | **string**| factTypeNames are the names of the Fact types in a comma-separated list. | 
 **optional** | ***FactTypesApiApiFactTypesFactTypeNamesFactsValuesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FactTypesApiApiFactTypesFactTypeNamesFactsValuesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityTypes** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiFactsGet**
> ApiFactsGet(ctx, )
Returns All Facts

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

