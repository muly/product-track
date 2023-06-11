build:
	go build -o product-track/product-track 
run:
	go run main.go scraping.go util.go api.go 
test:
	go test ./...