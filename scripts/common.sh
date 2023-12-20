#!/bin/sh
set -x
set -e


function generate_app_yml(){
    # generate app.yaml from template
    export PROJECT_NUMBER="149500152182" # project number, not the project id
    export COMMIT_HASH=$(git rev-parse HEAD)
    envsubst < app.yaml.tmpl > app.yaml
}
