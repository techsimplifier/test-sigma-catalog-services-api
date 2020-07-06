# \RelatedEntitiesApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRelatedEntitiesGet**](RelatedEntitiesApi.md#ApiRelatedEntitiesGet) | **Get** /api/RelatedEntities | Returns a set of entities that are related to a given entity.


# **ApiRelatedEntitiesGet**
> ApiRelatedEntitiesGet(ctx, optional)
Returns a set of entities that are related to a given entity.

Used for interrogating instances using the T_RelatedEntityRule pattern and the TRelated_Contextual_Entity_Rule.  <para></para>  There are two flavours of this operation. It can retrieve the information about the relations based upon two different conditions:  <para></para>  • Related entities based upon the Type(s) of the relation. This will return all elements which match the types passed.  <para></para>  • Related entities based upon element names. This will return the data from these specific elements only.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RelatedEntitiesApiApiRelatedEntitiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelatedEntitiesApiApiRelatedEntitiesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| entityId is the GUID of the entity being queried. | 
 **relationTypes** | **optional.String**| relationTypes is a comma separated list of types of relation or element names. | 
 **requestType** | **optional.String**| requestType indicates whether the relation types contains “types” or “elements” according to the value passed here. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

