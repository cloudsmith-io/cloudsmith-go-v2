# WorkspacesPoliciesActionsListRequest


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Page`                                         | `*int64`                                       | :heavy_minus_sign:                             | A page number within the paginated result set. |
| `PageSize`                                     | `*int64`                                       | :heavy_minus_sign:                             | Number of results to return per page.          |
| `Policy`                                       | `string`                                       | :heavy_check_mark:                             | The `slug_perm` (identifier) of the Policy.    |
| `Workspace`                                    | `string`                                       | :heavy_check_mark:                             | The `name` of the Workspace.                   |