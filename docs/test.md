# テスト

## 単体テスト

[gomock](https://github.com/golang/mock) を利用。

```
% GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0
% mockgen -version
v1.5.0
```

以下でモックのコードを生成。

```
% mockgen -source app/usecase/user_repository.go -destination app/usecase/mock_usecase/mock_user_repository.go
% mockgen -source app/usecase/user.go -destination app/usecase/mock_usecase/mock_user.go
```

以下で単体テストを実行。

```zsh
% go test ./...

% go test -v -cover ./...
% go test -coverprofile=./test/coverage/cover.out ./...
% go tool cover -html=./test/coverage/cover.out -o ./test/coverage/cover.html
% open ./test/coverage/cover.html
```

## 結合テスト

postman / newman を利用してテストを実行する。

```
% brew install postman --cask
% brew install newman
```

以下で実行。

```
% newman run -e test/postman/env/docker-compose.postman_environment.json test/postman/collection/ca-app.postman_collection.json
newman

ca-app

→ Get Users
  GET localhost:8080/api/v1/users [200 OK, 177B, 101ms]
  ✓  Status code is 200

┌─────────────────────────┬────────────────────┬───────────────────┐
│                         │           executed │            failed │
├─────────────────────────┼────────────────────┼───────────────────┤
│              iterations │                  1 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│                requests │                  1 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│            test-scripts │                  2 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│      prerequest-scripts │                  1 │                 0 │
├─────────────────────────┼────────────────────┼───────────────────┤
│              assertions │                  1 │                 0 │
├─────────────────────────┴────────────────────┴───────────────────┤
│ total run duration: 166ms                                        │
├──────────────────────────────────────────────────────────────────┤
│ total data received: 3B (approx)                                 │
├──────────────────────────────────────────────────────────────────┤
│ average response time: 101ms [min: 101ms, max: 101ms, s.d.: 0µs] │
└──────────────────────────────────────────────────────────────────┘
```

### テストコードの作成

Postman の GUI でテストを作成する。  
設定項目は以下。

|項目|内容|
|:---|:---|
|Authorization|Basic AuthやOAuth2など、リクエストの認証方法を指定する|
|Headers|content-typeやAuthorizationなど、リクエストヘッダを指定する|
|Body|PostやPUTなどの場合に指定するリクエストボディ|
|Pre-request Script|リクエスト前に実行するJSコード|
|Tests|レスポンスペイロードの妥当性を検証するJSコード|

作成した内容は JOSN 形式で Export 可能。

### 環境変数の設定

同じく Postman の GUI で設定可能。  
設定方法は右上の歯車マークから「 MANAGE ENVIRONMENTS 」を選択。  
作成して名前を付けた環境情報は JSON 形式で 「 Download Environment 」から取得可能。