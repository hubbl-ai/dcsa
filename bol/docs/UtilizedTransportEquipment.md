# UtilizedTransportEquipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equipment** | [**Equipment**](Equipment.md) |  | 
**CargoGrossWeightUnit** | [**WeightUnit**](WeightUnit.md) |  | 
**CargoGrossWeight** | **float32** | The grand total weight of the cargo and weight per container(s) including packaging items being carried, which can be expressed in imperial or metric terms, as provided by the shipper. Excludes the tare weight of the container(s). | 
**ActiveReeferSettings** | Pointer to **map[string]interface{}** | specifies the settings for an active reefer container used to a shipment. | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) |  | [optional] 

## Methods

### NewUtilizedTransportEquipment

`func NewUtilizedTransportEquipment(equipment Equipment, cargoGrossWeightUnit WeightUnit, cargoGrossWeight float32, ) *UtilizedTransportEquipment`

NewUtilizedTransportEquipment instantiates a new UtilizedTransportEquipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizedTransportEquipmentWithDefaults

`func NewUtilizedTransportEquipmentWithDefaults() *UtilizedTransportEquipment`

NewUtilizedTransportEquipmentWithDefaults instantiates a new UtilizedTransportEquipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipment

`func (o *UtilizedTransportEquipment) GetEquipment() Equipment`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *UtilizedTransportEquipment) GetEquipmentOk() (*Equipment, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *UtilizedTransportEquipment) SetEquipment(v Equipment)`

SetEquipment sets Equipment field to given value.


### GetCargoGrossWeightUnit

`func (o *UtilizedTransportEquipment) GetCargoGrossWeightUnit() WeightUnit`

GetCargoGrossWeightUnit returns the CargoGrossWeightUnit field if non-nil, zero value otherwise.

### GetCargoGrossWeightUnitOk

`func (o *UtilizedTransportEquipment) GetCargoGrossWeightUnitOk() (*WeightUnit, bool)`

GetCargoGrossWeightUnitOk returns a tuple with the CargoGrossWeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoGrossWeightUnit

`func (o *UtilizedTransportEquipment) SetCargoGrossWeightUnit(v WeightUnit)`

SetCargoGrossWeightUnit sets CargoGrossWeightUnit field to given value.


### GetCargoGrossWeight

`func (o *UtilizedTransportEquipment) GetCargoGrossWeight() float32`

GetCargoGrossWeight returns the CargoGrossWeight field if non-nil, zero value otherwise.

### GetCargoGrossWeightOk

`func (o *UtilizedTransportEquipment) GetCargoGrossWeightOk() (*float32, bool)`

GetCargoGrossWeightOk returns a tuple with the CargoGrossWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoGrossWeight

`func (o *UtilizedTransportEquipment) SetCargoGrossWeight(v float32)`

SetCargoGrossWeight sets CargoGrossWeight field to given value.


### GetActiveReeferSettings

`func (o *UtilizedTransportEquipment) GetActiveReeferSettings() map[string]interface{}`

GetActiveReeferSettings returns the ActiveReeferSettings field if non-nil, zero value otherwise.

### GetActiveReeferSettingsOk

`func (o *UtilizedTransportEquipment) GetActiveReeferSettingsOk() (*map[string]interface{}, bool)`

GetActiveReeferSettingsOk returns a tuple with the ActiveReeferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveReeferSettings

`func (o *UtilizedTransportEquipment) SetActiveReeferSettings(v map[string]interface{})`

SetActiveReeferSettings sets ActiveReeferSettings field to given value.

### HasActiveReeferSettings

`func (o *UtilizedTransportEquipment) HasActiveReeferSettings() bool`

HasActiveReeferSettings returns a boolean if a field has been set.

### GetSeals

`func (o *UtilizedTransportEquipment) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *UtilizedTransportEquipment) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *UtilizedTransportEquipment) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *UtilizedTransportEquipment) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


