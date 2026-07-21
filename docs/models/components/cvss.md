# Cvss

CVSS details including vector, base_score, and parsed metric groups. Only populated for CVSS severity types.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Vector`                                            | optionalnullable.OptionalNullable[`string`]         | :heavy_minus_sign:                                  | The original CVSS vector string.                    |
| `BaseScore`                                         | optionalnullable.OptionalNullable[`float64`]        | :heavy_minus_sign:                                  | The numerical CVSS base score.                      |
| `Version`                                           | optionalnullable.OptionalNullable[`string`]         | :heavy_minus_sign:                                  | The CVSS version (e.g. '2.0', '3.0', '3.1', '4.0'). |
| `Groups`                                            | optionalnullable.OptionalNullable[[]`any`]          | :heavy_minus_sign:                                  | Parsed metric groups for display.                   |