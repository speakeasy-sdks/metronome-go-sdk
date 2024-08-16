# Invoices
(*Invoices*)

## Overview

[Invoices](https://docs.metronome.com/invoicing/) reflect how much a customer spent during a period, which is the basis for billing. Metronome automatically generates invoices based upon your pricing, packaging, and usage events. Use these endpoints to retrieve invoices.

### Available Operations

* [RegenerateInvoice](#regenerateinvoice) - Regenerate an invoice

## RegenerateInvoice

Regenerate a voided contract invoice

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
    var request *operations.RegenerateInvoiceRequestBody = &operations.RegenerateInvoiceRequestBody{
        ID: "6a37bb88-8538-48c5-b37b-a41c836328bd",
    }
    ctx := context.Background()
    res, err := s.Invoices.RegenerateInvoice(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RegenerateInvoiceRequestBody](../../models/operations/regenerateinvoicerequestbody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.RegenerateInvoiceResponse](../../models/operations/regenerateinvoiceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
