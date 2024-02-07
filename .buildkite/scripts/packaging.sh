#!/usr/bin/env bash

source .buildkite/scripts/install_tools.sh

set -euo pipefail

echo "--- Run Packaging for $BEATS_PROJECT_NAME"
pushd "${BEATS_PROJECT_NAME}" > /dev/null

mage package

popd > /dev/null
