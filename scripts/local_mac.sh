#!/bin/sh
set -x
set -e
emulatorPort="8090"

go mod vendor

go get ./...

go fmt ./...

source ./secrets/gmail_pwd.sh
if [[ -z "$GMAIL_PASSWORD" ]]; then
    echo "Must provide GMAIL_PASSWORD in environment" 1>&2
    exit 1
fi

rm -f product-track 
go build -o product-track

if [ $(lsof -i:$emulatorPort -t) ]
then
    kill -9 $(lsof -i:$emulatorPort -t);
fi

gcloud beta emulators firestore start --host-port=localhost:$emulatorPort &
sleep 10

PORT="8006" GCP_PROJECT="pt-emulator-project" FIRESTORE_EMULATOR_HOST="localhost:$emulatorPort" bash -c './product-track'
