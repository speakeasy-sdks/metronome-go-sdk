# Products
(*Products*)

### Available Operations

* [Get](#get) - Get a product
* [List](#list) - List products
* [Create](#create) - Create a product
* [Update](#update) - Update a product
* [Archive](#archive) - Archive a product

## Get

Get a specific product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.GetProductRequestBody = &operations.GetProductRequestBody{
        ID: "d84e7f4e-7a70-4fe4-be02-7a5027beffcc",
    }
    ctx := context.Background()
    res, err := s.Products.Get(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetProductRequestBody](../../models/operations/getproductrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |


### Response

**[*operations.GetProductResponse](../../models/operations/getproductresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.NotFound | 404                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List products


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"context"
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



### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `limit`                                                                                   | **int64*                                                                                  | :heavy_minus_sign:                                                                        | Max number of results that should be returned                                             |
| `nextPage`                                                                                | **string*                                                                                 | :heavy_minus_sign:                                                                        | Cursor that indicates where the next page of results should start.                        |
| `requestBody`                                                                             | [*operations.ListProductsRequestBody](../../models/operations/listproductsrequestbody.md) | :heavy_minus_sign:                                                                        | Get list of products                                                                      |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |


### Response

**[*operations.ListProductsResponse](../../models/operations/listproductsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

Create a new product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.CreateProductRequestBody = &operations.CreateProductRequestBody{
        Name: "My Product",
        Type: operations.TypeUsageUpper,
        BillableMetricID: metronomegosdk.String("13117714-3f05-48e5-a6e9-a66093f13b4d"),
    }
    ctx := context.Background()
    res, err := s.Products.Create(ctx, request)
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
| `request`                                                                                  | [operations.CreateProductRequestBody](../../models/operations/createproductrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.CreateProductResponse](../../models/operations/createproductresponse.md), error**
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## Update

Update a product


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
    var request *operations.UpdateProductRequestBody = &operations.UpdateProductRequestBody{
        ProductID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
        Name: metronomegosdk.String("My Updated Product"),
        StartingAt: types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
    }
    ctx := context.Background()
    res, err := s.Products.Update(ctx, request)
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
| `request`                                                                                  | [operations.UpdateProductRequestBody](../../models/operations/updateproductrequestbody.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |


### Response

**[*operations.UpdateProductResponse](../../models/operations/updateproductresponse.md), error**
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |

## Archive

Archive a product


### Example Usage

```go
package main

import(
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"context"
	"log"
)

func main() {
    s := metronomegosdk.New(
        metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var request *operations.ArchiveProductListItemRequestBody = &operations.ArchiveProductListItemRequestBody{
        ProductID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
    }
    ctx := context.Background()
    res, err := s.Products.Archive(ctx, request)
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
| `request`                                                                                                    | [operations.ArchiveProductListItemRequestBody](../../models/operations/archiveproductlistitemrequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |


### Response

**[*operations.ArchiveProductListItemResponse](../../models/operations/archiveproductlistitemresponse.md), error**
| Error Object         | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| sdkerrors.BadRequest | 400                  | application/json     |
| sdkerrors.NotFound   | 404                  | application/json     |
| sdkerrors.SDKError   | 4xx-5xx              | */*                  |
