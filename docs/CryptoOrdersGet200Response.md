# CryptoOrdersGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | Number of total orders | [optional] 
**Next** | Pointer to **string** | next page of orders | [optional] 
**Previous** | Pointer to **string** | previous page of orders | [optional] 
**Results** | Pointer to [**[]CryptoCurrencyOrder**](CryptoCurrencyOrder.md) |  | [optional] 

## Methods

### NewCryptoOrdersGet200Response

`func NewCryptoOrdersGet200Response() *CryptoOrdersGet200Response`

NewCryptoOrdersGet200Response instantiates a new CryptoOrdersGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoOrdersGet200ResponseWithDefaults

`func NewCryptoOrdersGet200ResponseWithDefaults() *CryptoOrdersGet200Response`

NewCryptoOrdersGet200ResponseWithDefaults instantiates a new CryptoOrdersGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CryptoOrdersGet200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CryptoOrdersGet200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CryptoOrdersGet200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CryptoOrdersGet200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *CryptoOrdersGet200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CryptoOrdersGet200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CryptoOrdersGet200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *CryptoOrdersGet200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *CryptoOrdersGet200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *CryptoOrdersGet200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *CryptoOrdersGet200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *CryptoOrdersGet200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *CryptoOrdersGet200Response) GetResults() []CryptoCurrencyOrder`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CryptoOrdersGet200Response) GetResultsOk() (*[]CryptoCurrencyOrder, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CryptoOrdersGet200Response) SetResults(v []CryptoCurrencyOrder)`

SetResults sets Results field to given value.

### HasResults

`func (o *CryptoOrdersGet200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


