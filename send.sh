#/usr/bin/env bash

NAMESPACE=$1
COLORS=$2
ENDPOINT="http://localhost:8123"
echo "${COLORS}" | tr ' ' ',' | tr -d '#' | xargs -I{} curl -XPOST -d"colors={}" "${ENDPOINT}"

