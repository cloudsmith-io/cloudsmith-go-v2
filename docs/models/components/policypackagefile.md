# PolicyPackageFile

The package files for policy purposes.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ChecksumMd5`                                         | optionalnullable.OptionalNullable[`string`]           | :heavy_minus_sign:                                    | N/A                                                   |
| `ChecksumSha1`                                        | optionalnullable.OptionalNullable[`string`]           | :heavy_minus_sign:                                    | N/A                                                   |
| `ChecksumSha256`                                      | optionalnullable.OptionalNullable[`string`]           | :heavy_minus_sign:                                    | N/A                                                   |
| `ChecksumSha512`                                      | optionalnullable.OptionalNullable[`string`]           | :heavy_minus_sign:                                    | N/A                                                   |
| `Filename`                                            | `string`                                              | :heavy_check_mark:                                    | N/A                                                   |
| `IsPrimary`                                           | `bool`                                                | :heavy_check_mark:                                    | Check if this is a primary package file (has no tag). |
| `Size`                                                | `*int64`                                              | :heavy_minus_sign:                                    | The calculated size of the file.                      |
| `SlugPerm`                                            | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `Tag`                                                 | optionalnullable.OptionalNullable[`string`]           | :heavy_minus_sign:                                    | Freeform descriptor that describes what the file is.  |