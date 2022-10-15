# Transport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadLocation** | Pointer to [**Location**](location.md) |  | [optional] 
**DischargeLocation** | Pointer to [**Location**](location.md) |  | [optional] 
**PlannedDepartureDate** | Pointer to **string** | The planned date of departure.  | [optional] 
**PlannedArrivalDate** | Pointer to **string** | The planned date of arrival.  | [optional] 
**ModeOfTransport** | Pointer to [**ModeOfTransport**](ModeOfTransport.md) |  | [optional] 
**VesselIMONumber** | Pointer to **string** | The unique reference for a registered Vessel. The reference is the International Maritime Organisation (IMO) number, also sometimes known as the Lloyd&amp;apos;s register code, which does not change during the lifetime of the vessel  | [optional] 
**CarrierVoyageNumber** | Pointer to **string** | The vessel operator-specific identifier of the Voyage. | [optional] 
**IsUnderShippersResponsibility** | Pointer to **bool** | Indicator whether mode of transportation for pre-carriage (e.g. truck, barge, rail) is under shipper&#39;s responsibility  | [optional] 

## Methods

### NewTransport

`func NewTransport() *Transport`

NewTransport instantiates a new Transport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportWithDefaults

`func NewTransportWithDefaults() *Transport`

NewTransportWithDefaults instantiates a new Transport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadLocation

`func (o *Transport) GetLoadLocation() Location`

GetLoadLocation returns the LoadLocation field if non-nil, zero value otherwise.

### GetLoadLocationOk

`func (o *Transport) GetLoadLocationOk() (*Location, bool)`

GetLoadLocationOk returns a tuple with the LoadLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadLocation

`func (o *Transport) SetLoadLocation(v Location)`

SetLoadLocation sets LoadLocation field to given value.

### HasLoadLocation

`func (o *Transport) HasLoadLocation() bool`

HasLoadLocation returns a boolean if a field has been set.

### GetDischargeLocation

`func (o *Transport) GetDischargeLocation() Location`

GetDischargeLocation returns the DischargeLocation field if non-nil, zero value otherwise.

### GetDischargeLocationOk

`func (o *Transport) GetDischargeLocationOk() (*Location, bool)`

GetDischargeLocationOk returns a tuple with the DischargeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargeLocation

`func (o *Transport) SetDischargeLocation(v Location)`

SetDischargeLocation sets DischargeLocation field to given value.

### HasDischargeLocation

`func (o *Transport) HasDischargeLocation() bool`

HasDischargeLocation returns a boolean if a field has been set.

### GetPlannedDepartureDate

`func (o *Transport) GetPlannedDepartureDate() string`

GetPlannedDepartureDate returns the PlannedDepartureDate field if non-nil, zero value otherwise.

### GetPlannedDepartureDateOk

`func (o *Transport) GetPlannedDepartureDateOk() (*string, bool)`

GetPlannedDepartureDateOk returns a tuple with the PlannedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDepartureDate

`func (o *Transport) SetPlannedDepartureDate(v string)`

SetPlannedDepartureDate sets PlannedDepartureDate field to given value.

### HasPlannedDepartureDate

`func (o *Transport) HasPlannedDepartureDate() bool`

HasPlannedDepartureDate returns a boolean if a field has been set.

### GetPlannedArrivalDate

`func (o *Transport) GetPlannedArrivalDate() string`

GetPlannedArrivalDate returns the PlannedArrivalDate field if non-nil, zero value otherwise.

### GetPlannedArrivalDateOk

`func (o *Transport) GetPlannedArrivalDateOk() (*string, bool)`

GetPlannedArrivalDateOk returns a tuple with the PlannedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedArrivalDate

`func (o *Transport) SetPlannedArrivalDate(v string)`

SetPlannedArrivalDate sets PlannedArrivalDate field to given value.

### HasPlannedArrivalDate

`func (o *Transport) HasPlannedArrivalDate() bool`

HasPlannedArrivalDate returns a boolean if a field has been set.

### GetModeOfTransport

`func (o *Transport) GetModeOfTransport() ModeOfTransport`

GetModeOfTransport returns the ModeOfTransport field if non-nil, zero value otherwise.

### GetModeOfTransportOk

`func (o *Transport) GetModeOfTransportOk() (*ModeOfTransport, bool)`

GetModeOfTransportOk returns a tuple with the ModeOfTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeOfTransport

`func (o *Transport) SetModeOfTransport(v ModeOfTransport)`

SetModeOfTransport sets ModeOfTransport field to given value.

### HasModeOfTransport

`func (o *Transport) HasModeOfTransport() bool`

HasModeOfTransport returns a boolean if a field has been set.

### GetVesselIMONumber

`func (o *Transport) GetVesselIMONumber() string`

GetVesselIMONumber returns the VesselIMONumber field if non-nil, zero value otherwise.

### GetVesselIMONumberOk

`func (o *Transport) GetVesselIMONumberOk() (*string, bool)`

GetVesselIMONumberOk returns a tuple with the VesselIMONumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselIMONumber

`func (o *Transport) SetVesselIMONumber(v string)`

SetVesselIMONumber sets VesselIMONumber field to given value.

### HasVesselIMONumber

`func (o *Transport) HasVesselIMONumber() bool`

HasVesselIMONumber returns a boolean if a field has been set.

### GetCarrierVoyageNumber

`func (o *Transport) GetCarrierVoyageNumber() string`

GetCarrierVoyageNumber returns the CarrierVoyageNumber field if non-nil, zero value otherwise.

### GetCarrierVoyageNumberOk

`func (o *Transport) GetCarrierVoyageNumberOk() (*string, bool)`

GetCarrierVoyageNumberOk returns a tuple with the CarrierVoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierVoyageNumber

`func (o *Transport) SetCarrierVoyageNumber(v string)`

SetCarrierVoyageNumber sets CarrierVoyageNumber field to given value.

### HasCarrierVoyageNumber

`func (o *Transport) HasCarrierVoyageNumber() bool`

HasCarrierVoyageNumber returns a boolean if a field has been set.

### GetIsUnderShippersResponsibility

`func (o *Transport) GetIsUnderShippersResponsibility() bool`

GetIsUnderShippersResponsibility returns the IsUnderShippersResponsibility field if non-nil, zero value otherwise.

### GetIsUnderShippersResponsibilityOk

`func (o *Transport) GetIsUnderShippersResponsibilityOk() (*bool, bool)`

GetIsUnderShippersResponsibilityOk returns a tuple with the IsUnderShippersResponsibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnderShippersResponsibility

`func (o *Transport) SetIsUnderShippersResponsibility(v bool)`

SetIsUnderShippersResponsibility sets IsUnderShippersResponsibility field to given value.

### HasIsUnderShippersResponsibility

`func (o *Transport) HasIsUnderShippersResponsibility() bool`

HasIsUnderShippersResponsibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


