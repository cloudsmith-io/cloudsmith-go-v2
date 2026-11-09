# Format

Response format. Use `json` (default) for a paginated JSON response, or `csv` to download the results as CSV.

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/operations"
)

value := operations.FormatCsv
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `FormatCsv`  | csv          |
| `FormatJSON` | json         |