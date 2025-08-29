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

Follow the [main repository instructions](../../README.md#-quick-start) to build and run this example.
