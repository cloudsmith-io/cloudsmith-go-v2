# PolicyOSVSeverityTypeEnum

* `Ubuntu` - An "Ubuntu" severity type indicates that the associated score is a lowercased string
    representing the Ubuntu priority.

    If a severity has this type, the associated "score" will be one of "negligible",
    "low", "medium", "high", or "critical".

    See https://ubuntu.com/security/cves/about#priority for more information.
    
* `CVSS_V2` - A "CVSS_V2" severity type indicates that the associated score is a CVSS vector
    string using a version of the Common Vulnerability Scoring System notation that
    is == 2.0 (e.g."AV:L/AC:M/Au:N/C:N/I:P/A:C").
    
* `CVSS_V3` - A "CVSS_V3" severity type indicates that the associated score is a CVSS vector string
    using a version of the Common Vulnerability Scoring System notation that
    is >= 3.0 and < 4.0 (e.g."CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:C/C:H/I:N/A:N").
    
* `CVSS_V4` - A "CVSS_V4" severity type indicates that the associated score is a CVSS vector string
    using a version of the Common Vulnerability Scoring System notation that
    is >= 4.0 and < 5.0 (e.g. "CVSS:4.0/AV:N/AC:L/AT:N/PR:H/UI:N/VC:L/VI:L/VA:N/SC:N/SI:N/SA:N").
    

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.PolicyOSVSeverityTypeEnumUbuntu

// Open enum: custom values can be created with a direct type cast
custom := components.PolicyOSVSeverityTypeEnum("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PolicyOSVSeverityTypeEnumUbuntu` | Ubuntu                            |
| `PolicyOSVSeverityTypeEnumCvssV2` | CVSS_V2                           |
| `PolicyOSVSeverityTypeEnumCvssV3` | CVSS_V3                           |
| `PolicyOSVSeverityTypeEnumCvssV4` | CVSS_V4                           |