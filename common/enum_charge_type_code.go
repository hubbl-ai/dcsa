/*
DCSA Components

Common components for DCSA APIs and Domains. This Domain maintains common SimpleTypes, Headers and parameters used throughout all DCSA APIs and Domains.  For a history of changes in this domain - please check [GitHub](https://github.com/dcsaorg/DCSA-OpenAPI/tree/master/domain/dcsa#v202). Please also [create a GitHub issue](https://github.com/dcsaorg/DCSA-OpenAPI/issues/new) if you have any questions/comments.

API version: 2.0.2
Contact: info@dcsa.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package common

import (
	"encoding/json"
	"fmt"
)

// ChargeTypeCode Description of the charge type applied.
type ChargeTypeCode string

// List of chargeTypeCode
const (
	ChargeTypeOriginCFR                   ChargeTypeCode = "OCFR"
	ChargeTypeOriginTerminalHandling      ChargeTypeCode = "OTHC"
	ChargeTypeDestinationTerminalHandling ChargeTypeCode = "DTHC"
	ChargeTypeOriginDF                    ChargeTypeCode = "ODF"
	ChargeTypeDestinationDF               ChargeTypeCode = "DDF"
)

// All allowed values of ChargeTypeCode enum
var AllowedChargeTypeCodeEnumValues = []ChargeTypeCode{
	ChargeTypeOriginCFR,
	ChargeTypeOriginTerminalHandling,
	ChargeTypeDestinationTerminalHandling,
	ChargeTypeOriginDF,
	ChargeTypeDestinationDF,
}

func (v *ChargeTypeCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChargeTypeCode(value)
	for _, existing := range AllowedChargeTypeCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChargeTypeCode", value)
}

// NewChargeTypeCodeFromValue returns a pointer to a valid ChargeTypeCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChargeTypeCodeFromValue(v string) (*ChargeTypeCode, error) {
	ev := ChargeTypeCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChargeTypeCode: valid values are %v", v, AllowedChargeTypeCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChargeTypeCode) IsValid() bool {
	for _, existing := range AllowedChargeTypeCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to chargeTypeCode value
func (v ChargeTypeCode) Ptr() *ChargeTypeCode {
	return &v
}

type NullableChargeTypeCode struct {
	value *ChargeTypeCode
	isSet bool
}

func (v NullableChargeTypeCode) Get() *ChargeTypeCode {
	return v.value
}

func (v *NullableChargeTypeCode) Set(val *ChargeTypeCode) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeTypeCode) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeTypeCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeTypeCode(val *ChargeTypeCode) *NullableChargeTypeCode {
	return &NullableChargeTypeCode{value: val, isSet: true}
}

func (v NullableChargeTypeCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeTypeCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
