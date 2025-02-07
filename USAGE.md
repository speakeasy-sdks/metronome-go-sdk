<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"github.com/speakeasy-sdks/metronome-go-sdk/types"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.SetCustomerBillableStatusRequestBody = &operations.SetCustomerBillableStatusRequestBody{
		CustomerID:     "04ca7e72-4229-4a6e-ab11-9f7376fccbcb",
		BillableStatus: operations.BillableStatusBillable,
		EffectiveAt:    types.MustTimeFromString("2021-01-01T00:00:00Z"),
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
<!-- End SDK Example Usage [usage] -->