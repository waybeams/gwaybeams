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

export PATH=$PATH:$BASEDIR/script

# Export the GOPATH
# export GOPATH=$BASEDIR/vendor:$BASEDIR
# export PATH=$PATH:$BASEDIR/vendor/bin

