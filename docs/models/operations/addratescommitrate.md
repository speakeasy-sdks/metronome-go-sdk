# AddRatesCommitRate

A distinct rate on the rate card. You can choose to use this rate rather than list rate when consuming a credit or commit.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `RateType`                                                                                   | [operations.AddRatesRateCardsRateType](../../models/operations/addratesratecardsratetype.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `Price`                                                                                      | **float64*                                                                                   | :heavy_minus_sign:                                                                           | Commit rate price. For FLAT rate_type, this must be >=0.                                     |
| `Tiers`                                                                                      | [][operations.AddRatesRateCardsTiers](../../models/operations/addratesratecardstiers.md)     | :heavy_minus_sign:                                                                           | Only set for TIERED rate_type.                                                               |