# ErrorDetail

A serializer for displaying a detail error message.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `Detail`                                                                                         | `string`                                                                                         | :heavy_check_mark:                                                                               | An extended message for the response.                                                            |
| `Fields`                                                                                         | map[string][]`string`                                                                            | :heavy_minus_sign:                                                                               | A Dictionary of related errors where key: Field and value: Array of Errors related to that field |
| `HTTPMeta`                                                                                       | [components.HTTPMetadata](../../models/components/httpmetadata.md)                               | :heavy_check_mark:                                                                               | N/A                                                                                              |