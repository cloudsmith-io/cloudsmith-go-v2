package hooks

import (
	"context"
	"os"
	"strings"

	"github.com/cloudsmith-io/cloudsmith-go-v2/internal/config"
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

const (
	envAPIHost = "CLOUDSMITH_API_HOST"
	envAPIKey  = "CLOUDSMITH_API_KEY"
)

// envConfigHook applies environment-variable defaults at SDK initialization.
//
// Precedence is: explicit option (WithServerURL / WithSecurity) > environment
// variable > SDK default. The hook therefore only applies an env value when
// the corresponding config field is still at its zero value (server URL
// empty, security source nil).
//
//   - CLOUDSMITH_API_HOST, if set and WithServerURL was not used, overrides
//     the server URL. Trailing slashes are normalised.
//   - CLOUDSMITH_API_KEY, if set and WithSecurity / WithSecuritySource was
//     not used, installs a security source backed by the env var.
type envConfigHook struct{}

func (envConfigHook) SDKInit(c config.SDKConfiguration) config.SDKConfiguration {
	if c.ServerURL == "" {
		if host := strings.TrimRight(os.Getenv(envAPIHost), "/"); host != "" {
			c.ServerURL = host + "/"
		}
	}

	if c.Security == nil {
		if key := os.Getenv(envAPIKey); key != "" {
			c.Security = func(_ context.Context) (interface{}, error) {
				return components.Security{Apikey: key}, nil
			}
		}
	}

	return c
}
