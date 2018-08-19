DEBUG_FLAG = $(if $(DEBUG),-debug)

GO=go1.11rc1

goast.wasm: *.go
	$(GO) get -d -t ./...
	GOOS=js GOARCH=wasm $(GO) build -o goast.wasm *.go
