# Universal Manifest Sync CLI

Cross-platform tool that scans development environments and generates manifests for:
- Docker
- winget
- vcpkg
- Homebrew
- apt
- pip
- npm

## Installation
```bash
go install github.com/yourusername/universal-manifest-sync-cli@latest
```

## Usage
```bash
# Scan environment
ums scan

# Generate manifest
ums generate docker  # or winget, vcpkg, brew, etc.
```
