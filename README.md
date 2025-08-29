# Router Examples

Customizable examples for extending the [Cosmo Router](https://github.com/wundergraph/cosmo) with custom modules and middleware.

## 🚀 Quick Start

Get up and running with a custom Cosmo Router in minutes. This guide walks you through building and running a router with custom modules using Docker and connecting it to WunderGraph Cloud for configuration management. Each example can be copied and modified to create your own router.

1. **Clone the repository**

   ```bash
   git clone https://github.com/wundergraph/router-examples.git
   cd router-examples
   ```

2. **Navigate to an example**

   ```bash
   cd examples/simple
   ```

3. **Build the router**

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

   The dockerfile is multi-arch, so it will build for the correct architecture for your machine.
   If you want to build for a different architecture, you can replace `TARGETOS` and `TARGETARCH` with your operating system and architecture. For example, if you are on a Mac, you can use `darwin` and `arm64`.

4. **Run the router**

   After following the [Cosmo Cloud Onboarding](https://cosmo-docs.wundergraph.com/getting-started/cosmo-cloud-onboarding) guide, you generated a `GRAPH_API_TOKEN` that you pass to the router as an environment variable. This will download the execution configuration from the WunderGraph Cloud.

   ```bash
   docker run \
   --name cosmo-router \
   --rm \
   -p 3002:3002 \
   --add-host=host.docker.internal:host-gateway \
   --pull always \
   -e DEV_MODE=true \
   -e DEMO_MODE=true \
   -e LISTEN_ADDR=0.0.0.0:3002 \
   -e GRAPH_API_TOKEN=<graph-api-token> \
   ghcr.io/wundergraph/cosmo/router:latest
   ```

   Visit [http://localhost:3002/](http://localhost:3002/) to see the router in action.

## 📁 Examples

### Hello World

A comprehensive example demonstrating how to:

- Create custom router modules
- Handle GraphQL operations
- Add request/response middleware
- Configure modules via YAML
- Access request context and logging

**Location:** [`examples/simple/`](examples/simple/)

## 🔧 Creating Your Own Module

Each example includes:

- `main.go` - Router entry point with module imports
- `module.go` - Custom module implementation
- `config.yaml` - Module configuration
- `go.mod` - Go module dependencies

## 📚 Documentation

For comprehensive router documentation and configuration options, visit:

- [Cosmo Router Documentation](https://cosmo-docs.wundergraph.com/router)
- [Router Configuration Reference](https://cosmo-docs.wundergraph.com/router/configuration)

## 🔄 Upgrade

To upgrade to the latest router version:

1. Visit the [Cosmo Router releases page](https://github.com/wundergraph/cosmo/releases?q=router%40&expanded=false)
2. Copy the commit SHA from your desired release
3. **Navigate to your router directory** (where `go.mod` is located):
   ```bash
   cd examples/simple  # or your specific example directory
   ```
4. **Update your dependencies**:
   ```bash
   go get github.com/wundergraph/cosmo/router@<commit-sha>
   ```

> ⚠️ **Important:** Always run the `go get` command from within the directory containing your `go.mod` file.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit examples, improvements, and bug fixes.

## 📄 License

Cosmo is licensed under the [Apache License, Version 2.0](LICENSE).
