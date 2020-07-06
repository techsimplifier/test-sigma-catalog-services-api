# \ClassificationApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClassificationsClassificationtypeGet**](ClassificationApi.md#ApiClassificationsClassificationtypeGet) | **Get** /api/classifications({classificationtype}) | Returns the complete classification structure for a specified classification type.
[**ApiClassificationtypesGet**](ClassificationApi.md#ApiClassificationtypesGet) | **Get** /api/classificationtypes | Returns all the data types configured as classifications.


# **ApiClassificationsClassificationtypeGet**
> ApiClassificationsClassificationtypeGet(ctx, classificationType)
Returns the complete classification structure for a specified classification type.

Includes all instances and parent instances of the specified classification type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **classificationType** | **string**| classificationType is the type of classification for which you want to retrieve the instance hierarchy. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiClassificationtypesGet**
> ApiClassificationtypesGet(ctx, )
Returns all the data types configured as classifications.

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

