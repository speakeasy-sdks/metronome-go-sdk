# Contracts
(*Contracts*)

## Overview

Contracts provide an alternative to plans for provisioning and invoicing customers. Use these endpoints to create and update contracts data.

### Available Operations

* [UpdateInvoiceIssueDate](#updateinvoiceissuedate) - Update invoice issue date
* [CreateHistoricalContractUsageInvoices](#createhistoricalcontractusageinvoices) - Create historical invoices

## UpdateInvoiceIssueDate

Update the issue date of a scheduled contract invoice


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
    var request *operations.UpdateInvoiceIssueDateRequestBody = &operations.UpdateInvoiceIssueDateRequestBody{
        InvoiceID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        IssueDate: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Contracts.UpdateInvoiceIssueDate(ctx, request)
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
| `request`                                                                                                    | [operations.UpdateInvoiceIssueDateRequestBody](../../models/operations/updateinvoiceissuedaterequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.UpdateInvoiceIssueDateResponse](../../models/operations/updateinvoiceissuedateresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.UpdateInvoiceIssueDateResponseBody          | 400                                                   | application/json                                      |
| sdkerrors.UpdateInvoiceIssueDateContractsResponseBody | 404                                                   | application/json                                      |
| sdkerrors.SDKError                                    | 4xx-5xx                                               | */*                                                   |

## CreateHistoricalContractUsageInvoices

Creates historical usage invoices for a contract

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
    var request *operations.CreateHistoricalContractUsageInvoicesRequestBody = &operations.CreateHistoricalContractUsageInvoicesRequestBody{
        Invoices: []operations.Invoices{
            operations.Invoices{
                CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
                ContractID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
                CreditTypeID: "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
                InclusiveStartDate: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
                ExclusiveEndDate: types.MustTimeFromString("2020-02-01T00:00:00.000Z"),
                IssueDate: types.MustTimeFromString("2020-02-01T00:00:00.000Z"),
                UsageLineItems: []operations.UsageLineItems{
                    operations.UsageLineItems{
                        ProductID: "f14d6729-6a44-4b13-9908-9387f1918790",
                        InclusiveStartDate: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
                        ExclusiveEndDate: types.MustTimeFromString("2020-02-01T00:00:00.000Z"),
                        Quantity: metronomegosdk.Float64(100),
                    },
                },
            },
        },
        Preview: false,
    }
    ctx := context.Background()
    res, err := s.Contracts.CreateHistoricalContractUsageInvoices(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.CreateHistoricalContractUsageInvoicesRequestBody](../../models/operations/createhistoricalcontractusageinvoicesrequestbody.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |


### Response

**[*operations.CreateHistoricalContractUsageInvoicesResponse](../../models/operations/createhistoricalcontractusageinvoicesresponse.md), error**
| Error Object                                                | Status Code                                                 | Content Type                                                |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| sdkerrors.CreateHistoricalContractUsageInvoicesResponseBody | 400                                                         | application/json                                            |
| sdkerrors.SDKError                                          | 4xx-5xx                                                     | */*                                                         |
