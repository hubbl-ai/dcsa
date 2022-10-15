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

// CarrierCodeListProvider The provider used for identifying the carrier Code
type CarrierCodeListProvider string

// List of carrierCodeListProvider
const (
	CarrierCodeListProviderSMDG  CarrierCodeListProvider = "SMDG"
	CarrierCodeListProviderNMFTA CarrierCodeListProvider = "NMFTA"
)

// All allowed values of CarrierCodeListProvider enum
var AllowedCarrierCodeListProviderEnumValues = []CarrierCodeListProvider{
	CarrierCodeListProviderSMDG,
	CarrierCodeListProviderNMFTA,
}

func (v *CarrierCodeListProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CarrierCodeListProvider(value)
	for _, existing := range AllowedCarrierCodeListProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CarrierCodeListProvider", value)
}

// NewCarrierCodeListProviderFromValue returns a pointer to a valid CarrierCodeListProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarrierCodeListProviderFromValue(v string) (*CarrierCodeListProvider, error) {
	ev := CarrierCodeListProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarrierCodeListProvider: valid values are %v", v, AllowedCarrierCodeListProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarrierCodeListProvider) IsValid() bool {
	for _, existing := range AllowedCarrierCodeListProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to carrierCodeListProvider value
func (v CarrierCodeListProvider) Ptr() *CarrierCodeListProvider {
	return &v
}

type NullableCarrierCodeListProvider struct {
	value *CarrierCodeListProvider
	isSet bool
}

func (v NullableCarrierCodeListProvider) Get() *CarrierCodeListProvider {
	return v.value
}

func (v *NullableCarrierCodeListProvider) Set(val *CarrierCodeListProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierCodeListProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierCodeListProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierCodeListProvider(val *CarrierCodeListProvider) *NullableCarrierCodeListProvider {
	return &NullableCarrierCodeListProvider{value: val, isSet: true}
}

func (v NullableCarrierCodeListProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierCodeListProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
