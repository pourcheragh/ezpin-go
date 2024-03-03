# OrdersGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | Number of total orders | [optional] 
**Next** | Pointer to **string** | next page of order list | [optional] 
**Previous** | Pointer to **string** | previous page of order list | [optional] 
**Results** | Pointer to [**[]Order**](Order.md) |  | [optional] 

## Methods

### NewOrdersGet200Response

`func NewOrdersGet200Response() *OrdersGet200Response`

NewOrdersGet200Response instantiates a new OrdersGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersGet200ResponseWithDefaults

`func NewOrdersGet200ResponseWithDefaults() *OrdersGet200Response`

NewOrdersGet200ResponseWithDefaults instantiates a new OrdersGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OrdersGet200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OrdersGet200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OrdersGet200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OrdersGet200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *OrdersGet200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *OrdersGet200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *OrdersGet200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *OrdersGet200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *OrdersGet200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *OrdersGet200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *OrdersGet200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *OrdersGet200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetResults

`func (o *OrdersGet200Response) GetResults() []Order`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OrdersGet200Response) GetResultsOk() (*[]Order, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OrdersGet200Response) SetResults(v []Order)`

SetResults sets Results field to given value.

### HasResults

`func (o *OrdersGet200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


