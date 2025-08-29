# Complete Module Example

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
.
├── main.go                 # Router entry point with module imports
├── myModule/
│   └── module.go          # Custom module implementation
├── config.yaml            # Router & module configuration
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency checksums
└── Dockerfile             # Multi-arch container build configuration
```

## 🚀 Getting Started

Copy the example directory to your own repository.

### Build and Run

Next step is to build and run the router. Choose your preferred method below.

<details>
<summary><strong>🐳 Docker (Recommended)</strong></summary>

<br/>

**Build the image:**

```bash
docker build -t myrouter:latest .
```

**Run the router:**

```bash
docker run --name myrouter --rm -p 3002:3002 \
  -e DEV_MODE=true \
  -e DEMO_MODE=true \
  -e LISTEN_ADDR=0.0.0.0:3002 \
  myrouter:latest
```

> 💡 **Multi-arch builds:** The Dockerfile supports multiple architectures. For custom builds, set `TARGETOS` and `TARGETARCH` (e.g., `darwin/arm64` for Mac M1).

</details>

<details>
<summary><strong>Go Direct</strong></summary>

<br/>

**Requirements:** Go 1.24+

**Install & Run:**

```bash
go mod download
DEV_MODE=true DEMO_MODE=true LISTEN_ADDR=0.0.0.0:3002 go run main.go
```

</details>

---

### Test Your Router

Visit **[localhost:3002](http://localhost:3002)** to see your router in action!

**Next Steps:** Connect to [WunderGraph Cloud](https://cosmo-docs.wundergraph.com/getting-started/cosmo-cloud-onboarding) for production configuration.
