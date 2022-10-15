# ShippingInstructionNoID

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
**CargoItems** | Pointer to [**[]CargoItem**](CargoItem.md) |  | [optional] 
**UtilizedTransportEquipments** | Pointer to [**[]UtilizedTransportEquipment**](UtilizedTransportEquipment.md) |  | [optional] 
**DocumentParties** | Pointer to [**[]DocumentParty**](DocumentParty.md) |  | [optional] 
**ShipmentLocations** | Pointer to [**[]ShipmentLocation**](ShipmentLocation.md) |  | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 

## Methods

### NewShippingInstructionNoID

`func NewShippingInstructionNoID(transportDocumentType TransportDocumentType, isShippedOnboardType bool, invoicePayableAt interface{}, isElectronic bool, ) *ShippingInstructionNoID`

NewShippingInstructionNoID instantiates a new ShippingInstructionNoID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingInstructionNoIDWithDefaults

`func NewShippingInstructionNoIDWithDefaults() *ShippingInstructionNoID`

NewShippingInstructionNoIDWithDefaults instantiates a new ShippingInstructionNoID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportDocumentType

`func (o *ShippingInstructionNoID) GetTransportDocumentType() TransportDocumentType`

GetTransportDocumentType returns the TransportDocumentType field if non-nil, zero value otherwise.

### GetTransportDocumentTypeOk

`func (o *ShippingInstructionNoID) GetTransportDocumentTypeOk() (*TransportDocumentType, bool)`

GetTransportDocumentTypeOk returns a tuple with the TransportDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportDocumentType

`func (o *ShippingInstructionNoID) SetTransportDocumentType(v TransportDocumentType)`

SetTransportDocumentType sets TransportDocumentType field to given value.


### GetIsShippedOnboardType

`func (o *ShippingInstructionNoID) GetIsShippedOnboardType() bool`

GetIsShippedOnboardType returns the IsShippedOnboardType field if non-nil, zero value otherwise.

### GetIsShippedOnboardTypeOk

`func (o *ShippingInstructionNoID) GetIsShippedOnboardTypeOk() (*bool, bool)`

GetIsShippedOnboardTypeOk returns a tuple with the IsShippedOnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippedOnboardType

`func (o *ShippingInstructionNoID) SetIsShippedOnboardType(v bool)`

SetIsShippedOnboardType sets IsShippedOnboardType field to given value.


### GetNumberOfCopies

`func (o *ShippingInstructionNoID) GetNumberOfCopies() int32`

GetNumberOfCopies returns the NumberOfCopies field if non-nil, zero value otherwise.

### GetNumberOfCopiesOk

`func (o *ShippingInstructionNoID) GetNumberOfCopiesOk() (*int32, bool)`

GetNumberOfCopiesOk returns a tuple with the NumberOfCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCopies

`func (o *ShippingInstructionNoID) SetNumberOfCopies(v int32)`

SetNumberOfCopies sets NumberOfCopies field to given value.

### HasNumberOfCopies

`func (o *ShippingInstructionNoID) HasNumberOfCopies() bool`

HasNumberOfCopies returns a boolean if a field has been set.

### GetNumberOfOriginals

`func (o *ShippingInstructionNoID) GetNumberOfOriginals() int32`

GetNumberOfOriginals returns the NumberOfOriginals field if non-nil, zero value otherwise.

### GetNumberOfOriginalsOk

`func (o *ShippingInstructionNoID) GetNumberOfOriginalsOk() (*int32, bool)`

GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOriginals

`func (o *ShippingInstructionNoID) SetNumberOfOriginals(v int32)`

SetNumberOfOriginals sets NumberOfOriginals field to given value.

### HasNumberOfOriginals

`func (o *ShippingInstructionNoID) HasNumberOfOriginals() bool`

HasNumberOfOriginals returns a boolean if a field has been set.

### GetPreCarriageUnderShippersResponsibility

`func (o *ShippingInstructionNoID) GetPreCarriageUnderShippersResponsibility() PreCarriageUnderShippersResponsibility`

GetPreCarriageUnderShippersResponsibility returns the PreCarriageUnderShippersResponsibility field if non-nil, zero value otherwise.

### GetPreCarriageUnderShippersResponsibilityOk

`func (o *ShippingInstructionNoID) GetPreCarriageUnderShippersResponsibilityOk() (*PreCarriageUnderShippersResponsibility, bool)`

GetPreCarriageUnderShippersResponsibilityOk returns a tuple with the PreCarriageUnderShippersResponsibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCarriageUnderShippersResponsibility

`func (o *ShippingInstructionNoID) SetPreCarriageUnderShippersResponsibility(v PreCarriageUnderShippersResponsibility)`

SetPreCarriageUnderShippersResponsibility sets PreCarriageUnderShippersResponsibility field to given value.

### HasPreCarriageUnderShippersResponsibility

`func (o *ShippingInstructionNoID) HasPreCarriageUnderShippersResponsibility() bool`

HasPreCarriageUnderShippersResponsibility returns a boolean if a field has been set.

### GetInvoicePayableAt

`func (o *ShippingInstructionNoID) GetInvoicePayableAt() interface{}`

GetInvoicePayableAt returns the InvoicePayableAt field if non-nil, zero value otherwise.

### GetInvoicePayableAtOk

`func (o *ShippingInstructionNoID) GetInvoicePayableAtOk() (*interface{}, bool)`

GetInvoicePayableAtOk returns a tuple with the InvoicePayableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayableAt

`func (o *ShippingInstructionNoID) SetInvoicePayableAt(v interface{})`

SetInvoicePayableAt sets InvoicePayableAt field to given value.


### SetInvoicePayableAtNil

`func (o *ShippingInstructionNoID) SetInvoicePayableAtNil(b bool)`

 SetInvoicePayableAtNil sets the value for InvoicePayableAt to be an explicit nil

### UnsetInvoicePayableAt
`func (o *ShippingInstructionNoID) UnsetInvoicePayableAt()`

UnsetInvoicePayableAt ensures that no value is present for InvoicePayableAt, not even an explicit nil
### GetIsElectronic

`func (o *ShippingInstructionNoID) GetIsElectronic() bool`

GetIsElectronic returns the IsElectronic field if non-nil, zero value otherwise.

### GetIsElectronicOk

`func (o *ShippingInstructionNoID) GetIsElectronicOk() (*bool, bool)`

GetIsElectronicOk returns a tuple with the IsElectronic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsElectronic

`func (o *ShippingInstructionNoID) SetIsElectronic(v bool)`

SetIsElectronic sets IsElectronic field to given value.


### GetIsChargesDisplayed

`func (o *ShippingInstructionNoID) GetIsChargesDisplayed() bool`

GetIsChargesDisplayed returns the IsChargesDisplayed field if non-nil, zero value otherwise.

### GetIsChargesDisplayedOk

`func (o *ShippingInstructionNoID) GetIsChargesDisplayedOk() (*bool, bool)`

GetIsChargesDisplayedOk returns a tuple with the IsChargesDisplayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChargesDisplayed

`func (o *ShippingInstructionNoID) SetIsChargesDisplayed(v bool)`

SetIsChargesDisplayed sets IsChargesDisplayed field to given value.

### HasIsChargesDisplayed

`func (o *ShippingInstructionNoID) HasIsChargesDisplayed() bool`

HasIsChargesDisplayed returns a boolean if a field has been set.

### GetCarrierBookingReference

`func (o *ShippingInstructionNoID) GetCarrierBookingReference() interface{}`

GetCarrierBookingReference returns the CarrierBookingReference field if non-nil, zero value otherwise.

### GetCarrierBookingReferenceOk

`func (o *ShippingInstructionNoID) GetCarrierBookingReferenceOk() (*interface{}, bool)`

GetCarrierBookingReferenceOk returns a tuple with the CarrierBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierBookingReference

`func (o *ShippingInstructionNoID) SetCarrierBookingReference(v interface{})`

SetCarrierBookingReference sets CarrierBookingReference field to given value.

### HasCarrierBookingReference

`func (o *ShippingInstructionNoID) HasCarrierBookingReference() bool`

HasCarrierBookingReference returns a boolean if a field has been set.

### SetCarrierBookingReferenceNil

`func (o *ShippingInstructionNoID) SetCarrierBookingReferenceNil(b bool)`

 SetCarrierBookingReferenceNil sets the value for CarrierBookingReference to be an explicit nil

### UnsetCarrierBookingReference
`func (o *ShippingInstructionNoID) UnsetCarrierBookingReference()`

UnsetCarrierBookingReference ensures that no value is present for CarrierBookingReference, not even an explicit nil
### GetCargoItems

`func (o *ShippingInstructionNoID) GetCargoItems() []CargoItem`

GetCargoItems returns the CargoItems field if non-nil, zero value otherwise.

### GetCargoItemsOk

`func (o *ShippingInstructionNoID) GetCargoItemsOk() (*[]CargoItem, bool)`

GetCargoItemsOk returns a tuple with the CargoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoItems

`func (o *ShippingInstructionNoID) SetCargoItems(v []CargoItem)`

SetCargoItems sets CargoItems field to given value.

### HasCargoItems

`func (o *ShippingInstructionNoID) HasCargoItems() bool`

HasCargoItems returns a boolean if a field has been set.

### GetUtilizedTransportEquipments

`func (o *ShippingInstructionNoID) GetUtilizedTransportEquipments() []UtilizedTransportEquipment`

GetUtilizedTransportEquipments returns the UtilizedTransportEquipments field if non-nil, zero value otherwise.

### GetUtilizedTransportEquipmentsOk

`func (o *ShippingInstructionNoID) GetUtilizedTransportEquipmentsOk() (*[]UtilizedTransportEquipment, bool)`

GetUtilizedTransportEquipmentsOk returns a tuple with the UtilizedTransportEquipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizedTransportEquipments

`func (o *ShippingInstructionNoID) SetUtilizedTransportEquipments(v []UtilizedTransportEquipment)`

SetUtilizedTransportEquipments sets UtilizedTransportEquipments field to given value.

### HasUtilizedTransportEquipments

`func (o *ShippingInstructionNoID) HasUtilizedTransportEquipments() bool`

HasUtilizedTransportEquipments returns a boolean if a field has been set.

### GetDocumentParties

`func (o *ShippingInstructionNoID) GetDocumentParties() []DocumentParty`

GetDocumentParties returns the DocumentParties field if non-nil, zero value otherwise.

### GetDocumentPartiesOk

`func (o *ShippingInstructionNoID) GetDocumentPartiesOk() (*[]DocumentParty, bool)`

GetDocumentPartiesOk returns a tuple with the DocumentParties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentParties

`func (o *ShippingInstructionNoID) SetDocumentParties(v []DocumentParty)`

SetDocumentParties sets DocumentParties field to given value.

### HasDocumentParties

`func (o *ShippingInstructionNoID) HasDocumentParties() bool`

HasDocumentParties returns a boolean if a field has been set.

### GetShipmentLocations

`func (o *ShippingInstructionNoID) GetShipmentLocations() []ShipmentLocation`

GetShipmentLocations returns the ShipmentLocations field if non-nil, zero value otherwise.

### GetShipmentLocationsOk

`func (o *ShippingInstructionNoID) GetShipmentLocationsOk() (*[]ShipmentLocation, bool)`

GetShipmentLocationsOk returns a tuple with the ShipmentLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLocations

`func (o *ShippingInstructionNoID) SetShipmentLocations(v []ShipmentLocation)`

SetShipmentLocations sets ShipmentLocations field to given value.

### HasShipmentLocations

`func (o *ShippingInstructionNoID) HasShipmentLocations() bool`

HasShipmentLocations returns a boolean if a field has been set.

### GetReferences

`func (o *ShippingInstructionNoID) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShippingInstructionNoID) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShippingInstructionNoID) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShippingInstructionNoID) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


