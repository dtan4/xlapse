#!/bin/bash

TARBALL_NAME="xlapse_linux_amd64.tar.gz"

create_tarball() {
  echo "[INFO] Creating tarball"

  rm -rf dist
  mkdir -p dist

  while read -r function; do
    cp "bazel-bin/function/${function}/${function}_/${function}" "dist/xlapse-${function}"
  done < <(ls function)

  pushd dist > /dev/null 2>&1
  tar czvf "${TARBALL_NAME}" -- *
  popd > /dev/null 2>&1
  mv "dist/${TARBALL_NAME}" .
}

upload_tarball_to_github_release() {
  echo "[INFO] Uploading tarball to GitHub Release"

  gh release create "${GITHUB_REF}" "${TARBALL_NAME}"
}

main() {
  set -eu
  set -o pipefail

  create_tarball
  upload_tarball_to_github_release
}

main
