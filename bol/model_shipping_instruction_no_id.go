/*
DCSA OpenAPI specification for Electronic Bill of Lading

API specification issued by DCSA.org.  For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2020/12/20201208-DCSA-P1-DCSA-Information-Model-v3.0-FINAL.pdf'>Information Model v3.0</a>  It is possible to use this API as a standalone API. In order to do so it is necessary to use the poll-endPoint - /v1/events  in order to poll event information.  It is recomended to implement the <a href='https://app.swaggerhub.com/apis/dcsaorg/DOCUMENTATION_EVENT_HUB'>DCSA Documentation Event Hub</a> in order to use the push model. Here events are pushed as they occur.

API version: 1.0.0
Contact: info@dcsa.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bol

import (
	"encoding/json"

	"github.com/hubbl-ai/dcsa/common"
)

// ShippingInstructionNoID an enrichment to the original booking shared by the shipper to the carrier. The shipping instruction includes volume or weight, cargo items, shipping dates, origin, destination, and other special instructions. The information given by the shipper through the shipping instruction is the information required to create a Transport Document.
type ShippingInstructionNoID struct {
	TransportDocumentType common.TransportDocumentType `json:"transportDocumentType"`
	// Specifies whether the Transport document is a received for shipment, or shipped onboard.
	IsShippedOnboardType bool `json:"isShippedOnboardType"`
	// The requested number of copies of the Transport document to be issued by the carrier. Only applicable for physical documents
	NumberOfCopies *int32 `json:"numberOfCopies,omitempty"`
	// Number of original copies of the negotiable bill of lading that has been issued to the customer.
	NumberOfOriginals                      *int32                                  `json:"numberOfOriginals,omitempty"`
	PreCarriageUnderShippersResponsibility *PreCarriageUnderShippersResponsibility `json:"preCarriageUnderShippersResponsibility,omitempty"`
	// A location object - here with an example with only UN location code and City name present
	InvoicePayableAt interface{} `json:"invoicePayableAt"`
	// An indicator whether the transport document is electronically transferred.
	IsElectronic bool `json:"isElectronic"`
	// An indicator whether the transport document should include charges.
	IsChargesDisplayed *bool `json:"isChargesDisplayed,omitempty"`
	// The associated booking number provided by the carrier. Cannot be used in combination with Cargo Item level carrierBookingReference
	CarrierBookingReference     interface{}                  `json:"carrierBookingReference,omitempty"`
	CargoItems                  []CargoItem                  `json:"cargoItems,omitempty"`
	UtilizedTransportEquipments []UtilizedTransportEquipment `json:"utilizedTransportEquipments,omitempty"`
	DocumentParties             []DocumentParty              `json:"documentParties,omitempty"`
	ShipmentLocations           []ShipmentLocation           `json:"shipmentLocations,omitempty"`
	References                  []Reference                  `json:"references,omitempty"`
}

// NewShippingInstructionNoID instantiates a new ShippingInstructionNoID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingInstructionNoID(transportDocumentType common.TransportDocumentType, isShippedOnboardType bool, invoicePayableAt interface{}, isElectronic bool) *ShippingInstructionNoID {
	this := ShippingInstructionNoID{}
	this.TransportDocumentType = transportDocumentType
	this.IsShippedOnboardType = isShippedOnboardType
	this.InvoicePayableAt = invoicePayableAt
	this.IsElectronic = isElectronic
	return &this
}

// NewShippingInstructionNoIDWithDefaults instantiates a new ShippingInstructionNoID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingInstructionNoIDWithDefaults() *ShippingInstructionNoID {
	this := ShippingInstructionNoID{}
	return &this
}

// GetTransportDocumentType returns the TransportDocumentType field value
func (o *ShippingInstructionNoID) GetTransportDocumentType() common.TransportDocumentType {
	if o == nil {
		var ret common.TransportDocumentType
		return ret
	}

	return o.TransportDocumentType
}

// GetTransportDocumentTypeOk returns a tuple with the TransportDocumentType field value
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetTransportDocumentTypeOk() (*common.TransportDocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportDocumentType, true
}

// SetTransportDocumentType sets field value
func (o *ShippingInstructionNoID) SetTransportDocumentType(v common.TransportDocumentType) {
	o.TransportDocumentType = v
}

// GetIsShippedOnboardType returns the IsShippedOnboardType field value
func (o *ShippingInstructionNoID) GetIsShippedOnboardType() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsShippedOnboardType
}

// GetIsShippedOnboardTypeOk returns a tuple with the IsShippedOnboardType field value
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetIsShippedOnboardTypeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsShippedOnboardType, true
}

// SetIsShippedOnboardType sets field value
func (o *ShippingInstructionNoID) SetIsShippedOnboardType(v bool) {
	o.IsShippedOnboardType = v
}

// GetNumberOfCopies returns the NumberOfCopies field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetNumberOfCopies() int32 {
	if o == nil || o.NumberOfCopies == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfCopies
}

// GetNumberOfCopiesOk returns a tuple with the NumberOfCopies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetNumberOfCopiesOk() (*int32, bool) {
	if o == nil || o.NumberOfCopies == nil {
		return nil, false
	}
	return o.NumberOfCopies, true
}

// HasNumberOfCopies returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasNumberOfCopies() bool {
	if o != nil && o.NumberOfCopies != nil {
		return true
	}

	return false
}

// SetNumberOfCopies gets a reference to the given int32 and assigns it to the NumberOfCopies field.
func (o *ShippingInstructionNoID) SetNumberOfCopies(v int32) {
	o.NumberOfCopies = &v
}

// GetNumberOfOriginals returns the NumberOfOriginals field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetNumberOfOriginals() int32 {
	if o == nil || o.NumberOfOriginals == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfOriginals
}

// GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetNumberOfOriginalsOk() (*int32, bool) {
	if o == nil || o.NumberOfOriginals == nil {
		return nil, false
	}
	return o.NumberOfOriginals, true
}

// HasNumberOfOriginals returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasNumberOfOriginals() bool {
	if o != nil && o.NumberOfOriginals != nil {
		return true
	}

	return false
}

// SetNumberOfOriginals gets a reference to the given int32 and assigns it to the NumberOfOriginals field.
func (o *ShippingInstructionNoID) SetNumberOfOriginals(v int32) {
	o.NumberOfOriginals = &v
}

// GetPreCarriageUnderShippersResponsibility returns the PreCarriageUnderShippersResponsibility field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetPreCarriageUnderShippersResponsibility() PreCarriageUnderShippersResponsibility {
	if o == nil || o.PreCarriageUnderShippersResponsibility == nil {
		var ret PreCarriageUnderShippersResponsibility
		return ret
	}
	return *o.PreCarriageUnderShippersResponsibility
}

// GetPreCarriageUnderShippersResponsibilityOk returns a tuple with the PreCarriageUnderShippersResponsibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetPreCarriageUnderShippersResponsibilityOk() (*PreCarriageUnderShippersResponsibility, bool) {
	if o == nil || o.PreCarriageUnderShippersResponsibility == nil {
		return nil, false
	}
	return o.PreCarriageUnderShippersResponsibility, true
}

// HasPreCarriageUnderShippersResponsibility returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasPreCarriageUnderShippersResponsibility() bool {
	if o != nil && o.PreCarriageUnderShippersResponsibility != nil {
		return true
	}

	return false
}

// SetPreCarriageUnderShippersResponsibility gets a reference to the given PreCarriageUnderShippersResponsibility and assigns it to the PreCarriageUnderShippersResponsibility field.
func (o *ShippingInstructionNoID) SetPreCarriageUnderShippersResponsibility(v PreCarriageUnderShippersResponsibility) {
	o.PreCarriageUnderShippersResponsibility = &v
}

// GetInvoicePayableAt returns the InvoicePayableAt field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShippingInstructionNoID) GetInvoicePayableAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.InvoicePayableAt
}

// GetInvoicePayableAtOk returns a tuple with the InvoicePayableAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingInstructionNoID) GetInvoicePayableAtOk() (*interface{}, bool) {
	if o == nil || o.InvoicePayableAt == nil {
		return nil, false
	}
	return &o.InvoicePayableAt, true
}

// SetInvoicePayableAt sets field value
func (o *ShippingInstructionNoID) SetInvoicePayableAt(v interface{}) {
	o.InvoicePayableAt = v
}

// GetIsElectronic returns the IsElectronic field value
func (o *ShippingInstructionNoID) GetIsElectronic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsElectronic
}

// GetIsElectronicOk returns a tuple with the IsElectronic field value
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetIsElectronicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsElectronic, true
}

// SetIsElectronic sets field value
func (o *ShippingInstructionNoID) SetIsElectronic(v bool) {
	o.IsElectronic = v
}

// GetIsChargesDisplayed returns the IsChargesDisplayed field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetIsChargesDisplayed() bool {
	if o == nil || o.IsChargesDisplayed == nil {
		var ret bool
		return ret
	}
	return *o.IsChargesDisplayed
}

// GetIsChargesDisplayedOk returns a tuple with the IsChargesDisplayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetIsChargesDisplayedOk() (*bool, bool) {
	if o == nil || o.IsChargesDisplayed == nil {
		return nil, false
	}
	return o.IsChargesDisplayed, true
}

// HasIsChargesDisplayed returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasIsChargesDisplayed() bool {
	if o != nil && o.IsChargesDisplayed != nil {
		return true
	}

	return false
}

// SetIsChargesDisplayed gets a reference to the given bool and assigns it to the IsChargesDisplayed field.
func (o *ShippingInstructionNoID) SetIsChargesDisplayed(v bool) {
	o.IsChargesDisplayed = &v
}

// GetCarrierBookingReference returns the CarrierBookingReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingInstructionNoID) GetCarrierBookingReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CarrierBookingReference
}

// GetCarrierBookingReferenceOk returns a tuple with the CarrierBookingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingInstructionNoID) GetCarrierBookingReferenceOk() (*interface{}, bool) {
	if o == nil || o.CarrierBookingReference == nil {
		return nil, false
	}
	return &o.CarrierBookingReference, true
}

// HasCarrierBookingReference returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasCarrierBookingReference() bool {
	if o != nil && o.CarrierBookingReference != nil {
		return true
	}

	return false
}

// SetCarrierBookingReference gets a reference to the given interface{} and assigns it to the CarrierBookingReference field.
func (o *ShippingInstructionNoID) SetCarrierBookingReference(v interface{}) {
	o.CarrierBookingReference = v
}

// GetCargoItems returns the CargoItems field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetCargoItems() []CargoItem {
	if o == nil || o.CargoItems == nil {
		var ret []CargoItem
		return ret
	}
	return o.CargoItems
}

// GetCargoItemsOk returns a tuple with the CargoItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetCargoItemsOk() ([]CargoItem, bool) {
	if o == nil || o.CargoItems == nil {
		return nil, false
	}
	return o.CargoItems, true
}

// HasCargoItems returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasCargoItems() bool {
	if o != nil && o.CargoItems != nil {
		return true
	}

	return false
}

// SetCargoItems gets a reference to the given []CargoItem and assigns it to the CargoItems field.
func (o *ShippingInstructionNoID) SetCargoItems(v []CargoItem) {
	o.CargoItems = v
}

// GetUtilizedTransportEquipments returns the UtilizedTransportEquipments field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetUtilizedTransportEquipments() []UtilizedTransportEquipment {
	if o == nil || o.UtilizedTransportEquipments == nil {
		var ret []UtilizedTransportEquipment
		return ret
	}
	return o.UtilizedTransportEquipments
}

// GetUtilizedTransportEquipmentsOk returns a tuple with the UtilizedTransportEquipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetUtilizedTransportEquipmentsOk() ([]UtilizedTransportEquipment, bool) {
	if o == nil || o.UtilizedTransportEquipments == nil {
		return nil, false
	}
	return o.UtilizedTransportEquipments, true
}

// HasUtilizedTransportEquipments returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasUtilizedTransportEquipments() bool {
	if o != nil && o.UtilizedTransportEquipments != nil {
		return true
	}

	return false
}

// SetUtilizedTransportEquipments gets a reference to the given []UtilizedTransportEquipment and assigns it to the UtilizedTransportEquipments field.
func (o *ShippingInstructionNoID) SetUtilizedTransportEquipments(v []UtilizedTransportEquipment) {
	o.UtilizedTransportEquipments = v
}

// GetDocumentParties returns the DocumentParties field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetDocumentParties() []DocumentParty {
	if o == nil || o.DocumentParties == nil {
		var ret []DocumentParty
		return ret
	}
	return o.DocumentParties
}

// GetDocumentPartiesOk returns a tuple with the DocumentParties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetDocumentPartiesOk() ([]DocumentParty, bool) {
	if o == nil || o.DocumentParties == nil {
		return nil, false
	}
	return o.DocumentParties, true
}

// HasDocumentParties returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasDocumentParties() bool {
	if o != nil && o.DocumentParties != nil {
		return true
	}

	return false
}

// SetDocumentParties gets a reference to the given []DocumentParty and assigns it to the DocumentParties field.
func (o *ShippingInstructionNoID) SetDocumentParties(v []DocumentParty) {
	o.DocumentParties = v
}

// GetShipmentLocations returns the ShipmentLocations field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetShipmentLocations() []ShipmentLocation {
	if o == nil || o.ShipmentLocations == nil {
		var ret []ShipmentLocation
		return ret
	}
	return o.ShipmentLocations
}

// GetShipmentLocationsOk returns a tuple with the ShipmentLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetShipmentLocationsOk() ([]ShipmentLocation, bool) {
	if o == nil || o.ShipmentLocations == nil {
		return nil, false
	}
	return o.ShipmentLocations, true
}

// HasShipmentLocations returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasShipmentLocations() bool {
	if o != nil && o.ShipmentLocations != nil {
		return true
	}

	return false
}

// SetShipmentLocations gets a reference to the given []ShipmentLocation and assigns it to the ShipmentLocations field.
func (o *ShippingInstructionNoID) SetShipmentLocations(v []ShipmentLocation) {
	o.ShipmentLocations = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ShippingInstructionNoID) GetReferences() []Reference {
	if o == nil || o.References == nil {
		var ret []Reference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingInstructionNoID) GetReferencesOk() ([]Reference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ShippingInstructionNoID) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []Reference and assigns it to the References field.
func (o *ShippingInstructionNoID) SetReferences(v []Reference) {
	o.References = v
}

func (o ShippingInstructionNoID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transportDocumentType"] = o.TransportDocumentType
	}
	if true {
		toSerialize["isShippedOnboardType"] = o.IsShippedOnboardType
	}
	if o.NumberOfCopies != nil {
		toSerialize["numberOfCopies"] = o.NumberOfCopies
	}
	if o.NumberOfOriginals != nil {
		toSerialize["numberOfOriginals"] = o.NumberOfOriginals
	}
	if o.PreCarriageUnderShippersResponsibility != nil {
		toSerialize["preCarriageUnderShippersResponsibility"] = o.PreCarriageUnderShippersResponsibility
	}
	if o.InvoicePayableAt != nil {
		toSerialize["invoicePayableAt"] = o.InvoicePayableAt
	}
	if true {
		toSerialize["isElectronic"] = o.IsElectronic
	}
	if o.IsChargesDisplayed != nil {
		toSerialize["isChargesDisplayed"] = o.IsChargesDisplayed
	}
	if o.CarrierBookingReference != nil {
		toSerialize["carrierBookingReference"] = o.CarrierBookingReference
	}
	if o.CargoItems != nil {
		toSerialize["cargoItems"] = o.CargoItems
	}
	if o.UtilizedTransportEquipments != nil {
		toSerialize["utilizedTransportEquipments"] = o.UtilizedTransportEquipments
	}
	if o.DocumentParties != nil {
		toSerialize["documentParties"] = o.DocumentParties
	}
	if o.ShipmentLocations != nil {
		toSerialize["shipmentLocations"] = o.ShipmentLocations
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableShippingInstructionNoID struct {
	value *ShippingInstructionNoID
	isSet bool
}

func (v NullableShippingInstructionNoID) Get() *ShippingInstructionNoID {
	return v.value
}

func (v *NullableShippingInstructionNoID) Set(val *ShippingInstructionNoID) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingInstructionNoID) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingInstructionNoID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingInstructionNoID(val *ShippingInstructionNoID) *NullableShippingInstructionNoID {
	return &NullableShippingInstructionNoID{value: val, isSet: true}
}

func (v NullableShippingInstructionNoID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingInstructionNoID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
