# GetRatesRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Limit`                                                                           | **int64*                                                                          | :heavy_minus_sign:                                                                | Max number of results that should be returned                                     |
| `NextPage`                                                                        | **string*                                                                         | :heavy_minus_sign:                                                                | Cursor that indicates where the next page of results should start.                |
| `RequestBody`                                                                     | [*operations.GetRatesRequestBody](../../models/operations/getratesrequestbody.md) | :heavy_minus_sign:                                                                | Rate schedule filter options.                                                     |