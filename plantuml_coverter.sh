#!/usr/bin/env bash

set -x

for md_file in "$@"; do

  plantuml-converter "$md_file"
  return_val="$?"

  if [ $return_val -ne 0 ]; then
    exit $return_val
  fi

done

exit 0
