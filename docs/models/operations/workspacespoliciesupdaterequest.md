# WorkspacesPoliciesUpdateRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Policy`                                                             | `string`                                                             | :heavy_check_mark:                                                   | The `slug_perm` (identifier) of the Policy.                          |
| `Workspace`                                                          | `string`                                                             | :heavy_check_mark:                                                   | The `name` of the Workspace.                                         |
| `Body`                                                               | [components.PolicyRequest](../../models/components/policyrequest.md) | :heavy_check_mark:                                                   | N/A                                                                  |