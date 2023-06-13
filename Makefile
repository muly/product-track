build:
	go build -o product-track 

run:
	go run main.go scraping.go util.go api.go 

test:
	go test ./...

lint:
	go fmt ./...
	go mod verify

deploy:
	# TODO: add commands to deploy to gcp
