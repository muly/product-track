build:
	go build -o product-track 

run:
	go run main.go scraping.go util.go api.go 

test:
	go test ./...

lint:
	go fmt ./...
	go mod verify

deploy_local:
	./scripts/local.sh 

deploy_dev:
	gcloud app deploy
