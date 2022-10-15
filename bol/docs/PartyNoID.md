# PartyNoID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartyName** | Pointer to **string** | Name of the party. | [optional] 
**TaxReference1** | Pointer to **string** | The identifying number of the consignee or shipper (Individual or entity) used for tax purposes. | [optional] 
**TaxReference2** | Pointer to **string** | Optional second identifying number of the consignee or shipper (Individual or entity) used for tax purposes. | [optional] 
**PublicKey** | Pointer to **string** | The public key used for a digital signature. | [optional] 
**Address** | Pointer to [**Address**](address.md) |  | [optional] 
**NmftaCode** | Pointer to **string** | The Standard Carrier Alpha Code (SCAC) provided by NMFTA. | [optional] 

## Methods

### NewPartyNoID

`func NewPartyNoID() *PartyNoID`

NewPartyNoID instantiates a new PartyNoID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyNoIDWithDefaults

`func NewPartyNoIDWithDefaults() *PartyNoID`

NewPartyNoIDWithDefaults instantiates a new PartyNoID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartyName

`func (o *PartyNoID) GetPartyName() string`

GetPartyName returns the PartyName field if non-nil, zero value otherwise.

### GetPartyNameOk

`func (o *PartyNoID) GetPartyNameOk() (*string, bool)`

GetPartyNameOk returns a tuple with the PartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyName

`func (o *PartyNoID) SetPartyName(v string)`

SetPartyName sets PartyName field to given value.

### HasPartyName

`func (o *PartyNoID) HasPartyName() bool`

HasPartyName returns a boolean if a field has been set.

### GetTaxReference1

`func (o *PartyNoID) GetTaxReference1() string`

GetTaxReference1 returns the TaxReference1 field if non-nil, zero value otherwise.

### GetTaxReference1Ok

`func (o *PartyNoID) GetTaxReference1Ok() (*string, bool)`

GetTaxReference1Ok returns a tuple with the TaxReference1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxReference1

`func (o *PartyNoID) SetTaxReference1(v string)`

SetTaxReference1 sets TaxReference1 field to given value.

### HasTaxReference1

`func (o *PartyNoID) HasTaxReference1() bool`

HasTaxReference1 returns a boolean if a field has been set.

### GetTaxReference2

`func (o *PartyNoID) GetTaxReference2() string`

GetTaxReference2 returns the TaxReference2 field if non-nil, zero value otherwise.

### GetTaxReference2Ok

`func (o *PartyNoID) GetTaxReference2Ok() (*string, bool)`

GetTaxReference2Ok returns a tuple with the TaxReference2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxReference2

`func (o *PartyNoID) SetTaxReference2(v string)`

SetTaxReference2 sets TaxReference2 field to given value.

### HasTaxReference2

`func (o *PartyNoID) HasTaxReference2() bool`

HasTaxReference2 returns a boolean if a field has been set.

### GetPublicKey

`func (o *PartyNoID) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PartyNoID) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PartyNoID) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PartyNoID) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetAddress

`func (o *PartyNoID) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartyNoID) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartyNoID) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PartyNoID) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNmftaCode

`func (o *PartyNoID) GetNmftaCode() string`

GetNmftaCode returns the NmftaCode field if non-nil, zero value otherwise.

### GetNmftaCodeOk

`func (o *PartyNoID) GetNmftaCodeOk() (*string, bool)`

GetNmftaCodeOk returns a tuple with the NmftaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmftaCode

`func (o *PartyNoID) SetNmftaCode(v string)`

SetNmftaCode sets NmftaCode field to given value.

### HasNmftaCode

`func (o *PartyNoID) HasNmftaCode() bool`

HasNmftaCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


