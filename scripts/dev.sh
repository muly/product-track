#!/bin/sh
set -x
set -e

source ./scripts/common.sh

go mod vendor

generate_app_yml

gcloud config set project smuly-test-ground
gcloud app deploy app.yaml --quiet --stop-previous-version
gcloud app deploy cron.yaml --quiet 