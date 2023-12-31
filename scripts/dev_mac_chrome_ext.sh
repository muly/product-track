#!/bin/sh
set -x
set -e

source ./secrets/dev.sh
export CHROME_EXT_VERSION=$1 
envsubst < chrome-exten/manifest.json.tmpl > chrome-exten/manifest.json
zip chrome-exten.zip chrome-exten/*

