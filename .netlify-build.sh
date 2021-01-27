#!/bin/bash

if [ -n "${NETLIFY_BUILD_BASE}" ]; then
  NETLIFY_CACHE_DIR="${NETLIFY_BUILD_BASE}/cache"
  BAZEL="${NETLIFY_CACHE_DIR}/bazelisk-linux-amd64"
  BAZEL_OPT=(--output_user_root="${NETLIFY_CACHE_DIR}/bazel_user_root")

  if [ ! -e "$BAZEL" ]; then
    (cd "$NETLIFY_CACHE_DIR" && curl -LO "https://github.com/bazelbuild/bazelisk/releases/download/v1.1.0/bazelisk-linux-amd64")
    chmod +x "$BAZEL"
  fi
else
  BAZEL=bazel
  BAZEL_OPT=()
fi
"$BAZEL" "${BAZEL_OPT[@]}" build //draftcode.osak.jp:final_layout
