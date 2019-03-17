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
以下の階層を作る。

```
web/
├── LICENSE
├── cmd/
│   └── root.go
├── domain/
├── infrastructure/
│   ├── datastore/
│   └── server/
│       └── httpserver.go
├── interface/
│   ├── controller/
│   └── presenter/
├── main.go
├── registry/
└── usecase/
```

コマンド作る。

```bash
$ cobra add httpserver
```

```bash
$ cd web
$ go build
$ ./web httpserver
$ curl localhost:8080
Hello Go!
```