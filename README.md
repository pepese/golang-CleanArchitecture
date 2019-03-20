```bash
$ pwd # プロジェクトルートであることを確認
$ mkdir web
$ cobra init $(pwd)/web
Your Cobra application is ready at
/Users/tanakakns/go/src/github.com/pepese/golang-CleanArchitecture/web

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.
```

cobra コマンド利用の都合で `web` 階層設けたが本当はいらぬ。  
コマンド作る。

```bash
$ cobra add httpserver
```

以下の階層を作る。

```
web/
├── LICENSE
├── cmd
│   ├── httpserver.go
│   └── root.go
├── domain
│   ├── helloLogic.go
│   └── model
│       └── hello.go
├── infrastructure
│   ├── datastore
│   │   └── helloDS.go
│   └── server
│       └── httpserver.go
├── interface
│   ├── controller
│   │   └── httphandler.go
│   └── presenter
│       └── helloRepository.go
├── main.go
├── registry
└── usecase
    └── helloUC.go
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
$ go build
```

- `GO111MODULE`
    - `on` ： 常に module-aware mode で動作する
    - `off` ： 常に GOPATH mode で動作する
    - `auto` ： `$GOPATH` 配下では  GOPATH modeで，それ以外のディレクトリでは module-aware mode で動作する