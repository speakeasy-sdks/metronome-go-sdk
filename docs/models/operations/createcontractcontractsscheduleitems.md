# CreateContractContractsScheduleItems


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Amount`                                  | *float64*                                 | :heavy_check_mark:                        | N/A                                       |
| `StartingAt`                              | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | RFC 3339 timestamp (inclusive)            |
| `EndingBefore`                            | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | RFC 3339 timestamp (exclusive)            |