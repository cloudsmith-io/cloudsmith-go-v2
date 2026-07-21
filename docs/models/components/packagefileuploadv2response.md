# PackageFileUploadV2Response

Request/response body for the v2 file-upload create/validate endpoints.

``method`` is documented as a plain string rather than the inherited
``ChoiceField``: a second enum field named ``method`` collides with the
field-named ``MethodEnum`` component drf-spectacular derives for the logs
endpoint, renaming it — a non-additive change. Runtime validation is
unaffected (the viewset still uses the v1 serializer's ``ChoiceField``); this
class only shapes the generated schema, so the v1 serializer (and its enum
component) never enters the v2 spec.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Identifier`                                                 | `string`                                                     | :heavy_check_mark:                                           | The identifier for the file to use when creating packages    |
| `UploadFields`                                               | `any`                                                        | :heavy_check_mark:                                           | The dictionary of fields that must be sent with POST uploads |
| `UploadHeaders`                                              | `any`                                                        | :heavy_check_mark:                                           | The dictionary of headers that must be sent with uploads     |
| `UploadQuerystring`                                          | `*string`                                                    | :heavy_check_mark:                                           | The querystring to use for the next-step POST or PUT upload  |
| `UploadURL`                                                  | `string`                                                     | :heavy_check_mark:                                           | The URL to use for the next-step POST or PUT upload          |