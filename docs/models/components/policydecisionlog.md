# PolicyDecisionLog

Serializer for policy decision logs


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ID`                                                                 | `*string`                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `CorrelationID`                                                      | `string`                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `Policy`                                                             | [components.Policy](../../models/components/policy.md)               | :heavy_check_mark:                                                   | N/A                                                                  |
| `StartedAt`                                                          | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | N/A                                                                  |
| `EndedAt`                                                            | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | N/A                                                                  |
| `PolicyInput`                                                        | [components.PolicyInput](../../models/components/policyinput.md)     | :heavy_check_mark:                                                   | Serializer for policy input.                                         |
| `PolicyOutput`                                                       | `any`                                                                | :heavy_check_mark:                                                   | N/A                                                                  |
| `Actions`                                                            | [][components.PolicyAction](../../models/components/policyaction.md) | :heavy_check_mark:                                                   | The actions configured against the policy.                           |
| `ParsedActions`                                                      | `any`                                                                | :heavy_check_mark:                                                   | The actions which were invoked on this package.                      |