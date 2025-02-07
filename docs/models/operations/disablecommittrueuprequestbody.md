# DisableCommitTrueupRequestBody

Information to identify the commit


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `CustomerID`                                          | *string*                                              | :heavy_check_mark:                                    | ID of the customer whose commit is to be updated      |
| `CommitID`                                            | *string*                                              | :heavy_check_mark:                                    | ID of the commit to update                            |
| `ContractID`                                          | *string*                                              | :heavy_check_mark:                                    | ID of the contract that the commit is on              |
| `AmendmentID`                                         | **string*                                             | :heavy_minus_sign:                                    | If applicable, the amendment ID that the commit is on |