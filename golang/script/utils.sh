#!/usr/bin/env bash
# set -eo pipefail

# Add the provided directory(ies) to the system path if they are not already
# present.
add_to_path() {
  for d; do
    d=$(cd -- "$d" && { pwd -P || pwd; }) 2>/dev/null  # canonicalize symbolic links
    if [ -z "$d" ]; then continue; fi  # skip nonexistent directory
    case ":$PATH:" in
      *":$d:"*) :;;
      *) export PATH=$d:$PATH;;
    esac
  done
}

# Add the provided directory(ies) to the library path if they are not already
# present.
add_to_lib_path() {
  for d; do
    d=$(cd -- "$d" && { pwd -P || pwd; }) 2>/dev/null  # canonicalize symbolic links
    if [ -z "$d" ]; then continue; fi  # skip nonexistent directory
    case ":$LD_LIBRARY_PATH:" in
      *":$d:"*) :;;
      *) export LD_LIBRARY_PATH=$d:$LD_LIBRARY_PATH;;
    esac
  done
}

# Add the provided directory(ies) to the man path if they are not already
# present.
add_to_man_path() {
  for d; do
    d=$(cd -- "$d" && { pwd -P || pwd; }) 2>/dev/null  # canonicalize symbolic links
    if [ -z "$d" ]; then continue; fi  # skip nonexistent directory
    case ":$MANPATH:" in
      *":$d:"*) :;;
      *) export MANPATH=$d:$MANPATH;;
    esac
  done
}
