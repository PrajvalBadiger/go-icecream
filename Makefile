TARGET=go-icecream
MAIN=./cmd/icecream/main.go

build:
	@go build -o bin/${TARGET} ./cmd/icecream/main.go

run: build
	@./bin/${TARGET}

watch:
	@gow -c run ${MAIN}

test:
	@go test -v ./...
