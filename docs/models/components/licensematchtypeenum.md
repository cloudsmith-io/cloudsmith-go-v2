# LicenseMatchTypeEnum

* `NONE` - None
* `OSS_EXACT` - OSS exact match
* `OSS_FUZZY` - OSS fuzzy match
* `PROPRIETARY` - Proprietary match
* `UNKNOWN` - Unknown match
* `CURATED` - Curated match

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.LicenseMatchTypeEnumNone

// Open enum: custom values can be created with a direct type cast
custom := components.LicenseMatchTypeEnum("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `LicenseMatchTypeEnumNone`        | NONE                              |
| `LicenseMatchTypeEnumOssExact`    | OSS_EXACT                         |
| `LicenseMatchTypeEnumOssFuzzy`    | OSS_FUZZY                         |
| `LicenseMatchTypeEnumProprietary` | PROPRIETARY                       |
| `LicenseMatchTypeEnumUnknown`     | UNKNOWN                           |
| `LicenseMatchTypeEnumCurated`     | CURATED                           |