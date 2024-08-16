# RateCards
(*RateCards*)

### Available Operations

* [GetRateSchedule](#getrateschedule) - Get a rate schedule
* [GetRates](#getrates) - Get rates
* [Get](#get) - Get a rate card
* [List](#list) - List rate cards
* [Create](#create) - Create a rate card
* [Update](#update) - Update a rate card
* [AddRate](#addrate) - Add a rate
* [AddRates](#addrates) - Add rates
* [SetProductsOrder](#setproductsorder) - Set the rate card products order
* [MoveProducts](#moveproducts) - Update the rate card products order

## GetRateSchedule

Get a specific rate schedule including all rate card entries


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"github.com/speakeasy-sdks/metronome-go-sdk/types"
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
    res, err := s.RateCards.GetRateSchedule(ctx, nil, nil, requestBody)
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
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"github.com/speakeasy-sdks/metronome-go-sdk/types"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var limit *int64 = metronomegosdk.Int64(10)

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
    res, err := s.RateCards.GetRates(ctx, limit, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `limit`                                                                           | **int64*                                                                          | :heavy_minus_sign:                                                                | Max number of results that should be returned                                     | 10                                                                                |
| `nextPage`                                                                        | **string*                                                                         | :heavy_minus_sign:                                                                | Cursor that indicates where the next page of results should start.                |                                                                                   |
| `requestBody`                                                                     | [*operations.GetRatesRequestBody](../../models/operations/getratesrequestbody.md) | :heavy_minus_sign:                                                                | Rate schedule filter options.                                                     |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |


### Response

**[*operations.GetRatesResponse](../../models/operations/getratesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Get a specific rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
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
    res, err := s.RateCards.Get(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.NotFound | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List rate cards


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var requestBody *operations.ListRateCardsRequestBody = &operations.ListRateCardsRequestBody{}
    ctx := context.Background()
    res, err := s.RateCards.List(ctx, nil, nil, requestBody)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
                for {
            // handle items
        
            res, err = res.Next()
        
            if err != nil {
                // handle error
            }
        
            if res == nil {
                break
            }
        }
        
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

## Create

Create a new rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
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
    res, err := s.RateCards.Create(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## Update

Update a rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
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
    res, err := s.RateCards.Update(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## AddRate

Add a new rate


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"github.com/speakeasy-sdks/metronome-go-sdk/types"
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
    res, err := s.RateCards.AddRate(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## AddRates

Add new rates


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"github.com/speakeasy-sdks/metronome-go-sdk/types"
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
    res, err := s.RateCards.AddRates(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## SetProductsOrder

Sets the ordering of products within a rate card


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
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
    res, err := s.RateCards.SetProductsOrder(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## MoveProducts

Updates ordering of specified products


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
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
    res, err := s.RateCards.MoveProducts(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
