# GetRateScheduleData


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ProductID`                                        | *string*                                           | :heavy_check_mark:                                 | N/A                                                |
| `ProductName`                                      | *string*                                           | :heavy_check_mark:                                 | N/A                                                |
| `ProductTags`                                      | []*string*                                         | :heavy_check_mark:                                 | N/A                                                |
| `PricingGroupValues`                               | map[string]*string*                                | :heavy_minus_sign:                                 | N/A                                                |
| `StartingAt`                                       | [time.Time](https://pkg.go.dev/time#Time)          | :heavy_check_mark:                                 | N/A                                                |
| `EndingBefore`                                     | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |
| `Entitled`                                         | *bool*                                             | :heavy_check_mark:                                 | N/A                                                |
| `Rate`                                             | [operations.Rate](../../models/operations/rate.md) | :heavy_check_mark:                                 | N/A                                                |