# Charge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | **string** | Description of the charge type applied. | 
**CurrencyAmount** | **float32** | The monetary value of all freight and other service charges for a transport document, with a maximum of 2-digit decimals. | 
**CurrencyCode** | **string** | The currency for the charge, using a 3-character code (ISO 4217). | 
**PaymentTerm** | **string** | An indicator of whether a charge is prepaid or collect. | 
**CalculationBasis** | **string** | The code specifying the measure unit used for the corresponding unit price for this cost, such as per day, per ton, per square metre. | 
**UnitPrice** | **float32** | The unit price of this charge item in the currency of the charge. | 
**Quantity** | **float32** | The amount of unit for this charge item. | 

## Methods

### NewCharge

`func NewCharge(chargeType string, currencyAmount float32, currencyCode string, paymentTerm string, calculationBasis string, unitPrice float32, quantity float32, ) *Charge`

NewCharge instantiates a new Charge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeWithDefaults

`func NewChargeWithDefaults() *Charge`

NewChargeWithDefaults instantiates a new Charge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *Charge) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *Charge) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *Charge) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.


### GetCurrencyAmount

`func (o *Charge) GetCurrencyAmount() float32`

GetCurrencyAmount returns the CurrencyAmount field if non-nil, zero value otherwise.

### GetCurrencyAmountOk

`func (o *Charge) GetCurrencyAmountOk() (*float32, bool)`

GetCurrencyAmountOk returns a tuple with the CurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAmount

`func (o *Charge) SetCurrencyAmount(v float32)`

SetCurrencyAmount sets CurrencyAmount field to given value.


### GetCurrencyCode

`func (o *Charge) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Charge) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Charge) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPaymentTerm

`func (o *Charge) GetPaymentTerm() string`

GetPaymentTerm returns the PaymentTerm field if non-nil, zero value otherwise.

### GetPaymentTermOk

`func (o *Charge) GetPaymentTermOk() (*string, bool)`

GetPaymentTermOk returns a tuple with the PaymentTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerm

`func (o *Charge) SetPaymentTerm(v string)`

SetPaymentTerm sets PaymentTerm field to given value.


### GetCalculationBasis

`func (o *Charge) GetCalculationBasis() string`

GetCalculationBasis returns the CalculationBasis field if non-nil, zero value otherwise.

### GetCalculationBasisOk

`func (o *Charge) GetCalculationBasisOk() (*string, bool)`

GetCalculationBasisOk returns a tuple with the CalculationBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationBasis

`func (o *Charge) SetCalculationBasis(v string)`

SetCalculationBasis sets CalculationBasis field to given value.


### GetUnitPrice

`func (o *Charge) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Charge) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Charge) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetQuantity

`func (o *Charge) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Charge) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Charge) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


