# \OrdersApi

All URIs are relative to *http://localhost/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrders**](OrdersApi.md#GetOrders) | **Get** /orders/{uuid} | Retrieve a set of Orders and all of its Revisions by UUID
[**IndexOrders**](OrdersApi.md#IndexOrders) | **Get** /orders | Retrieve orders that match a particular search
[**PostRevision**](OrdersApi.md#PostRevision) | **Post** /orders | Submit a new set of orders, make an amendment to an existing set of orders, or cancel a set of orders.
[**PostRevisionToOrders**](OrdersApi.md#PostRevisionToOrders) | **Post** /orders/{uuid} | Make an amendment to or cancel an existing set of orders by UUID


# **GetOrders**
> Orders GetOrders(ctx, uuid)
Retrieve a set of Orders and all of its Revisions by UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | [**string**](.md)| UUID of the orders to return | 

### Return type

[**Orders**](Orders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndexOrders**
> []Orders IndexOrders(ctx, optional)
Retrieve orders that match a particular search

Returns all orders that match all of the supplied parameters. At least one query parameter must be provided. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordersNum** | **string**| The Orders number as it appears on the Orders | 
 **edipi** | **string**| Electronic Data Interchange Personal Identifier, AKA the 10 digit DoD ID Number of the member | 
 **latestOnly** | **bool**| If true, look only at the latest Revision (by seqNum) of any set of Orders when applying the other Revision-specific parameters. If false, search all Revisions.  Defaults to false if omitted.  | 
 **status** | **string**| Return only Orders where the status of the latest Revision of the Orders matches the supplied status. | 
 **issuingAuthority** | **string**| Name of the Issuing Authority of the Orders. | 

### Return type

[**[]Orders**](Orders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRevision**
> Orders PostRevision(ctx, ordersNum, edipi, issuingAuthority, revision)
Submit a new set of orders, make an amendment to an existing set of orders, or cancel a set of orders.

Creates a Revision of a set of orders.  ## New Orders The supplied Revision is considered part of a new set of Orders if the supplied ordersNum has never been seen before for the supplied issuingAuthority. A new UUID is created and associated with a new Orders, which is returned along with the supplied Revision.  ## Amended Orders If the system determines that the supplied Revision is an amendment to an existing set of Orders, usually because the ordersNum already exists and the seqNum is unique, then the supplied Revision is added to the existing Orders.  Specifying the existing Orders' UUID is optional; clients are not required to remember the UUIDs returned with the Orders by previous calls to this API. If it is omitted, the system will attempt to locate the correct Orders using the ordersNum and issuingAuthority fields.  ## Canceled, Rescinded, or Revoked Orders To cancel, rescind, or revoke Orders, POST a new Revision with the status set to \"canceled\".  # Errors It is an error to specify an already-created seqNum in the provided Revision for an existing set of Orders. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ordersNum** | **string**| The Orders number as it appears on the Orders | 
  **edipi** | **string**| Electronic Data Interchange Personal Identifier, AKA the 10 digit DoD ID Number of the member | 
  **issuingAuthority** | **string**| Name of the Issuing Authority of the Orders. | 
  **revision** | [**Revision**](Revision.md)|  | 

### Return type

[**Orders**](Orders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRevisionToOrders**
> Orders PostRevisionToOrders(ctx, uuid, revision)
Make an amendment to or cancel an existing set of orders by UUID

Creates a Revision of a set of orders. The Orders to be amended or canceled must already exist with the supplied UUID.  The seqNum in the supplied Revision must be unique, and not already present in the Orders. Nothing else is required to change in the Revision compared to any other Revision in the Orders.  ## Errors It is an error to specify a non-existent UUID.  It is an error to specify an already-created seqNum in the Revision. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **uuid** | [**string**](.md)| UUID of the orders to return | 
  **revision** | [**Revision**](Revision.md)|  | 

### Return type

[**Orders**](Orders.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

