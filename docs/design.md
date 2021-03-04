# 設計

<img src="./img/CleanArchitecture-Uncle-Bob.jpg" />
<img src="./img/Flow-Of-Control.jpg" />

勝手ながら各層に名前をつけて整理する。

- 第 1 層：青色：Frameworks & Drivers： **infrastructure** 層：外の世界とシステムを相互に繋ぐ層、最も具体的で詳細な実装、変更が多い
    - **server** ディレクトリ： FW に依存したリクエスト受信・返却の実装
        - view（HTMLとか）、api(JSONとかXML)
    - **batch** ディレクトリ：FW に依存したバッチ・ジョブの実装
- 第 2 層：緑色：interface Adapters： **interface** 層： infrastructure と usecase を相互に繋ぐ層
    - **controller** ディレクトリ：infrastructure からの呼び出しを usecase にマッピングする **handler** を実装する
	- **presenter** ディレクトリ：usecase層の実行結果をレスポンスを返却する形へ整形する
    - **gateway** ディレクトリ：datastore や httpclient のインターフェースとなるusecase層の **repository** インターフェースの実装
	    - **datastore** ディレクトリ：データベース製品に依存したアクセス、 repository の実装
		- **httpclient** ディレクトリ：他のマイクロサービスが外部サービスの API 呼び出し、 repository の実装
- 第 3 層：赤色：Applications Business Rules： **usecase** 層：アプリケーション独自のロジックを実装、domain 層のことだけ知っている（参照している）
    - interface 層より呼び出される
	- 1 つ以上のドメインモデルを操作・ドメインロジックを実行し、ビジネスロジックを成立させる
	- ビジネスロジックの結果を controller へ返却（ return ）する
- 第 4 層：黄色：Enterprise Business Rules： **domain** 層：最も抽象的な重要な実装、変更が少ない、他の層への参照が無く知らない
    - usecase 層によって扱われるドメインモデルとドメインロジック
    - ユースケースを実行するビジネスロジックは usecase 層に実装するが、ドメイン固有のルール・ロジックはここに実装する
	- **model** ディレクトリ
	    - ドメインモデルを入れる
- その他の層
    - **registry** ： DI とか
    - **cmd** ： cobra とかで作ったコマンド
    - **view** ： html template とか
    - **public** ： css とか js とか

今回は以下のようになった。

```
.
├── main.go
├── cmd
│   ├── root.go
│   └── server.go
├── infrastructure
│   └── server
│       ├── gin.go
│       ├── http.go
│       └── util.go
├── interface
│   ├── controller
│   │   ├── gin_router.go
│   │   └── http_router.go
│   ├── gateway
│   │   └── hello_datastore.go
│   └── presenter
│       └── hello.go
├── usecase
│   ├── hello.go
│   ├── hello_input.go
│   ├── hello_output.go
│   └── hello_repository.go
└── domain
    ├── hello_logic.go
    └── model
        └── hello.go
```

## 参考

- [CleanArchitectureでGolangらしいPackage構成を考える](https://qiita.com/inosy22/items/ce4a6ea7545c5cefd24b)
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [Go言語でClean Architectureを実現して、gomockでテストしてみた](https://qiita.com/ogady/items/34aae1b2af3080e0fec4)
- [GoにおけるDI](http://inukirom.hatenablog.com/entry/di-in-go)
- [Goのサーバサイド実装におけるレイヤ設計とレイヤ内実装について考える](https://www.slideshare.net/pospome/go-80591000)