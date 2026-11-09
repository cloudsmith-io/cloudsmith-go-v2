# PolicyOSVSeveritySourceEnum

* `NVD` - The score was assigned by the National Vulnerability Database (NVD).
* `CNA` - The score was assigned by a CVE Numbering Authority (CNA).
* `SELF` - The score was assigned by the record's own database.

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PolicyOSVSeveritySourceEnumNvd

// Open enum: custom values can be created with a direct type cast
custom := components.PolicyOSVSeveritySourceEnum("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PolicyOSVSeveritySourceEnumNvd`  | NVD                               |
| `PolicyOSVSeveritySourceEnumCna`  | CNA                               |
| `PolicyOSVSeveritySourceEnumSelf` | SELF                              |