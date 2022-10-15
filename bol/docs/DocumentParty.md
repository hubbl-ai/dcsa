# DocumentParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Party** | Pointer to [**PartyNoID**](PartyNoID.md) |  | [optional] 
**PartyFunction** | [**PartyFunction**](PartyFunction.md) |  | 
**DisplayedAddress** | Pointer to **[]string** |  | [optional] 
**PartyContactDetails** | **map[string]interface{}** |  | 
**IsToBeNotified** | **bool** | Used to decide whether the party will be notified of the arrival of the cargo. | 

## Methods

### NewDocumentParty

`func NewDocumentParty(partyFunction PartyFunction, partyContactDetails map[string]interface{}, isToBeNotified bool, ) *DocumentParty`

NewDocumentParty instantiates a new DocumentParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentPartyWithDefaults

`func NewDocumentPartyWithDefaults() *DocumentParty`

NewDocumentPartyWithDefaults instantiates a new DocumentParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParty

`func (o *DocumentParty) GetParty() PartyNoID`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *DocumentParty) GetPartyOk() (*PartyNoID, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *DocumentParty) SetParty(v PartyNoID)`

SetParty sets Party field to given value.

### HasParty

`func (o *DocumentParty) HasParty() bool`

HasParty returns a boolean if a field has been set.

### GetPartyFunction

`func (o *DocumentParty) GetPartyFunction() PartyFunction`

GetPartyFunction returns the PartyFunction field if non-nil, zero value otherwise.

### GetPartyFunctionOk

`func (o *DocumentParty) GetPartyFunctionOk() (*PartyFunction, bool)`

GetPartyFunctionOk returns a tuple with the PartyFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyFunction

`func (o *DocumentParty) SetPartyFunction(v PartyFunction)`

SetPartyFunction sets PartyFunction field to given value.


### GetDisplayedAddress

`func (o *DocumentParty) GetDisplayedAddress() []string`

GetDisplayedAddress returns the DisplayedAddress field if non-nil, zero value otherwise.

### GetDisplayedAddressOk

`func (o *DocumentParty) GetDisplayedAddressOk() (*[]string, bool)`

GetDisplayedAddressOk returns a tuple with the DisplayedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedAddress

`func (o *DocumentParty) SetDisplayedAddress(v []string)`

SetDisplayedAddress sets DisplayedAddress field to given value.

### HasDisplayedAddress

`func (o *DocumentParty) HasDisplayedAddress() bool`

HasDisplayedAddress returns a boolean if a field has been set.

### GetPartyContactDetails

`func (o *DocumentParty) GetPartyContactDetails() map[string]interface{}`

GetPartyContactDetails returns the PartyContactDetails field if non-nil, zero value otherwise.

### GetPartyContactDetailsOk

`func (o *DocumentParty) GetPartyContactDetailsOk() (*map[string]interface{}, bool)`

GetPartyContactDetailsOk returns a tuple with the PartyContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyContactDetails

`func (o *DocumentParty) SetPartyContactDetails(v map[string]interface{})`

SetPartyContactDetails sets PartyContactDetails field to given value.


### GetIsToBeNotified

`func (o *DocumentParty) GetIsToBeNotified() bool`

GetIsToBeNotified returns the IsToBeNotified field if non-nil, zero value otherwise.

### GetIsToBeNotifiedOk

`func (o *DocumentParty) GetIsToBeNotifiedOk() (*bool, bool)`

GetIsToBeNotifiedOk returns a tuple with the IsToBeNotified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsToBeNotified

`func (o *DocumentParty) SetIsToBeNotified(v bool)`

SetIsToBeNotified sets IsToBeNotified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


