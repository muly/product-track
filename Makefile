build:
	go build -o product-track 

test:
	go test ./... --cover

lint:
	go fmt ./...
	go mod verify

deploy_local_windows:
	./scripts/local_windows.sh

deploy_local_mac:
	./scripts/local_mac.sh

deploy_dev:
	gcloud config set project smuly-test-ground
	gcloud app deploy app.yaml --quiet --stop-previous-version
	gcloud app deploy cron.yaml --quiet 