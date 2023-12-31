#!/bin/sh
set -x
set -e

source ./secrets/local.sh
export CHROME_EXT_VERSION=$1 
envsubst < chrome-exten/manifest.json.tmpl > chrome-exten/manifest.json
