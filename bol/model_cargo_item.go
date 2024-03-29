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

// CargoItem addresses the cargo items to be stuffed into a container for a shipment. A cargo item cannot be split across containers.
type CargoItem struct {
	CargoLineItems []CargoLineItem `json:"cargoLineItems,omitempty"`
	// The associated booking number provided by the carrier for this cargo item. Cannot be used in combination with Shipping Instruction header level carrierBookingReference
	CarrierBookingReference interface{} `json:"carrierBookingReference,omitempty"`
	// The cargo description are details which accurately and properly describe the cargo being shipped in the container(s) as provided by the shipper.
	DescriptionOfGoods string `json:"descriptionOfGoods"`
	// Used by customs to classify the product being shipped.
	HSCode string `json:"HSCode"`
	// The total weight of the cargo including packaging items being carried in the container(s). Excludes the tare weight of the container(s).
	Weight float32 `json:"weight"`
	// Calculated by multiplying the width, height, and length of the packed cargo.
	Volume     *float32           `json:"volume,omitempty"`
	WeightUnit common.WeightUnit  `json:"weightUnit"`
	VolumeUnit *common.VolumeUnit `json:"volumeUnit,omitempty"`
	// Specifies the number of packages associated with this cargo item
	NumberOfPackages int32 `json:"numberOfPackages"`
	// The unique identifier for the package type
	PackageCode string `json:"packageCode"`
	// The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.
	EquipmentReference string `json:"equipmentReference"`
}

// NewCargoItem instantiates a new CargoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCargoItem(descriptionOfGoods string, hSCode string, weight float32, weightUnit common.WeightUnit, numberOfPackages int32, packageCode string, equipmentReference string) *CargoItem {
	this := CargoItem{}
	this.DescriptionOfGoods = descriptionOfGoods
	this.HSCode = hSCode
	this.Weight = weight
	this.WeightUnit = weightUnit
	this.NumberOfPackages = numberOfPackages
	this.PackageCode = packageCode
	this.EquipmentReference = equipmentReference
	return &this
}

// NewCargoItemWithDefaults instantiates a new CargoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCargoItemWithDefaults() *CargoItem {
	this := CargoItem{}
	return &this
}

// GetCargoLineItems returns the CargoLineItems field value if set, zero value otherwise.
func (o *CargoItem) GetCargoLineItems() []CargoLineItem {
	if o == nil || o.CargoLineItems == nil {
		var ret []CargoLineItem
		return ret
	}
	return o.CargoLineItems
}

// GetCargoLineItemsOk returns a tuple with the CargoLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CargoItem) GetCargoLineItemsOk() ([]CargoLineItem, bool) {
	if o == nil || o.CargoLineItems == nil {
		return nil, false
	}
	return o.CargoLineItems, true
}

// HasCargoLineItems returns a boolean if a field has been set.
func (o *CargoItem) HasCargoLineItems() bool {
	if o != nil && o.CargoLineItems != nil {
		return true
	}

	return false
}

// SetCargoLineItems gets a reference to the given []CargoLineItem and assigns it to the CargoLineItems field.
func (o *CargoItem) SetCargoLineItems(v []CargoLineItem) {
	o.CargoLineItems = v
}

// GetCarrierBookingReference returns the CarrierBookingReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CargoItem) GetCarrierBookingReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CarrierBookingReference
}

// GetCarrierBookingReferenceOk returns a tuple with the CarrierBookingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CargoItem) GetCarrierBookingReferenceOk() (*interface{}, bool) {
	if o == nil || o.CarrierBookingReference == nil {
		return nil, false
	}
	return &o.CarrierBookingReference, true
}

// HasCarrierBookingReference returns a boolean if a field has been set.
func (o *CargoItem) HasCarrierBookingReference() bool {
	if o != nil && o.CarrierBookingReference != nil {
		return true
	}

	return false
}

// SetCarrierBookingReference gets a reference to the given interface{} and assigns it to the CarrierBookingReference field.
func (o *CargoItem) SetCarrierBookingReference(v interface{}) {
	o.CarrierBookingReference = v
}

// GetDescriptionOfGoods returns the DescriptionOfGoods field value
func (o *CargoItem) GetDescriptionOfGoods() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DescriptionOfGoods
}

// GetDescriptionOfGoodsOk returns a tuple with the DescriptionOfGoods field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetDescriptionOfGoodsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionOfGoods, true
}

// SetDescriptionOfGoods sets field value
func (o *CargoItem) SetDescriptionOfGoods(v string) {
	o.DescriptionOfGoods = v
}

// GetHSCode returns the HSCode field value
func (o *CargoItem) GetHSCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HSCode
}

// GetHSCodeOk returns a tuple with the HSCode field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetHSCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HSCode, true
}

// SetHSCode sets field value
func (o *CargoItem) SetHSCode(v string) {
	o.HSCode = v
}

// GetWeight returns the Weight field value
func (o *CargoItem) GetWeight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetWeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *CargoItem) SetWeight(v float32) {
	o.Weight = v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *CargoItem) GetVolume() float32 {
	if o == nil || o.Volume == nil {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CargoItem) GetVolumeOk() (*float32, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *CargoItem) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *CargoItem) SetVolume(v float32) {
	o.Volume = &v
}

// GetWeightUnit returns the WeightUnit field value
func (o *CargoItem) GetWeightUnit() common.WeightUnit {
	if o == nil {
		var ret common.WeightUnit
		return ret
	}

	return o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetWeightUnitOk() (*common.WeightUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WeightUnit, true
}

// SetWeightUnit sets field value
func (o *CargoItem) SetWeightUnit(v common.WeightUnit) {
	o.WeightUnit = v
}

// GetVolumeUnit returns the VolumeUnit field value if set, zero value otherwise.
func (o *CargoItem) GetVolumeUnit() common.VolumeUnit {
	if o == nil || o.VolumeUnit == nil {
		var ret common.VolumeUnit
		return ret
	}
	return *o.VolumeUnit
}

// GetVolumeUnitOk returns a tuple with the VolumeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CargoItem) GetVolumeUnitOk() (*common.VolumeUnit, bool) {
	if o == nil || o.VolumeUnit == nil {
		return nil, false
	}
	return o.VolumeUnit, true
}

// HasVolumeUnit returns a boolean if a field has been set.
func (o *CargoItem) HasVolumeUnit() bool {
	if o != nil && o.VolumeUnit != nil {
		return true
	}

	return false
}

// SetVolumeUnit gets a reference to the given VolumeUnit and assigns it to the VolumeUnit field.
func (o *CargoItem) SetVolumeUnit(v common.VolumeUnit) {
	o.VolumeUnit = &v
}

// GetNumberOfPackages returns the NumberOfPackages field value
func (o *CargoItem) GetNumberOfPackages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfPackages
}

// GetNumberOfPackagesOk returns a tuple with the NumberOfPackages field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetNumberOfPackagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfPackages, true
}

// SetNumberOfPackages sets field value
func (o *CargoItem) SetNumberOfPackages(v int32) {
	o.NumberOfPackages = v
}

// GetPackageCode returns the PackageCode field value
func (o *CargoItem) GetPackageCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetPackageCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageCode, true
}

// SetPackageCode sets field value
func (o *CargoItem) SetPackageCode(v string) {
	o.PackageCode = v
}

// GetEquipmentReference returns the EquipmentReference field value
func (o *CargoItem) GetEquipmentReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EquipmentReference
}

// GetEquipmentReferenceOk returns a tuple with the EquipmentReference field value
// and a boolean to check if the value has been set.
func (o *CargoItem) GetEquipmentReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquipmentReference, true
}

// SetEquipmentReference sets field value
func (o *CargoItem) SetEquipmentReference(v string) {
	o.EquipmentReference = v
}

func (o CargoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CargoLineItems != nil {
		toSerialize["cargoLineItems"] = o.CargoLineItems
	}
	if o.CarrierBookingReference != nil {
		toSerialize["carrierBookingReference"] = o.CarrierBookingReference
	}
	if true {
		toSerialize["descriptionOfGoods"] = o.DescriptionOfGoods
	}
	if true {
		toSerialize["HSCode"] = o.HSCode
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if true {
		toSerialize["weightUnit"] = o.WeightUnit
	}
	if o.VolumeUnit != nil {
		toSerialize["volumeUnit"] = o.VolumeUnit
	}
	if true {
		toSerialize["numberOfPackages"] = o.NumberOfPackages
	}
	if true {
		toSerialize["packageCode"] = o.PackageCode
	}
	if true {
		toSerialize["equipmentReference"] = o.EquipmentReference
	}
	return json.Marshal(toSerialize)
}

type NullableCargoItem struct {
	value *CargoItem
	isSet bool
}

func (v NullableCargoItem) Get() *CargoItem {
	return v.value
}

func (v *NullableCargoItem) Set(val *CargoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCargoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCargoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCargoItem(val *CargoItem) *NullableCargoItem {
	return &NullableCargoItem{value: val, isSet: true}
}

func (v NullableCargoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCargoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
