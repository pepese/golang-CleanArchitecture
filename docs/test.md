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
% go test -coverprofile=./coverage/cover.out ./...
% go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
% open ./coverage/cover.html
```

## 結合テスト

postman / newman を利用してテストを実行する。

```
% brew install postman --cask
% brew install newman
```

以下で実行。

```
% % newman run test/postman/collection/ca-app\ test.postman_collection.json 
newman

ca-app test

→ Get Users
  GET localhost:8080/api/v1/users [200 OK, 177B, 91ms]
  ✓  Status code is 200

┌─────────────────────────┬──────────────────┬──────────────────┐
│                         │         executed │           failed │
├─────────────────────────┼──────────────────┼──────────────────┤
│              iterations │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│                requests │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│            test-scripts │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│      prerequest-scripts │                0 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│              assertions │                1 │                0 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 149ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 3B (approx)                              │
├───────────────────────────────────────────────────────────────┤
│ average response time: 91ms [min: 91ms, max: 91ms, s.d.: 0µs] │
└───────────────────────────────────────────────────────────────┘
```