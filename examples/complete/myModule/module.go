package myModule

import (
	"fmt"
	"net/http"

	"github.com/wundergraph/cosmo/router/core"
	"go.uber.org/zap"
)

func init() {
	// Register your module here
	core.RegisterModule(&MyModule{})
}

const myModuleID = "myModule"

// MyModule demonstrates GraphQL operation access and header manipulation across different router handlers.
// Config values are auto-populated from config.yaml under `modules.<name>` using mapstructure tags.
type MyModule struct {
	Value uint64 `mapstructure:"value"` // Config value from modules.myModule.value

	Logger *zap.Logger
}

// Provision validates config and initializes resources during module startup
func (m *MyModule) Provision(ctx *core.ModuleContext) error {
	if m.Value == 0 {
		ctx.Logger.Error("Value must be greater than 0")
		return fmt.Errorf("value must be greater than 0")
	}

	m.Logger = ctx.Logger // For non-request related logging

	return nil
}

// Cleanup releases resources and closes connections during module shutdown
func (m *MyModule) Cleanup() error {
	return nil
}

// RouterOnRequest handles early request processing before authentication and tracing
func (m *MyModule) RouterOnRequest(ctx core.RequestContext, next http.Handler) {
	logger := ctx.Logger()
	logger.Info("Test RouterOnRequest custom module logs")

	next.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

// OnOriginResponse processes subgraph responses for logging, caching, and modifications
func (m *MyModule) OnOriginResponse(response *http.Response, ctx core.RequestContext) *http.Response {
	// Access the custom value set in OnOriginRequest
	value := ctx.GetString("myValue")

	fmt.Println("SharedValue", value)
	fmt.Println("OnOriginResponse", response.Request.URL, response.StatusCode)

	return nil // Pass response to next handler
}

// OnOriginRequest modifies subgraph requests by adding headers and sharing context data
func (m *MyModule) OnOriginRequest(request *http.Request, ctx core.RequestContext) (*http.Request, *http.Response) {
	request.Header.Set("myHeader", ctx.GetString("myValue")) // Add custom header
	ctx.Set("myValue", "myValue")                            // Share data with OnOriginResponse

	return request, nil
}

// Middleware provides access to GraphQL operation details for logging and validation
func (m *MyModule) Middleware(ctx core.RequestContext, next http.Handler) {
	operation := ctx.Operation()

	logger := ctx.Logger()
	logger.Info("Test custom module logs")

	// Access GraphQL operation details
	fmt.Println(
		operation.Name(),
		operation.Type(),
		operation.Hash(),
		operation.Content(),
	)

	//
	// Return a GraphQL error response to the client.
	//
	// Never write errors directly to the http.ResponseWriter
	// because they will not be logged and tracked in the telemetry system.
	// or rendered as a GraphQL error response.
	//
	//core.WriteResponseError(ctx, fmt.Errorf("test error"))

	next.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
}

// Module returns registration info with unique ID and execution priority
func (m *MyModule) Module() core.ModuleInfo {
	return core.ModuleInfo{
		ID:       myModuleID, // Unique module identifier
		Priority: 1,          // Lower numbers execute first
		New: func() core.Module {
			return &MyModule{}
		},
	}
}

// Interface guard
var (
	_ core.RouterMiddlewareHandler = (*MyModule)(nil)
	_ core.RouterOnRequestHandler  = (*MyModule)(nil)
	_ core.EnginePreOriginHandler  = (*MyModule)(nil)
	_ core.EnginePostOriginHandler = (*MyModule)(nil)
	_ core.Provisioner             = (*MyModule)(nil)
	_ core.Cleaner                 = (*MyModule)(nil)
)
