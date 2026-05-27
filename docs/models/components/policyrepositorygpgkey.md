# PolicyRepositoryGpgKey

A GPG key used to sign repository contents.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Active`                                                 | `*bool`                                                  | :heavy_minus_sign:                                       | If selected this is the active key for this repository.  |
| `Comment`                                                | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `CreatedAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)               | :heavy_minus_sign:                                       | N/A                                                      |
| `Default`                                                | `*bool`                                                  | :heavy_minus_sign:                                       | If selected this is the default key for this repository. |
| `FingerprintShort`                                       | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `Fingerprint`                                            | `string`                                                 | :heavy_check_mark:                                       | The long identifier used by GPG for this key.            |
| `PublicKey`                                              | `string`                                                 | :heavy_check_mark:                                       | The public key given to repository users.                |