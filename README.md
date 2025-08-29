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

Choose your example below.

### Complete Module

> **Path:** [`examples/complete/`](examples/complete/) | **[Documentation](examples/complete/README.md)**

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
   cd examples/complete  # your example directory
   go get github.com/wundergraph/cosmo/router@<commit-sha>
   ```

> ⚠️ **Important:** Run `go get` from the directory containing `go.mod`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit examples, improvements, and bug fixes.

## License

Cosmo is licensed under the [Apache License, Version 2.0](LICENSE).
