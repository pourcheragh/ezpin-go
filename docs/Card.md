# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | Pointer to **string** | gift card number | [optional] 
**PinCode** | Pointer to **string** | gift card pin | [optional] 
**ClaimUrl** | Pointer to **string** | link of gift card | [optional] 

## Methods

### NewCard

`func NewCard() *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *Card) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *Card) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *Card) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *Card) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetPinCode

`func (o *Card) GetPinCode() string`

GetPinCode returns the PinCode field if non-nil, zero value otherwise.

### GetPinCodeOk

`func (o *Card) GetPinCodeOk() (*string, bool)`

GetPinCodeOk returns a tuple with the PinCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinCode

`func (o *Card) SetPinCode(v string)`

SetPinCode sets PinCode field to given value.

### HasPinCode

`func (o *Card) HasPinCode() bool`

HasPinCode returns a boolean if a field has been set.

### GetClaimUrl

`func (o *Card) GetClaimUrl() string`

GetClaimUrl returns the ClaimUrl field if non-nil, zero value otherwise.

### GetClaimUrlOk

`func (o *Card) GetClaimUrlOk() (*string, bool)`

GetClaimUrlOk returns a tuple with the ClaimUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimUrl

`func (o *Card) SetClaimUrl(v string)`

SetClaimUrl sets ClaimUrl field to given value.

### HasClaimUrl

`func (o *Card) HasClaimUrl() bool`

HasClaimUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


