# PolicyValidateFinding

A single finding emitted by the validate endpoint.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Severity`                                                                      | `string`                                                                        | :heavy_check_mark:                                                              | Finding severity. Currently always 'error'.                                     |
| `Code`                                                                          | `string`                                                                        | :heavy_check_mark:                                                              | Machine-readable finding code (e.g. 'rego_type_error', 'policy_compile_error'). |
| `Message`                                                                       | `string`                                                                        | :heavy_check_mark:                                                              | Human-readable finding message.                                                 |
| `Line`                                                                          | `*int64`                                                                        | :heavy_check_mark:                                                              | 1-indexed line number in the source, if known.                                  |