Build production version for:

1. MacOS ARM64

```
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64  go build -o polygon-edge -a -installsuffix cgo  main.go
```

2. Linux

```
CGO_ENABLED=0 GOOS=linux go build -o polygon-edge -a -installsuffix cgo  main.go
```
