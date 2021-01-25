#!/bin/bash

NETLIFY_CACHE_DIR="${NETLIFY_BUILD_BASE}/cache"
BAZEL="${NETLIFY_CACHE_DIR}/bazelisk-linux-amd64"

if [ ! -e "$BAZEL" ]; then
  (cd "$NETLIFY_CACHE_DIR" && curl -LO "https://github.com/bazelbuild/bazelisk/releases/download/v1.1.0/bazelisk-linux-amd64")
  chmod +x "$BAZEL"
fi
"$BAZEL" --batch build --disk_cache="${NETLIFY_CACHE_DIR}/bazel" //draftcode.osak.jp:final_layout
