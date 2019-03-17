# Overview
gRPC implementation of [Stars Wars API](https://swapi.co).

## Generate/Update GRPC
```
protoc -I /home/aaron/go/src/swapi-grpc swapi/swapi.proto --go_out=plugins=grpc:.
```

## Build
```
go build -o swapi
```

## Run
```
./swapi
```