# Customers
(*Customers*)

### Available Operations

* [SetBillableStatus](#setbillablestatus) - Set customer billable status

## SetBillableStatus

Set a customer's billable status. This endpoint's availability is dependent on your client's configuration.

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
    var request *operations.SetCustomerBillableStatusRequestBody = &operations.SetCustomerBillableStatusRequestBody{
        CustomerID: "04ca7e72-4229-4a6e-ab11-9f7376fccbcb",
        BillableStatus: operations.BillableStatusBillable,
        EffectiveAt: types.MustTimeFromString("2021-01-01T00:00:00Z"),
    }
    ctx := context.Background()
    res, err := s.Customers.SetBillableStatus(ctx, request)
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
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
