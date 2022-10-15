# TransportDocument

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
**ShippingInstruction** | [**ShippingInstruction**](shippingInstruction.md) |  | 
**Charges** | Pointer to [**[]Charge**](Charge.md) |  | [optional] 
**Clauses** | Pointer to [**[]Clause**](Clause.md) |  | [optional] 
**Transports** | [**[]Transport**](Transport.md) |  | 

## Methods

### NewTransportDocument

`func NewTransportDocument(transportDocumentReference string, issueDate string, termsAndConditions string, cargoMovementTypeAtOrigin CargoMovementTypeAtOrigin, cargoMovementTypeAtDestination CargoMovementTypeAtDestination, receiptDeliveryTypeAtOrigin ReceiptDeliveryTypeAtOrigin, receiptDeliveryTypeAtDestination ReceiptDeliveryTypeAtDestination, shippingInstruction ShippingInstruction, transports []Transport, ) *TransportDocument`

NewTransportDocument instantiates a new TransportDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportDocumentWithDefaults

`func NewTransportDocumentWithDefaults() *TransportDocument`

NewTransportDocumentWithDefaults instantiates a new TransportDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportDocumentReference

`func (o *TransportDocument) GetTransportDocumentReference() string`

GetTransportDocumentReference returns the TransportDocumentReference field if non-nil, zero value otherwise.

### GetTransportDocumentReferenceOk

`func (o *TransportDocument) GetTransportDocumentReferenceOk() (*string, bool)`

GetTransportDocumentReferenceOk returns a tuple with the TransportDocumentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportDocumentReference

`func (o *TransportDocument) SetTransportDocumentReference(v string)`

SetTransportDocumentReference sets TransportDocumentReference field to given value.


### GetPlaceOfIssue

`func (o *TransportDocument) GetPlaceOfIssue() Location`

GetPlaceOfIssue returns the PlaceOfIssue field if non-nil, zero value otherwise.

### GetPlaceOfIssueOk

`func (o *TransportDocument) GetPlaceOfIssueOk() (*Location, bool)`

GetPlaceOfIssueOk returns a tuple with the PlaceOfIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfIssue

`func (o *TransportDocument) SetPlaceOfIssue(v Location)`

SetPlaceOfIssue sets PlaceOfIssue field to given value.

### HasPlaceOfIssue

`func (o *TransportDocument) HasPlaceOfIssue() bool`

HasPlaceOfIssue returns a boolean if a field has been set.

### GetIssueDate

`func (o *TransportDocument) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *TransportDocument) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *TransportDocument) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetShippedOnBoardDate

`func (o *TransportDocument) GetShippedOnBoardDate() string`

GetShippedOnBoardDate returns the ShippedOnBoardDate field if non-nil, zero value otherwise.

### GetShippedOnBoardDateOk

`func (o *TransportDocument) GetShippedOnBoardDateOk() (*string, bool)`

GetShippedOnBoardDateOk returns a tuple with the ShippedOnBoardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedOnBoardDate

`func (o *TransportDocument) SetShippedOnBoardDate(v string)`

SetShippedOnBoardDate sets ShippedOnBoardDate field to given value.

### HasShippedOnBoardDate

`func (o *TransportDocument) HasShippedOnBoardDate() bool`

HasShippedOnBoardDate returns a boolean if a field has been set.

### GetReceivedForShipmentDate

`func (o *TransportDocument) GetReceivedForShipmentDate() string`

GetReceivedForShipmentDate returns the ReceivedForShipmentDate field if non-nil, zero value otherwise.

### GetReceivedForShipmentDateOk

`func (o *TransportDocument) GetReceivedForShipmentDateOk() (*string, bool)`

GetReceivedForShipmentDateOk returns a tuple with the ReceivedForShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedForShipmentDate

`func (o *TransportDocument) SetReceivedForShipmentDate(v string)`

SetReceivedForShipmentDate sets ReceivedForShipmentDate field to given value.

### HasReceivedForShipmentDate

`func (o *TransportDocument) HasReceivedForShipmentDate() bool`

HasReceivedForShipmentDate returns a boolean if a field has been set.

### GetTermsAndConditions

`func (o *TransportDocument) GetTermsAndConditions() string`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *TransportDocument) GetTermsAndConditionsOk() (*string, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *TransportDocument) SetTermsAndConditions(v string)`

SetTermsAndConditions sets TermsAndConditions field to given value.


### GetIssuerCode

`func (o *TransportDocument) GetIssuerCode() string`

GetIssuerCode returns the IssuerCode field if non-nil, zero value otherwise.

### GetIssuerCodeOk

`func (o *TransportDocument) GetIssuerCodeOk() (*string, bool)`

GetIssuerCodeOk returns a tuple with the IssuerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCode

`func (o *TransportDocument) SetIssuerCode(v string)`

SetIssuerCode sets IssuerCode field to given value.

### HasIssuerCode

`func (o *TransportDocument) HasIssuerCode() bool`

HasIssuerCode returns a boolean if a field has been set.

### GetIssuerCodeListProvider

`func (o *TransportDocument) GetIssuerCodeListProvider() IssuerCodeListProvider`

GetIssuerCodeListProvider returns the IssuerCodeListProvider field if non-nil, zero value otherwise.

### GetIssuerCodeListProviderOk

`func (o *TransportDocument) GetIssuerCodeListProviderOk() (*IssuerCodeListProvider, bool)`

GetIssuerCodeListProviderOk returns a tuple with the IssuerCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCodeListProvider

`func (o *TransportDocument) SetIssuerCodeListProvider(v IssuerCodeListProvider)`

SetIssuerCodeListProvider sets IssuerCodeListProvider field to given value.

### HasIssuerCodeListProvider

`func (o *TransportDocument) HasIssuerCodeListProvider() bool`

HasIssuerCodeListProvider returns a boolean if a field has been set.

### GetDeclaredValueCurrency

`func (o *TransportDocument) GetDeclaredValueCurrency() string`

GetDeclaredValueCurrency returns the DeclaredValueCurrency field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyOk

`func (o *TransportDocument) GetDeclaredValueCurrencyOk() (*string, bool)`

GetDeclaredValueCurrencyOk returns a tuple with the DeclaredValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrency

`func (o *TransportDocument) SetDeclaredValueCurrency(v string)`

SetDeclaredValueCurrency sets DeclaredValueCurrency field to given value.

### HasDeclaredValueCurrency

`func (o *TransportDocument) HasDeclaredValueCurrency() bool`

HasDeclaredValueCurrency returns a boolean if a field has been set.

### GetDeclaredValue

`func (o *TransportDocument) GetDeclaredValue() float32`

GetDeclaredValue returns the DeclaredValue field if non-nil, zero value otherwise.

### GetDeclaredValueOk

`func (o *TransportDocument) GetDeclaredValueOk() (*float32, bool)`

GetDeclaredValueOk returns a tuple with the DeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValue

`func (o *TransportDocument) SetDeclaredValue(v float32)`

SetDeclaredValue sets DeclaredValue field to given value.

### HasDeclaredValue

`func (o *TransportDocument) HasDeclaredValue() bool`

HasDeclaredValue returns a boolean if a field has been set.

### GetNumberOfRiderPages

`func (o *TransportDocument) GetNumberOfRiderPages() int32`

GetNumberOfRiderPages returns the NumberOfRiderPages field if non-nil, zero value otherwise.

### GetNumberOfRiderPagesOk

`func (o *TransportDocument) GetNumberOfRiderPagesOk() (*int32, bool)`

GetNumberOfRiderPagesOk returns a tuple with the NumberOfRiderPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRiderPages

`func (o *TransportDocument) SetNumberOfRiderPages(v int32)`

SetNumberOfRiderPages sets NumberOfRiderPages field to given value.

### HasNumberOfRiderPages

`func (o *TransportDocument) HasNumberOfRiderPages() bool`

HasNumberOfRiderPages returns a boolean if a field has been set.

### GetCargoMovementTypeAtOrigin

`func (o *TransportDocument) GetCargoMovementTypeAtOrigin() CargoMovementTypeAtOrigin`

GetCargoMovementTypeAtOrigin returns the CargoMovementTypeAtOrigin field if non-nil, zero value otherwise.

### GetCargoMovementTypeAtOriginOk

`func (o *TransportDocument) GetCargoMovementTypeAtOriginOk() (*CargoMovementTypeAtOrigin, bool)`

GetCargoMovementTypeAtOriginOk returns a tuple with the CargoMovementTypeAtOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoMovementTypeAtOrigin

`func (o *TransportDocument) SetCargoMovementTypeAtOrigin(v CargoMovementTypeAtOrigin)`

SetCargoMovementTypeAtOrigin sets CargoMovementTypeAtOrigin field to given value.


### GetCargoMovementTypeAtDestination

`func (o *TransportDocument) GetCargoMovementTypeAtDestination() CargoMovementTypeAtDestination`

GetCargoMovementTypeAtDestination returns the CargoMovementTypeAtDestination field if non-nil, zero value otherwise.

### GetCargoMovementTypeAtDestinationOk

`func (o *TransportDocument) GetCargoMovementTypeAtDestinationOk() (*CargoMovementTypeAtDestination, bool)`

GetCargoMovementTypeAtDestinationOk returns a tuple with the CargoMovementTypeAtDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoMovementTypeAtDestination

`func (o *TransportDocument) SetCargoMovementTypeAtDestination(v CargoMovementTypeAtDestination)`

SetCargoMovementTypeAtDestination sets CargoMovementTypeAtDestination field to given value.


### GetReceiptDeliveryTypeAtOrigin

`func (o *TransportDocument) GetReceiptDeliveryTypeAtOrigin() ReceiptDeliveryTypeAtOrigin`

GetReceiptDeliveryTypeAtOrigin returns the ReceiptDeliveryTypeAtOrigin field if non-nil, zero value otherwise.

### GetReceiptDeliveryTypeAtOriginOk

`func (o *TransportDocument) GetReceiptDeliveryTypeAtOriginOk() (*ReceiptDeliveryTypeAtOrigin, bool)`

GetReceiptDeliveryTypeAtOriginOk returns a tuple with the ReceiptDeliveryTypeAtOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDeliveryTypeAtOrigin

`func (o *TransportDocument) SetReceiptDeliveryTypeAtOrigin(v ReceiptDeliveryTypeAtOrigin)`

SetReceiptDeliveryTypeAtOrigin sets ReceiptDeliveryTypeAtOrigin field to given value.


### GetReceiptDeliveryTypeAtDestination

`func (o *TransportDocument) GetReceiptDeliveryTypeAtDestination() ReceiptDeliveryTypeAtDestination`

GetReceiptDeliveryTypeAtDestination returns the ReceiptDeliveryTypeAtDestination field if non-nil, zero value otherwise.

### GetReceiptDeliveryTypeAtDestinationOk

`func (o *TransportDocument) GetReceiptDeliveryTypeAtDestinationOk() (*ReceiptDeliveryTypeAtDestination, bool)`

GetReceiptDeliveryTypeAtDestinationOk returns a tuple with the ReceiptDeliveryTypeAtDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDeliveryTypeAtDestination

`func (o *TransportDocument) SetReceiptDeliveryTypeAtDestination(v ReceiptDeliveryTypeAtDestination)`

SetReceiptDeliveryTypeAtDestination sets ReceiptDeliveryTypeAtDestination field to given value.


### GetServiceContractReference

`func (o *TransportDocument) GetServiceContractReference() string`

GetServiceContractReference returns the ServiceContractReference field if non-nil, zero value otherwise.

### GetServiceContractReferenceOk

`func (o *TransportDocument) GetServiceContractReferenceOk() (*string, bool)`

GetServiceContractReferenceOk returns a tuple with the ServiceContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContractReference

`func (o *TransportDocument) SetServiceContractReference(v string)`

SetServiceContractReference sets ServiceContractReference field to given value.

### HasServiceContractReference

`func (o *TransportDocument) HasServiceContractReference() bool`

HasServiceContractReference returns a boolean if a field has been set.

### GetShippingInstruction

`func (o *TransportDocument) GetShippingInstruction() ShippingInstruction`

GetShippingInstruction returns the ShippingInstruction field if non-nil, zero value otherwise.

### GetShippingInstructionOk

`func (o *TransportDocument) GetShippingInstructionOk() (*ShippingInstruction, bool)`

GetShippingInstructionOk returns a tuple with the ShippingInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInstruction

`func (o *TransportDocument) SetShippingInstruction(v ShippingInstruction)`

SetShippingInstruction sets ShippingInstruction field to given value.


### GetCharges

`func (o *TransportDocument) GetCharges() []Charge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *TransportDocument) GetChargesOk() (*[]Charge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *TransportDocument) SetCharges(v []Charge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *TransportDocument) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetClauses

`func (o *TransportDocument) GetClauses() []Clause`

GetClauses returns the Clauses field if non-nil, zero value otherwise.

### GetClausesOk

`func (o *TransportDocument) GetClausesOk() (*[]Clause, bool)`

GetClausesOk returns a tuple with the Clauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauses

`func (o *TransportDocument) SetClauses(v []Clause)`

SetClauses sets Clauses field to given value.

### HasClauses

`func (o *TransportDocument) HasClauses() bool`

HasClauses returns a boolean if a field has been set.

### GetTransports

`func (o *TransportDocument) GetTransports() []Transport`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *TransportDocument) GetTransportsOk() (*[]Transport, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *TransportDocument) SetTransports(v []Transport)`

SetTransports sets Transports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


