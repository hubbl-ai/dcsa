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
)

// BaseShipmentEvent The shipment event entity is a specialization of the event entity to support specification of data that only applies to a shipment event.
type BaseShipmentEvent struct {
	EventType *string `json:"eventType,omitempty"`
	// Value for eventDateTime must be the same value as eventCreatedDateTime
	EventDateTime interface{} `json:"eventDateTime,omitempty"`
	// Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)
	EventClassifierCode   *string               `json:"eventClassifierCode,omitempty"`
	ShipmentEventTypeCode ShipmentEventTypeCode `json:"shipmentEventTypeCode"`
	// The id of the object defined by the documentTypeCode.
	DocumentID       string           `json:"documentID"`
	DocumentTypeCode DocumentTypeCode `json:"documentTypeCode"`
	// Deprecated
	// ShipmentInformationTypeCode *ShipmentInformationType `json:"shipmentInformationTypeCode,omitempty"`
	// Reason field in a Shipment event. This field can be used to explain why a specific event has been sent.
	Reason *string `json:"reason,omitempty"`
	// Unique identifier for Event Type Code. For shipment events this can be - RECE (Received) - CONF (Confirmed) - ISSU (Issued) - APPR (Approved) - SUBM (Submitted) - SURR (Surrendered) - REJE (Rejected) - PENA (Pending approval)  Deprecated - use shipmentEventTypeCode instead
	// Deprecated
	EventTypeCode *string `json:"eventTypeCode,omitempty"`
	// ID uniquely identifying a shipment.  Deprecated - this is replaced by documentID which can contain different values depending on the value of the documentTypeCode field
	// Deprecated
	// ShipmentID interface{} `json:"shipmentID,omitempty"`
	References []Reference `json:"references,omitempty"`
}

// NewBaseShipmentEvent instantiates a new BaseShipmentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShipmentEvent(shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode) *BaseShipmentEvent {
	this := BaseShipmentEvent{}
	this.ShipmentEventTypeCode = shipmentEventTypeCode
	this.DocumentID = documentID
	this.DocumentTypeCode = documentTypeCode
	return &this
}

// NewBaseShipmentEventWithDefaults instantiates a new BaseShipmentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShipmentEventWithDefaults() *BaseShipmentEvent {
	this := BaseShipmentEvent{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *BaseShipmentEvent) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *BaseShipmentEvent) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *BaseShipmentEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetEventDateTime returns the EventDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseShipmentEvent) GetEventDateTime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseShipmentEvent) GetEventDateTimeOk() (*interface{}, bool) {
	if o == nil || o.EventDateTime == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// HasEventDateTime returns a boolean if a field has been set.
func (o *BaseShipmentEvent) HasEventDateTime() bool {
	if o != nil && o.EventDateTime != nil {
		return true
	}

	return false
}

// SetEventDateTime gets a reference to the given interface{} and assigns it to the EventDateTime field.
func (o *BaseShipmentEvent) SetEventDateTime(v interface{}) {
	o.EventDateTime = v
}

// GetEventClassifierCode returns the EventClassifierCode field value if set, zero value otherwise.
func (o *BaseShipmentEvent) GetEventClassifierCode() string {
	if o == nil || o.EventClassifierCode == nil {
		var ret string
		return ret
	}
	return *o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil || o.EventClassifierCode == nil {
		return nil, false
	}
	return o.EventClassifierCode, true
}

// HasEventClassifierCode returns a boolean if a field has been set.
func (o *BaseShipmentEvent) HasEventClassifierCode() bool {
	if o != nil && o.EventClassifierCode != nil {
		return true
	}

	return false
}

// SetEventClassifierCode gets a reference to the given string and assigns it to the EventClassifierCode field.
func (o *BaseShipmentEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = &v
}

// GetShipmentEventTypeCode returns the ShipmentEventTypeCode field value
func (o *BaseShipmentEvent) GetShipmentEventTypeCode() ShipmentEventTypeCode {
	if o == nil {
		var ret ShipmentEventTypeCode
		return ret
	}

	return o.ShipmentEventTypeCode
}

// GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentEventTypeCode, true
}

// SetShipmentEventTypeCode sets field value
func (o *BaseShipmentEvent) SetShipmentEventTypeCode(v ShipmentEventTypeCode) {
	o.ShipmentEventTypeCode = v
}

// GetDocumentID returns the DocumentID field value
func (o *BaseShipmentEvent) GetDocumentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentID
}

// GetDocumentIDOk returns a tuple with the DocumentID field value
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetDocumentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentID, true
}

// SetDocumentID sets field value
func (o *BaseShipmentEvent) SetDocumentID(v string) {
	o.DocumentID = v
}

// GetDocumentTypeCode returns the DocumentTypeCode field value
func (o *BaseShipmentEvent) GetDocumentTypeCode() DocumentTypeCode {
	if o == nil {
		var ret DocumentTypeCode
		return ret
	}

	return o.DocumentTypeCode
}

// GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field value
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentTypeCode, true
}

// SetDocumentTypeCode sets field value
func (o *BaseShipmentEvent) SetDocumentTypeCode(v DocumentTypeCode) {
	o.DocumentTypeCode = v
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BaseShipmentEvent) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BaseShipmentEvent) SetReason(v string) {
	o.Reason = &v
}

// GetEventTypeCode returns the EventTypeCode field value if set, zero value otherwise.
// Deprecated
func (o *BaseShipmentEvent) GetEventTypeCode() string {
	if o == nil || o.EventTypeCode == nil {
		var ret string
		return ret
	}
	return *o.EventTypeCode
}

// GetEventTypeCodeOk returns a tuple with the EventTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BaseShipmentEvent) GetEventTypeCodeOk() (*string, bool) {
	if o == nil || o.EventTypeCode == nil {
		return nil, false
	}
	return o.EventTypeCode, true
}

// HasEventTypeCode returns a boolean if a field has been set.
func (o *BaseShipmentEvent) HasEventTypeCode() bool {
	if o != nil && o.EventTypeCode != nil {
		return true
	}

	return false
}

// SetEventTypeCode gets a reference to the given string and assigns it to the EventTypeCode field.
// Deprecated
func (o *BaseShipmentEvent) SetEventTypeCode(v string) {
	o.EventTypeCode = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BaseShipmentEvent) GetReferences() []Reference {
	if o == nil || o.References == nil {
		var ret []Reference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEvent) GetReferencesOk() ([]Reference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BaseShipmentEvent) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []Reference and assigns it to the References field.
func (o *BaseShipmentEvent) SetReferences(v []Reference) {
	o.References = v
}

func (o BaseShipmentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventDateTime != nil {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if o.EventClassifierCode != nil {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	if true {
		toSerialize["shipmentEventTypeCode"] = o.ShipmentEventTypeCode
	}
	if true {
		toSerialize["documentID"] = o.DocumentID
	}
	if true {
		toSerialize["documentTypeCode"] = o.DocumentTypeCode
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.EventTypeCode != nil {
		toSerialize["eventTypeCode"] = o.EventTypeCode
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableBaseShipmentEvent struct {
	value *BaseShipmentEvent
	isSet bool
}

func (v NullableBaseShipmentEvent) Get() *BaseShipmentEvent {
	return v.value
}

func (v *NullableBaseShipmentEvent) Set(val *BaseShipmentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShipmentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShipmentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShipmentEvent(val *BaseShipmentEvent) *NullableBaseShipmentEvent {
	return &NullableBaseShipmentEvent{value: val, isSet: true}
}

func (v NullableBaseShipmentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShipmentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}