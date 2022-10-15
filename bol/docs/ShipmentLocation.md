# ShipmentLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](location.md) |  | [optional] 
**DisplayedName** | Pointer to **string** | The address of the party to be displayed on the transport document. | [optional] 
**LocationType** | [**LocationType**](LocationType.md) |  | 

## Methods

### NewShipmentLocation

`func NewShipmentLocation(locationType LocationType, ) *ShipmentLocation`

NewShipmentLocation instantiates a new ShipmentLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentLocationWithDefaults

`func NewShipmentLocationWithDefaults() *ShipmentLocation`

NewShipmentLocationWithDefaults instantiates a new ShipmentLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ShipmentLocation) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ShipmentLocation) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ShipmentLocation) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ShipmentLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDisplayedName

`func (o *ShipmentLocation) GetDisplayedName() string`

GetDisplayedName returns the DisplayedName field if non-nil, zero value otherwise.

### GetDisplayedNameOk

`func (o *ShipmentLocation) GetDisplayedNameOk() (*string, bool)`

GetDisplayedNameOk returns a tuple with the DisplayedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedName

`func (o *ShipmentLocation) SetDisplayedName(v string)`

SetDisplayedName sets DisplayedName field to given value.

### HasDisplayedName

`func (o *ShipmentLocation) HasDisplayedName() bool`

HasDisplayedName returns a boolean if a field has been set.

### GetLocationType

`func (o *ShipmentLocation) GetLocationType() LocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *ShipmentLocation) GetLocationTypeOk() (*LocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *ShipmentLocation) SetLocationType(v LocationType)`

SetLocationType sets LocationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


