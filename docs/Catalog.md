# Catalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | Pointer to **float32** | sku of product | [optional] 
**Upc** | Pointer to **float32** | upc of product | [optional] 
**Title** | Pointer to **string** | title of product | [optional] 
**MinPrice** | Pointer to **float32** | product minimumm price | [optional] 
**MaxPrice** | Pointer to **float32** | product maximumm price | [optional] 
**PreOrder** | Pointer to **bool** | product has preorder feature or not. | [optional] 
**ActivationFee** | Pointer to **float32** | fixed activation fee of product | [optional] 
**PercentageOfBuyingPrice** | Pointer to **float32** | discount applied on buying price | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**Regions** | Pointer to [**[]Regions**](Regions.md) |  | [optional] 
**Image** | Pointer to **string** | image link of product | [optional] 
**Description** | Pointer to **string** | product discription | [optional] 
**ShowingPrice** | Pointer to [**CatalogShowingPrice**](CatalogShowingPrice.md) |  | [optional] 

## Methods

### NewCatalog

`func NewCatalog() *Catalog`

NewCatalog instantiates a new Catalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogWithDefaults

`func NewCatalogWithDefaults() *Catalog`

NewCatalogWithDefaults instantiates a new Catalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *Catalog) GetSku() float32`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Catalog) GetSkuOk() (*float32, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Catalog) SetSku(v float32)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Catalog) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUpc

`func (o *Catalog) GetUpc() float32`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *Catalog) GetUpcOk() (*float32, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *Catalog) SetUpc(v float32)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *Catalog) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetTitle

`func (o *Catalog) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Catalog) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Catalog) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Catalog) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMinPrice

`func (o *Catalog) GetMinPrice() float32`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *Catalog) GetMinPriceOk() (*float32, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *Catalog) SetMinPrice(v float32)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *Catalog) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *Catalog) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *Catalog) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *Catalog) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *Catalog) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetPreOrder

`func (o *Catalog) GetPreOrder() bool`

GetPreOrder returns the PreOrder field if non-nil, zero value otherwise.

### GetPreOrderOk

`func (o *Catalog) GetPreOrderOk() (*bool, bool)`

GetPreOrderOk returns a tuple with the PreOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOrder

`func (o *Catalog) SetPreOrder(v bool)`

SetPreOrder sets PreOrder field to given value.

### HasPreOrder

`func (o *Catalog) HasPreOrder() bool`

HasPreOrder returns a boolean if a field has been set.

### GetActivationFee

`func (o *Catalog) GetActivationFee() float32`

GetActivationFee returns the ActivationFee field if non-nil, zero value otherwise.

### GetActivationFeeOk

`func (o *Catalog) GetActivationFeeOk() (*float32, bool)`

GetActivationFeeOk returns a tuple with the ActivationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationFee

`func (o *Catalog) SetActivationFee(v float32)`

SetActivationFee sets ActivationFee field to given value.

### HasActivationFee

`func (o *Catalog) HasActivationFee() bool`

HasActivationFee returns a boolean if a field has been set.

### GetPercentageOfBuyingPrice

`func (o *Catalog) GetPercentageOfBuyingPrice() float32`

GetPercentageOfBuyingPrice returns the PercentageOfBuyingPrice field if non-nil, zero value otherwise.

### GetPercentageOfBuyingPriceOk

`func (o *Catalog) GetPercentageOfBuyingPriceOk() (*float32, bool)`

GetPercentageOfBuyingPriceOk returns a tuple with the PercentageOfBuyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfBuyingPrice

`func (o *Catalog) SetPercentageOfBuyingPrice(v float32)`

SetPercentageOfBuyingPrice sets PercentageOfBuyingPrice field to given value.

### HasPercentageOfBuyingPrice

`func (o *Catalog) HasPercentageOfBuyingPrice() bool`

HasPercentageOfBuyingPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *Catalog) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Catalog) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Catalog) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Catalog) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCategories

`func (o *Catalog) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Catalog) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Catalog) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Catalog) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetRegions

`func (o *Catalog) GetRegions() []Regions`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Catalog) GetRegionsOk() (*[]Regions, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Catalog) SetRegions(v []Regions)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Catalog) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetImage

`func (o *Catalog) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Catalog) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Catalog) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Catalog) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDescription

`func (o *Catalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Catalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Catalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Catalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShowingPrice

`func (o *Catalog) GetShowingPrice() CatalogShowingPrice`

GetShowingPrice returns the ShowingPrice field if non-nil, zero value otherwise.

### GetShowingPriceOk

`func (o *Catalog) GetShowingPriceOk() (*CatalogShowingPrice, bool)`

GetShowingPriceOk returns a tuple with the ShowingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowingPrice

`func (o *Catalog) SetShowingPrice(v CatalogShowingPrice)`

SetShowingPrice sets ShowingPrice field to given value.

### HasShowingPrice

`func (o *Catalog) HasShowingPrice() bool`

HasShowingPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


