# golang-CleanArchitecture

Go 言語でクリーアーキテクチャを実現する場合の構成を作ってみる。

<img src="https://paulovich.net/img/CleanArchitecture-Uncle-Bob.jpg" />
<img src="https://paulovich.net/img/Flow-Of-Control.png" />

## 実行

```zsh
% git clone
% go version
go version go1.13.3 darwin/amd64
% export GO111MODULE=on
% go mod download
% go build -o app
% ./app server
% ./app server --type gin # ./app server -t gin
```

## 実装メモ

```zsh
% pwd # プロジェクトルートであることを確認
% export GO111MODULE=on
% go mod init
% go get -u github.com/spf13/cobra/cobra
% cobra init $(pwd)
Your Cobra application is ready at
$HOME/go/src/github.com/pepese/golang-CleanArchitecture

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.
```

コマンド作る。

```zsh
% cobra add server
```

以下の階層を作る。

```zsh
.
├── LICENSE
├── README.md
├── cmd
│   ├── root.go
│   └── server.go
├── domain
│   ├── hello_logic.go
│   └── model
│       └── hello.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── datastore
│   │   └── hello_datastore.go
│   └── server
│       ├── ginserver.go
│       ├── httpserver.go
│       └── util.go
├── interface
│   ├── controller
│   │   ├── gin_router.go
│   │   └── http_router.go
│   └── presenter
│       └── hello_repository.go
├── main.go
├── usecase
│   └── hello_usecase.go
└── web
    └── registry
```

実行する。

```zsh
% go build -o app
% ./app server # ./app server -t gin
% curl localhost:8080
Hello Go!
```