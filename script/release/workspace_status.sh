#!/bin/bash

# Stamping with the workspace status script (replacement of -ldflags)
# https://github.com/bazelbuild/rules_go/blob/70b8365a90e226b98f66bc35d3abb3e335883d81/go/core.rst#stamping-with-the-workspace-status-script

main() {
  set -eu
  set -o pipefail

  # https://docs.github.com/en/actions/reference/environment-variables
  echo RELEASE_VERSION "$(printf "%s" "${GITHUB_REF##*/}" | sed -E "s/^v//")"
  echo RELEASE_COMMIT "${GITHUB_SHA}"
  echo RELEASE_DATE "$(TZ=UTC date "+%Y-%m-%dT%H:%M:%SZ")"
}

main
