<div align="center">
    <img src="https://avatars.githubusercontent.com/u/60759619?s=200&v=4">
    <h1>Golang SDK</h1>
        <p><strong>Billing that accelerates you</strong></p>
        <p>Meet the usage-based billing platform built for modern software companies. Launch products faster, offer any pricing model, and streamline finance workflows without writing code.</p>
    <a href="https://docs.metronome.com"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/Metronome-Industries/metronome-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Contracts](docs/sdks/contracts/README.md)

* [SetCustomerBillableStatus](docs/sdks/contracts/README.md#setcustomerbillablestatus) - Set customer billable status
* [GetProduct](docs/sdks/contracts/README.md#getproduct) - Get a product
* [ListProducts](docs/sdks/contracts/README.md#listproducts) - List products
* [CreateProduct](docs/sdks/contracts/README.md#createproduct) - Create a product
* [UpdateProduct](docs/sdks/contracts/README.md#updateproduct) - Update a product
* [ArchiveProductListItem](docs/sdks/contracts/README.md#archiveproductlistitem) - Archive a product
* [GetRateSchedule](docs/sdks/contracts/README.md#getrateschedule) - Get a rate schedule
* [GetRates](docs/sdks/contracts/README.md#getrates) - Get rates
* [GetRateCard](docs/sdks/contracts/README.md#getratecard) - Get a rate card
* [ListRateCards](docs/sdks/contracts/README.md#listratecards) - List rate cards
* [CreateRateCard](docs/sdks/contracts/README.md#createratecard) - Create a rate card
* [UpdateRateCard](docs/sdks/contracts/README.md#updateratecard) - Update a rate card
* [AddRate](docs/sdks/contracts/README.md#addrate) - Add a rate
* [AddRates](docs/sdks/contracts/README.md#addrates) - Add rates
* [SetRateCardProductsOrder](docs/sdks/contracts/README.md#setratecardproductsorder) - Set the rate card products order
* [MoveRateCardProducts](docs/sdks/contracts/README.md#moveratecardproducts) - Update the rate card products order
* [GetContract](docs/sdks/contracts/README.md#getcontract) - Get a contract
* [ListContracts](docs/sdks/contracts/README.md#listcontracts) - List customer contracts
* [CreateContract](docs/sdks/contracts/README.md#createcontract) - Create a contract
* [AmendContract](docs/sdks/contracts/README.md#amendcontract) - Amend a contract
* [ArchiveContract](docs/sdks/contracts/README.md#archivecontract) - Archive a contract
* [SetUsageFilter](docs/sdks/contracts/README.md#setusagefilter) - Set a contract usage filter
* [AddManualBalanceLedgerEntry](docs/sdks/contracts/README.md#addmanualbalanceledgerentry) - Add a manual balance entry
* [UpdateContractEndDate](docs/sdks/contracts/README.md#updatecontractenddate) - Update the contract end date
* [GetContractRateSchedule](docs/sdks/contracts/README.md#getcontractrateschedule) - Get the rate schedule for a contract
* [ListCustomerCommits](docs/sdks/contracts/README.md#listcustomercommits) - List commits
* [CreateCustomerCommit](docs/sdks/contracts/README.md#createcustomercommit) - Create a commit
* [UpdateCommitEndDate](docs/sdks/contracts/README.md#updatecommitenddate) - Update the commit end date
* [ListCustomerCredits](docs/sdks/contracts/README.md#listcustomercredits) - List credits
* [CreateCustomerCredit](docs/sdks/contracts/README.md#createcustomercredit) - Create a credit
* [UpdateCreditEndDate](docs/sdks/contracts/README.md#updatecreditenddate) - Update the credit end date
* [ListCustomerBalances](docs/sdks/contracts/README.md#listcustomerbalances) - List balances
* [ScheduleProServicesInvoice](docs/sdks/contracts/README.md#scheduleproservicesinvoice) - Schedule ProService invoice

### [Invoices](docs/sdks/invoices/README.md)

* [RegenerateInvoice](docs/sdks/invoices/README.md#regenerateinvoice) - Regenerate an invoice

### [NamedSchedules](docs/sdks/namedschedules/README.md)

* [GetCustomerNamedSchedule](docs/sdks/namedschedules/README.md#getcustomernamedschedule) - Get a customer's named schedule
* [UpdateCustomerNamedSchedule](docs/sdks/namedschedules/README.md#updatecustomernamedschedule) - Update a customer's named schedule
* [GetContractNamedSchedule](docs/sdks/namedschedules/README.md#getcontractnamedschedule) - Get a contract's named schedule
* [UpdateContractNamedSchedule](docs/sdks/namedschedules/README.md#updatecontractnamedschedule) - Update a contract's named schedule
* [GetRateCardNamedSchedule](docs/sdks/namedschedules/README.md#getratecardnamedschedule) - Get a rate card's named schedule
* [UpdateRateCardNamedSchedule](docs/sdks/namedschedules/README.md#updateratecardnamedschedule) - Update a rate card's named schedule
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/retry"
	"log"
	"models/operations"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/retry"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

### Example

```go
package main

import (
	"context"
	"errors"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"github.com/Metronome-Industries/metronome-go-sdk/models/sdkerrors"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
	if err != nil {

		var e *sdkerrors.BadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.NotFound
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.metronome.com/v1` | None |

#### Example

```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithServerIndex(0),
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithServerURL("https://api.metronome.com/v1"),
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Contracts.SetCustomerBillableStatus(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	metronomegosdk "github.com/Metronome-Industries/metronome-go-sdk"
	"github.com/Metronome-Industries/metronome-go-sdk/models/operations"
	"log"
)

func main() {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var requestBody *operations.ListProductsRequestBody = &operations.ListProductsRequestBody{
		ArchiveFilter: operations.ArchiveFilterNotArchived.ToPointer(),
	}
	ctx := context.Background()
	res, err := s.Contracts.ListProducts(ctx, nil, nil, requestBody)
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
<!-- End Pagination [pagination] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/metronome-industries/metronome-go-sdk&utm_campaign=go)
