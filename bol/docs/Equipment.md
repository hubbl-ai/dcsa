# Equipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentReference** | **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | 
**ISOEquipmentCode** | Pointer to **string** | Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**TareWeight** | Pointer to **float32** | The weight of an empty container (gross container weight). | [optional] 
**WeightUnit** | Pointer to [**WeightUnit**](WeightUnit.md) |  | [optional] 
**IsShipperOwned** | **bool** | Indicates whether the container is shipper owned (SOC). | 

## Methods

### NewEquipment

`func NewEquipment(equipmentReference string, isShipperOwned bool, ) *Equipment`

NewEquipment instantiates a new Equipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentWithDefaults

`func NewEquipmentWithDefaults() *Equipment`

NewEquipmentWithDefaults instantiates a new Equipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentReference

`func (o *Equipment) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *Equipment) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *Equipment) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.


### GetISOEquipmentCode

`func (o *Equipment) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *Equipment) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *Equipment) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *Equipment) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetTareWeight

`func (o *Equipment) GetTareWeight() float32`

GetTareWeight returns the TareWeight field if non-nil, zero value otherwise.

### GetTareWeightOk

`func (o *Equipment) GetTareWeightOk() (*float32, bool)`

GetTareWeightOk returns a tuple with the TareWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTareWeight

`func (o *Equipment) SetTareWeight(v float32)`

SetTareWeight sets TareWeight field to given value.

### HasTareWeight

`func (o *Equipment) HasTareWeight() bool`

HasTareWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *Equipment) GetWeightUnit() WeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *Equipment) GetWeightUnitOk() (*WeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *Equipment) SetWeightUnit(v WeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *Equipment) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetIsShipperOwned

`func (o *Equipment) GetIsShipperOwned() bool`

GetIsShipperOwned returns the IsShipperOwned field if non-nil, zero value otherwise.

### GetIsShipperOwnedOk

`func (o *Equipment) GetIsShipperOwnedOk() (*bool, bool)`

GetIsShipperOwnedOk returns a tuple with the IsShipperOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShipperOwned

`func (o *Equipment) SetIsShipperOwned(v bool)`

SetIsShipperOwned sets IsShipperOwned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


