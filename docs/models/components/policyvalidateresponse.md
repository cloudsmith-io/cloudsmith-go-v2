# PolicyValidateResponse

Response body for POST .../policies/validate/.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `IsValid`                                                                              | `bool`                                                                                 | :heavy_check_mark:                                                                     | True when no finding has severity 'error'.                                             |
| `Findings`                                                                             | [][components.PolicyValidateFinding](../../models/components/policyvalidatefinding.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `Raw`                                                                                  | `string`                                                                               | :heavy_check_mark:                                                                     | Full raw OPA output; empty when the policy is valid.                                   |