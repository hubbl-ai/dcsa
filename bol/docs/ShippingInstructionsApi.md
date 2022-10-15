# \ShippingInstructionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ShippingInstructionsPost**](ShippingInstructionsApi.md#V1ShippingInstructionsPost) | **Post** /v1/shipping-instructions | Post shipping instruction
[**V1ShippingInstructionsShippingInstructionIDGet**](ShippingInstructionsApi.md#V1ShippingInstructionsShippingInstructionIDGet) | **Get** /v1/shipping-instructions/{shippingInstructionID} | Get shipping instruction
[**V1ShippingInstructionsShippingInstructionIDPut**](ShippingInstructionsApi.md#V1ShippingInstructionsShippingInstructionIDPut) | **Put** /v1/shipping-instructions/{shippingInstructionID} | Put shipping instruction
[**V1ShippingInstructionsSummariesGet**](ShippingInstructionsApi.md#V1ShippingInstructionsSummariesGet) | **Get** /v1/shipping-instructions-summaries | Get shipping instruction



## V1ShippingInstructionsPost

> ShippingInstruction V1ShippingInstructionsPost(ctx).ShippingInstructionNoID(shippingInstructionNoID).APIVersion(aPIVersion).Execute()

Post shipping instruction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shippingInstructionNoID := *openapiclient.NewShippingInstructionNoID(openapiclient.transportDocumentType("BOL"), true, interface{}({"UNLocationCode":"USNYC","address":{"cityName":"København","stateRegion":"N/A","country":"Denmark"}}), true) // ShippingInstructionNoID | Parameters used to create the shipping instruction
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingInstructionsApi.V1ShippingInstructionsPost(context.Background()).ShippingInstructionNoID(shippingInstructionNoID).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingInstructionsApi.V1ShippingInstructionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ShippingInstructionsPost`: ShippingInstruction
    fmt.Fprintf(os.Stdout, "Response from `ShippingInstructionsApi.V1ShippingInstructionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ShippingInstructionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingInstructionNoID** | [**ShippingInstructionNoID**](ShippingInstructionNoID.md) | Parameters used to create the shipping instruction | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**ShippingInstruction**](ShippingInstruction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ShippingInstructionsShippingInstructionIDGet

> ShippingInstruction V1ShippingInstructionsShippingInstructionIDGet(ctx, shippingInstructionID).APIVersion(aPIVersion).Execute()

Get shipping instruction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shippingInstructionID := "shippingInstructionID_example" // string | 
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingInstructionsApi.V1ShippingInstructionsShippingInstructionIDGet(context.Background(), shippingInstructionID).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingInstructionsApi.V1ShippingInstructionsShippingInstructionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ShippingInstructionsShippingInstructionIDGet`: ShippingInstruction
    fmt.Fprintf(os.Stdout, "Response from `ShippingInstructionsApi.V1ShippingInstructionsShippingInstructionIDGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingInstructionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ShippingInstructionsShippingInstructionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**ShippingInstruction**](ShippingInstruction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ShippingInstructionsShippingInstructionIDPut

> ShippingInstruction V1ShippingInstructionsShippingInstructionIDPut(ctx, shippingInstructionID).ShippingInstruction(shippingInstruction).APIVersion(aPIVersion).Execute()

Put shipping instruction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    shippingInstructionID := "shippingInstructionID_example" // string | 
    shippingInstruction := *openapiclient.NewShippingInstruction("e0559d83-00e2-438e-afd9-fdd610c1a008", openapiclient.transportDocumentType("BOL"), true, interface{}({"UNLocationCode":"USNYC","address":{"cityName":"København","stateRegion":"N/A","country":"Denmark"}}), true) // ShippingInstruction | Parameters used to update the shipping instruction
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingInstructionsApi.V1ShippingInstructionsShippingInstructionIDPut(context.Background(), shippingInstructionID).ShippingInstruction(shippingInstruction).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingInstructionsApi.V1ShippingInstructionsShippingInstructionIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ShippingInstructionsShippingInstructionIDPut`: ShippingInstruction
    fmt.Fprintf(os.Stdout, "Response from `ShippingInstructionsApi.V1ShippingInstructionsShippingInstructionIDPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingInstructionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ShippingInstructionsShippingInstructionIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingInstruction** | [**ShippingInstruction**](ShippingInstruction.md) | Parameters used to update the shipping instruction | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**ShippingInstruction**](ShippingInstruction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ShippingInstructionsSummariesGet

> []ShippingInstructionHeader V1ShippingInstructionsSummariesGet(ctx).CarrierBookingReference(carrierBookingReference).DocumentStatus(documentStatus).EquipmentReference(equipmentReference).Limit(limit).Cursor(cursor).APIVersion(aPIVersion).Execute()

Get shipping instruction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    carrierBookingReference := "carrierBookingReference_example" // string | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  (optional)
    documentStatus := openapiclient.documentStatus("RECE") // DocumentStatus | Filter by the status of the document in the process. Possible values are: - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void)  (optional)
    equipmentReference := "equipmentReference_example" // string | Filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  (optional)
    limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
    cursor := "fE9mZnNldHw9MTAmbGltaXQ9MTA=" // string | A server generated value to specify a specific point in a collection result, used for pagination. (optional)
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingInstructionsApi.V1ShippingInstructionsSummariesGet(context.Background()).CarrierBookingReference(carrierBookingReference).DocumentStatus(documentStatus).EquipmentReference(equipmentReference).Limit(limit).Cursor(cursor).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingInstructionsApi.V1ShippingInstructionsSummariesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ShippingInstructionsSummariesGet`: []ShippingInstructionHeader
    fmt.Fprintf(os.Stdout, "Response from `ShippingInstructionsApi.V1ShippingInstructionsSummariesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ShippingInstructionsSummariesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  | 
 **documentStatus** | [**DocumentStatus**](DocumentStatus.md) | Filter by the status of the document in the process. Possible values are: - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void)  | 
 **equipmentReference** | **string** | Filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **cursor** | **string** | A server generated value to specify a specific point in a collection result, used for pagination. | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**[]ShippingInstructionHeader**](ShippingInstructionHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

