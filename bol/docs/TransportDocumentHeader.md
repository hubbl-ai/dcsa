# TransportDocumentHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportDocumentReference** | **string** | A unique number allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment.  | 
**PlaceOfIssue** | Pointer to [**Location**](location.md) |  | [optional] 
**IssueDate** | **string** | Date when the transport document has been issued | 
**ShippedOnBoardDate** | Pointer to **string** | Date when the last container that is linked to the transport document is physically loaded onboard the vessel indicated on the transport document. | [optional] 
**ReceivedForShipmentDate** | Pointer to **string** | Date when the last container linked to the transport document is physically in the terminal (customers cleared against the intended vessel). | [optional] 
**TermsAndConditions** | **string** | Carrier general terms and conditions for this transport document. | 
**IssuerCode** | Pointer to **string** | The code used for identifying the specific issuer.  | [optional] 
**IssuerCodeListProvider** | Pointer to [**IssuerCodeListProvider**](IssuerCodeListProvider.md) |  | [optional] 
**DeclaredValueCurrency** | Pointer to **string** | The currency used for the declared value, using the 3-character code defined by ISO 4217. | [optional] 
**DeclaredValue** | Pointer to **float32** | The value of the cargo that the shipper declares to avoid the carrier&amp;apos;s limitation of liability and \&quot;Ad Valorem\&quot; freight, i.e. freight which is calculated based on the value of the goods declared by the shipper. | [optional] 
**NumberOfRiderPages** | Pointer to **int32** | The number of additional pages required to contain the goods description on a transport document. Only applicable for physical transport documents. | [optional] 
**CargoMovementTypeAtOrigin** | [**CargoMovementTypeAtOrigin**](CargoMovementTypeAtOrigin.md) |  | 
**CargoMovementTypeAtDestination** | [**CargoMovementTypeAtDestination**](CargoMovementTypeAtDestination.md) |  | 
**ReceiptDeliveryTypeAtOrigin** | [**ReceiptDeliveryTypeAtOrigin**](ReceiptDeliveryTypeAtOrigin.md) |  | 
**ReceiptDeliveryTypeAtDestination** | [**ReceiptDeliveryTypeAtDestination**](ReceiptDeliveryTypeAtDestination.md) |  | 
**ServiceContractReference** | Pointer to **string** | Reference number for agreement between shipper and carrier through which the shipper commits to provide a certain minimum quantity of cargo over a fixed period, and the carrier commits to a certain rate or rate schedule. | [optional] 

## Methods

### NewTransportDocumentHeader

`func NewTransportDocumentHeader(transportDocumentReference string, issueDate string, termsAndConditions string, cargoMovementTypeAtOrigin CargoMovementTypeAtOrigin, cargoMovementTypeAtDestination CargoMovementTypeAtDestination, receiptDeliveryTypeAtOrigin ReceiptDeliveryTypeAtOrigin, receiptDeliveryTypeAtDestination ReceiptDeliveryTypeAtDestination, ) *TransportDocumentHeader`

NewTransportDocumentHeader instantiates a new TransportDocumentHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportDocumentHeaderWithDefaults

`func NewTransportDocumentHeaderWithDefaults() *TransportDocumentHeader`

NewTransportDocumentHeaderWithDefaults instantiates a new TransportDocumentHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportDocumentReference

`func (o *TransportDocumentHeader) GetTransportDocumentReference() string`

GetTransportDocumentReference returns the TransportDocumentReference field if non-nil, zero value otherwise.

### GetTransportDocumentReferenceOk

`func (o *TransportDocumentHeader) GetTransportDocumentReferenceOk() (*string, bool)`

GetTransportDocumentReferenceOk returns a tuple with the TransportDocumentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportDocumentReference

`func (o *TransportDocumentHeader) SetTransportDocumentReference(v string)`

SetTransportDocumentReference sets TransportDocumentReference field to given value.


### GetPlaceOfIssue

`func (o *TransportDocumentHeader) GetPlaceOfIssue() Location`

GetPlaceOfIssue returns the PlaceOfIssue field if non-nil, zero value otherwise.

### GetPlaceOfIssueOk

`func (o *TransportDocumentHeader) GetPlaceOfIssueOk() (*Location, bool)`

GetPlaceOfIssueOk returns a tuple with the PlaceOfIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfIssue

`func (o *TransportDocumentHeader) SetPlaceOfIssue(v Location)`

SetPlaceOfIssue sets PlaceOfIssue field to given value.

### HasPlaceOfIssue

`func (o *TransportDocumentHeader) HasPlaceOfIssue() bool`

HasPlaceOfIssue returns a boolean if a field has been set.

### GetIssueDate

`func (o *TransportDocumentHeader) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *TransportDocumentHeader) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *TransportDocumentHeader) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetShippedOnBoardDate

`func (o *TransportDocumentHeader) GetShippedOnBoardDate() string`

GetShippedOnBoardDate returns the ShippedOnBoardDate field if non-nil, zero value otherwise.

### GetShippedOnBoardDateOk

`func (o *TransportDocumentHeader) GetShippedOnBoardDateOk() (*string, bool)`

GetShippedOnBoardDateOk returns a tuple with the ShippedOnBoardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedOnBoardDate

`func (o *TransportDocumentHeader) SetShippedOnBoardDate(v string)`

SetShippedOnBoardDate sets ShippedOnBoardDate field to given value.

### HasShippedOnBoardDate

`func (o *TransportDocumentHeader) HasShippedOnBoardDate() bool`

HasShippedOnBoardDate returns a boolean if a field has been set.

### GetReceivedForShipmentDate

`func (o *TransportDocumentHeader) GetReceivedForShipmentDate() string`

GetReceivedForShipmentDate returns the ReceivedForShipmentDate field if non-nil, zero value otherwise.

### GetReceivedForShipmentDateOk

`func (o *TransportDocumentHeader) GetReceivedForShipmentDateOk() (*string, bool)`

GetReceivedForShipmentDateOk returns a tuple with the ReceivedForShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedForShipmentDate

`func (o *TransportDocumentHeader) SetReceivedForShipmentDate(v string)`

SetReceivedForShipmentDate sets ReceivedForShipmentDate field to given value.

### HasReceivedForShipmentDate

`func (o *TransportDocumentHeader) HasReceivedForShipmentDate() bool`

HasReceivedForShipmentDate returns a boolean if a field has been set.

### GetTermsAndConditions

`func (o *TransportDocumentHeader) GetTermsAndConditions() string`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *TransportDocumentHeader) GetTermsAndConditionsOk() (*string, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *TransportDocumentHeader) SetTermsAndConditions(v string)`

SetTermsAndConditions sets TermsAndConditions field to given value.


### GetIssuerCode

`func (o *TransportDocumentHeader) GetIssuerCode() string`

GetIssuerCode returns the IssuerCode field if non-nil, zero value otherwise.

### GetIssuerCodeOk

`func (o *TransportDocumentHeader) GetIssuerCodeOk() (*string, bool)`

GetIssuerCodeOk returns a tuple with the IssuerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCode

`func (o *TransportDocumentHeader) SetIssuerCode(v string)`

SetIssuerCode sets IssuerCode field to given value.

### HasIssuerCode

`func (o *TransportDocumentHeader) HasIssuerCode() bool`

HasIssuerCode returns a boolean if a field has been set.

### GetIssuerCodeListProvider

`func (o *TransportDocumentHeader) GetIssuerCodeListProvider() IssuerCodeListProvider`

GetIssuerCodeListProvider returns the IssuerCodeListProvider field if non-nil, zero value otherwise.

### GetIssuerCodeListProviderOk

`func (o *TransportDocumentHeader) GetIssuerCodeListProviderOk() (*IssuerCodeListProvider, bool)`

GetIssuerCodeListProviderOk returns a tuple with the IssuerCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCodeListProvider

`func (o *TransportDocumentHeader) SetIssuerCodeListProvider(v IssuerCodeListProvider)`

SetIssuerCodeListProvider sets IssuerCodeListProvider field to given value.

### HasIssuerCodeListProvider

`func (o *TransportDocumentHeader) HasIssuerCodeListProvider() bool`

HasIssuerCodeListProvider returns a boolean if a field has been set.

### GetDeclaredValueCurrency

`func (o *TransportDocumentHeader) GetDeclaredValueCurrency() string`

GetDeclaredValueCurrency returns the DeclaredValueCurrency field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyOk

`func (o *TransportDocumentHeader) GetDeclaredValueCurrencyOk() (*string, bool)`

GetDeclaredValueCurrencyOk returns a tuple with the DeclaredValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrency

`func (o *TransportDocumentHeader) SetDeclaredValueCurrency(v string)`

SetDeclaredValueCurrency sets DeclaredValueCurrency field to given value.

### HasDeclaredValueCurrency

`func (o *TransportDocumentHeader) HasDeclaredValueCurrency() bool`

HasDeclaredValueCurrency returns a boolean if a field has been set.

### GetDeclaredValue

`func (o *TransportDocumentHeader) GetDeclaredValue() float32`

GetDeclaredValue returns the DeclaredValue field if non-nil, zero value otherwise.

### GetDeclaredValueOk

`func (o *TransportDocumentHeader) GetDeclaredValueOk() (*float32, bool)`

GetDeclaredValueOk returns a tuple with the DeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValue

`func (o *TransportDocumentHeader) SetDeclaredValue(v float32)`

SetDeclaredValue sets DeclaredValue field to given value.

### HasDeclaredValue

`func (o *TransportDocumentHeader) HasDeclaredValue() bool`

HasDeclaredValue returns a boolean if a field has been set.

### GetNumberOfRiderPages

`func (o *TransportDocumentHeader) GetNumberOfRiderPages() int32`

GetNumberOfRiderPages returns the NumberOfRiderPages field if non-nil, zero value otherwise.

### GetNumberOfRiderPagesOk

`func (o *TransportDocumentHeader) GetNumberOfRiderPagesOk() (*int32, bool)`

GetNumberOfRiderPagesOk returns a tuple with the NumberOfRiderPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRiderPages

`func (o *TransportDocumentHeader) SetNumberOfRiderPages(v int32)`

SetNumberOfRiderPages sets NumberOfRiderPages field to given value.

### HasNumberOfRiderPages

`func (o *TransportDocumentHeader) HasNumberOfRiderPages() bool`

HasNumberOfRiderPages returns a boolean if a field has been set.

### GetCargoMovementTypeAtOrigin

`func (o *TransportDocumentHeader) GetCargoMovementTypeAtOrigin() CargoMovementTypeAtOrigin`

GetCargoMovementTypeAtOrigin returns the CargoMovementTypeAtOrigin field if non-nil, zero value otherwise.

### GetCargoMovementTypeAtOriginOk

`func (o *TransportDocumentHeader) GetCargoMovementTypeAtOriginOk() (*CargoMovementTypeAtOrigin, bool)`

GetCargoMovementTypeAtOriginOk returns a tuple with the CargoMovementTypeAtOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoMovementTypeAtOrigin

`func (o *TransportDocumentHeader) SetCargoMovementTypeAtOrigin(v CargoMovementTypeAtOrigin)`

SetCargoMovementTypeAtOrigin sets CargoMovementTypeAtOrigin field to given value.


### GetCargoMovementTypeAtDestination

`func (o *TransportDocumentHeader) GetCargoMovementTypeAtDestination() CargoMovementTypeAtDestination`

GetCargoMovementTypeAtDestination returns the CargoMovementTypeAtDestination field if non-nil, zero value otherwise.

### GetCargoMovementTypeAtDestinationOk

`func (o *TransportDocumentHeader) GetCargoMovementTypeAtDestinationOk() (*CargoMovementTypeAtDestination, bool)`

GetCargoMovementTypeAtDestinationOk returns a tuple with the CargoMovementTypeAtDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoMovementTypeAtDestination

`func (o *TransportDocumentHeader) SetCargoMovementTypeAtDestination(v CargoMovementTypeAtDestination)`

SetCargoMovementTypeAtDestination sets CargoMovementTypeAtDestination field to given value.


### GetReceiptDeliveryTypeAtOrigin

`func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtOrigin() ReceiptDeliveryTypeAtOrigin`

GetReceiptDeliveryTypeAtOrigin returns the ReceiptDeliveryTypeAtOrigin field if non-nil, zero value otherwise.

### GetReceiptDeliveryTypeAtOriginOk

`func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtOriginOk() (*ReceiptDeliveryTypeAtOrigin, bool)`

GetReceiptDeliveryTypeAtOriginOk returns a tuple with the ReceiptDeliveryTypeAtOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDeliveryTypeAtOrigin

`func (o *TransportDocumentHeader) SetReceiptDeliveryTypeAtOrigin(v ReceiptDeliveryTypeAtOrigin)`

SetReceiptDeliveryTypeAtOrigin sets ReceiptDeliveryTypeAtOrigin field to given value.


### GetReceiptDeliveryTypeAtDestination

`func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtDestination() ReceiptDeliveryTypeAtDestination`

GetReceiptDeliveryTypeAtDestination returns the ReceiptDeliveryTypeAtDestination field if non-nil, zero value otherwise.

### GetReceiptDeliveryTypeAtDestinationOk

`func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtDestinationOk() (*ReceiptDeliveryTypeAtDestination, bool)`

GetReceiptDeliveryTypeAtDestinationOk returns a tuple with the ReceiptDeliveryTypeAtDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDeliveryTypeAtDestination

`func (o *TransportDocumentHeader) SetReceiptDeliveryTypeAtDestination(v ReceiptDeliveryTypeAtDestination)`

SetReceiptDeliveryTypeAtDestination sets ReceiptDeliveryTypeAtDestination field to given value.


### GetServiceContractReference

`func (o *TransportDocumentHeader) GetServiceContractReference() string`

GetServiceContractReference returns the ServiceContractReference field if non-nil, zero value otherwise.

### GetServiceContractReferenceOk

`func (o *TransportDocumentHeader) GetServiceContractReferenceOk() (*string, bool)`

GetServiceContractReferenceOk returns a tuple with the ServiceContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContractReference

`func (o *TransportDocumentHeader) SetServiceContractReference(v string)`

SetServiceContractReference sets ServiceContractReference field to given value.

### HasServiceContractReference

`func (o *TransportDocumentHeader) HasServiceContractReference() bool`

HasServiceContractReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


