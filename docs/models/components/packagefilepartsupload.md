# PackageFilePartsUpload

Serializer for package file part requests.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Identifier`                                         | `string`                                             | :heavy_check_mark:                                   | The identifier for the file to use uploading parts.  |
| `UploadQuerystring`                                  | `*string`                                            | :heavy_check_mark:                                   | The querystring to use for the next-step PUT upload. |
| `UploadURL`                                          | `string`                                             | :heavy_check_mark:                                   | The URL to use for the next-step PUT upload          |