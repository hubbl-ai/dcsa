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
	"time"
)

// BaseEventBody The Event entity is described as a generalization of all the specific event categories. An event always takes place in relation to a shipment and can additionally be linked to a transport or an equipment
type BaseEventBody struct {
	EventType string `json:"eventType"`
	// Code for the event classifier. Values can vary depending on eventType
	EventClassifierCode string `json:"eventClassifierCode"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format.
	EventDateTime time.Time `json:"eventDateTime"`
}

// NewBaseEventBody instantiates a new BaseEventBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEventBody(eventType string, eventClassifierCode string, eventDateTime time.Time) *BaseEventBody {
	this := BaseEventBody{}
	this.EventType = eventType
	this.EventClassifierCode = eventClassifierCode
	this.EventDateTime = eventDateTime
	return &this
}

// NewBaseEventBodyWithDefaults instantiates a new BaseEventBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventBodyWithDefaults() *BaseEventBody {
	this := BaseEventBody{}
	return &this
}

// GetEventType returns the EventType field value
func (o *BaseEventBody) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *BaseEventBody) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *BaseEventBody) SetEventType(v string) {
	o.EventType = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *BaseEventBody) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *BaseEventBody) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *BaseEventBody) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *BaseEventBody) GetEventDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *BaseEventBody) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *BaseEventBody) SetEventDateTime(v time.Time) {
	o.EventDateTime = v
}

func (o BaseEventBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	if true {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEventBody struct {
	value *BaseEventBody
	isSet bool
}

func (v NullableBaseEventBody) Get() *BaseEventBody {
	return v.value
}

func (v *NullableBaseEventBody) Set(val *BaseEventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEventBody(val *BaseEventBody) *NullableBaseEventBody {
	return &NullableBaseEventBody{value: val, isSet: true}
}

func (v NullableBaseEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}