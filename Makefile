build:
	go build -o product-track 

test: # unit and functional tests
	go test ./... --cover

lint:
	go mod verify

deploy_local_windows:
	./scripts/local_windows.sh

deploy_local_mac:
	./scripts/local_mac.sh

deploy_dev:
	./scripts/dev.sh


prepare_chrome_ext_for_local_deployment:
	./scripts/local_mac_chrome_ext.sh 1.1.1
	
prepare_chrome_ext_for_dev_deployment:
	./scripts/dev_mac_chrome_ext.sh 1.1.3
