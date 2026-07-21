# PackageFileUploadV2Response

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

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Identifier`                                                 | `string`                                                     | :heavy_check_mark:                                           | The identifier of the new package file.                      |
| `UploadFields`                                               | `any`                                                        | :heavy_check_mark:                                           | The dictionary of fields that must be sent with POST uploads |
| `UploadHeaders`                                              | `any`                                                        | :heavy_check_mark:                                           | The dictionary of headers that must be sent with uploads     |
| `UploadQuerystring`                                          | `*string`                                                    | :heavy_check_mark:                                           | The querystring to use for the next-step POST or PUT upload  |
| `UploadURL`                                                  | `string`                                                     | :heavy_check_mark:                                           | The URL to use for the next-step POST or PUT upload          |