# AuthenticationMethodEnum

* `NONE` - none
* `TOKEN` - token
* `SIGNED` - signed
* `BASIC_API_KEY` - basic_api_key
* `BASIC_TOKEN` - basic_token

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.AuthenticationMethodEnumNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationMethodEnum("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AuthenticationMethodEnumNone`        | NONE                                  |
| `AuthenticationMethodEnumToken`       | TOKEN                                 |
| `AuthenticationMethodEnumSigned`      | SIGNED                                |
| `AuthenticationMethodEnumBasicAPIKey` | BASIC_API_KEY                         |
| `AuthenticationMethodEnumBasicToken`  | BASIC_TOKEN                           |