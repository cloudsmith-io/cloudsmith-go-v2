# PackageFileUploadV2

Request/response body for the v2 file-upload create/validate endpoints.

``method`` is documented as a plain string rather than the inherited
``ChoiceField``: a second enum field named ``method`` collides with the
field-named ``MethodEnum`` component drf-spectacular derives for the logs
endpoint, renaming it — a non-additive change. Runtime validation is
unaffected (the viewset still uses the v1 serializer's ``ChoiceField``); this
class only shapes the generated schema, so the v1 serializer (and its enum
component) never enters the v2 spec.

``identifier`` is re-declared as a plain string: the inherited v1 field is a
``URLField``, but the response value is ``package_file.slug_perm`` (a slug,
never a URL), so documenting it as ``format: uri`` would mislead SDKs.


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `Filename`                                                                                                               | `string`                                                                                                                 | :heavy_check_mark:                                                                                                       | Filename for the package file upload.                                                                                    |
| `Method`                                                                                                                 | `*string`                                                                                                                | :heavy_minus_sign:                                                                                                       | The method to use for package file upload: one of `put_parts` (preferred), `put`, `post`, `presigned` or `unsigned_put`. |
| `Md5Checksum`                                                                                                            | `*string`                                                                                                                | :heavy_minus_sign:                                                                                                       | MD5 checksum for a POST-based package file upload.                                                                       |
| `Sha256Checksum`                                                                                                         | `*string`                                                                                                                | :heavy_minus_sign:                                                                                                       | SHA256 checksum for a PUT-based package file upload.                                                                     |