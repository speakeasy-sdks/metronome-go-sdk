# ArchiveContractRequestBody

Archive a contract


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `CustomerID`                                                                          | *string*                                                                              | :heavy_check_mark:                                                                    | ID of the customer whose contract is to be archived                                   |
| `ContractID`                                                                          | *string*                                                                              | :heavy_check_mark:                                                                    | ID of the contract to archive                                                         |
| `VoidInvoices`                                                                        | *bool*                                                                                | :heavy_check_mark:                                                                    | If false, the existing finalized invoices will remain after the contract is archived. |