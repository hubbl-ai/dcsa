# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationName** | Pointer to **string** | The name of the location. | [optional] 
**Latitude** | Pointer to **string** | Geographic coordinate that specifies the north–south position of a point on the Earth&amp;apos;s surface. | [optional] 
**Longitude** | Pointer to **string** | Geographic coordinate that specifies the east–west position of a point on the Earth&amp;apos;s surface. | [optional] 
**UNLocationCode** | Pointer to **string** | The UN Location code specifying where the place is located. | [optional] 
**Address** | Pointer to [**Address**](address.md) | Address related information | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationName

`func (o *Location) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *Location) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *Location) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *Location) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetLatitude

`func (o *Location) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Location) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *Location) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Location) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetUNLocationCode

`func (o *Location) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *Location) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *Location) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.

### HasUNLocationCode

`func (o *Location) HasUNLocationCode() bool`

HasUNLocationCode returns a boolean if a field has been set.

### GetAddress

`func (o *Location) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Location) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Location) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Location) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


