# CargoItemAllOf9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | [optional] 

## Methods

### NewCargoItemAllOf9

`func NewCargoItemAllOf9() *CargoItemAllOf9`

NewCargoItemAllOf9 instantiates a new CargoItemAllOf9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCargoItemAllOf9WithDefaults

`func NewCargoItemAllOf9WithDefaults() *CargoItemAllOf9`

NewCargoItemAllOf9WithDefaults instantiates a new CargoItemAllOf9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentReference

`func (o *CargoItemAllOf9) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *CargoItemAllOf9) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *CargoItemAllOf9) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *CargoItemAllOf9) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


