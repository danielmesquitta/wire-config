# Wire Config

A Go tool for managing dependency injection configurations across different environments using Google's Wire framework.

## Overview

Wire Config is a utility that helps you manage multiple dependency injection configurations for different environments (Development, Staging, Testing, Production) in your Go applications. It works in conjunction with Google's Wire framework to generate environment-specific dependency injection code.

## Features

- Support for multiple environments (Dev, Staging, Test, Prod)
- Automatic generation of environment-specific wire configurations
- Merging of default providers with environment-specific providers
- Clean and maintainable configuration management
- Go code formatting for generated files

## Installation

```bash
go install github.com/danielmesquitta/wire-config@latest
```

## Usage

1. Create a `wire_config.go` file in your project with your dependency configurations:

```go
package main

import (
    "github.com/google/wire"
    // your other imports
)

var providers = []any{
    // your default providers
}

var devProviders = []any{
    // development-specific providers
}

var prodProviders = []any{
    // production-specific providers
}

func params() (
    // your parameter declarations
) {
    // parameter initialization
}
```

2. Run the wire-config tool to generate the wire.go file:

```bash
wire-config -c wire_config.go -o wire.go -p your/package/import/path
```

## Command Line Options

- `-c, --config`: Path to the wire configuration file (default: "wire_config.go")
- `-o, --output`: Path to the output wire.go file (default: "wire.go")
- `-p, --package`: Main package import path
- `-e, --environments`: Comma-separated list of environment names (default: "Prod")

## Example

Given a configuration file:

```go
// wire_config.go
package main

import "github.com/google/wire"

var providers = []any{
    NewDatabase,
    NewLogger,
}

var devProviders = []any{
    NewDevDatabase,
}

var prodProviders = []any{
    NewProdDatabase,
}

func params() (
    dbConfig *Config,
    logLevel string,
) {
    // parameter initialization
}
```

Running:

```bash
wire-config -c wire_config.go -o wire.go -p github.com/your/package -e Dev,Prod
```

Will generate a wire.go file with environment-specific dependency injection functions.
