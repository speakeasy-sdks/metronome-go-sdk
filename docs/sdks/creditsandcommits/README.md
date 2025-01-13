# CreditsAndCommits
(*CreditsAndCommits*)

## Overview

Credits and commits are used to manage customer balances.

### Available Operations

* [DisableCommitTrueup](#disablecommittrueup) - Disable trueup for commit

## DisableCommitTrueup

Disable trueup invoice for a POSTPAID commit


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
    var request *operations.DisableCommitTrueupRequestBody = &operations.DisableCommitTrueupRequestBody{
        CustomerID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
        CommitID: "6162d87b-e5db-4a33-b7f2-76ce6ead4e85",
        ContractID: "7526bacf-f08a-47af-b473-bc57b88890e1",
    }
    ctx := context.Background()
    res, err := s.CreditsAndCommits.DisableCommitTrueup(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```



### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DisableCommitTrueupRequestBody](../../models/operations/disablecommittrueuprequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |


### Response

**[*operations.DisableCommitTrueupResponse](../../models/operations/disablecommittrueupresponse.md), error**
| Error Object                                               | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| sdkerrors.DisableCommitTrueupResponseBody                  | 400                                                        | application/json                                           |
| sdkerrors.DisableCommitTrueupCreditsAndCommitsResponseBody | 404                                                        | application/json                                           |
| sdkerrors.SDKError                                         | 4xx-5xx                                                    | */*                                                        |
