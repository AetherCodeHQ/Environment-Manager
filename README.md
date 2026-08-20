# Environment Manager

![CI](https://github.com/Qyroxen/Environment-Manager/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Environment-Manager/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Environment-Manager?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Environment-Manager)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Environment-Manager)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Environment-Manager?style=social)](https://github.com/Qyroxen/Environment-Manager/stargazers)

## What is it?

Environment Manager is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Environment-Manager.git
cd Environment-Manager
go build -o environmentmanager .

# Run
./environmentmanager --help
```

## CLI Usage

```bash
# Basic usage
./environmentmanager

# With flags
./environmentmanager --verbose --output json

# Get help
./environmentmanager --help
```

## Examples

```bash
# Example 1
./environmentmanager example1

# Example 2
./environmentmanager example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o environmentmanager .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Environment-Manager/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Environment-Manager?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Environment-Manager/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Environment-Manager?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Environment-Manager/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Environment-Manager" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Environment-Manager/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Environment-Manager" alt="Pull Requests">
  </a>
</p>
