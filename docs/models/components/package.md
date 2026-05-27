# Package

Identifies the affected code library or command provided by the package.

The OSV package field is described in detail at https://ossf.github.io/osv-schema/#affectedpackage-field.


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Ecosystem`                                                                                                   | `string`                                                                                                      | :heavy_check_mark:                                                                                            | The overall library ecosystem which package belongs to.                                                       |
| `Name`                                                                                                        | `string`                                                                                                      | :heavy_check_mark:                                                                                            | The library within the overall ecosystem. Note that the semantics of this field depend on the ecosystem.      |
| `Purl`                                                                                                        | `*string`                                                                                                     | :heavy_check_mark:                                                                                            | A string following the Package URL specification that identifies the package, without the @version component. |