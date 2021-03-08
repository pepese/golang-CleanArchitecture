# gRPC

```zsh
% mkdir app/interface/controller/grpc_gen
% protoc -I=./api --go_out=plugins=grpc:./app/interface/controller/grpc_gen api/user.proto
```