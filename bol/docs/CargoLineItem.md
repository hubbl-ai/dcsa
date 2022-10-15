# CargoLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CargoLineItemID** | **string** | Identifies the cargo line item (package) within the cargo. The cargo line item ID is provided by the shipper and is used to define the stuffing. Cargo line items belonging to the same cargo items are stuffed in the same container. | 
**ShippingMarks** | **string** | The identifying details of a package or the actual markings that appear on the package(s). This information is provided by the shipper. | 

## Methods

### NewCargoLineItem

`func NewCargoLineItem(cargoLineItemID string, shippingMarks string, ) *CargoLineItem`

NewCargoLineItem instantiates a new CargoLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCargoLineItemWithDefaults

`func NewCargoLineItemWithDefaults() *CargoLineItem`

NewCargoLineItemWithDefaults instantiates a new CargoLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCargoLineItemID

`func (o *CargoLineItem) GetCargoLineItemID() string`

GetCargoLineItemID returns the CargoLineItemID field if non-nil, zero value otherwise.

### GetCargoLineItemIDOk

`func (o *CargoLineItem) GetCargoLineItemIDOk() (*string, bool)`

GetCargoLineItemIDOk returns a tuple with the CargoLineItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoLineItemID

`func (o *CargoLineItem) SetCargoLineItemID(v string)`

SetCargoLineItemID sets CargoLineItemID field to given value.


### GetShippingMarks

`func (o *CargoLineItem) GetShippingMarks() string`

GetShippingMarks returns the ShippingMarks field if non-nil, zero value otherwise.

### GetShippingMarksOk

`func (o *CargoLineItem) GetShippingMarksOk() (*string, bool)`

GetShippingMarksOk returns a tuple with the ShippingMarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMarks

`func (o *CargoLineItem) SetShippingMarks(v string)`

SetShippingMarks sets ShippingMarks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


