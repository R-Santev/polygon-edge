# Build procedure

Information about building and using the application

## Build production version for:

1. MacOS ARM64

```
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64  go build -o polygon-edge -a -installsuffix cgo  main.go
```

2. Linux

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o polygon-edge -a -installsuffix cgo  main.go
```

## Move to path

1. Linux

```
sudo mv polygon-edge /usr/local/bin
```
