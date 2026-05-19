# ArtifactMetadataValidate

Serializer for POST /metadata/validate.

Same shape and validation rules as the write serializer but without
fields tied to persistence (e.g. ``source_identity``). Used by clients
that want to fail fast on a metadata payload before they have a
package slug to attach it to.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Content`                                                          | `any`                                                              | :heavy_check_mark:                                                 | The metadata payload as a JSON object.                             |
| `ContentType`                                                      | `string`                                                           | :heavy_check_mark:                                                 | The content type describing the structure of the metadata payload. |