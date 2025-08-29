<p align="center">
  <a href="https://github.com/wundergraph/router-examples">
    <img src="img.png" width="500px" alt="gRPC Plugin Demo" />
  </a>
</p>

<p align="center">Customizable examples for extending the Cosmo Router with custom modules and middleware.</p>

<p align="center">
  <a href="https://cosmo-docs.wundergraph.com/router">Router Documentation</a> •
  <a href="https://cosmo-docs.wundergraph.com/router/custom-modules">Module Configuration</a>
</p>

## 🚀 Quick Start

Get your custom Cosmo Router running in under 5 minutes! Choose your preferred approach below.

### Prerequisites

```bash
# Clone and navigate to the example
git clone https://github.com/wundergraph/router-examples.git
cd router-examples/examples/simple
```

### Choose Your Method

Next step is to build and run the router. You can run these commands in every example directory.

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

## 📁 Examples

### Simple Module

> **Path:** [`examples/simple/`](examples/simple/) | **[Documentation](examples/simple/README.md)**

Comprehensive example showcasing advanced router customization:

- Custom module creation with configuration validation
- Request/response interceptors and middleware chains
- GraphQL operation handling and context management
- YAML-based configuration with live reloading

## 🔄 Upgrade Router

**Update to the latest version:**

1. Get the commit SHA from [releases page](https://github.com/wundergraph/cosmo/releases?q=router%40&expanded=false)
2. Navigate to your router directory (where `go.mod` exists)
3. Update dependencies:
   ```bash
   cd examples/simple  # your example directory
   go get github.com/wundergraph/cosmo/router@<commit-sha>
   ```

> ⚠️ **Important:** Run `go get` from the directory containing `go.mod`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit examples, improvements, and bug fixes.

## License

Cosmo is licensed under the [Apache License, Version 2.0](LICENSE).
