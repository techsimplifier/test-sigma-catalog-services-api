# \CategoriesApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCategoriesCategoriesEntitiesGet**](CategoriesApi.md#ApiCategoriesCategoriesEntitiesGet) | **Get** /api/categories({categories})/entities | Returns a list of all entities that reside in the specified categories.
[**ApiCategoriesGet**](CategoriesApi.md#ApiCategoriesGet) | **Get** /api/categories | Returns all the categories in the datastore.


# **ApiCategoriesCategoriesEntitiesGet**
> ApiCategoriesCategoriesEntitiesGet(ctx, categories, optional)
Returns a list of all entities that reside in the specified categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categories** | **string**| categories is a comma-separated list of category IDs. | 
 **optional** | ***CategoriesApiApiCategoriesCategoriesEntitiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiApiCategoriesCategoriesEntitiesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceTypeName** | **optional.String**| InstanceTypeName is the name of the instance types to be returned. For example, Component, Charge. | 
 **includeChildCategories** | **optional.String**| includeChildCategories indicates whether entities should be returned that live in categories that are children of the category ids specified in the Categories parameter. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCategoriesGet**
> ApiCategoriesGet(ctx, )
Returns all the categories in the datastore.

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

