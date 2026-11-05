#!/usr/bin/env bash
# Build the pinned open-source Speakeasy generator into a local cache.
set -euo pipefail

GENERATOR_REPO="https://github.com/speakeasy-api/openapi-generation.git"
GENERATOR_VERSION="v2.959.1"
GENERATOR_COMMIT="0be1696ac92bf925b616669fc50743032748f17c"

generator_dir="${GENERATOR_DIR:-${XDG_CACHE_HOME:-$HOME/.cache}/openapi-generation/$GENERATOR_VERSION}"

if [[ -x "$generator_dir/generate" && -x "$generator_dir/overlay" ]]; then
  exit 0
fi

clone_dir="$(mktemp -d)"
trap 'rm -rf "$clone_dir"' EXIT

git clone --quiet --depth 1 --branch "$GENERATOR_VERSION" "$GENERATOR_REPO" "$clone_dir"
actual_commit="$(git -C "$clone_dir" rev-parse HEAD)"
if [[ "$actual_commit" != "$GENERATOR_COMMIT" ]]; then
  echo "Tag $GENERATOR_VERSION is at $actual_commit, expected $GENERATOR_COMMIT" >&2
  exit 1
fi
mkdir -p "$generator_dir"
(cd "$clone_dir" && go build -o "$generator_dir/generate" ./cmd/generate)
(cd "$clone_dir" && go build -o "$generator_dir/overlay" ./cmd/overlay)
