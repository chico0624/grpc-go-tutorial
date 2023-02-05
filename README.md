# gPRC Tutorial for a Golang 

## 参考
[作ってわかる！はじめてのgPRC](https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/rpc)


## 事前準備
参考：https://grpc.io/docs/languages/go/quickstart/

- インストール
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
- $GOPATH/binにPATHを通す

- grpcurl（直接gRPCサービスメソッドを叩くために必要）
`brew install grpcurl`


## protoc生成&再生成
```
cd api
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	hello.proto
```

## 実行
- Server起動
`go run ./cmd/server/main.go`
- grpc操作

```
# grpcサービス一覧
grpcurl -plaintext localhost:8080 list
# サービスメソッド一覧
grpcurl -plaintext localhost:8080 list myapp.GreetingService
# 呼び出し
grpcurl -plaintext -d '{"name": "chico"}' localhost:8080 myapp.GreetingService.Hello
```