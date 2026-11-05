#!/usr/bin/env bash
# Replace .speakeasy/openapi.yaml with the current Cloudsmith v2 OpenAPI document.
set -euo pipefail

SPEC_URL="https://api.cloudsmith.io/v2/openapi/"

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

curl -sSfL --max-time 60 -o "$root_dir/.speakeasy/openapi.yaml" "$SPEC_URL"
