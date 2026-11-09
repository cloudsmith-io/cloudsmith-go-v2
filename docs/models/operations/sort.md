# Sort

Field to sort by. Prepend `-` for descending order. `severity` sorts by threat level, then CVSS score; within each threat level, vulnerabilities without a CVSS score come last when descending and first when ascending. `epss` sorts by the EPSS score of the vulnerability's CVE (`cve_id`); vulnerabilities without an EPSS score always come last, and equal scores are ordered most severe first in either direction.

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.SortMinusEpss
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `SortMinusEpss`     | -epss               |
| `SortMinusSeverity` | -severity           |
| `SortEpss`          | epss                |
| `SortSeverity`      | severity            |