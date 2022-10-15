# ShippingInstructionHeaderNoID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportDocumentType** | [**TransportDocumentType**](TransportDocumentType.md) |  | 
**IsShippedOnboardType** | **bool** | Specifies whether the Transport document is a received for shipment, or shipped onboard. | 
**NumberOfCopies** | Pointer to **int32** | The requested number of copies of the Transport document to be issued by the carrier. Only applicable for physical documents | [optional] 
**NumberOfOriginals** | Pointer to **int32** | Number of original copies of the negotiable bill of lading that has been issued to the customer. | [optional] 
**PreCarriageUnderShippersResponsibility** | Pointer to [**PreCarriageUnderShippersResponsibility**](PreCarriageUnderShippersResponsibility.md) |  | [optional] 
**InvoicePayableAt** | **interface{}** | A location object - here with an example with only UN location code and City name present | 
**IsElectronic** | **bool** | An indicator whether the transport document is electronically transferred. | 
**IsChargesDisplayed** | Pointer to **bool** | An indicator whether the transport document should include charges. | [optional] 
**CarrierBookingReference** | Pointer to **interface{}** | The associated booking number provided by the carrier. Cannot be used in combination with Cargo Item level carrierBookingReference | [optional] 

## Methods

### NewShippingInstructionHeaderNoID

`func NewShippingInstructionHeaderNoID(transportDocumentType TransportDocumentType, isShippedOnboardType bool, invoicePayableAt interface{}, isElectronic bool, ) *ShippingInstructionHeaderNoID`

NewShippingInstructionHeaderNoID instantiates a new ShippingInstructionHeaderNoID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingInstructionHeaderNoIDWithDefaults

`func NewShippingInstructionHeaderNoIDWithDefaults() *ShippingInstructionHeaderNoID`

NewShippingInstructionHeaderNoIDWithDefaults instantiates a new ShippingInstructionHeaderNoID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportDocumentType

`func (o *ShippingInstructionHeaderNoID) GetTransportDocumentType() TransportDocumentType`

GetTransportDocumentType returns the TransportDocumentType field if non-nil, zero value otherwise.

### GetTransportDocumentTypeOk

`func (o *ShippingInstructionHeaderNoID) GetTransportDocumentTypeOk() (*TransportDocumentType, bool)`

GetTransportDocumentTypeOk returns a tuple with the TransportDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportDocumentType

`func (o *ShippingInstructionHeaderNoID) SetTransportDocumentType(v TransportDocumentType)`

SetTransportDocumentType sets TransportDocumentType field to given value.


### GetIsShippedOnboardType

`func (o *ShippingInstructionHeaderNoID) GetIsShippedOnboardType() bool`

GetIsShippedOnboardType returns the IsShippedOnboardType field if non-nil, zero value otherwise.

### GetIsShippedOnboardTypeOk

`func (o *ShippingInstructionHeaderNoID) GetIsShippedOnboardTypeOk() (*bool, bool)`

GetIsShippedOnboardTypeOk returns a tuple with the IsShippedOnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippedOnboardType

`func (o *ShippingInstructionHeaderNoID) SetIsShippedOnboardType(v bool)`

SetIsShippedOnboardType sets IsShippedOnboardType field to given value.


### GetNumberOfCopies

`func (o *ShippingInstructionHeaderNoID) GetNumberOfCopies() int32`

GetNumberOfCopies returns the NumberOfCopies field if non-nil, zero value otherwise.

### GetNumberOfCopiesOk

`func (o *ShippingInstructionHeaderNoID) GetNumberOfCopiesOk() (*int32, bool)`

GetNumberOfCopiesOk returns a tuple with the NumberOfCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCopies

`func (o *ShippingInstructionHeaderNoID) SetNumberOfCopies(v int32)`

SetNumberOfCopies sets NumberOfCopies field to given value.

### HasNumberOfCopies

`func (o *ShippingInstructionHeaderNoID) HasNumberOfCopies() bool`

HasNumberOfCopies returns a boolean if a field has been set.

### GetNumberOfOriginals

`func (o *ShippingInstructionHeaderNoID) GetNumberOfOriginals() int32`

GetNumberOfOriginals returns the NumberOfOriginals field if non-nil, zero value otherwise.

### GetNumberOfOriginalsOk

`func (o *ShippingInstructionHeaderNoID) GetNumberOfOriginalsOk() (*int32, bool)`

GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOriginals

`func (o *ShippingInstructionHeaderNoID) SetNumberOfOriginals(v int32)`

SetNumberOfOriginals sets NumberOfOriginals field to given value.

### HasNumberOfOriginals

`func (o *ShippingInstructionHeaderNoID) HasNumberOfOriginals() bool`

HasNumberOfOriginals returns a boolean if a field has been set.

### GetPreCarriageUnderShippersResponsibility

`func (o *ShippingInstructionHeaderNoID) GetPreCarriageUnderShippersResponsibility() PreCarriageUnderShippersResponsibility`

GetPreCarriageUnderShippersResponsibility returns the PreCarriageUnderShippersResponsibility field if non-nil, zero value otherwise.

### GetPreCarriageUnderShippersResponsibilityOk

`func (o *ShippingInstructionHeaderNoID) GetPreCarriageUnderShippersResponsibilityOk() (*PreCarriageUnderShippersResponsibility, bool)`

GetPreCarriageUnderShippersResponsibilityOk returns a tuple with the PreCarriageUnderShippersResponsibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCarriageUnderShippersResponsibility

`func (o *ShippingInstructionHeaderNoID) SetPreCarriageUnderShippersResponsibility(v PreCarriageUnderShippersResponsibility)`

SetPreCarriageUnderShippersResponsibility sets PreCarriageUnderShippersResponsibility field to given value.

### HasPreCarriageUnderShippersResponsibility

`func (o *ShippingInstructionHeaderNoID) HasPreCarriageUnderShippersResponsibility() bool`

HasPreCarriageUnderShippersResponsibility returns a boolean if a field has been set.

### GetInvoicePayableAt

`func (o *ShippingInstructionHeaderNoID) GetInvoicePayableAt() interface{}`

GetInvoicePayableAt returns the InvoicePayableAt field if non-nil, zero value otherwise.

### GetInvoicePayableAtOk

`func (o *ShippingInstructionHeaderNoID) GetInvoicePayableAtOk() (*interface{}, bool)`

GetInvoicePayableAtOk returns a tuple with the InvoicePayableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayableAt

`func (o *ShippingInstructionHeaderNoID) SetInvoicePayableAt(v interface{})`

SetInvoicePayableAt sets InvoicePayableAt field to given value.


### SetInvoicePayableAtNil

`func (o *ShippingInstructionHeaderNoID) SetInvoicePayableAtNil(b bool)`

 SetInvoicePayableAtNil sets the value for InvoicePayableAt to be an explicit nil

### UnsetInvoicePayableAt
`func (o *ShippingInstructionHeaderNoID) UnsetInvoicePayableAt()`

UnsetInvoicePayableAt ensures that no value is present for InvoicePayableAt, not even an explicit nil
### GetIsElectronic

`func (o *ShippingInstructionHeaderNoID) GetIsElectronic() bool`

GetIsElectronic returns the IsElectronic field if non-nil, zero value otherwise.

### GetIsElectronicOk

`func (o *ShippingInstructionHeaderNoID) GetIsElectronicOk() (*bool, bool)`

GetIsElectronicOk returns a tuple with the IsElectronic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsElectronic

`func (o *ShippingInstructionHeaderNoID) SetIsElectronic(v bool)`

SetIsElectronic sets IsElectronic field to given value.


### GetIsChargesDisplayed

`func (o *ShippingInstructionHeaderNoID) GetIsChargesDisplayed() bool`

GetIsChargesDisplayed returns the IsChargesDisplayed field if non-nil, zero value otherwise.

### GetIsChargesDisplayedOk

`func (o *ShippingInstructionHeaderNoID) GetIsChargesDisplayedOk() (*bool, bool)`

GetIsChargesDisplayedOk returns a tuple with the IsChargesDisplayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChargesDisplayed

`func (o *ShippingInstructionHeaderNoID) SetIsChargesDisplayed(v bool)`

SetIsChargesDisplayed sets IsChargesDisplayed field to given value.

### HasIsChargesDisplayed

`func (o *ShippingInstructionHeaderNoID) HasIsChargesDisplayed() bool`

HasIsChargesDisplayed returns a boolean if a field has been set.

### GetCarrierBookingReference

`func (o *ShippingInstructionHeaderNoID) GetCarrierBookingReference() interface{}`

GetCarrierBookingReference returns the CarrierBookingReference field if non-nil, zero value otherwise.

### GetCarrierBookingReferenceOk

`func (o *ShippingInstructionHeaderNoID) GetCarrierBookingReferenceOk() (*interface{}, bool)`

GetCarrierBookingReferenceOk returns a tuple with the CarrierBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierBookingReference

`func (o *ShippingInstructionHeaderNoID) SetCarrierBookingReference(v interface{})`

SetCarrierBookingReference sets CarrierBookingReference field to given value.

### HasCarrierBookingReference

`func (o *ShippingInstructionHeaderNoID) HasCarrierBookingReference() bool`

HasCarrierBookingReference returns a boolean if a field has been set.

### SetCarrierBookingReferenceNil

`func (o *ShippingInstructionHeaderNoID) SetCarrierBookingReferenceNil(b bool)`

 SetCarrierBookingReferenceNil sets the value for CarrierBookingReference to be an explicit nil

### UnsetCarrierBookingReference
`func (o *ShippingInstructionHeaderNoID) UnsetCarrierBookingReference()`

UnsetCarrierBookingReference ensures that no value is present for CarrierBookingReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


