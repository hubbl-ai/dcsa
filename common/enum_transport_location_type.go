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

// TransportLocationType Code of a location to which a certain transport event is corresponding (Berth, Pilot Boarding Place)
type TransportLocationType string

// List of transportLocationType
const (
	BERTH TransportLocationType = "BERTH"
	PBP   TransportLocationType = "PBP"
)

// All allowed values of TransportLocationType enum
var AllowedTransportLocationTypeEnumValues = []TransportLocationType{
	"BERTH",
	"PBP",
}

func (v *TransportLocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransportLocationType(value)
	for _, existing := range AllowedTransportLocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransportLocationType", value)
}

// NewTransportLocationTypeFromValue returns a pointer to a valid TransportLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransportLocationTypeFromValue(v string) (*TransportLocationType, error) {
	ev := TransportLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransportLocationType: valid values are %v", v, AllowedTransportLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransportLocationType) IsValid() bool {
	for _, existing := range AllowedTransportLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to transportLocationType value
func (v TransportLocationType) Ptr() *TransportLocationType {
	return &v
}

type NullableTransportLocationType struct {
	value *TransportLocationType
	isSet bool
}

func (v NullableTransportLocationType) Get() *TransportLocationType {
	return v.value
}

func (v *NullableTransportLocationType) Set(val *TransportLocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLocationType(val *TransportLocationType) *NullableTransportLocationType {
	return &NullableTransportLocationType{value: val, isSet: true}
}

func (v NullableTransportLocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}