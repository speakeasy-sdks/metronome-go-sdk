# CustomerCredits
(*CustomerCredits*)

### Available Operations

* [List](#list) - List credits
* [Create](#create) - Create a credit
* [UpdateEndDate](#updateenddate) - Update the credit end date

## List

List credits.


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
    var request *operations.ListCustomerCreditsRequestBody = &operations.ListCustomerCreditsRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CreditID: metronomegosdk.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
        IncludeLedgers: metronomegosdk.Bool(true),
    }
    ctx := context.Background()
    res, err := s.CustomerCredits.List(ctx, request)
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

## Create

Create a new credit at the customer level.


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
    res, err := s.CustomerCredits.Create(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## UpdateEndDate

Update the end date of a credit


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
    var request *operations.UpdateCreditEndDateRequestBody = &operations.UpdateCreditEndDateRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CreditID: "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
        AccessEndingBefore: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.CustomerCredits.UpdateEndDate(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
