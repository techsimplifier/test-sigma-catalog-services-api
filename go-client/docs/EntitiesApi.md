# \EntitiesApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEntitiesEntityIdGet**](EntitiesApi.md#ApiEntitiesEntityIdGet) | **Get** /api/entities({entityId}) | Returns entity with the specified business ID or GUID.
[**ApiEntitiesEntityInContextPost**](EntitiesApi.md#ApiEntitiesEntityInContextPost) | **Post** /api/entities/EntityInContext | Returns an entity in context of a product specification based on a guid path or business path
[**ApiEntitiesEntityTypeFactelementfilterElementFilterListGet**](EntitiesApi.md#ApiEntitiesEntityTypeFactelementfilterElementFilterListGet) | **Get** /api/entities({entityType})/factelementfilter({elementFilterList}) | Returns entities that use given fact values according to the filter.
[**ApiEntitiesGuidParentsGet**](EntitiesApi.md#ApiEntitiesGuidParentsGet) | **Get** /api/entities({guid})/parents | Returns a list of all entities that are parents of the supplied entity.
[**ApiEntitiesInputGuidProductcandidateGet**](EntitiesApi.md#ApiEntitiesInputGuidProductcandidateGet) | **Get** /api/entities({inputGuid})/productcandidate | In order to assist with creation of Product Candidate structures CPQ provides a resource that will return, for any entity, a product candidate focused structure.
[**ApiEntitiesInstanceTypeNameClassificationsClassificationsGet**](EntitiesApi.md#ApiEntitiesInstanceTypeNameClassificationsClassificationsGet) | **Get** /api/entities({instanceTypeName})/Classifications({classifications}) | Returns a list of all entities that are of, or derive from, a given type and are classified or not classified by a given set of classification GUIDs.
[**ApiEntitiesSchemaGet**](EntitiesApi.md#ApiEntitiesSchemaGet) | **Get** /api/entities/Schema | Return the schema currently used by the Entities.
[**ApiEntitiesWithClassificationsAndCharacteristicsGet**](EntitiesApi.md#ApiEntitiesWithClassificationsAndCharacteristicsGet) | **Get** /api/entities/WithClassificationsAndCharacteristics | Returns a list of all entities that are of a given type or derive from a given type, are classified by a given set of classification Guids and have a given set of characteristics
[**ApiEntitiesWithGet**](EntitiesApi.md#ApiEntitiesWithGet) | **Get** /api/entities/With | Returns a list of all entities that: are of, or derive from, a given type   and that are classified by a given set of classification Guids  and/or have a given set of characteristics and/or have given fact values.


# **ApiEntitiesEntityIdGet**
> ApiEntitiesEntityIdGet(ctx, entityId, optional)
Returns entity with the specified business ID or GUID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityId** | **string**| entityId is the GUID or BusinessId of the entity. | 
 **optional** | ***EntitiesApiApiEntitiesEntityIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiApiEntitiesEntityIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.String**| idType is either “GUID” or “Business” to indicate the type of ID being passed into the query. Note: The IDType parameter is not required to return data when EntityID is a GUID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesEntityInContextPost**
> ApiEntitiesEntityInContextPost(ctx, optional)
Returns an entity in context of a product specification based on a guid path or business path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntitiesApiApiEntitiesEntityInContextPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiApiEntitiesEntityInContextPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parameters** | [**optional.Interface of EntityInContextRequest**](EntityInContextRequest.md)| parameters are the GUID or BUID path of a entity within a product specification. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesEntityTypeFactelementfilterElementFilterListGet**
> ApiEntitiesEntityTypeFactelementfilterElementFilterListGet(ctx, entityType, elementFilterList)
Returns entities that use given fact values according to the filter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | **string**| Entity Type is a list the entity types to look for comma separated e.g. Component, Charge. | 
  **elementFilterList** | **string**| &lt;main&gt;A list of &#39;fact value&#39; packets which define what fact-values you want to match on.  You can pass multiple packets into this query. Each packet is in the following:&lt;/main&gt;  &lt;sub&gt;1. [FACTVALUE; FACTVALUEGUIDS; MATCHALLFACTVALUES; CANNEVEREXIST; NOT]&lt;/sub&gt;  &lt;sub&gt;2. FACTVALUE - The schema element name that contains the fact values to filter on.&lt;/sub&gt;  &lt;sub&gt;3. FACTVALUEGUIDS - is a comma seperated list that represents all the fact values on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLFACTVALUES - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt;  &lt;sub&gt;6. NOT - an optional parameter to allow the filter to be inverted&lt;/sub&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesGuidParentsGet**
> ApiEntitiesGuidParentsGet(ctx, guid, optional)
Returns a list of all entities that are parents of the supplied entity.

The list of parents is filtered based on the parent type parameter. If the query does not specify a type, all the  parent entities are returned regardless of type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **guid** | **string**| guid is the GUID of the entity for which the query finds parents. | 
 **optional** | ***EntitiesApiApiEntitiesGuidParentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiApiEntitiesGuidParentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentTypes** | **optional.String**| parentTypes is a comma separated list of parent types, for example, Component,Package_Template. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesInputGuidProductcandidateGet**
> ApiEntitiesInputGuidProductcandidateGet(ctx, inputGuid)
In order to assist with creation of Product Candidate structures CPQ provides a resource that will return, for any entity, a product candidate focused structure.

Inside the response structure there will be a single instance for every child entity that exists within the product specification. This is irrespective of the cardinality on the relations involved.   This means that the structure will in most cases not be valid for the validation service. The data though is in the correct format to be worked with to construct the input by duplicating or removing entries from the XML.  <para>  Within each child entity all of the configurable data items (those that inherit from TConfigurable_Fact and TSpecCharUse) will also be included as CharacteristicUse items. Again all values will be shown which will make the structure invalid in most cases.  </para><para>  Also in order to assist with the understanding of the produced structure there are a number of attributes that are added to the output that are not part of the Product Candidate xsd structure. These elements are all marked with the attribute ‘Remove = “true”’  to allow them to be easily taken out of the structure. These attributes provide the name and typing of the entities involved and also provide the cardinalities of the relations and configurable items alongside the dates for any relations.  </para>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inputGuid** | **string**| inputGuid is the GUID of the product specification that a candidate is required for. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesInstanceTypeNameClassificationsClassificationsGet**
> ApiEntitiesInstanceTypeNameClassificationsClassificationsGet(ctx, instanceTypeName, classifications, optional)
Returns a list of all entities that are of, or derive from, a given type and are classified or not classified by a given set of classification GUIDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **instanceTypeName** | **string**| instanceTypeName is the name of the type that contains the classification element of interest,              for example, Package, Component, or Bundle. | 
  **classifications** | **string**| &lt;main&gt;A comma separated list of classification packets that define the classification on which to match.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [CLASSIFICATIONELEMENTNAME; CLASSIFICATIONGUIDS; MATCHALLCLASSIFICATION; CANNEVEREXIST]&lt;/sub&gt;  &lt;sub&gt;2. CLASSIFICATIONELEMENTNAME - is the name of the schema element that holds the classifications on which they query is filtering.&lt;/sub&gt;  &lt;sub&gt;3. CLASSIFICATIONGUIDS - is a comma-separated list that represents all the classifications on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLCLASSIFICATION - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt; | 
 **optional** | ***EntitiesApiApiEntitiesInstanceTypeNameClassificationsClassificationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiApiEntitiesInstanceTypeNameClassificationsClassificationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **operationType** | **optional.String**| operationType is either “classified” or “notClassified” to indicate whether or not to classify              the entity by the list of classifications. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesSchemaGet**
> ApiEntitiesSchemaGet(ctx, )
Return the schema currently used by the Entities.

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

# **ApiEntitiesWithClassificationsAndCharacteristicsGet**
> ApiEntitiesWithClassificationsAndCharacteristicsGet(ctx, optional)
Returns a list of all entities that are of a given type or derive from a given type, are classified by a given set of classification Guids and have a given set of characteristics

Multiple characteristic, facts and classification packets can be added by appending each element set inside square brackets []. Each additional packet will be combined with the others using “AND” logic so all must match.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntitiesApiApiEntitiesWithClassificationsAndCharacteristicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiApiEntitiesWithClassificationsAndCharacteristicsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classifications** | **optional.String**| &lt;main&gt;A comma separated list of classification packets that define the classification on which to match.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [CLASSIFICATIONELEMENTNAME; CLASSIFICATIONGUIDS; MATCHALLCLASSIFICATION; CANNEVEREXIST; NOT]&lt;/sub&gt;  &lt;sub&gt;2. CLASSIFICATIONELEMENTNAME - is the name of the schema element that holds the classifications on which they query is filtering.&lt;/sub&gt;  &lt;sub&gt;3. CLASSIFICATIONGUIDS - is a comma-separated list that represents all the classifications on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLCLASSIFICATION - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt;  &lt;sub&gt;6. NOT - an optional parameter to allow the filter to be inverted&lt;/sub&gt; | 
 **instanceTypeName** | **optional.String**| instanceTypeName is the name of the type that contains the classification element of interest,              for example, Package, Component, or Bundle. | 
 **characteristics** | **optional.String**| &lt;main&gt;A comma separated list of characteristics packets that define the characteristics on which to match.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [CHARACTERISTICELEMENTNAME; CHARACTERISTICGUIDS; MATCHALLCHARACTERISTICS; CANNEVEREXIST; NOT]&lt;/sub&gt;  &lt;sub&gt;2. CHARACTERISTICELEMENTNAME - is the name of the schema element that holds the characteristics on which they query is filtering.&lt;/sub&gt;  &lt;sub&gt;3. CHARACTERISTICNGUIDS - is a comma-separated list that represents all the characteristics on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLCHARACTERISTICS - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt;  &lt;sub&gt;6. NOT - an optional parameter to allow the filter to be inverted&lt;/sub&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiEntitiesWithGet**
> ApiEntitiesWithGet(ctx, optional)
Returns a list of all entities that: are of, or derive from, a given type   and that are classified by a given set of classification Guids  and/or have a given set of characteristics and/or have given fact values.

Each of the parameters on the query is optional with the exception of the InstanceTypeName so this query can be used to search for any and all types of entity restricting it by any of the properties available.  Multiple characteristic, facts and classification packets can be added by appending each element set inside square brackets []. Each additional packet will be combined with the others using “AND” logic so all must match.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EntitiesApiApiEntitiesWithGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EntitiesApiApiEntitiesWithGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classifications** | **optional.String**| &lt;main&gt;A comma separated list of classification packets that define the classification on which to match.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [CLASSIFICATIONELEMENTNAME; CLASSIFICATIONGUIDS; MATCHALLCLASSIFICATION; CANNEVEREXIST; NOT]&lt;/sub&gt;  &lt;sub&gt;2. CLASSIFICATIONELEMENTNAME - is the name of the schema element that holds the classifications on which they query is filtering.&lt;/sub&gt;  &lt;sub&gt;3. CLASSIFICATIONGUIDS - is a comma-separated list that represents all the classifications on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLCLASSIFICATION - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt;  &lt;sub&gt;6. NOT - an optional parameter to allow the filter to be inverted&lt;/sub&gt; | 
 **instanceTypeName** | **optional.String**| instanceTypeName is the name of the type that contains the classification element of interest, for example, Package, Component, or Bundle. | 
 **characteristics** | **optional.String**| &lt;main&gt;A comma separated list of characteristics packets that define the characteristics on which to match.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [CHARACTERISTICELEMENTNAME; CHARACTERISTICGUIDS; MATCHALLCHARACTERISTICS; CANNEVEREXIST; NOT]&lt;/sub&gt;  &lt;sub&gt;2. CHARACTERISTICELEMENTNAME - is the name of the schema element that holds the characteristics on which they query is filtering.&lt;/sub&gt;  &lt;sub&gt;3. CHARACTERISTICNGUIDS - is a comma-separated list that represents all the characteristics on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLCHARACTERISTICS - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt;  &lt;sub&gt;6. NOT - an optional parameter to allow the filter to be inverted&lt;/sub&gt; | 
 **factElementFilter** | **optional.String**| &lt;main&gt;A list of &#39;fact value&#39; packets which define what fact-values you want to match on.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [FACTVALUE; FACTVALUEGUIDS; MATCHALLFACTVALUES; CANNEVEREXIST; NOT]&lt;/sub&gt;  &lt;sub&gt;2. FACTVALUE - The schema element name that contains the fact values to filter on.&lt;/sub&gt;  &lt;sub&gt;3. FACTVALUEGUIDS - is a comma seperated list that represents all the fact values on which the query is filtering.&lt;/sub&gt;  &lt;sub&gt;4. MATCHALLFACTVALUES - is either TRUE or FALSE to specify whether the query results need to match ALL of these or just ANY.&lt;/sub&gt;  &lt;sub&gt;5. CANNEVEREXIST - either true or false to specify whether we want to return entities that lack the Element name (Optional)&lt;/sub&gt;  &lt;sub&gt;6. NOT - an optional parameter to allow the filter to be inverted&lt;/sub&gt; | 
 **categories** | **optional.String**| &lt;main&gt;A list of &#39;categories value&#39; packets which define what categories you want to match on.  You may pass multiple packets and each packet has the following format:&lt;/main&gt;  &lt;sub&gt;1. [Category Ids; IncludeChildCategories]&lt;/sub&gt;  &lt;sub&gt;2. Category Ids - the comma separated list of category ids.&lt;/sub&gt;  &lt;sub&gt;3. IncludeChildCategories - boolean value indicating if filter apply on child categories as well&lt;/sub&gt; | 
 **availableDate** | **optional.String**| Target date. If supplied, only entities available on this date are returned. Otherwise results are not date filtered | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

