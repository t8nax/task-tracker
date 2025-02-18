clear:
	rm tracker tasks.json

build:
	go build -o tracker cmd/main.go

test:
	go test -v -count=1 ./...

test100:
	go test -v -count=100 ./...

.PHONY: gen
gen-repo:
	mockgen -source=internal/task/repository.go \
	-destination=internal/tsak/repository/mock_repository.go

.PHONY: gen-ucase
gen-ucase:
	mockgen -source=internal/task/usecase.go \
	-destination=internal/task/usecase/mock_usecase.go

.PHONY: cover
cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
