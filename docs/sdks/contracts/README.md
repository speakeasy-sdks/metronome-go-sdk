# Contracts
(*Contracts*)

### Available Operations

* [Get](#get) - Get a contract
* [List](#list) - List customer contracts
* [Create](#create) - Create a contract
* [Amend](#amend) - Amend a contract
* [Archive](#archive) - Archive a contract
* [SetUsageFilter](#setusagefilter) - Set a contract usage filter
* [AddManualBalanceEntry](#addmanualbalanceentry) - Add a manual balance entry
* [UpdateEndDate](#updateenddate) - Update the contract end date
* [GetRateSchedule](#getrateschedule) - Get the rate schedule for a contract
* [ScheduleProServicesInvoice](#scheduleproservicesinvoice) - Schedule ProService invoice

## Get

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
    res, err := s.Contracts.Get(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.NotFound | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

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
    res, err := s.Contracts.List(ctx, request)
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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.NotFound | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

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
    res, err := s.Contracts.Create(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## Amend

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
    res, err := s.Contracts.Amend(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## Archive

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
    res, err := s.Contracts.Archive(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## AddManualBalanceEntry

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
    res, err := s.Contracts.AddManualBalanceEntry(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## UpdateEndDate

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
    res, err := s.Contracts.UpdateEndDate(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## GetRateSchedule

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
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.NotFound | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
