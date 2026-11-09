# DomainType

The purpose the domain serves. A native API type names the package format, and the interface for a format served over OCI (e.g. `NATIVE_HELM_OCI`).

* `DOWNLOAD` - Downloads
* `UPLOAD` - Uploads
* `CLOUDSMITH_API` - Cloudsmith API
* `NATIVE_SKILLS` - Native API: AI Skills
* `NATIVE_CARGO` - Native API: Cargo
* `NATIVE_COMPOSER` - Native API: Composer
* `NATIVE_CONAN` - Native API: Conan
* `NATIVE_CONDA` - Native API: Conda
* `NATIVE_DART` - Native API: Dart
* `NATIVE_DOCKER` - Native API: Docker
* `NATIVE_GENERIC` - Native API: Generic
* `NATIVE_GO` - Native API: Go
* `NATIVE_HELM_OCI` - Native API: Helm (OCI)
* `NATIVE_HEX` - Native API: Hex
* `NATIVE_HUGGINGFACE` - Native API: HuggingFace
* `NATIVE_MAVEN` - Native API: Maven
* `NATIVE_NIX` - Native API: Nix
* `NATIVE_NPM` - Native API: npm
* `NATIVE_NUGET` - Native API: NuGet
* `NATIVE_PYTHON` - Native API: Python
* `NATIVE_RUBY` - Native API: Ruby
* `NATIVE_SWIFT` - Native API: Swift
* `NATIVE_TERRAFORM` - Native API: Terraform
* `NATIVE_VSX` - Native API: VSX

## Example Usage

```go
import (
	"github.com/cloudsmith-io/cloudsmith-go-v2/models/components"
)

value := components.DomainTypeDownload

// Open enum: custom values can be created with a direct type cast
custom := components.DomainType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `DomainTypeDownload`          | DOWNLOAD                      |
| `DomainTypeUpload`            | UPLOAD                        |
| `DomainTypeCloudsmithAPI`     | CLOUDSMITH_API                |
| `DomainTypeNativeSkills`      | NATIVE_SKILLS                 |
| `DomainTypeNativeCargo`       | NATIVE_CARGO                  |
| `DomainTypeNativeComposer`    | NATIVE_COMPOSER               |
| `DomainTypeNativeConan`       | NATIVE_CONAN                  |
| `DomainTypeNativeConda`       | NATIVE_CONDA                  |
| `DomainTypeNativeDart`        | NATIVE_DART                   |
| `DomainTypeNativeDocker`      | NATIVE_DOCKER                 |
| `DomainTypeNativeGeneric`     | NATIVE_GENERIC                |
| `DomainTypeNativeGo`          | NATIVE_GO                     |
| `DomainTypeNativeHelmOci`     | NATIVE_HELM_OCI               |
| `DomainTypeNativeHex`         | NATIVE_HEX                    |
| `DomainTypeNativeHuggingface` | NATIVE_HUGGINGFACE            |
| `DomainTypeNativeMaven`       | NATIVE_MAVEN                  |
| `DomainTypeNativeNix`         | NATIVE_NIX                    |
| `DomainTypeNativeNpm`         | NATIVE_NPM                    |
| `DomainTypeNativeNuget`       | NATIVE_NUGET                  |
| `DomainTypeNativePython`      | NATIVE_PYTHON                 |
| `DomainTypeNativeRuby`        | NATIVE_RUBY                   |
| `DomainTypeNativeSwift`       | NATIVE_SWIFT                  |
| `DomainTypeNativeTerraform`   | NATIVE_TERRAFORM              |
| `DomainTypeNativeVsx`         | NATIVE_VSX                    |