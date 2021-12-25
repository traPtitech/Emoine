# Emoine

## 開発環境

golang 1.17
node 16

download https://github.com/protocolbuffers/protobuf/releases

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
```

初回のみ必要
```shell
$ make proto
```

Windowsの場合、上記コマンドが実行できないので、代わりに以下を実行する
```shell
protoc -I=./docs --go_out=. ./docs/comment.proto
protoc -I=./docs --go_out=. ./docs/reaction.proto
protoc -I=./docs --go_out=. ./docs/state.proto
protoc -I=./docs --go_out=. ./docs/viewer.proto
protoc -I=./docs --go_out=. ./docs/message.proto
```

起動

```shell
$ make up
```

終了
```shell
$ make down
```
