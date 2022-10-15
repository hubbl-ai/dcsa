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

// DocumentStatus The status of the document in the process. Possible values are:
// - RECE (Received)
// - DRFT (Drafted)
// - PENA (Pending Approval)
// - PENU (Pending Update)
// - REJE (Rejected)
// - APPR (Approved)
// - ISSU (Issued)
// - SURR (Surrendered)
// - SUBM (Submitted)
// - VOID (Void)
type DocumentStatus string

// List of documentStatus
const (
	DocumentStatusReceived        DocumentStatus = "RECE"
	DocumentStatusDrafted         DocumentStatus = "DRFT"
	DocumentStatusPendingApproval DocumentStatus = "PENA"
	DocumentStatusPendingUpdate   DocumentStatus = "PENU"
	DocumentStatusRejected        DocumentStatus = "REJE"
	DocumentStatusApproved        DocumentStatus = "APPR"
	DocumentStatusIssued          DocumentStatus = "ISSU"
	DocumentStatusSurrendered     DocumentStatus = "SURR"
	DocumentStatusSubmitted       DocumentStatus = "SUBM"
	DocumentStatusVoid            DocumentStatus = "VOID"
)

// All allowed values of DocumentStatus enum
var AllowedDocumentStatusEnumValues = []DocumentStatus{
	DocumentStatusReceived,
	DocumentStatusDrafted,
	DocumentStatusPendingApproval,
	DocumentStatusPendingUpdate,
	DocumentStatusRejected,
	DocumentStatusApproved,
	DocumentStatusIssued,
	DocumentStatusSurrendered,
	DocumentStatusSubmitted,
	DocumentStatusVoid,
}

func (v *DocumentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentStatus(value)
	for _, existing := range AllowedDocumentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentStatus", value)
}

// NewDocumentStatusFromValue returns a pointer to a valid DocumentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentStatusFromValue(v string) (*DocumentStatus, error) {
	ev := DocumentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentStatus: valid values are %v", v, AllowedDocumentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentStatus) IsValid() bool {
	for _, existing := range AllowedDocumentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to documentStatus value
func (v DocumentStatus) Ptr() *DocumentStatus {
	return &v
}

type NullableDocumentStatus struct {
	value *DocumentStatus
	isSet bool
}

func (v NullableDocumentStatus) Get() *DocumentStatus {
	return v.value
}

func (v *NullableDocumentStatus) Set(val *DocumentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentStatus(val *DocumentStatus) *NullableDocumentStatus {
	return &NullableDocumentStatus{value: val, isSet: true}
}

func (v NullableDocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
