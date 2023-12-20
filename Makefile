build:
	go build -o product-track 

test:
	go test ./... --cover

lint:
	go mod verify

deploy_local_windows:
	go mod vendor
	./scripts/local_windows.sh

deploy_local_mac:
	go mod vendor
	./scripts/local_mac.sh

prepare_chrome_ext_for_local_deployment:
	./scripts/local_mac_chrome_ext.sh 1.1.1
	
prepare_chrome_ext_for_dev_deployment:
	./scripts/dev_mac_chrome_ext.sh 1.1.1

deploy_dev:
	go mod vendor
	gcloud config set project smuly-test-ground
	gcloud app deploy app.yaml --quiet --stop-previous-version
	gcloud app deploy cron.yaml --quiet 