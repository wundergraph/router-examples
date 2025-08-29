# Simple Module Example

A comprehensive example demonstrating advanced Cosmo Router customization with custom modules and middleware.

## 🔧 What This Example Demonstrates

- **Custom Module Creation**: Complete module implementation with configuration validation
- **Request/Response Interceptors**: Handle both origin requests and responses
- **GraphQL Operation Context**: Access operation details (name, type, hash, content)
- **Middleware Chain**: Multiple middleware handlers with proper execution order
- **Configuration Management**: YAML-based module configuration with validation
- **Request Context**: Share data between handlers and access logging
- **Header Manipulation**: Add custom headers to origin requests

## 📁 Project Structure

```
examples/simple/
├── main.go                 # Router entry point with module imports
├── myModule/
│   └── module.go          # Custom module implementation
├── config.yaml            # Router & module configuration
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
└── Dockerfile             # Multi-arch container build configuration
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- Docker (for containerized deployment)

### Build and Run

1. **Build the Router**

   ```bash
   docker build \
   --platform linux/amd64 \
   --build-arg TARGETOS=linux \
   --build-arg TARGETARCH=amd64 \
   --build-arg VERSION=$(git describe --tags --always --dirty) \
   --build-arg COMMIT=$(git rev-parse HEAD) \
   --build-arg DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
   -t myrouter:latest .
   ```

   The dockerfile is multi-arch compatible. For different architectures, replace `TARGETOS` and `TARGETARCH` accordingly (e.g., `darwin` and `arm64` for Mac).

2. **Run the Router**

   ```bash
   docker run \
   --name myrouter \
   --rm \
   -p 3002:3002 \
   -e DEV_MODE=true \
   -e DEMO_MODE=true \
   -e LISTEN_ADDR=0.0.0.0:3002 \
   myrouter:latest
   ```

   Visit [http://localhost:3002/](http://localhost:3002/) to see the router in action.

   > **Note**: For production use, connect the router to WunderGraph Cloud. Follow the [Cosmo Cloud Onboarding](https://cosmo-docs.wundergraph.com/getting-started/cosmo-cloud-onboarding) guide.

## 🔍 Module Deep Dive

### MyModule Implementation

The `myModule` demonstrates several key router extension patterns:

#### Configuration Validation

```go
func (m *MyModule) Provision(ctx *core.ModuleContext) error {
    if m.Value == 0 {
        return fmt.Errorf("value must be greater than 0")
    }
    return nil
}
```

#### Request/Response Interception

- **OnOriginRequest**: Modify outgoing requests to GraphQL origins
- **OnOriginResponse**: Handle responses from GraphQL origins
- **Middleware**: GraphQL operation-aware middleware
- **RouterOnRequest**: General HTTP request middleware

#### Context Sharing

```go
// Set data in one handler
ctx.Set("myValue", "myValue")

// Access in another handler
value := ctx.GetString("myValue")
```

#### GraphQL Operation Access

```go
operation := ctx.Operation()
fmt.Println(
    operation.Name(),    // Operation name
    operation.Type(),    // Query/Mutation/Subscription
    operation.Hash(),    // Operation hash
    operation.Content(), // Full operation string
)
```

## ⚙️ Configuration

The `config.yaml` demonstrates module configuration:

```yaml
version: "1"

modules:
  myModule:
    value: 10 # Custom module configuration
```

Configuration is automatically mapped to struct fields using `mapstructure` tags.

## 🔄 Customization

To create your own module:

1. **Implement Core Interfaces**: Choose from available handler interfaces based on your needs
2. **Add Configuration**: Define struct fields with `mapstructure` tags
3. **Register Module**: Import your module in `main.go`
4. **Configure**: Add module configuration to `config.yaml`

### Available Handler Interfaces

- `core.RouterMiddlewareHandler` - GraphQL operation middleware
- `core.RouterOnRequestHandler` - HTTP request middleware
- `core.EnginePreOriginHandler` - Pre-origin request handling
- `core.EnginePostOriginHandler` - Post-origin response handling
- `core.Provisioner` - Module initialization
- `core.Cleaner` - Module cleanup

## 📚 Additional Resources

- [Cosmo Router Documentation](https://cosmo-docs.wundergraph.com/router)
- [Router Configuration Reference](https://cosmo-docs.wundergraph.com/router/configuration)
- [Main Repository](https://github.com/wundergraph/router-examples)
