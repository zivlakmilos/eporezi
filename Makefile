VERSION=`git describe --tags --match v[0-9]* 2> /dev/null`
EXE=ePorezi

.PHONY: all
all: run

.PHONY: run
run: build
	@./build/$(EXE)

.PHONY: build
build:
	@go build -o build/$(EXE) ./cmd/eporezi/main.go

.PHONY: dist
dist: dist-linux

.PHONY: dist-linux
dist-linux: dist-linux-amd64

.PHONY: dist-linux-amd64
dist-linux-amd64:
	@rm -f build/$(EXE) && GOARCH=amd64 GOOS=linux go build -o build/$(EXE) ./cmd/eporezi/main.go && tar -czvf "gophoria-${VERSION}-linux-amd64.tar.gz" build/$(EXE)

.PHONY: clean
clean:
	@rm -Rf build/*
