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

// PartyFunction Specifies the role of the party in the context of the given Shipping Instruction.
// - OS (Original shipper)
// - CN (Consignee)
// - COW (Invoice payer on behalf of the consignor (shipper))
// - COX (Invoice payer on behalf of the consignee)
// - MS (Document/message issuer/sender)
// - N1 (First Notify Party)
// - N2 (Second Notify Party)
// - NI (Other Notify Party)
// - DDR (Consignor's freight forwarder)
// - DDS (Consignee's freight forwarder)
type PartyFunction string

// List of partyFunction
const (
	PartyFunctionOriginalShipper   PartyFunction = "OS"  // Original shipper
	PartyFunctionConsignee         PartyFunction = "CN"  // Consignee
	PartyFunctionConsignorBehalf   PartyFunction = "COW" // Invoice payer on behalf of the consignor (shipper)
	PartyFunctionConsigneeBehalf   PartyFunction = "COX" // Invoice payer on behalf of the consignee
	PartyFunctionIssuer            PartyFunction = "MS"  // Document/message issuer/sender
	PartyFunctionFirstNotifyParty                = "N1"  // First Notify Party
	PartyFunctionSecondNotifyParty               = "N2"  // Second Notify Party
	PartyFunctionOtherNotifyParty                = "NI"  // Other Notify Party
	PartyFunctionExportForwarder                 = "DDR" // Consignor's freight forwarder
	PartyFunctionImportForwarder                 = "DDS" // Consignee's freight forwarder
)

// All allowed values of PartyFunction enum
var AllowedPartyFunctionEnumValues = []PartyFunction{
	PartyFunctionOriginalShipper,
	PartyFunctionConsignee,
	PartyFunctionConsignorBehalf,
	PartyFunctionConsigneeBehalf,
	PartyFunctionIssuer,
	PartyFunctionFirstNotifyParty,
	PartyFunctionSecondNotifyParty,
	PartyFunctionOtherNotifyParty,
	PartyFunctionExportForwarder,
	PartyFunctionImportForwarder,
}

func (v *PartyFunction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartyFunction(value)
	for _, existing := range AllowedPartyFunctionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartyFunction", value)
}

// NewPartyFunctionFromValue returns a pointer to a valid PartyFunction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartyFunctionFromValue(v string) (*PartyFunction, error) {
	ev := PartyFunction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartyFunction: valid values are %v", v, AllowedPartyFunctionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartyFunction) IsValid() bool {
	for _, existing := range AllowedPartyFunctionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to partyFunction value
func (v PartyFunction) Ptr() *PartyFunction {
	return &v
}

type NullablePartyFunction struct {
	value *PartyFunction
	isSet bool
}

func (v NullablePartyFunction) Get() *PartyFunction {
	return v.value
}

func (v *NullablePartyFunction) Set(val *PartyFunction) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyFunction) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyFunction(val *PartyFunction) *NullablePartyFunction {
	return &NullablePartyFunction{value: val, isSet: true}
}

func (v NullablePartyFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}