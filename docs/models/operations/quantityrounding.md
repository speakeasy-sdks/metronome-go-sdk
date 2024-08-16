# QuantityRounding

Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `RoundingMethod`                                                       | [operations.RoundingMethod](../../models/operations/roundingmethod.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `DecimalPlaces`                                                        | *float64*                                                              | :heavy_check_mark:                                                     | N/A                                                                    |