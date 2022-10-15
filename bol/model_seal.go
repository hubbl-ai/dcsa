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

	"github.com/hubbl-ai/dcsa/common"
)

// Seal addresses the seal-related information associated with the shipment equipment. A seal is put on a shipment equipment once it is loaded. This seal is meant to stay on until the shipment equipment reaches its final destination.
type Seal struct {
	// Identifies a seal affixed to the container.
	SealNumber string             `json:"sealNumber"`
	SealSource *common.SealSource `json:"sealSource,omitempty"`
	SealType   common.SealType    `json:"sealType"`
}

// NewSeal instantiates a new Seal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeal(sealNumber string, sealType common.SealType) *Seal {
	this := Seal{}
	this.SealNumber = sealNumber
	this.SealType = sealType
	return &this
}

// NewSealWithDefaults instantiates a new Seal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSealWithDefaults() *Seal {
	this := Seal{}
	return &this
}

// GetSealNumber returns the SealNumber field value
func (o *Seal) GetSealNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SealNumber
}

// GetSealNumberOk returns a tuple with the SealNumber field value
// and a boolean to check if the value has been set.
func (o *Seal) GetSealNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SealNumber, true
}

// SetSealNumber sets field value
func (o *Seal) SetSealNumber(v string) {
	o.SealNumber = v
}

// GetSealSource returns the SealSource field value if set, zero value otherwise.
func (o *Seal) GetSealSource() common.SealSource {
	if o == nil || o.SealSource == nil {
		var ret common.SealSource
		return ret
	}
	return *o.SealSource
}

// GetSealSourceOk returns a tuple with the SealSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seal) GetSealSourceOk() (*common.SealSource, bool) {
	if o == nil || o.SealSource == nil {
		return nil, false
	}
	return o.SealSource, true
}

// HasSealSource returns a boolean if a field has been set.
func (o *Seal) HasSealSource() bool {
	if o != nil && o.SealSource != nil {
		return true
	}

	return false
}

// SetSealSource gets a reference to the given SealSource and assigns it to the SealSource field.
func (o *Seal) SetSealSource(v common.SealSource) {
	o.SealSource = &v
}

// GetSealType returns the SealType field value
func (o *Seal) GetSealType() common.SealType {
	if o == nil {
		var ret common.SealType
		return ret
	}

	return o.SealType
}

// GetSealTypeOk returns a tuple with the SealType field value
// and a boolean to check if the value has been set.
func (o *Seal) GetSealTypeOk() (*common.SealType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SealType, true
}

// SetSealType sets field value
func (o *Seal) SetSealType(v common.SealType) {
	o.SealType = v
}

func (o Seal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sealNumber"] = o.SealNumber
	}
	if o.SealSource != nil {
		toSerialize["sealSource"] = o.SealSource
	}
	if true {
		toSerialize["sealType"] = o.SealType
	}
	return json.Marshal(toSerialize)
}

type NullableSeal struct {
	value *Seal
	isSet bool
}

func (v NullableSeal) Get() *Seal {
	return v.value
}

func (v *NullableSeal) Set(val *Seal) {
	v.value = val
	v.isSet = true
}

func (v NullableSeal) IsSet() bool {
	return v.isSet
}

func (v *NullableSeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeal(val *Seal) *NullableSeal {
	return &NullableSeal{value: val, isSet: true}
}

func (v NullableSeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
