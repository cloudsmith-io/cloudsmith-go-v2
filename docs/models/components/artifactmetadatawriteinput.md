# ArtifactMetadataWriteInput

Serializer for POST create and PATCH update on artifact metadata.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Content`                                                                     | `any`                                                                         | :heavy_check_mark:                                                            | The metadata payload as a JSON object.                                        |
| `ContentType`                                                                 | `string`                                                                      | :heavy_check_mark:                                                            | The content type describing the structure of the metadata payload.            |
| `SourceIdentity`                                                              | `string`                                                                      | :heavy_check_mark:                                                            | The identifier for the source of this metadata, for example 'upstream:npmjs'. |