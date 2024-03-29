# Go API client for ezpin

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import ezpin "github.com/pourcheragh/ezpin-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `ezpin.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), ezpin.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `ezpin.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), ezpin.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `ezpin.ContextOperationServerIndices` and `ezpin.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), ezpin.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), ezpin.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ezpaypin.com/vendors/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**AuthTokenPost**](docs/DefaultAPI.md#authtokenpost) | **Post** /auth/token/ | 
*DefaultAPI* | [**BalanceGet**](docs/DefaultAPI.md#balanceget) | **Get** /balance/ | 
*DefaultAPI* | [**CardsActivatePost**](docs/DefaultAPI.md#cardsactivatepost) | **Post** /cards/activate/ | 
*DefaultAPI* | [**CardsCheckGet**](docs/DefaultAPI.md#cardscheckget) | **Get** /cards/check/ | 
*DefaultAPI* | [**CardsReferenceCodeGet**](docs/DefaultAPI.md#cardsreferencecodeget) | **Get** /cards/{reference_code}/ | 
*DefaultAPI* | [**CatalogsGet**](docs/DefaultAPI.md#catalogsget) | **Get** /catalogs/ | 
*DefaultAPI* | [**CatalogsProductSkuAvailabilityGet**](docs/DefaultAPI.md#catalogsproductskuavailabilityget) | **Get** /catalogs/{product_sku}/availability/ | 
*DefaultAPI* | [**CryptoCatalogGet**](docs/DefaultAPI.md#cryptocatalogget) | **Get** /crypto/catalog/ | 
*DefaultAPI* | [**CryptoOrdersGet**](docs/DefaultAPI.md#cryptoordersget) | **Get** /crypto/orders/ | 
*DefaultAPI* | [**CryptoOrdersPost**](docs/DefaultAPI.md#cryptoorderspost) | **Post** /crypto/orders/ | 
*DefaultAPI* | [**CryptoOrdersReferenceCodeFinalizePost**](docs/DefaultAPI.md#cryptoordersreferencecodefinalizepost) | **Post** /crypto/orders/{reference_code}/finalize/ | 
*DefaultAPI* | [**CryptoOrdersReferenceCodeGet**](docs/DefaultAPI.md#cryptoordersreferencecodeget) | **Get** /crypto/orders/{reference_code}/ | 
*DefaultAPI* | [**ExchangeRatesGet**](docs/DefaultAPI.md#exchangeratesget) | **Get** /exchange_rates/ | 
*DefaultAPI* | [**NotificationConfigGet**](docs/DefaultAPI.md#notificationconfigget) | **Get** /notification_config/ | 
*DefaultAPI* | [**NotificationConfigPost**](docs/DefaultAPI.md#notificationconfigpost) | **Post** /notification_config/ | 
*DefaultAPI* | [**OrdersGet**](docs/DefaultAPI.md#ordersget) | **Get** /orders/ | 
*DefaultAPI* | [**OrdersPost**](docs/DefaultAPI.md#orderspost) | **Post** /orders/ | 
*DefaultAPI* | [**OrdersReferenceCodeCardsGet**](docs/DefaultAPI.md#ordersreferencecodecardsget) | **Get** /orders/{reference_code}/cards/ | 
*DefaultAPI* | [**OrdersReferenceCodeGet**](docs/DefaultAPI.md#ordersreferencecodeget) | **Get** /orders/{reference_code}/ | 


## Documentation For Models

 - [AuthTokenPost200Response](docs/AuthTokenPost200Response.md)
 - [AuthTokenPost400Response](docs/AuthTokenPost400Response.md)
 - [AuthTokenPostRequest](docs/AuthTokenPostRequest.md)
 - [Balance](docs/Balance.md)
 - [Card](docs/Card.md)
 - [CardsActivatePostRequest](docs/CardsActivatePostRequest.md)
 - [CardsCheckGet200Response](docs/CardsCheckGet200Response.md)
 - [Catalog](docs/Catalog.md)
 - [CatalogShowingPrice](docs/CatalogShowingPrice.md)
 - [CatalogsGet200Response](docs/CatalogsGet200Response.md)
 - [CatalogsProductSkuAvailabilityGet200Response](docs/CatalogsProductSkuAvailabilityGet200Response.md)
 - [Category](docs/Category.md)
 - [CryptoCatalogGet200Response](docs/CryptoCatalogGet200Response.md)
 - [CryptoCurrencyCatalog](docs/CryptoCurrencyCatalog.md)
 - [CryptoCurrencyOrder](docs/CryptoCurrencyOrder.md)
 - [CryptoCurrencyOrderCryptoCurrencyData](docs/CryptoCurrencyOrderCryptoCurrencyData.md)
 - [CryptoCurrencyOrderCurrencyDaya](docs/CryptoCurrencyOrderCurrencyDaya.md)
 - [CryptoOrdersGet200Response](docs/CryptoOrdersGet200Response.md)
 - [CryptoOrdersPostRequest](docs/CryptoOrdersPostRequest.md)
 - [CryptoOrdersReferenceCodeFinalizePost200Response](docs/CryptoOrdersReferenceCodeFinalizePost200Response.md)
 - [CryptoOrdersReferenceCodeFinalizePostRequest](docs/CryptoOrdersReferenceCodeFinalizePostRequest.md)
 - [Currency](docs/Currency.md)
 - [ExchangeRatesGet200Response](docs/ExchangeRatesGet200Response.md)
 - [ExchangeRatesGet200ResponseResultsInner](docs/ExchangeRatesGet200ResponseResultsInner.md)
 - [ExchangeRatesGet200ResponseResultsInnerFromCurrency](docs/ExchangeRatesGet200ResponseResultsInnerFromCurrency.md)
 - [ExchangeRatesGet200ResponseResultsInnerToCurrency](docs/ExchangeRatesGet200ResponseResultsInnerToCurrency.md)
 - [Notification](docs/Notification.md)
 - [NotificationConfigPost200Response](docs/NotificationConfigPost200Response.md)
 - [Order](docs/Order.md)
 - [OrdersGet200Response](docs/OrdersGet200Response.md)
 - [OrdersPost400Response](docs/OrdersPost400Response.md)
 - [OrdersPost406Response](docs/OrdersPost406Response.md)
 - [OrdersPostRequest](docs/OrdersPostRequest.md)
 - [OrdersReferenceCodeCardsGet200Response](docs/OrdersReferenceCodeCardsGet200Response.md)
 - [Product](docs/Product.md)
 - [Regions](docs/Regions.md)
 - [ShowingCurrency](docs/ShowingCurrency.md)
 - [Transaction](docs/Transaction.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), ezpin.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `Ptrdecimal.Decimal`
* `PtrString`
* `PtrTime`

## Author



