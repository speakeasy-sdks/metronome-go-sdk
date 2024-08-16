<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Customers.SetBillableStatus(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->