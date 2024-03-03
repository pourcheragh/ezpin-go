# CatalogShowingPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **decimal.Decimal** | showing price | [optional] 
**ShowingCurrency** | Pointer to [**[]ShowingCurrency**](ShowingCurrency.md) |  | [optional] 

## Methods

### NewCatalogShowingPrice

`func NewCatalogShowingPrice() *CatalogShowingPrice`

NewCatalogShowingPrice instantiates a new CatalogShowingPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogShowingPriceWithDefaults

`func NewCatalogShowingPriceWithDefaults() *CatalogShowingPrice`

NewCatalogShowingPriceWithDefaults instantiates a new CatalogShowingPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *CatalogShowingPrice) GetPrice() decimal.Decimal`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogShowingPrice) GetPriceOk() (*decimal.Decimal, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogShowingPrice) SetPrice(v decimal.Decimal)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogShowingPrice) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetShowingCurrency

`func (o *CatalogShowingPrice) GetShowingCurrency() []ShowingCurrency`

GetShowingCurrency returns the ShowingCurrency field if non-nil, zero value otherwise.

### GetShowingCurrencyOk

`func (o *CatalogShowingPrice) GetShowingCurrencyOk() (*[]ShowingCurrency, bool)`

GetShowingCurrencyOk returns a tuple with the ShowingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowingCurrency

`func (o *CatalogShowingPrice) SetShowingCurrency(v []ShowingCurrency)`

SetShowingCurrency sets ShowingCurrency field to given value.

### HasShowingCurrency

`func (o *CatalogShowingPrice) HasShowingCurrency() bool`

HasShowingCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


