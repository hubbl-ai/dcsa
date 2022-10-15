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

// TransportDocumentHeader struct for TransportDocumentHeader
type TransportDocumentHeader struct {
	// A unique number allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment.
	TransportDocumentReference string    `json:"transportDocumentReference"`
	PlaceOfIssue               *Location `json:"placeOfIssue,omitempty"`
	// Date when the transport document has been issued
	IssueDate string `json:"issueDate"`
	// Date when the last container that is linked to the transport document is physically loaded onboard the vessel indicated on the transport document.
	ShippedOnBoardDate *string `json:"shippedOnBoardDate,omitempty"`
	// Date when the last container linked to the transport document is physically in the terminal (customers cleared against the intended vessel).
	ReceivedForShipmentDate *string `json:"receivedForShipmentDate,omitempty"`
	// Carrier general terms and conditions for this transport document.
	TermsAndConditions string `json:"termsAndConditions"`
	// The code used for identifying the specific issuer.
	IssuerCode             *string                 `json:"issuerCode,omitempty"`
	IssuerCodeListProvider *IssuerCodeListProvider `json:"issuerCodeListProvider,omitempty"`
	// The currency used for the declared value, using the 3-character code defined by ISO 4217.
	DeclaredValueCurrency *string `json:"declaredValueCurrency,omitempty"`
	// The value of the cargo that the shipper declares to avoid the carrier&apos;s limitation of liability and \"Ad Valorem\" freight, i.e. freight which is calculated based on the value of the goods declared by the shipper.
	DeclaredValue *float32 `json:"declaredValue,omitempty"`
	// The number of additional pages required to contain the goods description on a transport document. Only applicable for physical transport documents.
	NumberOfRiderPages               *int32                   `json:"numberOfRiderPages,omitempty"`
	CargoMovementTypeAtOrigin        common.CargoMovementType `json:"cargoMovementTypeAtOrigin"`
	CargoMovementTypeAtDestination   common.CargoMovementType `json:"cargoMovementTypeAtDestination"`
	ReceiptDeliveryTypeAtOrigin      ReceiptDeliveryType      `json:"receiptDeliveryTypeAtOrigin"`
	ReceiptDeliveryTypeAtDestination ReceiptDeliveryType      `json:"receiptDeliveryTypeAtDestination"`
	// Reference number for agreement between shipper and carrier through which the shipper commits to provide a certain minimum quantity of cargo over a fixed period, and the carrier commits to a certain rate or rate schedule.
	ServiceContractReference *string `json:"serviceContractReference,omitempty"`
}

// NewTransportDocumentHeader instantiates a new TransportDocumentHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportDocumentHeader(transportDocumentReference string, issueDate string, termsAndConditions string, cargoMovementTypeAtOrigin common.CargoMovementType, cargoMovementTypeAtDestination common.CargoMovementType, receiptDeliveryTypeAtOrigin ReceiptDeliveryType, receiptDeliveryTypeAtDestination ReceiptDeliveryType) *TransportDocumentHeader {
	this := TransportDocumentHeader{}
	this.TransportDocumentReference = transportDocumentReference
	this.IssueDate = issueDate
	this.TermsAndConditions = termsAndConditions
	this.CargoMovementTypeAtOrigin = cargoMovementTypeAtOrigin
	this.CargoMovementTypeAtDestination = cargoMovementTypeAtDestination
	this.ReceiptDeliveryTypeAtOrigin = receiptDeliveryTypeAtOrigin
	this.ReceiptDeliveryTypeAtDestination = receiptDeliveryTypeAtDestination
	return &this
}

// NewTransportDocumentHeaderWithDefaults instantiates a new TransportDocumentHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportDocumentHeaderWithDefaults() *TransportDocumentHeader {
	this := TransportDocumentHeader{}
	return &this
}

// GetTransportDocumentReference returns the TransportDocumentReference field value
func (o *TransportDocumentHeader) GetTransportDocumentReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransportDocumentReference
}

// GetTransportDocumentReferenceOk returns a tuple with the TransportDocumentReference field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetTransportDocumentReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportDocumentReference, true
}

// SetTransportDocumentReference sets field value
func (o *TransportDocumentHeader) SetTransportDocumentReference(v string) {
	o.TransportDocumentReference = v
}

// GetPlaceOfIssue returns the PlaceOfIssue field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetPlaceOfIssue() Location {
	if o == nil || o.PlaceOfIssue == nil {
		var ret Location
		return ret
	}
	return *o.PlaceOfIssue
}

// GetPlaceOfIssueOk returns a tuple with the PlaceOfIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetPlaceOfIssueOk() (*Location, bool) {
	if o == nil || o.PlaceOfIssue == nil {
		return nil, false
	}
	return o.PlaceOfIssue, true
}

// HasPlaceOfIssue returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasPlaceOfIssue() bool {
	if o != nil && o.PlaceOfIssue != nil {
		return true
	}

	return false
}

// SetPlaceOfIssue gets a reference to the given Location and assigns it to the PlaceOfIssue field.
func (o *TransportDocumentHeader) SetPlaceOfIssue(v Location) {
	o.PlaceOfIssue = &v
}

// GetIssueDate returns the IssueDate field value
func (o *TransportDocumentHeader) GetIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetIssueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *TransportDocumentHeader) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetShippedOnBoardDate returns the ShippedOnBoardDate field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetShippedOnBoardDate() string {
	if o == nil || o.ShippedOnBoardDate == nil {
		var ret string
		return ret
	}
	return *o.ShippedOnBoardDate
}

// GetShippedOnBoardDateOk returns a tuple with the ShippedOnBoardDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetShippedOnBoardDateOk() (*string, bool) {
	if o == nil || o.ShippedOnBoardDate == nil {
		return nil, false
	}
	return o.ShippedOnBoardDate, true
}

// HasShippedOnBoardDate returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasShippedOnBoardDate() bool {
	if o != nil && o.ShippedOnBoardDate != nil {
		return true
	}

	return false
}

// SetShippedOnBoardDate gets a reference to the given string and assigns it to the ShippedOnBoardDate field.
func (o *TransportDocumentHeader) SetShippedOnBoardDate(v string) {
	o.ShippedOnBoardDate = &v
}

// GetReceivedForShipmentDate returns the ReceivedForShipmentDate field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetReceivedForShipmentDate() string {
	if o == nil || o.ReceivedForShipmentDate == nil {
		var ret string
		return ret
	}
	return *o.ReceivedForShipmentDate
}

// GetReceivedForShipmentDateOk returns a tuple with the ReceivedForShipmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetReceivedForShipmentDateOk() (*string, bool) {
	if o == nil || o.ReceivedForShipmentDate == nil {
		return nil, false
	}
	return o.ReceivedForShipmentDate, true
}

// HasReceivedForShipmentDate returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasReceivedForShipmentDate() bool {
	if o != nil && o.ReceivedForShipmentDate != nil {
		return true
	}

	return false
}

// SetReceivedForShipmentDate gets a reference to the given string and assigns it to the ReceivedForShipmentDate field.
func (o *TransportDocumentHeader) SetReceivedForShipmentDate(v string) {
	o.ReceivedForShipmentDate = &v
}

// GetTermsAndConditions returns the TermsAndConditions field value
func (o *TransportDocumentHeader) GetTermsAndConditions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TermsAndConditions
}

// GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetTermsAndConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermsAndConditions, true
}

// SetTermsAndConditions sets field value
func (o *TransportDocumentHeader) SetTermsAndConditions(v string) {
	o.TermsAndConditions = v
}

// GetIssuerCode returns the IssuerCode field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetIssuerCode() string {
	if o == nil || o.IssuerCode == nil {
		var ret string
		return ret
	}
	return *o.IssuerCode
}

// GetIssuerCodeOk returns a tuple with the IssuerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetIssuerCodeOk() (*string, bool) {
	if o == nil || o.IssuerCode == nil {
		return nil, false
	}
	return o.IssuerCode, true
}

// HasIssuerCode returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasIssuerCode() bool {
	if o != nil && o.IssuerCode != nil {
		return true
	}

	return false
}

// SetIssuerCode gets a reference to the given string and assigns it to the IssuerCode field.
func (o *TransportDocumentHeader) SetIssuerCode(v string) {
	o.IssuerCode = &v
}

// GetIssuerCodeListProvider returns the IssuerCodeListProvider field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetIssuerCodeListProvider() IssuerCodeListProvider {
	if o == nil || o.IssuerCodeListProvider == nil {
		var ret IssuerCodeListProvider
		return ret
	}
	return *o.IssuerCodeListProvider
}

// GetIssuerCodeListProviderOk returns a tuple with the IssuerCodeListProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetIssuerCodeListProviderOk() (*IssuerCodeListProvider, bool) {
	if o == nil || o.IssuerCodeListProvider == nil {
		return nil, false
	}
	return o.IssuerCodeListProvider, true
}

// HasIssuerCodeListProvider returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasIssuerCodeListProvider() bool {
	if o != nil && o.IssuerCodeListProvider != nil {
		return true
	}

	return false
}

// SetIssuerCodeListProvider gets a reference to the given IssuerCodeListProvider and assigns it to the IssuerCodeListProvider field.
func (o *TransportDocumentHeader) SetIssuerCodeListProvider(v IssuerCodeListProvider) {
	o.IssuerCodeListProvider = &v
}

// GetDeclaredValueCurrency returns the DeclaredValueCurrency field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetDeclaredValueCurrency() string {
	if o == nil || o.DeclaredValueCurrency == nil {
		var ret string
		return ret
	}
	return *o.DeclaredValueCurrency
}

// GetDeclaredValueCurrencyOk returns a tuple with the DeclaredValueCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetDeclaredValueCurrencyOk() (*string, bool) {
	if o == nil || o.DeclaredValueCurrency == nil {
		return nil, false
	}
	return o.DeclaredValueCurrency, true
}

// HasDeclaredValueCurrency returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasDeclaredValueCurrency() bool {
	if o != nil && o.DeclaredValueCurrency != nil {
		return true
	}

	return false
}

// SetDeclaredValueCurrency gets a reference to the given string and assigns it to the DeclaredValueCurrency field.
func (o *TransportDocumentHeader) SetDeclaredValueCurrency(v string) {
	o.DeclaredValueCurrency = &v
}

// GetDeclaredValue returns the DeclaredValue field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetDeclaredValue() float32 {
	if o == nil || o.DeclaredValue == nil {
		var ret float32
		return ret
	}
	return *o.DeclaredValue
}

// GetDeclaredValueOk returns a tuple with the DeclaredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetDeclaredValueOk() (*float32, bool) {
	if o == nil || o.DeclaredValue == nil {
		return nil, false
	}
	return o.DeclaredValue, true
}

// HasDeclaredValue returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasDeclaredValue() bool {
	if o != nil && o.DeclaredValue != nil {
		return true
	}

	return false
}

// SetDeclaredValue gets a reference to the given float32 and assigns it to the DeclaredValue field.
func (o *TransportDocumentHeader) SetDeclaredValue(v float32) {
	o.DeclaredValue = &v
}

// GetNumberOfRiderPages returns the NumberOfRiderPages field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetNumberOfRiderPages() int32 {
	if o == nil || o.NumberOfRiderPages == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfRiderPages
}

// GetNumberOfRiderPagesOk returns a tuple with the NumberOfRiderPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetNumberOfRiderPagesOk() (*int32, bool) {
	if o == nil || o.NumberOfRiderPages == nil {
		return nil, false
	}
	return o.NumberOfRiderPages, true
}

// HasNumberOfRiderPages returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasNumberOfRiderPages() bool {
	if o != nil && o.NumberOfRiderPages != nil {
		return true
	}

	return false
}

// SetNumberOfRiderPages gets a reference to the given int32 and assigns it to the NumberOfRiderPages field.
func (o *TransportDocumentHeader) SetNumberOfRiderPages(v int32) {
	o.NumberOfRiderPages = &v
}

// GetCargoMovementTypeAtOrigin returns the CargoMovementTypeAtOrigin field value
func (o *TransportDocumentHeader) GetCargoMovementTypeAtOrigin() common.CargoMovementType {
	if o == nil {
		var ret common.CargoMovementType
		return ret
	}

	return o.CargoMovementTypeAtOrigin
}

// GetCargoMovementTypeAtOriginOk returns a tuple with the CargoMovementTypeAtOrigin field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetCargoMovementTypeAtOriginOk() (*common.CargoMovementType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CargoMovementTypeAtOrigin, true
}

// SetCargoMovementTypeAtOrigin sets field value
func (o *TransportDocumentHeader) SetCargoMovementTypeAtOrigin(v common.CargoMovementType) {
	o.CargoMovementTypeAtOrigin = v
}

// GetCargoMovementTypeAtDestination returns the CargoMovementTypeAtDestination field value
func (o *TransportDocumentHeader) GetCargoMovementTypeAtDestination() common.CargoMovementType {
	if o == nil {
		var ret common.CargoMovementType
		return ret
	}

	return o.CargoMovementTypeAtDestination
}

// GetCargoMovementTypeAtDestinationOk returns a tuple with the CargoMovementTypeAtDestination field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetCargoMovementTypeAtDestinationOk() (*common.CargoMovementType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CargoMovementTypeAtDestination, true
}

// SetCargoMovementTypeAtDestination sets field value
func (o *TransportDocumentHeader) SetCargoMovementTypeAtDestination(v common.CargoMovementType) {
	o.CargoMovementTypeAtDestination = v
}

// GetReceiptDeliveryTypeAtOrigin returns the ReceiptDeliveryTypeAtOrigin field value
func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtOrigin() ReceiptDeliveryType {
	if o == nil {
		var ret ReceiptDeliveryType
		return ret
	}

	return o.ReceiptDeliveryTypeAtOrigin
}

// GetReceiptDeliveryTypeAtOriginOk returns a tuple with the ReceiptDeliveryTypeAtOrigin field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtOriginOk() (*ReceiptDeliveryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptDeliveryTypeAtOrigin, true
}

// SetReceiptDeliveryTypeAtOrigin sets field value
func (o *TransportDocumentHeader) SetReceiptDeliveryTypeAtOrigin(v ReceiptDeliveryType) {
	o.ReceiptDeliveryTypeAtOrigin = v
}

// GetReceiptDeliveryTypeAtDestination returns the ReceiptDeliveryTypeAtDestination field value
func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtDestination() ReceiptDeliveryType {
	if o == nil {
		var ret ReceiptDeliveryType
		return ret
	}

	return o.ReceiptDeliveryTypeAtDestination
}

// GetReceiptDeliveryTypeAtDestinationOk returns a tuple with the ReceiptDeliveryTypeAtDestination field value
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetReceiptDeliveryTypeAtDestinationOk() (*ReceiptDeliveryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptDeliveryTypeAtDestination, true
}

// SetReceiptDeliveryTypeAtDestination sets field value
func (o *TransportDocumentHeader) SetReceiptDeliveryTypeAtDestination(v ReceiptDeliveryType) {
	o.ReceiptDeliveryTypeAtDestination = v
}

// GetServiceContractReference returns the ServiceContractReference field value if set, zero value otherwise.
func (o *TransportDocumentHeader) GetServiceContractReference() string {
	if o == nil || o.ServiceContractReference == nil {
		var ret string
		return ret
	}
	return *o.ServiceContractReference
}

// GetServiceContractReferenceOk returns a tuple with the ServiceContractReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportDocumentHeader) GetServiceContractReferenceOk() (*string, bool) {
	if o == nil || o.ServiceContractReference == nil {
		return nil, false
	}
	return o.ServiceContractReference, true
}

// HasServiceContractReference returns a boolean if a field has been set.
func (o *TransportDocumentHeader) HasServiceContractReference() bool {
	if o != nil && o.ServiceContractReference != nil {
		return true
	}

	return false
}

// SetServiceContractReference gets a reference to the given string and assigns it to the ServiceContractReference field.
func (o *TransportDocumentHeader) SetServiceContractReference(v string) {
	o.ServiceContractReference = &v
}

func (o TransportDocumentHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transportDocumentReference"] = o.TransportDocumentReference
	}
	if o.PlaceOfIssue != nil {
		toSerialize["placeOfIssue"] = o.PlaceOfIssue
	}
	if true {
		toSerialize["issueDate"] = o.IssueDate
	}
	if o.ShippedOnBoardDate != nil {
		toSerialize["shippedOnBoardDate"] = o.ShippedOnBoardDate
	}
	if o.ReceivedForShipmentDate != nil {
		toSerialize["receivedForShipmentDate"] = o.ReceivedForShipmentDate
	}
	if true {
		toSerialize["termsAndConditions"] = o.TermsAndConditions
	}
	if o.IssuerCode != nil {
		toSerialize["issuerCode"] = o.IssuerCode
	}
	if o.IssuerCodeListProvider != nil {
		toSerialize["issuerCodeListProvider"] = o.IssuerCodeListProvider
	}
	if o.DeclaredValueCurrency != nil {
		toSerialize["declaredValueCurrency"] = o.DeclaredValueCurrency
	}
	if o.DeclaredValue != nil {
		toSerialize["declaredValue"] = o.DeclaredValue
	}
	if o.NumberOfRiderPages != nil {
		toSerialize["numberOfRiderPages"] = o.NumberOfRiderPages
	}
	if true {
		toSerialize["cargoMovementTypeAtOrigin"] = o.CargoMovementTypeAtOrigin
	}
	if true {
		toSerialize["cargoMovementTypeAtDestination"] = o.CargoMovementTypeAtDestination
	}
	if true {
		toSerialize["receiptDeliveryTypeAtOrigin"] = o.ReceiptDeliveryTypeAtOrigin
	}
	if true {
		toSerialize["receiptDeliveryTypeAtDestination"] = o.ReceiptDeliveryTypeAtDestination
	}
	if o.ServiceContractReference != nil {
		toSerialize["serviceContractReference"] = o.ServiceContractReference
	}
	return json.Marshal(toSerialize)
}

type NullableTransportDocumentHeader struct {
	value *TransportDocumentHeader
	isSet bool
}

func (v NullableTransportDocumentHeader) Get() *TransportDocumentHeader {
	return v.value
}

func (v *NullableTransportDocumentHeader) Set(val *TransportDocumentHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportDocumentHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportDocumentHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportDocumentHeader(val *TransportDocumentHeader) *NullableTransportDocumentHeader {
	return &NullableTransportDocumentHeader{value: val, isSet: true}
}

func (v NullableTransportDocumentHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportDocumentHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
