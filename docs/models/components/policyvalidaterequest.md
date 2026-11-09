# PolicyValidateRequest

Request body for POST .../policies/validate/.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Rego`                                                             | `string`                                                           | :heavy_check_mark:                                                 | The Rego policy source to validate. Maximum 256 KB. Not persisted. |