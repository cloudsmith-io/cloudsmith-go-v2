# PackageFileUploadInfoV2

Response body for a v2 file-upload request: where to send the file bytes.

Only ``identifier`` and ``upload_url`` are always present; the other fields are
populated per upload method (POST form fields, PUT headers or a presigned
querystring), so they are optional.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Identifier`                                | `string`                                    | :heavy_check_mark:                          | The identifier of the new package file.     |
| `UploadURL`                                 | `string`                                    | :heavy_check_mark:                          | The URL to send the file bytes to.          |
| `UploadFields`                              | optionalnullable.OptionalNullable[`any`]    | :heavy_minus_sign:                          | Form fields to include with a POST upload.  |
| `UploadHeaders`                             | optionalnullable.OptionalNullable[`any`]    | :heavy_minus_sign:                          | Headers to include with a PUT upload.       |
| `UploadQuerystring`                         | optionalnullable.OptionalNullable[`string`] | :heavy_minus_sign:                          | Query string to append to the upload URL.   |