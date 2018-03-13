#!/usr/bin/env bash

# Run this in terminal with `source setup-env.sh.
if [ -n "${ZSH_VERSION}" ]; then
  BASEDIR="$( cd $( dirname "${(%):-%N}" ) && pwd )"
elif [ -n "${BASH_VERSION}" ]; then
  if [[ "$(basename -- "$0")" == "setup-env.sh" ]]; then
    echo "Don't run $0, source it (see README.md)" >&2
    exit 1
  fi
  BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
else
  echo "Unsupported shell, use bash or zsh."
  exit 2
fi

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source "${BASEDIR}/script/utils.sh"

# Add script to the path.
add_to_path "${BASEDIR}/script"
echo "Updated PATH with ${BASEDIR}/script"

# Add go lang bin to the path.
add_to_path "${BASEDIR}/lib/go-1.10/bin"
echo "Updated PATH with ${BASEDIR}/lib/go-1.10/bin"

# Add project bin to the path.
add_to_path "${BASEDIR}/bin/"
echo "Updated PATH with ${BASEDIR}/bin"

# Add vendor bin to the path.
add_to_path "${BASEDIR}/vendor/bin/"
echo "Updated PATH with ${BASEDIR}/vendor/bin"

# Add Google depot_tools to path.
add_to_path "${BASEDIR}/lib/depot_tools"
echo "Updated PATH with ${BASEDIR}/lib/depot_tools"

# Set the custom GOROOT value to the local Go installation.
export GOROOT="${BASEDIR}/lib/go-1.10"
echo "Set GOROOT to ${BASEDIR}/lib/go-1.10"

# Set the gopath for this project
export GOPATH="${BASEDIR}:${BASEDIR}/vendor"
echo "Set GOPATH to $GOPATH"

# Add skia SO file(s) to LD_LIBRARY_PATH
add_to_lib_path "${BASEDIR}/lib/skia/out/Shared"
echo "Updated LD_LIBRARY_PATH with ${BASEDIR}/lib/skia/out/Shared"

