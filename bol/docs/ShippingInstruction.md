# ShippingInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingInstructionID** | **string** | The identifier for a shipping instruction provided by the carrier for system purposes.  | 
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

### NewShippingInstruction

`func NewShippingInstruction(shippingInstructionID string, transportDocumentType TransportDocumentType, isShippedOnboardType bool, invoicePayableAt interface{}, isElectronic bool, ) *ShippingInstruction`

NewShippingInstruction instantiates a new ShippingInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingInstructionWithDefaults

`func NewShippingInstructionWithDefaults() *ShippingInstruction`

NewShippingInstructionWithDefaults instantiates a new ShippingInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingInstructionID

`func (o *ShippingInstruction) GetShippingInstructionID() string`

GetShippingInstructionID returns the ShippingInstructionID field if non-nil, zero value otherwise.

### GetShippingInstructionIDOk

`func (o *ShippingInstruction) GetShippingInstructionIDOk() (*string, bool)`

GetShippingInstructionIDOk returns a tuple with the ShippingInstructionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInstructionID

`func (o *ShippingInstruction) SetShippingInstructionID(v string)`

SetShippingInstructionID sets ShippingInstructionID field to given value.


### GetTransportDocumentType

`func (o *ShippingInstruction) GetTransportDocumentType() TransportDocumentType`

GetTransportDocumentType returns the TransportDocumentType field if non-nil, zero value otherwise.

### GetTransportDocumentTypeOk

`func (o *ShippingInstruction) GetTransportDocumentTypeOk() (*TransportDocumentType, bool)`

GetTransportDocumentTypeOk returns a tuple with the TransportDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportDocumentType

`func (o *ShippingInstruction) SetTransportDocumentType(v TransportDocumentType)`

SetTransportDocumentType sets TransportDocumentType field to given value.


### GetIsShippedOnboardType

`func (o *ShippingInstruction) GetIsShippedOnboardType() bool`

GetIsShippedOnboardType returns the IsShippedOnboardType field if non-nil, zero value otherwise.

### GetIsShippedOnboardTypeOk

`func (o *ShippingInstruction) GetIsShippedOnboardTypeOk() (*bool, bool)`

GetIsShippedOnboardTypeOk returns a tuple with the IsShippedOnboardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippedOnboardType

`func (o *ShippingInstruction) SetIsShippedOnboardType(v bool)`

SetIsShippedOnboardType sets IsShippedOnboardType field to given value.


### GetNumberOfCopies

`func (o *ShippingInstruction) GetNumberOfCopies() int32`

GetNumberOfCopies returns the NumberOfCopies field if non-nil, zero value otherwise.

### GetNumberOfCopiesOk

`func (o *ShippingInstruction) GetNumberOfCopiesOk() (*int32, bool)`

GetNumberOfCopiesOk returns a tuple with the NumberOfCopies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCopies

`func (o *ShippingInstruction) SetNumberOfCopies(v int32)`

SetNumberOfCopies sets NumberOfCopies field to given value.

### HasNumberOfCopies

`func (o *ShippingInstruction) HasNumberOfCopies() bool`

HasNumberOfCopies returns a boolean if a field has been set.

### GetNumberOfOriginals

`func (o *ShippingInstruction) GetNumberOfOriginals() int32`

GetNumberOfOriginals returns the NumberOfOriginals field if non-nil, zero value otherwise.

### GetNumberOfOriginalsOk

`func (o *ShippingInstruction) GetNumberOfOriginalsOk() (*int32, bool)`

GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOriginals

`func (o *ShippingInstruction) SetNumberOfOriginals(v int32)`

SetNumberOfOriginals sets NumberOfOriginals field to given value.

### HasNumberOfOriginals

`func (o *ShippingInstruction) HasNumberOfOriginals() bool`

HasNumberOfOriginals returns a boolean if a field has been set.

### GetPreCarriageUnderShippersResponsibility

`func (o *ShippingInstruction) GetPreCarriageUnderShippersResponsibility() PreCarriageUnderShippersResponsibility`

GetPreCarriageUnderShippersResponsibility returns the PreCarriageUnderShippersResponsibility field if non-nil, zero value otherwise.

### GetPreCarriageUnderShippersResponsibilityOk

`func (o *ShippingInstruction) GetPreCarriageUnderShippersResponsibilityOk() (*PreCarriageUnderShippersResponsibility, bool)`

GetPreCarriageUnderShippersResponsibilityOk returns a tuple with the PreCarriageUnderShippersResponsibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCarriageUnderShippersResponsibility

`func (o *ShippingInstruction) SetPreCarriageUnderShippersResponsibility(v PreCarriageUnderShippersResponsibility)`

SetPreCarriageUnderShippersResponsibility sets PreCarriageUnderShippersResponsibility field to given value.

### HasPreCarriageUnderShippersResponsibility

`func (o *ShippingInstruction) HasPreCarriageUnderShippersResponsibility() bool`

HasPreCarriageUnderShippersResponsibility returns a boolean if a field has been set.

### GetInvoicePayableAt

`func (o *ShippingInstruction) GetInvoicePayableAt() interface{}`

GetInvoicePayableAt returns the InvoicePayableAt field if non-nil, zero value otherwise.

### GetInvoicePayableAtOk

`func (o *ShippingInstruction) GetInvoicePayableAtOk() (*interface{}, bool)`

GetInvoicePayableAtOk returns a tuple with the InvoicePayableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePayableAt

`func (o *ShippingInstruction) SetInvoicePayableAt(v interface{})`

SetInvoicePayableAt sets InvoicePayableAt field to given value.


### SetInvoicePayableAtNil

`func (o *ShippingInstruction) SetInvoicePayableAtNil(b bool)`

 SetInvoicePayableAtNil sets the value for InvoicePayableAt to be an explicit nil

### UnsetInvoicePayableAt
`func (o *ShippingInstruction) UnsetInvoicePayableAt()`

UnsetInvoicePayableAt ensures that no value is present for InvoicePayableAt, not even an explicit nil
### GetIsElectronic

`func (o *ShippingInstruction) GetIsElectronic() bool`

GetIsElectronic returns the IsElectronic field if non-nil, zero value otherwise.

### GetIsElectronicOk

`func (o *ShippingInstruction) GetIsElectronicOk() (*bool, bool)`

GetIsElectronicOk returns a tuple with the IsElectronic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsElectronic

`func (o *ShippingInstruction) SetIsElectronic(v bool)`

SetIsElectronic sets IsElectronic field to given value.


### GetIsChargesDisplayed

`func (o *ShippingInstruction) GetIsChargesDisplayed() bool`

GetIsChargesDisplayed returns the IsChargesDisplayed field if non-nil, zero value otherwise.

### GetIsChargesDisplayedOk

`func (o *ShippingInstruction) GetIsChargesDisplayedOk() (*bool, bool)`

GetIsChargesDisplayedOk returns a tuple with the IsChargesDisplayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChargesDisplayed

`func (o *ShippingInstruction) SetIsChargesDisplayed(v bool)`

SetIsChargesDisplayed sets IsChargesDisplayed field to given value.

### HasIsChargesDisplayed

`func (o *ShippingInstruction) HasIsChargesDisplayed() bool`

HasIsChargesDisplayed returns a boolean if a field has been set.

### GetCarrierBookingReference

`func (o *ShippingInstruction) GetCarrierBookingReference() interface{}`

GetCarrierBookingReference returns the CarrierBookingReference field if non-nil, zero value otherwise.

### GetCarrierBookingReferenceOk

`func (o *ShippingInstruction) GetCarrierBookingReferenceOk() (*interface{}, bool)`

GetCarrierBookingReferenceOk returns a tuple with the CarrierBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierBookingReference

`func (o *ShippingInstruction) SetCarrierBookingReference(v interface{})`

SetCarrierBookingReference sets CarrierBookingReference field to given value.

### HasCarrierBookingReference

`func (o *ShippingInstruction) HasCarrierBookingReference() bool`

HasCarrierBookingReference returns a boolean if a field has been set.

### SetCarrierBookingReferenceNil

`func (o *ShippingInstruction) SetCarrierBookingReferenceNil(b bool)`

 SetCarrierBookingReferenceNil sets the value for CarrierBookingReference to be an explicit nil

### UnsetCarrierBookingReference
`func (o *ShippingInstruction) UnsetCarrierBookingReference()`

UnsetCarrierBookingReference ensures that no value is present for CarrierBookingReference, not even an explicit nil
### GetCargoItems

`func (o *ShippingInstruction) GetCargoItems() []CargoItem`

GetCargoItems returns the CargoItems field if non-nil, zero value otherwise.

### GetCargoItemsOk

`func (o *ShippingInstruction) GetCargoItemsOk() (*[]CargoItem, bool)`

GetCargoItemsOk returns a tuple with the CargoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoItems

`func (o *ShippingInstruction) SetCargoItems(v []CargoItem)`

SetCargoItems sets CargoItems field to given value.

### HasCargoItems

`func (o *ShippingInstruction) HasCargoItems() bool`

HasCargoItems returns a boolean if a field has been set.

### GetUtilizedTransportEquipments

`func (o *ShippingInstruction) GetUtilizedTransportEquipments() []UtilizedTransportEquipment`

GetUtilizedTransportEquipments returns the UtilizedTransportEquipments field if non-nil, zero value otherwise.

### GetUtilizedTransportEquipmentsOk

`func (o *ShippingInstruction) GetUtilizedTransportEquipmentsOk() (*[]UtilizedTransportEquipment, bool)`

GetUtilizedTransportEquipmentsOk returns a tuple with the UtilizedTransportEquipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizedTransportEquipments

`func (o *ShippingInstruction) SetUtilizedTransportEquipments(v []UtilizedTransportEquipment)`

SetUtilizedTransportEquipments sets UtilizedTransportEquipments field to given value.

### HasUtilizedTransportEquipments

`func (o *ShippingInstruction) HasUtilizedTransportEquipments() bool`

HasUtilizedTransportEquipments returns a boolean if a field has been set.

### GetDocumentParties

`func (o *ShippingInstruction) GetDocumentParties() []DocumentParty`

GetDocumentParties returns the DocumentParties field if non-nil, zero value otherwise.

### GetDocumentPartiesOk

`func (o *ShippingInstruction) GetDocumentPartiesOk() (*[]DocumentParty, bool)`

GetDocumentPartiesOk returns a tuple with the DocumentParties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentParties

`func (o *ShippingInstruction) SetDocumentParties(v []DocumentParty)`

SetDocumentParties sets DocumentParties field to given value.

### HasDocumentParties

`func (o *ShippingInstruction) HasDocumentParties() bool`

HasDocumentParties returns a boolean if a field has been set.

### GetShipmentLocations

`func (o *ShippingInstruction) GetShipmentLocations() []ShipmentLocation`

GetShipmentLocations returns the ShipmentLocations field if non-nil, zero value otherwise.

### GetShipmentLocationsOk

`func (o *ShippingInstruction) GetShipmentLocationsOk() (*[]ShipmentLocation, bool)`

GetShipmentLocationsOk returns a tuple with the ShipmentLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLocations

`func (o *ShippingInstruction) SetShipmentLocations(v []ShipmentLocation)`

SetShipmentLocations sets ShipmentLocations field to given value.

### HasShipmentLocations

`func (o *ShippingInstruction) HasShipmentLocations() bool`

HasShipmentLocations returns a boolean if a field has been set.

### GetReferences

`func (o *ShippingInstruction) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShippingInstruction) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShippingInstruction) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShippingInstruction) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


