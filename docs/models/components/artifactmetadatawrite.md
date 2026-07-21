# ArtifactMetadataWrite

Fields for creating or updating package metadata.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Content`                                                                     | map[string]`any`                                                              | :heavy_check_mark:                                                            | The metadata payload as a JSON object.                                        |
| `ContentType`                                                                 | `string`                                                                      | :heavy_check_mark:                                                            | The content type describing the structure of the metadata payload.            |
| `SourceIdentity`                                                              | `string`                                                                      | :heavy_check_mark:                                                            | The identifier for the source of this metadata, for example 'upstream:npmjs'. |