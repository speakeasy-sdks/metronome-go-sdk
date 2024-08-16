# MetronomeRateCardsNamedSchedules
(*RateCards.NamedSchedules*)

### Available Operations

* [Get](#get) - Get a rate card's named schedule
* [Update](#update) - Update a rate card's named schedule

## Get

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
    res, err := s.RateCards.NamedSchedules.Get(ctx, request)
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

## Update

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
    res, err := s.RateCards.NamedSchedules.Update(ctx, request)
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
