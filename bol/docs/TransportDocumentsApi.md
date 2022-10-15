# \TransportDocumentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TransportDocumentsSummariesGet**](TransportDocumentsApi.md#V1TransportDocumentsSummariesGet) | **Get** /v1/transport-documents-summaries | Get Transport Documents
[**V1TransportDocumentsTransportDocumentReferenceApprovalsPost**](TransportDocumentsApi.md#V1TransportDocumentsTransportDocumentReferenceApprovalsPost) | **Post** /v1/transport-documents/{transportDocumentReference}/approvals | Approve a Transport Document.
[**V1TransportDocumentsTransportDocumentReferenceGet**](TransportDocumentsApi.md#V1TransportDocumentsTransportDocumentReferenceGet) | **Get** /v1/transport-documents/{transportDocumentReference} | Get Transport Document



## V1TransportDocumentsSummariesGet

> []TransportDocumentHeader V1TransportDocumentsSummariesGet(ctx).Limit(limit).Cursor(cursor).CarrierBookingReference(carrierBookingReference).DocumentStatus(documentStatus).EquipmentReference(equipmentReference).APIVersion(aPIVersion).Execute()

Get Transport Documents



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
    limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
    cursor := "fE9mZnNldHw9MTAmbGltaXQ9MTA=" // string | A server generated value to specify a specific point in a collection result, used for pagination. (optional)
    carrierBookingReference := "carrierBookingReference_example" // string | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  (optional)
    documentStatus := openapiclient.documentStatus("RECE") // DocumentStatus | Filter by the status of the document in the process. Possible values are: - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void)  (optional)
    equipmentReference := "equipmentReference_example" // string | Filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  (optional)
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportDocumentsApi.V1TransportDocumentsSummariesGet(context.Background()).Limit(limit).Cursor(cursor).CarrierBookingReference(carrierBookingReference).DocumentStatus(documentStatus).EquipmentReference(equipmentReference).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentsApi.V1TransportDocumentsSummariesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransportDocumentsSummariesGet`: []TransportDocumentHeader
    fmt.Fprintf(os.Stdout, "Response from `TransportDocumentsApi.V1TransportDocumentsSummariesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1TransportDocumentsSummariesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **cursor** | **string** | A server generated value to specify a specific point in a collection result, used for pagination. | 
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  | 
 **documentStatus** | [**DocumentStatus**](DocumentStatus.md) | Filter by the status of the document in the process. Possible values are: - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void)  | 
 **equipmentReference** | **string** | Filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**[]TransportDocumentHeader**](TransportDocumentHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TransportDocumentsTransportDocumentReferenceApprovalsPost

> TransportDocument V1TransportDocumentsTransportDocumentReferenceApprovalsPost(ctx, transportDocumentReference).DocumentSignature(documentSignature).APIVersion(aPIVersion).Execute()

Approve a Transport Document.



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
    transportDocumentReference := "transportDocumentReference_example" // string | The transportDocumentReference of the transportDocument
    documentSignature := "documentSignature_example" // string | The signature of the Transport Document to approve
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportDocumentsApi.V1TransportDocumentsTransportDocumentReferenceApprovalsPost(context.Background(), transportDocumentReference).DocumentSignature(documentSignature).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentsApi.V1TransportDocumentsTransportDocumentReferenceApprovalsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransportDocumentsTransportDocumentReferenceApprovalsPost`: TransportDocument
    fmt.Fprintf(os.Stdout, "Response from `TransportDocumentsApi.V1TransportDocumentsTransportDocumentReferenceApprovalsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportDocumentReference** | **string** | The transportDocumentReference of the transportDocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TransportDocumentsTransportDocumentReferenceApprovalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentSignature** | **string** | The signature of the Transport Document to approve | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**TransportDocument**](TransportDocument.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1TransportDocumentsTransportDocumentReferenceGet

> TransportDocument V1TransportDocumentsTransportDocumentReferenceGet(ctx, transportDocumentReference).APIVersion(aPIVersion).Execute()

Get Transport Document



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
    transportDocumentReference := "transportDocumentReference_example" // string | The transportDocumentReference of the transportDocument
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportDocumentsApi.V1TransportDocumentsTransportDocumentReferenceGet(context.Background(), transportDocumentReference).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportDocumentsApi.V1TransportDocumentsTransportDocumentReferenceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1TransportDocumentsTransportDocumentReferenceGet`: TransportDocument
    fmt.Fprintf(os.Stdout, "Response from `TransportDocumentsApi.V1TransportDocumentsTransportDocumentReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportDocumentReference** | **string** | The transportDocumentReference of the transportDocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1TransportDocumentsTransportDocumentReferenceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**TransportDocument**](TransportDocument.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

