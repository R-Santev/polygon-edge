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

## Build devnet node docker image

1. Build node source code

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o polygon-edge -a -installsuffix cgo  main.go
```

2. Build node image

```
docker build --platform linux/amd64 -t rsantev/polygon-edge:latest -f Dockerfile.release .
```

3. Push node image to DockerHub

Use Docker Desktop or:

```
docker push rsantev/polygon-edge:latest
```

4. Build hydrag devnet image

```
docker build --platform linux/amd64 -t rsantev/hydrag-devnet:latest ./h_devnet
```

5. Push hydrag devnet image to DockerHub

```
docker push rsantev/hydrag-devnet:latest
```
