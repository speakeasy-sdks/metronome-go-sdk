# NamedSchedules
(*NamedSchedules*)

## Overview

Named schedules are used for storing custom data that can change over time. Named schedules are often used in custom pricing logic.

### Available Operations

* [GetCustomerNamedSchedule](#getcustomernamedschedule) - Get a customer's named schedule
* [UpdateCustomerNamedSchedule](#updatecustomernamedschedule) - Update a customer's named schedule
* [GetContractNamedSchedule](#getcontractnamedschedule) - Get a contract's named schedule
* [UpdateContractNamedSchedule](#updatecontractnamedschedule) - Update a contract's named schedule
* [GetRateCardNamedSchedule](#getratecardnamedschedule) - Get a rate card's named schedule
* [UpdateRateCardNamedSchedule](#updateratecardnamedschedule) - Update a rate card's named schedule

## GetCustomerNamedSchedule

Get a named schedule for the given customer. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.GetCustomerNamedScheduleRequestBody = &operations.GetCustomerNamedScheduleRequestBody{
        CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
        ScheduleName: "my-schedule",
        CoveringDate: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
    }
    ctx := context.Background()
    res, err := s.NamedSchedules.GetCustomerNamedSchedule(ctx, request)
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
| `request`                                                                                                        | [operations.GetCustomerNamedScheduleRequestBody](../../models/operations/getcustomernamedschedulerequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.GetCustomerNamedScheduleResponse](../../models/operations/getcustomernamedscheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateCustomerNamedSchedule

Update a named schedule for the given customer. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.UpdateCustomerNamedScheduleRequestBody = &operations.UpdateCustomerNamedScheduleRequestBody{
        CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
        ScheduleName: "my-schedule",
        StartingAt: types.MustTimeFromString("2022-02-01T00:00:00Z"),
        EndingBefore: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
        Value: map[string]any{
        "my_key": "my_value",
    },
    }
    ctx := context.Background()
    res, err := s.NamedSchedules.UpdateCustomerNamedSchedule(ctx, request)
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
| `request`                                                                                                              | [operations.UpdateCustomerNamedScheduleRequestBody](../../models/operations/updatecustomernamedschedulerequestbody.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.UpdateCustomerNamedScheduleResponse](../../models/operations/updatecustomernamedscheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetContractNamedSchedule

Get a named schedule for the given contract. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.GetContractNamedScheduleRequestBody = &operations.GetContractNamedScheduleRequestBody{
        CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ScheduleName: "my-schedule",
        CoveringDate: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
    }
    ctx := context.Background()
    res, err := s.NamedSchedules.GetContractNamedSchedule(ctx, request)
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
| `request`                                                                                                        | [operations.GetContractNamedScheduleRequestBody](../../models/operations/getcontractnamedschedulerequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.GetContractNamedScheduleResponse](../../models/operations/getcontractnamedscheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateContractNamedSchedule

Update a named schedule for the given contract. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.UpdateContractNamedScheduleRequestBody = &operations.UpdateContractNamedScheduleRequestBody{
        CustomerID: "9b85c1c1-5238-4f2a-a409-61412905e1e1",
        ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ScheduleName: "my-schedule",
        StartingAt: types.MustTimeFromString("2022-02-01T00:00:00Z"),
        EndingBefore: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
        Value: map[string]any{
        "my_key": "my_value",
    },
    }
    ctx := context.Background()
    res, err := s.NamedSchedules.UpdateContractNamedSchedule(ctx, request)
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
| `request`                                                                                                              | [operations.UpdateContractNamedScheduleRequestBody](../../models/operations/updatecontractnamedschedulerequestbody.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.UpdateContractNamedScheduleResponse](../../models/operations/updatecontractnamedscheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRateCardNamedSchedule

Get a named schedule for the given rate card. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.GetRateCardNamedScheduleRequestBody = &operations.GetRateCardNamedScheduleRequestBody{
        RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ScheduleName: "my-schedule",
        CoveringDate: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
    }
    ctx := context.Background()
    res, err := s.NamedSchedules.GetRateCardNamedSchedule(ctx, request)
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
| `request`                                                                                                        | [operations.GetRateCardNamedScheduleRequestBody](../../models/operations/getratecardnamedschedulerequestbody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |


### Response

**[*operations.GetRateCardNamedScheduleResponse](../../models/operations/getratecardnamedscheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateRateCardNamedSchedule

Update a named schedule for the given rate card. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.UpdateRateCardNamedScheduleRequestBody = &operations.UpdateRateCardNamedScheduleRequestBody{
        RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        ScheduleName: "my-schedule",
        StartingAt: types.MustTimeFromString("2022-02-01T00:00:00Z"),
        EndingBefore: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
        Value: map[string]any{
        "my_key": "my_value",
    },
    }
    ctx := context.Background()
    res, err := s.NamedSchedules.UpdateRateCardNamedSchedule(ctx, request)
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
| `request`                                                                                                              | [operations.UpdateRateCardNamedScheduleRequestBody](../../models/operations/updateratecardnamedschedulerequestbody.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.UpdateRateCardNamedScheduleResponse](../../models/operations/updateratecardnamedscheduleresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
