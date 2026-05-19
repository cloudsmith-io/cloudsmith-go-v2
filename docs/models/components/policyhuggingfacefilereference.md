# PolicyHuggingfaceFileReference

The files in a Huggingface model or dataset for policy purposes.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `CreatedAt`                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                   | :heavy_minus_sign:                                                           | N/A                                                                          |
| `FileExtension`                                                              | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `Filename`                                                                   | optionalnullable.OptionalNullable[`string`]                                  | :heavy_minus_sign:                                                           | The filename i.e 'model.safetensors'.                                        |
| `Oid`                                                                        | `string`                                                                     | :heavy_check_mark:                                                           | The content-addressed object ID (OID) of this file, typically a SHA256 hash. |
| `Size`                                                                       | `int64`                                                                      | :heavy_check_mark:                                                           | The calculated storage size for the file.                                    |
| `Path`                                                                       | `string`                                                                     | :heavy_check_mark:                                                           | The full path of the file within the package i.e 'models/model.safetensors'. |