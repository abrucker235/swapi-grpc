## Generate/Update GRPC
```
protoc -I /home/aaron/go/src/swapi-grpc swapi/swapi.proto --go_out=plugins=grpc:.
```