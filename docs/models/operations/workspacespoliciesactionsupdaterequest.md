# WorkspacesPoliciesActionsUpdateRequest


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Action`                                                                      | `string`                                                                      | :heavy_check_mark:                                                            | The `slug_perm` (identifier) of the Action.                                   |
| `Policy`                                                                      | `string`                                                                      | :heavy_check_mark:                                                            | The `slug_perm` (identifier) of the Policy.                                   |
| `Workspace`                                                                   | `string`                                                                      | :heavy_check_mark:                                                            | The `name` of the Workspace.                                                  |
| `Body`                                                                        | [*components.PolicyActionInput](../../models/components/policyactioninput.md) | :heavy_minus_sign:                                                            | N/A                                                                           |