# PackageFormat

Format of the requested package.

* `DEB` - Debian
* `RPM` - RedHat
* `RUBY` - Ruby
* `PYTHON` - Python
* `MAVEN` - Maven
* `BOWER` - Bower
* `DOCKER` - Docker
* `RAW` - Raw
* `CHOCOLATEY` - Chocolatey
* `NPM` - npm
* `NUGET` - NuGet
* `VAGRANT` - Vagrant
* `COMPOSER` - Composer
* `ALPINE` - Alpine
* `HELM` - Helm
* `CONAN` - Conan
* `CARGO` - Cargo
* `LUAROCKS` - LuaRocks
* `CRAN` - CRAN
* `GO` - Go
* `DART` - Dart
* `COCOAPODS` - CocoaPods
* `TERRAFORM` - Terraform
* `P2` - P2
* `CONDA` - Conda
* `HEX` - Hex
* `SWIFT` - Swift
* `HUGGINGFACE` - HuggingFace
* `GENERIC` - Generic
* `VSX` - VSX
* `MCP` - MCP
* `NIX` - Nix
* `DEFAULT` - Default


## Supported Types

### LogsResponsePackageFormatEnum

```go
packageFormat := components.CreatePackageFormatLogsResponsePackageFormatEnum(components.LogsResponsePackageFormatEnum{/* values here */})
```

### 

```go
packageFormat := components.CreatePackageFormatAny(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch packageFormat.Type {
	case components.PackageFormatTypeLogsResponsePackageFormatEnum:
		// packageFormat.LogsResponsePackageFormatEnum is populated
	case components.PackageFormatTypeAny:
		// packageFormat.Any is populated
	default:
		// Unknown type - use packageFormat.GetUnknownRaw() for raw JSON
}
```
