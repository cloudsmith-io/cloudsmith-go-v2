# WorkspacesRepositoriesFilesInfoRetrieveRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Filename`                                         | `string`                                           | :heavy_check_mark:                                 | The filename of the file being uploaded.           |
| `Identifier`                                       | `string`                                           | :heavy_check_mark:                                 | The `slug_perm` (identifier) of the package file.  |
| `PartNumber`                                       | `*int64`                                           | :heavy_minus_sign:                                 | The part number to be uploaded next.               |
| `Repository`                                       | `string`                                           | :heavy_check_mark:                                 | The `slug` of the repository within the workspace. |
| `Workspace`                                        | `string`                                           | :heavy_check_mark:                                 | The `slug` of the workspace.                       |