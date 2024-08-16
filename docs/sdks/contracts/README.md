# Contracts
(*Contracts*)

## Overview

Contracts provide an alternative to plans for provisioning and invoicing customers. Use these endpoints to create and update contracts data.

### Available Operations

* [SetCustomerBillableStatus](#setcustomerbillablestatus) - Set customer billable status
* [GetProduct](#getproduct) - Get a product
* [ListProducts](#listproducts) - List products
* [CreateProduct](#createproduct) - Create a product
* [UpdateProduct](#updateproduct) - Update a product
* [ArchiveProductListItem](#archiveproductlistitem) - Archive a product
* [GetRateSchedule](#getrateschedule) - Get a rate schedule
* [GetRates](#getrates) - Get rates
* [GetRateCard](#getratecard) - Get a rate card
* [ListRateCards](#listratecards) - List rate cards
* [CreateRateCard](#createratecard) - Create a rate card
* [UpdateRateCard](#updateratecard) - Update a rate card
* [AddRate](#addrate) - Add a rate
* [AddRates](#addrates) - Add rates
* [SetRateCardProductsOrder](#setratecardproductsorder) - Set the rate card products order
* [MoveRateCardProducts](#moveratecardproducts) - Update the rate card products order
* [GetContract](#getcontract) - Get a contract
* [ListContracts](#listcontracts) - List customer contracts
* [CreateContract](#createcontract) - Create a contract
* [AmendContract](#amendcontract) - Amend a contract
* [ArchiveContract](#archivecontract) - Archive a contract
* [SetUsageFilter](#setusagefilter) - Set a contract usage filter
* [AddManualBalanceLedgerEntry](#addmanualbalanceledgerentry) - Add a manual balance entry
* [UpdateContractEndDate](#updatecontractenddate) - Update the contract end date
* [GetContractRateSchedule](#getcontractrateschedule) - Get the rate schedule for a contract
* [ListCustomerCommits](#listcustomercommits) - List commits
* [CreateCustomerCommit](#createcustomercommit) - Create a commit
* [UpdateCommitEndDate](#updatecommitenddate) - Update the commit end date
* [ListCustomerCredits](#listcustomercredits) - List credits
* [CreateCustomerCredit](#createcustomercredit) - Create a credit
* [UpdateCreditEndDate](#updatecreditenddate) - Update the credit end date
* [ListCustomerBalances](#listcustomerbalances) - List balances
* [ScheduleProServicesInvoice](#scheduleproservicesinvoice) - Schedule ProService invoice

## SetCustomerBillableStatus

Set a customer's billable status. This endpoint's availability is dependent on your client's configuration.

### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.SetCustomerBillableStatusRequestBody](../../models/operations/setcustomerbillablestatusrequestbody.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.SetCustomerBillableStatusResponse](../../models/operations/setcustomerbillablestatusresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.SetCustomerBillableStatusResponseBody          | 400                                                      | application/json                                         |
| sdkerrors.SetCustomerBillableStatusContractsResponseBody | 404                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 4xx-5xx                                                  | */*                                                      |

## GetProduct

Get a specific product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.GetProductRequestBody = &operations.GetProductRequestBody{
        ID: "d84e7f4e-7a70-4fe4-be02-7a5027beffcc",
    }
    ctx := context.Background()
    res, err := s.Contracts.GetProduct(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetProductRequestBody](../../models/operations/getproductrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.GetProductResponse](../../models/operations/getproductresponse.md), error**
| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.GetProductResponseBody | 404                              | application/json                 |
| sdkerrors.SDKError               | 4xx-5xx                          | */*                              |

## ListProducts

List products


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var requestBody *operations.ListProductsRequestBody = &operations.ListProductsRequestBody{
        ArchiveFilter: operations.ArchiveFilterNotArchived.ToPointer(),
    }
    ctx := context.Background()
    res, err := s.Contracts.ListProducts(ctx, nil, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `limit`                                                                                   | **int64*                                                                                  | :heavy_minus_sign:                                                                        | Max number of results that should be returned                                             |
| `nextPage`                                                                                | **string*                                                                                 | :heavy_minus_sign:                                                                        | Cursor that indicates where the next page of results should start.                        |
| `requestBody`                                                                             | [*operations.ListProductsRequestBody](../../models/operations/listproductsrequestbody.md) | :heavy_minus_sign:                                                                        | Get list of products                                                                      |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |


### Response

**[*operations.ListProductsResponse](../../models/operations/listproductsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateProduct

Create a new product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.CreateProductRequestBody = &operations.CreateProductRequestBody{
        Name: "My Product",
        Type: operations.TypeUsageUpper,
        BillableMetricID: metronomegosdk.String("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    }
    ctx := context.Background()
    res, err := s.Contracts.CreateProduct(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateProductRequestBody](../../models/operations/createproductrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.CreateProductResponse](../../models/operations/createproductresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.CreateProductResponseBody | 400                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## UpdateProduct

Update a product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.UpdateProductRequestBody = &operations.UpdateProductRequestBody{
        ProductID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        Name: metronomegosdk.String("My Updated Product"),
        StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.UpdateProduct(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateProductRequestBody](../../models/operations/updateproductrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.UpdateProductResponse](../../models/operations/updateproductresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.UpdateProductResponseBody          | 400                                          | application/json                             |
| sdkerrors.UpdateProductContractsResponseBody | 404                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |

## ArchiveProductListItem

Archive a product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ArchiveProductListItemRequestBody = &operations.ArchiveProductListItemRequestBody{
        ProductID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    }
    ctx := context.Background()
    res, err := s.Contracts.ArchiveProductListItem(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ArchiveProductListItemRequestBody](../../models/operations/archiveproductlistitemrequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.ArchiveProductListItemResponse](../../models/operations/archiveproductlistitemresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.ArchiveProductListItemResponseBody          | 400                                                   | application/json                                      |
| sdkerrors.ArchiveProductListItemContractsResponseBody | 404                                                   | application/json                                      |
| sdkerrors.SDKError                                    | 4xx-5xx                                               | */*                                                   |

## GetRateSchedule

Get a specific rate schedule including all rate card entries


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var requestBody *operations.GetRateScheduleRequestBody = &operations.GetRateScheduleRequestBody{
        RateCardID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
        StartingAt: types.MustTimeFromString("2024-01-01T00:00:00.000Z"),
        Selectors: []operations.Selectors{
            operations.Selectors{
                ProductID: metronomegosdk.String("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
                PartialPricingGroupValues: map[string]string{
                    "region": "us-west-2",
                    "cloud": "aws",
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.GetRateSchedule(ctx, nil, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `limit`                                                                                         | **int64*                                                                                        | :heavy_minus_sign:                                                                              | Max number of results that should be returned                                                   |
| `nextPage`                                                                                      | **string*                                                                                       | :heavy_minus_sign:                                                                              | Cursor that indicates where the next page of results should start.                              |
| `requestBody`                                                                                   | [*operations.GetRateScheduleRequestBody](../../models/operations/getrateschedulerequestbody.md) | :heavy_minus_sign:                                                                              | Rate schedule filter options.                                                                   |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |


### Response

**[*operations.GetRateScheduleResponse](../../models/operations/getratescheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRates

Get rate card rates for a specific time.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var requestBody *operations.GetRatesRequestBody = &operations.GetRatesRequestBody{
        RateCardID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
        At: types.MustTimeFromString("2024-01-01T00:00:00.000Z"),
        Selectors: []operations.GetRatesSelectors{
            operations.GetRatesSelectors{
                ProductID: metronomegosdk.String("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
                PartialPricingGroupValues: map[string]string{
                    "region": "us-west-2",
                    "cloud": "aws",
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.GetRates(ctx, nil, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `limit`                                                                           | **int64*                                                                          | :heavy_minus_sign:                                                                | Max number of results that should be returned                                     |
| `nextPage`                                                                        | **string*                                                                         | :heavy_minus_sign:                                                                | Cursor that indicates where the next page of results should start.                |
| `requestBody`                                                                     | [*operations.GetRatesRequestBody](../../models/operations/getratesrequestbody.md) | :heavy_minus_sign:                                                                | Rate schedule filter options.                                                     |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |


### Response

**[*operations.GetRatesResponse](../../models/operations/getratesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRateCard

Get a specific rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.GetRateCardRequestBody = &operations.GetRateCardRequestBody{
        ID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
    }
    ctx := context.Background()
    res, err := s.Contracts.GetRateCard(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetRateCardRequestBody](../../models/operations/getratecardrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetRateCardResponse](../../models/operations/getratecardresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetRateCardResponseBody | 404                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## ListRateCards

List rate cards


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var requestBody *operations.ListRateCardsRequestBody = &operations.ListRateCardsRequestBody{}
    ctx := context.Background()
    res, err := s.Contracts.ListRateCards(ctx, nil, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `limit`                                                                                     | **int64*                                                                                    | :heavy_minus_sign:                                                                          | Max number of results that should be returned                                               |
| `nextPage`                                                                                  | **string*                                                                                   | :heavy_minus_sign:                                                                          | Cursor that indicates where the next page of results should start.                          |
| `requestBody`                                                                               | [*operations.ListRateCardsRequestBody](../../models/operations/listratecardsrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |


### Response

**[*operations.ListRateCardsResponse](../../models/operations/listratecardsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateRateCard

Create a new rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.CreateRateCardRequestBody = &operations.CreateRateCardRequestBody{
        Name: "My Rate Card",
        Description: metronomegosdk.String("My Rate Card Description"),
        FiatCreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
        CreditTypeConversions: []operations.CreditTypeConversions{
            operations.CreditTypeConversions{
                CustomCreditTypeID: "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
                FiatPerCustomCredit: 2,
            },
        },
        Aliases: []operations.Aliases{
            operations.Aliases{
                Name: "my-rate-card",
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.CreateRateCard(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateRateCardRequestBody](../../models/operations/createratecardrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.CreateRateCardResponse](../../models/operations/createratecardresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.CreateRateCardResponseBody | 400                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## UpdateRateCard

Update a rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.UpdateRateCardRequestBody = &operations.UpdateRateCardRequestBody{
        RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        Name: metronomegosdk.String("My Updated Rate Card"),
        Description: metronomegosdk.String("My Updated Rate Card Description"),
    }
    ctx := context.Background()
    res, err := s.Contracts.UpdateRateCard(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateRateCardRequestBody](../../models/operations/updateratecardrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.UpdateRateCardResponse](../../models/operations/updateratecardresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.UpdateRateCardResponseBody          | 400                                           | application/json                              |
| sdkerrors.UpdateRateCardContractsResponseBody | 404                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## AddRate

Add a new rate


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.AddRateRequestBody = &operations.AddRateRequestBody{
        RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
        Entitled: true,
        RateType: operations.RateTypeFlatUpper,
        Price: metronomegosdk.Float64(100),
        CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
    }
    ctx := context.Background()
    res, err := s.Contracts.AddRate(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.AddRateRequestBody](../../models/operations/addraterequestbody.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |


### Response

**[*operations.AddRateResponse](../../models/operations/addrateresponse.md), error**
| Error Object                           | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| sdkerrors.AddRateResponseBody          | 400                                    | application/json                       |
| sdkerrors.AddRateContractsResponseBody | 404                                    | application/json                       |
| sdkerrors.SDKError                     | 4xx-5xx                                | */*                                    |

## AddRates

Add new rates


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.AddRatesRequestBody = &operations.AddRatesRequestBody{
        RateCardID: metronomegosdk.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
        Rates: []operations.Rates{
            operations.Rates{
                ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
                PricingGroupValues: map[string]string{
                    "region": "us-west-2",
                    "cloud": "aws",
                },
                StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
                Entitled: true,
                RateType: operations.AddRatesRateTypeFlatUpper,
                Price: metronomegosdk.Float64(100),
                CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
            },
            operations.Rates{
                ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
                PricingGroupValues: map[string]string{
                    "region": "us-east-2",
                    "cloud": "aws",
                },
                StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
                Entitled: true,
                RateType: operations.AddRatesRateTypeFlatUpper,
                Price: metronomegosdk.Float64(120),
                CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.AddRates(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.AddRatesRequestBody](../../models/operations/addratesrequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.AddRatesResponse](../../models/operations/addratesresponse.md), error**
| Error Object                            | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| sdkerrors.AddRatesResponseBody          | 400                                     | application/json                        |
| sdkerrors.AddRatesContractsResponseBody | 404                                     | application/json                        |
| sdkerrors.SDKError                      | 4xx-5xx                                 | */*                                     |

## SetRateCardProductsOrder

Sets the ordering of products within a rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.SetRateCardProductsOrderRequestBody = &operations.SetRateCardProductsOrderRequestBody{
        RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ProductOrder: []string{
            "13117714-3f05-48e5-a6e9-a66093f13b4d",
            "b086f2f4-9851-4466-9ca0-30d53e6a42ac",
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.SetRateCardProductsOrder(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.SetRateCardProductsOrderRequestBody](../../models/operations/setratecardproductsorderrequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.SetRateCardProductsOrderResponse](../../models/operations/setratecardproductsorderresponse.md), error**
| Error Object                                            | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| sdkerrors.SetRateCardProductsOrderResponseBody          | 400                                                     | application/json                                        |
| sdkerrors.SetRateCardProductsOrderContractsResponseBody | 404                                                     | application/json                                        |
| sdkerrors.SDKError                                      | 4xx-5xx                                                 | */*                                                     |

## MoveRateCardProducts

Updates ordering of specified products


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.MoveRateCardProductsRequestBody = &operations.MoveRateCardProductsRequestBody{
        RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ProductMoves: []operations.ProductMoves{
            operations.ProductMoves{
                ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
                Position: 0,
            },
            operations.ProductMoves{
                ProductID: "b086f2f4-9851-4466-9ca0-30d53e6a42ac",
                Position: 1,
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.MoveRateCardProducts(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.MoveRateCardProductsRequestBody](../../models/operations/moveratecardproductsrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.MoveRateCardProductsResponse](../../models/operations/moveratecardproductsresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.MoveRateCardProductsResponseBody          | 400                                                 | application/json                                    |
| sdkerrors.MoveRateCardProductsContractsResponseBody | 404                                                 | application/json                                    |
| sdkerrors.SDKError                                  | 4xx-5xx                                             | */*                                                 |

## GetContract

Get a specific contract


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.GetContractRequestBody = &operations.GetContractRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    }
    ctx := context.Background()
    res, err := s.Contracts.GetContract(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetContractRequestBody](../../models/operations/getcontractrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |


### Response

**[*operations.GetContractResponse](../../models/operations/getcontractresponse.md), error**
| Error Object                      | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| sdkerrors.GetContractResponseBody | 404                               | application/json                  |
| sdkerrors.SDKError                | 4xx-5xx                           | */*                               |

## ListContracts

List all contracts for a customer

### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ListContractsRequestBody = &operations.ListContractsRequestBody{
        CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
    }
    ctx := context.Background()
    res, err := s.Contracts.ListContracts(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListContractsRequestBody](../../models/operations/listcontractsrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.ListContractsResponse](../../models/operations/listcontractsresponse.md), error**
| Error Object                        | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.ListContractsResponseBody | 404                                 | application/json                    |
| sdkerrors.SDKError                  | 4xx-5xx                             | */*                                 |

## CreateContract

Create a new contract


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.CreateContractRequestBody = &operations.CreateContractRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        RateCardID: metronomegosdk.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
        StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.CreateContract(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateContractRequestBody](../../models/operations/createcontractrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.CreateContractResponse](../../models/operations/createcontractresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.CreateContractResponseBody          | 400                                           | application/json                              |
| sdkerrors.CreateContractContractsResponseBody | 404                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## AmendContract

Amend a contract


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.AmendContractRequestBody = &operations.AmendContractRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.AmendContract(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.AmendContractRequestBody](../../models/operations/amendcontractrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.AmendContractResponse](../../models/operations/amendcontractresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.AmendContractResponseBody          | 400                                          | application/json                             |
| sdkerrors.AmendContractContractsResponseBody | 404                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |

## ArchiveContract

Archive a contract

### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ArchiveContractRequestBody = &operations.ArchiveContractRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        VoidInvoices: true,
    }
    ctx := context.Background()
    res, err := s.Contracts.ArchiveContract(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ArchiveContractRequestBody](../../models/operations/archivecontractrequestbody.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |


### Response

**[*operations.ArchiveContractResponse](../../models/operations/archivecontractresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.ArchiveContractResponseBody          | 400                                            | application/json                               |
| sdkerrors.ArchiveContractContractsResponseBody | 404                                            | application/json                               |
| sdkerrors.SDKError                             | 4xx-5xx                                        | */*                                            |

## SetUsageFilter

Set usage filter for a contract


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.SetUsageFilterRequestBody = &operations.SetUsageFilterRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        GroupKey: "business_subscription_id",
        GroupValues: []string{
            "ID-1",
            "ID-2",
        },
        StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.SetUsageFilter(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SetUsageFilterRequestBody](../../models/operations/setusagefilterrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |


### Response

**[*operations.SetUsageFilterResponse](../../models/operations/setusagefilterresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.SetUsageFilterResponseBody          | 400                                           | application/json                              |
| sdkerrors.SetUsageFilterContractsResponseBody | 404                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## AddManualBalanceLedgerEntry

Add a manual balance entry


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.AddManualBalanceLedgerEntryRequestBody = &operations.AddManualBalanceLedgerEntryRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: metronomegosdk.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
        ID: "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
        SegmentID: "66368e29-3f97-4d15-a6e9-120897f0070a",
        Amount: -1000,
        Reason: "Reason for entry",
    }
    ctx := context.Background()
    res, err := s.Contracts.AddManualBalanceLedgerEntry(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AddManualBalanceLedgerEntryRequestBody](../../models/operations/addmanualbalanceledgerentryrequestbody.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.AddManualBalanceLedgerEntryResponse](../../models/operations/addmanualbalanceledgerentryresponse.md), error**
| Error Object                                               | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| sdkerrors.AddManualBalanceLedgerEntryResponseBody          | 400                                                        | application/json                                           |
| sdkerrors.AddManualBalanceLedgerEntryContractsResponseBody | 404                                                        | application/json                                           |
| sdkerrors.SDKError                                         | 4xx-5xx                                                    | */*                                                        |

## UpdateContractEndDate

Update the end date of a contract


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.UpdateContractEndDateRequestBody = &operations.UpdateContractEndDateRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        EndingBefore: types.MustNewTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.UpdateContractEndDate(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateContractEndDateRequestBody](../../models/operations/updatecontractenddaterequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |


### Response

**[*operations.UpdateContractEndDateResponse](../../models/operations/updatecontractenddateresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.UpdateContractEndDateResponseBody          | 400                                                  | application/json                                     |
| sdkerrors.UpdateContractEndDateContractsResponseBody | 404                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 4xx-5xx                                              | */*                                                  |

## GetContractRateSchedule

Get the rate schedule for the rate card on a given contract.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var requestBody *operations.GetContractRateScheduleRequestBody = &operations.GetContractRateScheduleRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        At: types.MustNewTimeFromString("2020-01-01T00:00:00.000Z"),
        Selectors: []operations.GetContractRateScheduleSelectors{
            operations.GetContractRateScheduleSelectors{
                ProductID: metronomegosdk.String("d6300dbb-882e-4d2d-8dec-5125d16b65d0"),
                PartialPricingGroupValues: map[string]string{
                    "region": "us-west-2",
                    "cloud": "aws",
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.GetContractRateSchedule(ctx, nil, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `limit`                                                                                                         | **int64*                                                                                                        | :heavy_minus_sign:                                                                                              | Max number of results that should be returned                                                                   |
| `nextPage`                                                                                                      | **string*                                                                                                       | :heavy_minus_sign:                                                                                              | Cursor that indicates where the next page of results should start.                                              |
| `requestBody`                                                                                                   | [*operations.GetContractRateScheduleRequestBody](../../models/operations/getcontractrateschedulerequestbody.md) | :heavy_minus_sign:                                                                                              | Contract rate schedule filter options.                                                                          |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |


### Response

**[*operations.GetContractRateScheduleResponse](../../models/operations/getcontractratescheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCustomerCommits

List commits.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ListCustomerCommitsRequestBody = &operations.ListCustomerCommitsRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CommitID: metronomegosdk.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
        IncludeLedgers: metronomegosdk.Bool(true),
    }
    ctx := context.Background()
    res, err := s.Contracts.ListCustomerCommits(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCustomerCommitsRequestBody](../../models/operations/listcustomercommitsrequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.ListCustomerCommitsResponse](../../models/operations/listcustomercommitsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCustomerCommit

Create a new commit at the customer level.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.CreateCustomerCommitRequestBody = &operations.CreateCustomerCommitRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        Type: operations.CreateCustomerCommitTypePrepaidLower,
        Name: metronomegosdk.String("My Commit"),
        Priority: 100,
        ProductID: "f14d6729-6a44-4b13-9908-9387f1918790",
        AccessSchedule: operations.CreateCustomerCommitAccessSchedule{
            CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
            ScheduleItems: []operations.CreateCustomerCommitScheduleItems{
                operations.CreateCustomerCommitScheduleItems{
                    Amount: 1000,
                    StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
                    EndingBefore: types.MustTimeFromString("2020-02-01T00:00:00.000Z"),
                },
            },
        },
        InvoiceSchedule: &operations.CreateCustomerCommitInvoiceSchedule{
            CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
            ScheduleItems: []operations.CreateCustomerCommitContractsScheduleItems{
                operations.CreateCustomerCommitContractsScheduleItems{
                    UnitPrice: metronomegosdk.Float64(10000000),
                    Quantity: metronomegosdk.Float64(1),
                    Amount: metronomegosdk.Float64(10000000),
                    Timestamp: types.MustTimeFromString("2020-03-01T00:00:00.000Z"),
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.CreateCustomerCommit(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateCustomerCommitRequestBody](../../models/operations/createcustomercommitrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.CreateCustomerCommitResponse](../../models/operations/createcustomercommitresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.CreateCustomerCommitResponseBody          | 400                                                 | application/json                                    |
| sdkerrors.CreateCustomerCommitContractsResponseBody | 404                                                 | application/json                                    |
| sdkerrors.SDKError                                  | 4xx-5xx                                             | */*                                                 |

## UpdateCommitEndDate

Update the end date of a PREPAID commit


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.UpdateCommitEndDateRequestBody = &operations.UpdateCommitEndDateRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CommitID: "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
        AccessEndingBefore: types.MustNewTimeFromString("2020-01-01T00:00:00.000Z"),
        InvoicesEndingBefore: types.MustNewTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.UpdateCommitEndDate(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateCommitEndDateRequestBody](../../models/operations/updatecommitenddaterequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.UpdateCommitEndDateResponse](../../models/operations/updatecommitenddateresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.UpdateCommitEndDateResponseBody          | 400                                                | application/json                                   |
| sdkerrors.UpdateCommitEndDateContractsResponseBody | 404                                                | application/json                                   |
| sdkerrors.SDKError                                 | 4xx-5xx                                            | */*                                                |

## ListCustomerCredits

List credits.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ListCustomerCreditsRequestBody = &operations.ListCustomerCreditsRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CreditID: metronomegosdk.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
        IncludeLedgers: metronomegosdk.Bool(true),
    }
    ctx := context.Background()
    res, err := s.Contracts.ListCustomerCredits(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.ListCustomerCreditsRequestBody](../../models/operations/listcustomercreditsrequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.ListCustomerCreditsResponse](../../models/operations/listcustomercreditsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateCustomerCredit

Create a new credit at the customer level.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.CreateCustomerCreditRequestBody = &operations.CreateCustomerCreditRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        Name: metronomegosdk.String("My Credit"),
        Priority: 100,
        ProductID: "f14d6729-6a44-4b13-9908-9387f1918790",
        AccessSchedule: operations.CreateCustomerCreditAccessSchedule{
            CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
            ScheduleItems: []operations.CreateCustomerCreditScheduleItems{
                operations.CreateCustomerCreditScheduleItems{
                    Amount: 1000,
                    StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
                    EndingBefore: types.MustTimeFromString("2020-02-01T00:00:00.000Z"),
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.Contracts.CreateCustomerCredit(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateCustomerCreditRequestBody](../../models/operations/createcustomercreditrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.CreateCustomerCreditResponse](../../models/operations/createcustomercreditresponse.md), error**
| Error Object                                        | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| sdkerrors.CreateCustomerCreditResponseBody          | 400                                                 | application/json                                    |
| sdkerrors.CreateCustomerCreditContractsResponseBody | 404                                                 | application/json                                    |
| sdkerrors.SDKError                                  | 4xx-5xx                                             | */*                                                 |

## UpdateCreditEndDate

Update the end date of a credit


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.UpdateCreditEndDateRequestBody = &operations.UpdateCreditEndDateRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CreditID: "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
        AccessEndingBefore: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.UpdateCreditEndDate(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateCreditEndDateRequestBody](../../models/operations/updatecreditenddaterequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.UpdateCreditEndDateResponse](../../models/operations/updatecreditenddateresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.UpdateCreditEndDateResponseBody          | 400                                                | application/json                                   |
| sdkerrors.UpdateCreditEndDateContractsResponseBody | 404                                                | application/json                                   |
| sdkerrors.SDKError                                 | 4xx-5xx                                            | */*                                                |

## ListCustomerBalances

List balances (commits and credits).


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ListCustomerBalancesRequestBody = &operations.ListCustomerBalancesRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        ID: metronomegosdk.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
        IncludeLedgers: metronomegosdk.Bool(true),
    }
    ctx := context.Background()
    res, err := s.Contracts.ListCustomerBalances(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListCustomerBalancesRequestBody](../../models/operations/listcustomerbalancesrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |


### Response

**[*operations.ListCustomerBalancesResponse](../../models/operations/listcustomerbalancesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ScheduleProServicesInvoice

Create a new, scheduled invoice for Professional Services terms on a contract. This endpoint's availability is dependent on your client's configuration.


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Contracts.ScheduleProServicesInvoice(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ScheduleProServicesInvoiceRequestBody](../../models/operations/scheduleproservicesinvoicerequestbody.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |


### Response

**[*operations.ScheduleProServicesInvoiceResponse](../../models/operations/scheduleproservicesinvoiceresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.ScheduleProServicesInvoiceResponseBody | 404                                              | application/json                                 |
| sdkerrors.SDKError                               | 4xx-5xx                                          | */*                                              |
