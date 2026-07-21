# WorkspacesPoliciesActionsCreateRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Policy`                                                                          | `string`                                                                          | :heavy_check_mark:                                                                | The `slug_perm` (identifier) of the Policy.                                       |
| `Workspace`                                                                       | `string`                                                                          | :heavy_check_mark:                                                                | The `name` of the Workspace.                                                      |
| `Body`                                                                            | [*components.PolicyActionRequest](../../models/components/policyactionrequest.md) | :heavy_minus_sign:                                                                | N/A                                                                               |