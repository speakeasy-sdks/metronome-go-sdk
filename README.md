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
go get github.com/speakeasy-sdks/metronome-go-sdk
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Customers](docs/sdks/customers/README.md)

* [SetBillableStatus](docs/sdks/customers/README.md#setbillablestatus) - Set customer billable status

### [Customers.NamedSchedules](docs/sdks/namedschedules/README.md)

* [Get](docs/sdks/namedschedules/README.md#get) - Get a customer's named schedule
* [Update](docs/sdks/namedschedules/README.md#update) - Update a customer's named schedule

### [Invoices](docs/sdks/invoices/README.md)

* [Regenerate](docs/sdks/invoices/README.md#regenerate) - Regenerate an invoice

### [Products](docs/sdks/products/README.md)

* [Get](docs/sdks/products/README.md#get) - Get a product
* [List](docs/sdks/products/README.md#list) - List products
* [Create](docs/sdks/products/README.md#create) - Create a product
* [Update](docs/sdks/products/README.md#update) - Update a product
* [Archive](docs/sdks/products/README.md#archive) - Archive a product

### [RateCards](docs/sdks/ratecards/README.md)

* [GetRateSchedule](docs/sdks/ratecards/README.md#getrateschedule) - Get a rate schedule
* [GetRates](docs/sdks/ratecards/README.md#getrates) - Get rates
* [Get](docs/sdks/ratecards/README.md#get) - Get a rate card
* [List](docs/sdks/ratecards/README.md#list) - List rate cards
* [Create](docs/sdks/ratecards/README.md#create) - Create a rate card
* [Update](docs/sdks/ratecards/README.md#update) - Update a rate card
* [AddRate](docs/sdks/ratecards/README.md#addrate) - Add a rate
* [AddRates](docs/sdks/ratecards/README.md#addrates) - Add rates
* [SetProductsOrder](docs/sdks/ratecards/README.md#setproductsorder) - Set the rate card products order
* [MoveProducts](docs/sdks/ratecards/README.md#moveproducts) - Update the rate card products order

### [RateCards.NamedSchedules](docs/sdks/metronomeratecardsnamedschedules/README.md)

* [Get](docs/sdks/metronomeratecardsnamedschedules/README.md#get) - Get a rate card's named schedule
* [Update](docs/sdks/metronomeratecardsnamedschedules/README.md#update) - Update a rate card's named schedule

### [Contracts](docs/sdks/contracts/README.md)

* [Get](docs/sdks/contracts/README.md#get) - Get a contract
* [List](docs/sdks/contracts/README.md#list) - List customer contracts
* [Create](docs/sdks/contracts/README.md#create) - Create a contract
* [Amend](docs/sdks/contracts/README.md#amend) - Amend a contract
* [Archive](docs/sdks/contracts/README.md#archive) - Archive a contract
* [SetUsageFilter](docs/sdks/contracts/README.md#setusagefilter) - Set a contract usage filter
* [AddManualBalanceEntry](docs/sdks/contracts/README.md#addmanualbalanceentry) - Add a manual balance entry
* [UpdateEndDate](docs/sdks/contracts/README.md#updateenddate) - Update the contract end date
* [GetRateSchedule](docs/sdks/contracts/README.md#getrateschedule) - Get the rate schedule for a contract
* [ScheduleProServicesInvoice](docs/sdks/contracts/README.md#scheduleproservicesinvoice) - Schedule ProService invoice

### [Contracts.NamedSchedules](docs/sdks/metronomenamedschedules/README.md)

* [Get](docs/sdks/metronomenamedschedules/README.md#get) - Get a contract's named schedule
* [Update](docs/sdks/metronomenamedschedules/README.md#update) - Update a contract's named schedule

### [CustomerCommits](docs/sdks/customercommits/README.md)

* [List](docs/sdks/customercommits/README.md#list) - List commits
* [Create](docs/sdks/customercommits/README.md#create) - Create a commit
* [UpdateEndDate](docs/sdks/customercommits/README.md#updateenddate) - Update the commit end date

### [CustomerCredits](docs/sdks/customercredits/README.md)

* [List](docs/sdks/customercredits/README.md#list) - List credits
* [Create](docs/sdks/customercredits/README.md#create) - Create a credit
* [UpdateEndDate](docs/sdks/customercredits/README.md#updateenddate) - Update the credit end date

### [CustomerBalances](docs/sdks/customerbalances/README.md)

* [List](docs/sdks/customerbalances/README.md#list) - List balances
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
	res, err := s.Customers.SetBillableStatus(ctx, nil, operations.WithRetries(
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
	res, err := s.Customers.SetBillableStatus(ctx, nil)
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
	res, err := s.Customers.SetBillableStatus(ctx, nil)
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
	res, err := s.Customers.SetBillableStatus(ctx, nil)
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
	res, err := s.Customers.SetBillableStatus(ctx, nil)
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
	res, err := s.Customers.SetBillableStatus(ctx, nil)
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
	res, err := s.Products.List(ctx, nil, nil, requestBody)
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
