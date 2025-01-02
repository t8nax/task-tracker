test:
	go test -v -count=1 ./...

.PHONY: gen
gen:
	mockgen -source=internal/storage/storage.go \
	-destination=internal/storage/mocks/mock_storage.go

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
