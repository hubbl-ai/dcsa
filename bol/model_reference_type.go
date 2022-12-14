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
	"fmt"
)

// ReferenceType The reference type codes defined by DCSA. - FF (Freight Forwarder’s Reference) - SI (Shipper’s Reference) - PO (Purchase Order Reference) - CR (Customer’s Reference) - AAO (Consignee’s Reference)
type ReferenceType string

// List of referenceType
const (
	FF  ReferenceType = "FF"
	SI  ReferenceType = "SI"
	PO  ReferenceType = "PO"
	CR  ReferenceType = "CR"
	AAO ReferenceType = "AAO"
)

// All allowed values of ReferenceType enum
var AllowedReferenceTypeEnumValues = []ReferenceType{
	"FF",
	"SI",
	"PO",
	"CR",
	"AAO",
}

func (v *ReferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReferenceType(value)
	for _, existing := range AllowedReferenceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReferenceType", value)
}

// NewReferenceTypeFromValue returns a pointer to a valid ReferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReferenceTypeFromValue(v string) (*ReferenceType, error) {
	ev := ReferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReferenceType: valid values are %v", v, AllowedReferenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReferenceType) IsValid() bool {
	for _, existing := range AllowedReferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to referenceType value
func (v ReferenceType) Ptr() *ReferenceType {
	return &v
}

type NullableReferenceType struct {
	value *ReferenceType
	isSet bool
}

func (v NullableReferenceType) Get() *ReferenceType {
	return v.value
}

func (v *NullableReferenceType) Set(val *ReferenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceType(val *ReferenceType) *NullableReferenceType {
	return &NullableReferenceType{value: val, isSet: true}
}

func (v NullableReferenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
