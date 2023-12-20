#!/bin/sh
set -x
set -e

export CLIENT_ID="" 
export CHROME_EXT_VERSION=$1

envsubst < chrome-exten/manifest.json.tmpl > chrome-exten/manifest.json
zip chrome-exten.zip chrome-exten/*
