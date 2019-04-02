```bash
$ pwd # プロジェクトルートであることを確認
$ cobra init $(pwd)
Your Cobra application is ready at
/Users/tanakakns/go/src/github.com/pepese/golang-CleanArchitecture

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.
```

コマンド作る。

```bash
$ cobra add httpserver
```

以下の階層を作る。

```
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

```bash
$ cd web
$ go build
$ ./web httpserver
$ curl localhost:8080
Hello Go!
```

- [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules)
- [Go 1.11 Modules](https://qiita.com/sky0621/items/9af758c7df5403caa991)

```bash
$ go version
go version go1.11.5 darwin/amd64
$ export GO111MODULE=on
$ go mod init
$ go build -o app
$ ./app httpserver
```

- `GO111MODULE`
    - `on` ： 常に module-aware mode で動作する
    - `off` ： 常に GOPATH mode で動作する
    - `auto` ： `$GOPATH` 配下では  GOPATH modeで，それ以外のディレクトリでは module-aware mode で動作する
- importするリポジトリのソースコードは `$GOPATH/pkg/mod` にバージョンを固定してDLされる

```bash
$ git clone
$ go version
go version go1.11.5 darwin/amd64
$ export GO111MODULE=on
$ go mod download
$ go build -o app
$ ./app server
$ ./app server --type gin
```