# PackagesVulnerabilitiesOsvListThreatLevel

Filter by threat level. Matches the single best classified severity. Accepts normalized threat levels (critical, high, medium, low, unknown, malware) and vendor aliases (important, moderate, negligible, none, unimportant).

* `critical` - critical
* `high` - high
* `important` - important
* `low` - low
* `malware` - malware
* `medium` - medium
* `moderate` - moderate
* `negligible` - negligible
* `none` - none
* `unimportant` - unimportant
* `unknown` - unknown

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.PackagesVulnerabilitiesOsvListThreatLevelCritical
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `PackagesVulnerabilitiesOsvListThreatLevelCritical`    | critical                                               |
| `PackagesVulnerabilitiesOsvListThreatLevelHigh`        | high                                                   |
| `PackagesVulnerabilitiesOsvListThreatLevelImportant`   | important                                              |
| `PackagesVulnerabilitiesOsvListThreatLevelLow`         | low                                                    |
| `PackagesVulnerabilitiesOsvListThreatLevelMalware`     | malware                                                |
| `PackagesVulnerabilitiesOsvListThreatLevelMedium`      | medium                                                 |
| `PackagesVulnerabilitiesOsvListThreatLevelModerate`    | moderate                                               |
| `PackagesVulnerabilitiesOsvListThreatLevelNegligible`  | negligible                                             |
| `PackagesVulnerabilitiesOsvListThreatLevelNone`        | none                                                   |
| `PackagesVulnerabilitiesOsvListThreatLevelUnimportant` | unimportant                                            |
| `PackagesVulnerabilitiesOsvListThreatLevelUnknown`     | unknown                                                |