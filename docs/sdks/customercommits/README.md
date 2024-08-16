# CustomerCommits
(*CustomerCommits*)

### Available Operations

* [List](#list) - List commits
* [Create](#create) - Create a commit
* [UpdateEndDate](#updateenddate) - Update the commit end date

## List

List commits.


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
    var request *operations.ListCustomerCommitsRequestBody = &operations.ListCustomerCommitsRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CommitID: metronomegosdk.String("6162d87b-e5db-4a33-b7f2-76ce6ead4e85"),
        IncludeLedgers: metronomegosdk.Bool(true),
    }
    ctx := context.Background()
    res, err := s.CustomerCommits.List(ctx, request)
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
| `request`                                                                                              | [operations.ListCustomerCommitsRequestBody](../../models/operations/listcustomercommitsrequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.ListCustomerCommitsResponse](../../models/operations/listcustomercommitsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create a new commit at the customer level.


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
            ScheduleItems: []operations.CreateCustomerCommitCustomerCommitsScheduleItems{
                operations.CreateCustomerCommitCustomerCommitsScheduleItems{
                    UnitPrice: metronomegosdk.Float64(10000000),
                    Quantity: metronomegosdk.Float64(1),
                    Amount: metronomegosdk.Float64(10000000),
                    Timestamp: types.MustTimeFromString("2020-03-01T00:00:00.000Z"),
                },
            },
        },
    }
    ctx := context.Background()
    res, err := s.CustomerCommits.Create(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## UpdateEndDate

Update the end date of a PREPAID commit


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
    var request *operations.UpdateCommitEndDateRequestBody = &operations.UpdateCommitEndDateRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CommitID: "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
        AccessEndingBefore: types.MustNewTimeFromString("2020-01-01T00:00:00.000Z"),
        InvoicesEndingBefore: types.MustNewTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.CustomerCommits.UpdateEndDate(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
