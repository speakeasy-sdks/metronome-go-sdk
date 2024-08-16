# GetContractRequestBody

Contract and customer IDs


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `CustomerID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `ContractID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `IncludeLedgers`                                                                            | **bool*                                                                                     | :heavy_minus_sign:                                                                          | Include commit ledgers in the response. Setting this flag may cause the query to be slower. |