# grpc gateway


## install protobuf
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

export PATH="$PATH:$(go env GOPATH)/bin"
```


## compile protoc
```bash
protoc --proto_path=./proto \
   --go_out=./proto --go_opt=paths=source_relative \
  --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=./proto --grpc-gateway_opt=paths=source_relative \
  ./proto/hello.proto
```

you may got problem like this:
```bash
protoc --proto_path=./proto \
>    --go_out=./proto --go_opt=paths=source_relative \
>   --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
>   --grpc-gateway_out=./proto --grpc-gateway_opt=paths=source_relative \
>   ./proto/hello.proto
google/api/annotations.proto: File not found.
```
just add [googleapi](https://github.com/googleapis/googleapis/tree/master/google/api) to your `/usr/local/include` dictionary and compile again
```
$ pwd
/usr/local/include/google/api
$
$ ls
BUILD.bazel			httpbody.proto
README.md			label.proto
annotations.proto		launch_stage.proto
apikeys				log.proto
auth.proto			logging.proto
backend.proto			metric.proto
billing.proto			monitored_resource.proto
client.proto			monitoring.proto
cloudquotas			policy.proto
config_change.proto		quota.proto
consumer.proto			resource.proto
context.proto			routing.proto
control.proto			service.proto
distribution.proto		serviceconfig.yaml
documentation.proto		servicecontrol
endpoint.proto			servicemanagement
error_reason.proto		serviceusage
expr				source_info.proto
field_behavior.proto		system_parameter.proto
field_info.proto		usage.proto
http.proto			visibility.proto
```

## run server
```bash
$ go run server/server.go
2024/06/07 15:49:18 Serving http on 0.0.0.0:8080
2024/06/07 15:49:18 Serving gRPC on 0.0.0.0:50051
```

## run client
```bash
# grpc client
$ go run client/client.go
2024/06/07 16:09:56 Greeting: hello world

# http client
$ curl localhost:8080/v1/greeter/sayhello?name=world
{"message":"hello world"}
```

## reference
[1] https://stackoverflow.com/questions/70586511/protoc-gen-go-unable-to-determine-go-import-path-for-simple-proto

[2] https://stackoverflow.com/questions/66168350/import-google-api-annotations-proto-was-not-found-or-had-errors-how-do-i-add

[3] https://stackoverflow.com/questions/71733519/golang-google-protobuf-empty-proto-file-not-found

[4] https://www.lixueduan.com/posts/grpc/07-grpc-gateway/