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
	./scripts/local_mac_chrome_ext.sh

prepare_chrome_ext_for_dev_deployment:
	CLIENT_ID="" CHROME_EXT_VERSION=1.1.0 envsubst < chrome-exten/manifest.json.tmpl > chrome-exten/manifest.json
	zip chrome-exten.zip chrome-exten/*

deploy_dev:
	go mod vendor
	gcloud config set project smuly-test-ground
	gcloud app deploy app.yaml --quiet --stop-previous-version
	gcloud app deploy cron.yaml --quiet 