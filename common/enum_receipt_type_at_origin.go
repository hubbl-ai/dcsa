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

// ReceiptTypeAtOrigin Indicates the type of service offered at Origin. Options are defined in the Receipt/Delivery entity. - CY (Container yard (incl. rail ramp)) - SD (Store Door) - CFS (Container Freight Station)
type ReceiptTypeAtOrigin string

// List of receiptTypeAtOrigin
const (
	CY  ReceiptTypeAtOrigin = "CY"
	SD  ReceiptTypeAtOrigin = "SD"
	CFS ReceiptTypeAtOrigin = "CFS"
)

// All allowed values of ReceiptTypeAtOrigin enum
var AllowedReceiptTypeAtOriginEnumValues = []ReceiptTypeAtOrigin{
	"CY",
	"SD",
	"CFS",
}

func (v *ReceiptTypeAtOrigin) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReceiptTypeAtOrigin(value)
	for _, existing := range AllowedReceiptTypeAtOriginEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReceiptTypeAtOrigin", value)
}

// NewReceiptTypeAtOriginFromValue returns a pointer to a valid ReceiptTypeAtOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReceiptTypeAtOriginFromValue(v string) (*ReceiptTypeAtOrigin, error) {
	ev := ReceiptTypeAtOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReceiptTypeAtOrigin: valid values are %v", v, AllowedReceiptTypeAtOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReceiptTypeAtOrigin) IsValid() bool {
	for _, existing := range AllowedReceiptTypeAtOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to receiptTypeAtOrigin value
func (v ReceiptTypeAtOrigin) Ptr() *ReceiptTypeAtOrigin {
	return &v
}

type NullableReceiptTypeAtOrigin struct {
	value *ReceiptTypeAtOrigin
	isSet bool
}

func (v NullableReceiptTypeAtOrigin) Get() *ReceiptTypeAtOrigin {
	return v.value
}

func (v *NullableReceiptTypeAtOrigin) Set(val *ReceiptTypeAtOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptTypeAtOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptTypeAtOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptTypeAtOrigin(val *ReceiptTypeAtOrigin) *NullableReceiptTypeAtOrigin {
	return &NullableReceiptTypeAtOrigin{value: val, isSet: true}
}

func (v NullableReceiptTypeAtOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptTypeAtOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
