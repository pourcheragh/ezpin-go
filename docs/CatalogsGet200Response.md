# CatalogsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of total catalogs | [optional] 
**Results** | Pointer to [**[]Catalog**](Catalog.md) |  | [optional] 

## Methods

### NewCatalogsGet200Response

`func NewCatalogsGet200Response() *CatalogsGet200Response`

NewCatalogsGet200Response instantiates a new CatalogsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogsGet200ResponseWithDefaults

`func NewCatalogsGet200ResponseWithDefaults() *CatalogsGet200Response`

NewCatalogsGet200ResponseWithDefaults instantiates a new CatalogsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CatalogsGet200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CatalogsGet200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CatalogsGet200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CatalogsGet200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *CatalogsGet200Response) GetResults() []Catalog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CatalogsGet200Response) GetResultsOk() (*[]Catalog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CatalogsGet200Response) SetResults(v []Catalog)`

SetResults sets Results field to given value.

### HasResults

`func (o *CatalogsGet200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


