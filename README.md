** Компиляция GRPC proto-файлов с использованием относительных путей:
```
protoc --proto_path=api/proto --go_out=paths=source_relative:api/proto/pb --go-grpc_out=paths=source_relative:api/proto/pb api/proto/service.proto

```