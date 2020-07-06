# \EntityTypesApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEntitytypesEntityTypesCharacteristicsGet**](EntityTypesApi.md#ApiEntitytypesEntityTypesCharacteristicsGet) | **Get** /api/entitytypes({entityTypes})/characteristics | Returns all of the characteristics used by the Entity Types passed.
[**ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationValuesGet**](EntityTypesApi.md#ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationValuesGet) | **Get** /api/entitytypes({usedByTypeNames})/classificationtypes/classification-values | Returns list of all the classifications with values that are being used for the supplied Entity/Type
[**ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationsGet**](EntityTypesApi.md#ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationsGet) | **Get** /api/entitytypes({usedByTypeNames})/classificationtypes/classifications | Reports on the \&quot;Classification\&quot; used by a specific type. The Type could be an entity or otherwise.
[**ApiEntitytypesUsedByTypeNamesFacttypesFactsGet**](EntityTypesApi.md#ApiEntitytypesUsedByTypeNamesFacttypesFactsGet) | **Get** /api/entitytypes({usedByTypeNames})/facttypes/facts | Returns all Facts, including where they are used in the data model and in which element context.
[**ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesElementsElementNamesGet**](EntityTypesApi.md#ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesElementsElementNamesGet) | **Get** /api/entitytypes({usedByTypeNames})/facttypes/facts-values/elements({elementNames}) | Returns list of all the facts with values that are being used for the supplied Entity/Type.
[**ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesGet**](EntityTypesApi.md#ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesGet) | **Get** /api/entitytypes({usedByTypeNames})/facttypes/facts-values | Reports on the \&quot;Facts\&quot; used by a specific type. The Type could be an entity or otherwise.
[**ApiTypesEntityTypesEntitiesGet**](EntityTypesApi.md#ApiTypesEntityTypesEntitiesGet) | **Get** /api/types({entityTypes})/entities | Returns a list of all entities that are of, or derive from, a given type or types.


# **ApiEntitytypesEntityTypesCharacteristicsGet**
> ApiEntitytypesEntityTypesCharacteristicsGet(ctx, entityTypes)
Returns all of the characteristics used by the Entity Types passed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTypes** | **string**| entityTypes is a comma-separated list of Entity types. If an entity type has inheriting types, the               elements of those types are not analysed by this query. Elements inherited from Parent types are               included. i.e. When querying for type &#39;Product&#39;, elements on &#39;Launch Entity&#39; are included,               whereas elements on the type &#39;Package&#39; are not included. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationValuesGet**
> ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationValuesGet(ctx, usedByTypeNames)
Returns list of all the classifications with values that are being used for the supplied Entity/Type

This query provides additional filtering in that only values that have actually been used against an entity will be included in the set of values returned.   If no entities of the type (or a type inheriting from the type provided) use a particular value then it will not be returned in the result set.   Therefore every value returned will be a classification about an entity of the type supplied. This is useful as a method to only display options that have been used.  As the type parameter is made more specific (i.e. further down the inheritance tree) the number of classification elements will increase as more are include from inheritance.   However the number of values will decrease as the number of entities involved decreases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usedByTypeNames** | **string**| usedByTypeNames is a comma separated list of the Entity type names, for example, &#39;Product,Package&#39;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationsGet**
> ApiEntitytypesUsedByTypeNamesClassificationtypesClassificationsGet(ctx, usedByTypeNames)
Reports on the \"Classification\" used by a specific type. The Type could be an entity or otherwise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usedByTypeNames** | **string**| usedByTypeNames is a comma-separated list of the Fact type names, for example, Product,Package | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitytypesUsedByTypeNamesFacttypesFactsGet**
> ApiEntitytypesUsedByTypeNamesFacttypesFactsGet(ctx, usedByTypeNames)
Returns all Facts, including where they are used in the data model and in which element context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usedByTypeNames** | **string**| usedByTypeNames is a comma-separated list of the Fact type names. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesElementsElementNamesGet**
> ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesElementsElementNamesGet(ctx, usedByTypeNames, elementNames)
Returns list of all the facts with values that are being used for the supplied Entity/Type.

This query provides additional filtering in that only values that have actually been used against an entity will be included in the set of values returned.   If no entities of the type (or a type inheriting from the type provided) use a particular value then it will not be returned in the result set.   Therefore every value returned will be a fact about an entity of the type supplied. This is useful as a method to only display options that have been used.  As the type parameter is made more specific (i.e. further down the inheritance tree) the number of fact elements will increase as more are include from inheritance.   However the number of values will decrease as the number of entities involved decreases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usedByTypeNames** | **string**| usedByTypeNames is a comma separated list of the Entity type names, for example, &#39;Product,Package&#39; | 
  **elementNames** | **string**| elementNames are the names of the types of elements in a comma-separated list. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesGet**
> ApiEntitytypesUsedByTypeNamesFacttypesFactsValuesGet(ctx, usedByTypeNames)
Reports on the \"Facts\" used by a specific type. The Type could be an entity or otherwise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **usedByTypeNames** | **string**| usedByTypeNames is a comma-separated list of the Fact type names, for example, Lookup_Speed,Lookup_Color | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiTypesEntityTypesEntitiesGet**
> ApiTypesEntityTypesEntitiesGet(ctx, entityTypes)
Returns a list of all entities that are of, or derive from, a given type or types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTypes** | **string**| entityTypes is a comma-separated list of entity types, for example, Component, Charge. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

