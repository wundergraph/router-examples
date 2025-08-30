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

## 📁 Examples

Each example is a self-contained router configuration with a custom module. You can copy the example directory to your own repository and start customizing it.

| Example             | Path                                       | Documentation                            | Description                                                      |
| ------------------- | ------------------------------------------ | ---------------------------------------- | ---------------------------------------------------------------- |
| **Complete Module** | [`examples/complete/`](examples/complete/) | [📖 README](examples/complete/README.md) | Advanced router customization with custom modules and middleware |

## 🔌 Go Production

### Packaging and Running the Router

There are two way to produce a deployable artifact:

<details>
<summary><strong>🐳 Docker Image (Recommended)</strong></summary>

<br/>

**Build the image:**

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

**Multi-arch builds:** The Dockerfile supports multiple architectures. For custom builds, set `TARGETOS` and `TARGETARCH` (e.g., `darwin/arm64` for Mac M1).

**Run the router:**

```bash
docker run --name myrouter --rm -p 3002:3002 \
  -e LISTEN_ADDR=0.0.0.0:3002 \
  -e DEMO_MODE=true \
  myrouter:latest
```

Visit **[localhost:3002](http://localhost:3002)** to see your router in action!

</details>

<details>
<summary><strong>Go Binary (Requires Go 1.24+)</strong></summary>

<br/>

**Requirements:** Go 1.24+

**Build & Run:**

```bash
go mod download
go build -o router main.go && chmod +x router
./router
```

Visit **[localhost:3002](http://localhost:3002)** to see your router in action!

</details>

---

### Connecting to WunderGraph Cloud

The router requires a graph api token to serve your supergraph.
Follow the instructions in [WunderGraph Cloud Onboarding](https://cosmo-docs.wundergraph.com/getting-started/cosmo-cloud-onboarding) to get your token.

## 🔄 Upgrade Router

You can upgrade the router in each example directory by following the instructions below.

1. Get the commit SHA from [releases page](https://github.com/wundergraph/cosmo/releases?q=router%40&expanded=false)
2. Navigate to your router directory (where `go.mod` exists)
3. Update dependencies:
   ```bash
   cd examples/complete  # your example directory
   go get github.com/wundergraph/cosmo/router@<commit-sha>
   ```

> [!NOTE]
> Run `go get` from the directory containing `go.mod`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit examples, improvements, and bug fixes.

## License

Cosmo is licensed under the [Apache License, Version 2.0](LICENSE).
