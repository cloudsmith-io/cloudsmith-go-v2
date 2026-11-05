#!/usr/bin/env bash
# Regenerate the SDK with the pinned open-source Speakeasy generator.
set -euo pipefail

OVERLAYS=(cloudsmith-v2 cloudsmith-retries cloudsmith-pagination cloudsmith-auth)

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
generator_version="$(sed -n 's/^GENERATOR_VERSION="\(.*\)"$/\1/p' "$root_dir/scripts/install-generator.sh")"
generator_dir="${GENERATOR_DIR:-${XDG_CACHE_HOME:-$HOME/.cache}/openapi-generation/$generator_version}"
work_dir="$(mktemp -d)"
trap 'rm -rf "$work_dir"' EXIT

# The generator sends usage telemetry unless this is set.
export SPEAKEASY_DISABLE_TELEMETRY=true

"$root_dir/scripts/install-generator.sh"

overlay_args=()
for overlay in "${OVERLAYS[@]}"; do
  overlay_args+=(-overlay "$root_dir/.speakeasy/overlays/$overlay.yaml")
done
"$generator_dir/overlay" -s "$root_dir/.speakeasy/openapi.yaml" "${overlay_args[@]}" -out "$work_dir/openapi.combined.yaml"

"$generator_dir/generate" -s "$work_dir/openapi.combined.yaml" -o "$root_dir" -l go --license agpl-3.0-only
