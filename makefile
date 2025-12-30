# Run the application
run:
	go run ./main.go
# compile for execution
bundle:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags="-s -w" -trimpath -o bosh_admin_linux ./main.go

# Smaller binary with UPX compression (requires UPX to be installed)
bundle-tiny: bundle
	upx --best --lzma bosh_admin_linux