# \EventsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EventsGet**](EventsApi.md#V1EventsGet) | **Get** /v1/events/ | Get Shipment events



## V1EventsGet

> []ShipmentEvent V1EventsGet(ctx, transportDocumentReference).ShipmentEventTypeCode(shipmentEventTypeCode).CarrierBookingReference(carrierBookingReference).TransportDocumentTypeCode(transportDocumentTypeCode).Limit(limit).Cursor(cursor).APIVersion(aPIVersion).Execute()

Get Shipment events



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
    shipmentEventTypeCode := []openapiclient.ShipmentEventTypeCode{openapiclient.shipmentEventTypeCode("RECE")} // []ShipmentEventTypeCode | The status of the document in the process to filter by. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example <i>shipmentEventTypeCode=RECE,DRFT</i>  Matches <b>both</b> Received (RECE) and Drafted (DRFT) shipment events.  Default is all shipmentEventTypeCodes.  This filter is only relevant when filtering on ShipmentEvents  (optional)
    carrierBookingReference := "carrierBookingReference_example" // string | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  (optional)
    transportDocumentTypeCode := []openapiclient.TransportDocumentType{openapiclient.transportDocumentType("BOL")} // []TransportDocumentType | Specifies the type of the transport document (a Bill of Lading (BOL) or a Sea Waybill (SWB)) to filter by.  Default is both.  (optional)
    limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
    cursor := "fE9mZnNldHw9MTAmbGltaXQ9MTA=" // string | A server generated value to specify a specific point in a collection result, used for pagination. (optional)
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.V1EventsGet(context.Background(), transportDocumentReference).ShipmentEventTypeCode(shipmentEventTypeCode).CarrierBookingReference(carrierBookingReference).TransportDocumentTypeCode(transportDocumentTypeCode).Limit(limit).Cursor(cursor).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.V1EventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EventsGet`: []ShipmentEvent
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.V1EventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transportDocumentReference** | **string** | The transportDocumentReference of the transportDocument | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shipmentEventTypeCode** | [**[]ShipmentEventTypeCode**](ShipmentEventTypeCode.md) | The status of the document in the process to filter by. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example &lt;i&gt;shipmentEventTypeCode&#x3D;RECE,DRFT&lt;/i&gt;  Matches &lt;b&gt;both&lt;/b&gt; Received (RECE) and Drafted (DRFT) shipment events.  Default is all shipmentEventTypeCodes.  This filter is only relevant when filtering on ShipmentEvents  | 
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking.  Specifying this filter will only return events related to this particular carrierBookingReference.  | 
 **transportDocumentTypeCode** | [**[]TransportDocumentType**](TransportDocumentType.md) | Specifies the type of the transport document (a Bill of Lading (BOL) or a Sea Waybill (SWB)) to filter by.  Default is both.  | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **cursor** | **string** | A server generated value to specify a specific point in a collection result, used for pagination. | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**[]ShipmentEvent**](ShipmentEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

