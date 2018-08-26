DEBUG_FLAG = $(if $(DEBUG),-debug)

GO=go

goast.wasm: *.go
	$(GO) get -d -t ./...
	GOOS=js GOARCH=wasm $(GO) build -o goast.wasm *.go
