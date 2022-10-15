/*
DCSA OpenAPI specification for Electronic Bill of Lading

API specification issued by DCSA.org.  For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2020/12/20201208-DCSA-P1-DCSA-Information-Model-v3.0-FINAL.pdf'>Information Model v3.0</a>  It is possible to use this API as a standalone API. In order to do so it is necessary to use the poll-endPoint - /v1/events  in order to poll event information.  It is recomended to implement the <a href='https://app.swaggerhub.com/apis/dcsaorg/DOCUMENTATION_EVENT_HUB'>DCSA Documentation Event Hub</a> in order to use the push model. Here events are pushed as they occur.

API version: 1.0.0
Contact: info@dcsa.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package common

import (
	"encoding/json"
	"fmt"
)

// ShipmentInformationType The code to identify the type of information documentID points to. Can be one of the following values
// - BOK (Booking) deprecated use BKG instead
// - BKG (Booking)
// - SHI (Shipping Instruction)
// - VGM (Verified Gross Mass)
// - SRM (Shipment Release Message)
// - TRD (Transport Document)
// - ARN (Arrival Notice)  <b>Deprecated</b> use documentTypeCode instead
type ShipmentInformationType string

// List of shipmentInformationType
const (
	_ShipmentInformationTypeBooking               ShipmentInformationType = "BOK" // Deprecated
	ShipmentInformationTypeBooking                ShipmentInformationType = "BKG" // Booking
	ShipmentInformationTypeShippingInstruction    ShipmentInformationType = "SHI" // Shipping instruction
	ShipmentInformationTypeVerifiedGrossMass      ShipmentInformationType = "VGM" // Verified gross mass
	ShipmentInformationTypeShipmentReleaseMessage ShipmentInformationType = "SRM" // Shipment release message
	ShipmentInformationTypeTransportDocument      ShipmentInformationType = "TRD" // Transport document
	_ShipmentInformationTypeArrivalNotice         ShipmentInformationType = "ARN" // Deprecated
)

// All allowed values of ShipmentInformationType enum
var AllowedShipmentInformationTypeEnumValues = []ShipmentInformationType{
	_ShipmentInformationTypeBooking,
	ShipmentInformationTypeBooking,
	ShipmentInformationTypeShippingInstruction,
	ShipmentInformationTypeVerifiedGrossMass,
	ShipmentInformationTypeShipmentReleaseMessage,
	ShipmentInformationTypeTransportDocument,
	_ShipmentInformationTypeArrivalNotice,
}

func (v *ShipmentInformationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipmentInformationType(value)
	for _, existing := range AllowedShipmentInformationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipmentInformationType", value)
}

// NewShipmentInformationTypeFromValue returns a pointer to a valid ShipmentInformationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentInformationTypeFromValue(v string) (*ShipmentInformationType, error) {
	ev := ShipmentInformationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentInformationType: valid values are %v", v, AllowedShipmentInformationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentInformationType) IsValid() bool {
	for _, existing := range AllowedShipmentInformationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shipmentInformationType value
func (v ShipmentInformationType) Ptr() *ShipmentInformationType {
	return &v
}

type NullableShipmentInformationType struct {
	value *ShipmentInformationType
	isSet bool
}

func (v NullableShipmentInformationType) Get() *ShipmentInformationType {
	return v.value
}

func (v *NullableShipmentInformationType) Set(val *ShipmentInformationType) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentInformationType) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentInformationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentInformationType(val *ShipmentInformationType) *NullableShipmentInformationType {
	return &NullableShipmentInformationType{value: val, isSet: true}
}

func (v NullableShipmentInformationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentInformationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
