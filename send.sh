#/usr/bin/env bash

COLORS=$1
ENDPOINT="https://localhost:8123"
echo "${COLORS}" | tr ' ' ',' | tr -d '#' | xargs -I{} curl -XPOST -d"colors={}" "localhost:8123"

