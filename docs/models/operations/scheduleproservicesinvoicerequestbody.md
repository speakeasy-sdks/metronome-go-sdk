# ScheduleProServicesInvoiceRequestBody

schedule an invoice for the specified Professional Services terms on a contract


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CustomerID`                                                   | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `ContractID`                                                   | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `IssuedAt`                                                     | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | The date the invoice is issued                                 |
| `NetsuiteInvoiceHeaderStart`                                   | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | The start date of the invoice header in Netsuite               |
| `NetsuiteInvoiceHeaderEnd`                                     | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | The end date of the invoice header in Netsuite                 |
| `LineItems`                                                    | [][operations.LineItems](../../models/operations/lineitems.md) | :heavy_check_mark:                                             | Each line requires an amount or both unit_price and quantity.  |