# モック

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