build:
	go build -o product-track 

run:
	go run main.go scraping.go util.go api.go 

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
	gcloud app deploy