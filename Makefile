clear:
	rm tracker tasks.json

build:
	go build -o tracker cmd/main.go

test:
	go test -v -count=1 ./...

test100:
	go test -v -count=100 ./...

.PHONY: gen
gen:
	mockgen -source=internal/storage/storage.go \
	-destination=internal/storage/mocks/mock_storage.go

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
