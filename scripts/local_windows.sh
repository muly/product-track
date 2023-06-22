#!/bin/sh
set -x
set -e
emulatorPort="8090"

 go get ./...

 go fmt ./...

 rm -f product-track 
 go build -o product-track
 
if [[ $(netstat -ano -p tcp |grep  $emulatorPort) ]]  
then
    tskill $(netstat -ano -p tcp | grep $emulatorPort |  xargs | tr ' ' '\n' | tail -1)
fi

 gcloud beta emulators firestore start --host-port=localhost:$emulatorPort &
 sleep 10 

 PORT="8006" GCP_PROJECT="pt-emulator-project" FIRESTORE_EMULATOR_HOST="localhost:$emulatorPort" bash -c './product-track'
