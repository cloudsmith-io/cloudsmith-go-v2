# PatchedArtifactMetadataWrite

Serializer for POST create and PATCH update on artifact metadata.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Content`                                                                     | `any`                                                                         | :heavy_minus_sign:                                                            | The metadata payload as a JSON object.                                        |
| `ContentType`                                                                 | `*string`                                                                     | :heavy_minus_sign:                                                            | The content type describing the structure of the metadata payload.            |
| `SourceIdentity`                                                              | `*string`                                                                     | :heavy_minus_sign:                                                            | The identifier for the source of this metadata, for example 'upstream:npmjs'. |