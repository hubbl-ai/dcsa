# CargoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CargoLineItems** | Pointer to [**[]CargoLineItem**](CargoLineItem.md) |  | [optional] 
**CarrierBookingReference** | Pointer to **interface{}** | The associated booking number provided by the carrier for this cargo item. Cannot be used in combination with Shipping Instruction header level carrierBookingReference | [optional] 
**DescriptionOfGoods** | **string** | The cargo description are details which accurately and properly describe the cargo being shipped in the container(s) as provided by the shipper. | 
**HSCode** | **string** | Used by customs to classify the product being shipped. | 
**Weight** | **float32** | The total weight of the cargo including packaging items being carried in the container(s). Excludes the tare weight of the container(s). | 
**Volume** | Pointer to **float32** | Calculated by multiplying the width, height, and length of the packed cargo. | [optional] 
**WeightUnit** | [**WeightUnit**](WeightUnit.md) |  | 
**VolumeUnit** | Pointer to [**VolumeUnit**](VolumeUnit.md) |  | [optional] 
**NumberOfPackages** | **int32** | Specifies the number of packages associated with this cargo item | 
**PackageCode** | **string** | The unique identifier for the package type | 
**EquipmentReference** | **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | 

## Methods

### NewCargoItem

`func NewCargoItem(descriptionOfGoods string, hSCode string, weight float32, weightUnit WeightUnit, numberOfPackages int32, packageCode string, equipmentReference string, ) *CargoItem`

NewCargoItem instantiates a new CargoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCargoItemWithDefaults

`func NewCargoItemWithDefaults() *CargoItem`

NewCargoItemWithDefaults instantiates a new CargoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCargoLineItems

`func (o *CargoItem) GetCargoLineItems() []CargoLineItem`

GetCargoLineItems returns the CargoLineItems field if non-nil, zero value otherwise.

### GetCargoLineItemsOk

`func (o *CargoItem) GetCargoLineItemsOk() (*[]CargoLineItem, bool)`

GetCargoLineItemsOk returns a tuple with the CargoLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoLineItems

`func (o *CargoItem) SetCargoLineItems(v []CargoLineItem)`

SetCargoLineItems sets CargoLineItems field to given value.

### HasCargoLineItems

`func (o *CargoItem) HasCargoLineItems() bool`

HasCargoLineItems returns a boolean if a field has been set.

### GetCarrierBookingReference

`func (o *CargoItem) GetCarrierBookingReference() interface{}`

GetCarrierBookingReference returns the CarrierBookingReference field if non-nil, zero value otherwise.

### GetCarrierBookingReferenceOk

`func (o *CargoItem) GetCarrierBookingReferenceOk() (*interface{}, bool)`

GetCarrierBookingReferenceOk returns a tuple with the CarrierBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierBookingReference

`func (o *CargoItem) SetCarrierBookingReference(v interface{})`

SetCarrierBookingReference sets CarrierBookingReference field to given value.

### HasCarrierBookingReference

`func (o *CargoItem) HasCarrierBookingReference() bool`

HasCarrierBookingReference returns a boolean if a field has been set.

### SetCarrierBookingReferenceNil

`func (o *CargoItem) SetCarrierBookingReferenceNil(b bool)`

 SetCarrierBookingReferenceNil sets the value for CarrierBookingReference to be an explicit nil

### UnsetCarrierBookingReference
`func (o *CargoItem) UnsetCarrierBookingReference()`

UnsetCarrierBookingReference ensures that no value is present for CarrierBookingReference, not even an explicit nil
### GetDescriptionOfGoods

`func (o *CargoItem) GetDescriptionOfGoods() string`

GetDescriptionOfGoods returns the DescriptionOfGoods field if non-nil, zero value otherwise.

### GetDescriptionOfGoodsOk

`func (o *CargoItem) GetDescriptionOfGoodsOk() (*string, bool)`

GetDescriptionOfGoodsOk returns a tuple with the DescriptionOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOfGoods

`func (o *CargoItem) SetDescriptionOfGoods(v string)`

SetDescriptionOfGoods sets DescriptionOfGoods field to given value.


### GetHSCode

`func (o *CargoItem) GetHSCode() string`

GetHSCode returns the HSCode field if non-nil, zero value otherwise.

### GetHSCodeOk

`func (o *CargoItem) GetHSCodeOk() (*string, bool)`

GetHSCodeOk returns a tuple with the HSCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSCode

`func (o *CargoItem) SetHSCode(v string)`

SetHSCode sets HSCode field to given value.


### GetWeight

`func (o *CargoItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CargoItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CargoItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetVolume

`func (o *CargoItem) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CargoItem) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CargoItem) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *CargoItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWeightUnit

`func (o *CargoItem) GetWeightUnit() WeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *CargoItem) GetWeightUnitOk() (*WeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *CargoItem) SetWeightUnit(v WeightUnit)`

SetWeightUnit sets WeightUnit field to given value.


### GetVolumeUnit

`func (o *CargoItem) GetVolumeUnit() VolumeUnit`

GetVolumeUnit returns the VolumeUnit field if non-nil, zero value otherwise.

### GetVolumeUnitOk

`func (o *CargoItem) GetVolumeUnitOk() (*VolumeUnit, bool)`

GetVolumeUnitOk returns a tuple with the VolumeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUnit

`func (o *CargoItem) SetVolumeUnit(v VolumeUnit)`

SetVolumeUnit sets VolumeUnit field to given value.

### HasVolumeUnit

`func (o *CargoItem) HasVolumeUnit() bool`

HasVolumeUnit returns a boolean if a field has been set.

### GetNumberOfPackages

`func (o *CargoItem) GetNumberOfPackages() int32`

GetNumberOfPackages returns the NumberOfPackages field if non-nil, zero value otherwise.

### GetNumberOfPackagesOk

`func (o *CargoItem) GetNumberOfPackagesOk() (*int32, bool)`

GetNumberOfPackagesOk returns a tuple with the NumberOfPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPackages

`func (o *CargoItem) SetNumberOfPackages(v int32)`

SetNumberOfPackages sets NumberOfPackages field to given value.


### GetPackageCode

`func (o *CargoItem) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *CargoItem) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *CargoItem) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.


### GetEquipmentReference

`func (o *CargoItem) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *CargoItem) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *CargoItem) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


