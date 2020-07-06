# \ImportApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiImportPost**](ImportApi.md#ApiImportPost) | **Post** /api/import | Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the request body.
[**ApiManageRefreshfromlocaldatastoreGet**](ImportApi.md#ApiManageRefreshfromlocaldatastoreGet) | **Get** /api/manage/refreshfromlocaldatastore | Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the appsettings.json - Server:DataStoreRefresh:DatasetPath
[**ApiManageRefreshfromlocaldatastorePost**](ImportApi.md#ApiManageRefreshfromlocaldatastorePost) | **Post** /api/manage/refreshfromlocaldatastore | Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the request body.
[**ApiStatusRefreshfromlocaldatastoreGet**](ImportApi.md#ApiStatusRefreshfromlocaldatastoreGet) | **Get** /api/status/refreshfromlocaldatastore | Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the appsettings.json - Server:DataStoreRefresh:DatasetPath


# **ApiImportPost**
> ApiImportPost(ctx, optional)
Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportApiApiImportPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportApiApiImportPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ImportTarget**](ImportTarget.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiManageRefreshfromlocaldatastoreGet**
> ApiManageRefreshfromlocaldatastoreGet(ctx, )
Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the appsettings.json - Server:DataStoreRefresh:DatasetPath

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

# **ApiManageRefreshfromlocaldatastorePost**
> ApiManageRefreshfromlocaldatastorePost(ctx, optional)
Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportApiApiManageRefreshfromlocaldatastorePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportApiApiManageRefreshfromlocaldatastorePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ImportTarget**](ImportTarget.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiStatusRefreshfromlocaldatastoreGet**
> ApiStatusRefreshfromlocaldatastoreGet(ctx, )
Requests a refresh of the Catalog Services Data from the local file system.  The path to the dataset is defined in the appsettings.json - Server:DataStoreRefresh:DatasetPath

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

