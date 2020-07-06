# \RatesApi

All URIs are relative to *https://localhost/CS3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRatesChargePathGet**](RatesApi.md#ApiRatesChargePathGet) | **Get** /api/rates({chargePath}) | Return the details of the charge for the supplied charge path.


# **ApiRatesChargePathGet**
> ApiRatesChargePathGet(ctx, chargePath)
Return the details of the charge for the supplied charge path.

Responses vary depending on whether the supplied charge is an Event Charge or a Non-Event Charge.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chargePath** | **string**| chargePath is a comma-separated list of GUIDs corresponding to the path to a charge. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

