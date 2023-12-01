build:
	go build -o product-track 

test:
	go test ./... --cover

lint:
	go mod verify

deploy_local_windows:
	./scripts/local_windows.sh

deploy_local_mac:
	./scripts/local_mac.sh

deploy_dev:
	./scripts/dev.sh
