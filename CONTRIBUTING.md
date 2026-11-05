# Contributing to This Repository

Thank you for your interest in contributing to this repository. Please note that this repository contains generated code. As such, we do not accept direct changes or pull requests. Instead, we encourage you to follow the guidelines below to report issues and suggest improvements.

## How to Report Issues

If you encounter any bugs or have suggestions for improvements, please open an issue on GitHub. When reporting an issue, please provide as much detail as possible to help us reproduce the problem. This includes:

- A clear and descriptive title
- Steps to reproduce the issue
- Expected and actual behavior
- Any relevant logs, screenshots, or error messages
- Information about your environment (e.g., operating system, software versions)
    - For example can be collected using the `npx envinfo` command from your terminal if you have Node.js installed

## Issue Triage and Upstream Fixes

We will review and triage issues as quickly as possible. Our goal is to address bugs and incorporate improvements in the upstream source code. Fixes will be included in the next generation of the generated code.

## Regenerating the SDK

The `Generate SDK` workflow runs `scripts/generate.sh` on every push to `main`. The script:

1. Builds the pinned [open-source Speakeasy generator](https://github.com/speakeasy-api/openapi-generation) with `scripts/install-generator.sh` (tag and commit in that script).
2. Downloads the OpenAPI document and applies `.speakeasy/overlays/*.yaml` in strict mode. An overlay action that matches nothing fails the run.
3. Generates the SDK under the AGPL-3.0-only license election, with telemetry disabled.

The generator applies `.speakeasy/patches/<generated file>.patch` to each generated file that carries hand edits. It applies the patch before it compiles and lints the SDK.

To run it locally, install the tools with `mise install` and run `./scripts/generate.sh`.

If a patch does not apply, the generator writes conflict markers into the file and the run fails. To update the patch:

1. Move the patch file out of `.speakeasy/patches` and run the script again.
2. Stage the unpatched file with `git add <file>`.
3. Apply the hand edit again.
4. Write the new patch with `git diff -- <file> > .speakeasy/patches/<file>.patch`.

Hand-written files that the generator does not track (`internal/resolvers`, `internal/pagination`, `models/apierrors/status.go`, `internal/hooks/registration.go`, `cloudsmith_test.go`) do not need patches.

## Contact

If you have any questions or need further assistance, please feel free to reach out by opening an issue.

Thank you for your understanding and cooperation!

The Maintainers
