# DomainScope

The scope of the domain. `GLOBAL` for a Cloudsmith-owned domain, `WORKSPACE` for a custom domain without a repository, and `REPOSITORY` for a custom domain restricted to `repository`.

* `GLOBAL` - Global
* `WORKSPACE` - Workspace
* `REPOSITORY` - Repository

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.DomainScopeGlobal

// Open enum: custom values can be created with a direct type cast
custom := components.DomainScope("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `DomainScopeGlobal`     | GLOBAL                  |
| `DomainScopeWorkspace`  | WORKSPACE               |
| `DomainScopeRepository` | REPOSITORY              |