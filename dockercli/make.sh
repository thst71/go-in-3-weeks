GOOS=darwin GOARCH=amd64 go build -o dockercli-darwin-amd64 dockercli.go
GOOS=darwin GOARCH=arm64 go build -o dockercli-darwin-arm64 dockercli.go

lipo dockercli-darwin-amd64 dockercli-darwin-arm64 -create -output dockercli-mac
