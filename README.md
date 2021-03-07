# golang-CleanArchitecture

Go 言語でクリーアーキテクチャを実現する構成を作ってみる。

## 実行

```zsh
% docker-compose up --build -d
% curl localhost:8080/api/v1/hello
Hello Go!
% docker-compose down
```

>以下、動かないがメモ程度に。
>
>### アプリ単体
>
>```zsh
>% go version
>go version go1.15.6 darwin/amd64
>% go mod tidy
>% go build -o ca-app
>% ./app server # ./app server --type gin # ./app server -t gin
>% curl localhost:8080/api/v1/hello
>Hello Go!
>```
>
>### Docker単体
>
>```zsh
>% docker build -t ca-app .
>% docker run -d -p 8080:8080 --name ca-app ca-app
>% curl localhost:8080/api/v1/hello
>Hello Go!
>% docker container stop ca-app
>```

## 疎通

```zsh
% curl -X POST -H 'Content-Type:application/json' -d '{"first_name":"first","last_name":"last"}' localhost:8080/api/v1/users | jq .
{
  "id": 1,
  "first_name": "first",
  "last_name": "last",
  "CreatedAt": "2019-11-11T18:55:04.015057+09:00",
  "UpdatedAt": "2019-11-11T18:55:04.015057+09:00",
  "DeletedAt": null
}

% curl localhost:8080/api/v1/users | jq .
[
  {
    "id": 1,
    "first_name": "first",
    "last_name": "last",
    "CreatedAt": "2019-11-11T18:55:04+09:00",
    "UpdatedAt": "2019-11-11T18:55:04+09:00",
    "DeletedAt": null
  }
]

% curl -X PUT -H 'Content-Type:application/json' -d '{"first_name":"First","last_name":"Last"}' localhost:8080/api/v1/users/1 | jq .
{
  "id": 1,
  "first_name": "First",
  "last_name": "Last",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "2019-11-11T18:55:53.262974+09:00",
  "DeletedAt": null
}

% curl localhost:8080/api/v1/users/1 | jq .
{
  "id": 1,
  "first_name": "First",
  "last_name": "Last",
  "CreatedAt": "2019-11-11T18:55:04+09:00",
  "UpdatedAt": "2019-11-11T18:55:53+09:00",
  "DeletedAt": null
}

% curl -X DELETE localhost:8080/api/v1/users/1 | jq .
{
  "id": 1,
  "first_name": "",
  "last_name": "",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdatedAt": "0001-01-01T00:00:00Z",
  "DeletedAt": null
}

% curl localhost:8080/api/v1/users | jq .
[]
```

### テーブルの確認

```zsh
% docker exec -it <container_id> /bin/bash
/# mysql
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| testdb             |
+--------------------+
5 rows in set (0.00 sec)

mysql> use testdb;
mysql> show tables;
+------------------+
| Tables_in_testdb |
+------------------+
| users            |
+------------------+
1 row in set (0.00 sec)

mysql> explain users;
+------------+------------------+------+-----+---------+----------------+
| Field      | Type             | Null | Key | Default | Extra          |
+------------+------------------+------+-----+---------+----------------+
| id         | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | timestamp        | YES  |     | NULL    |                |
| updated_at | timestamp        | YES  |     | NULL    |                |
| deleted_at | timestamp        | YES  | MUL | NULL    |                |
| first_name | varchar(255)     | YES  |     | NULL    |                |
| last_name  | varchar(255)     | YES  |     | NULL    |                |
+------------+------------------+------+-----+---------+----------------+
6 rows in set (0.01 sec)

mysql> quit;
/# exit
```

## 実装メモ

```zsh
% pwd # プロジェクトルートであることを確認
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

## [命名規則](./docs/naming.md)

## [設計](./docs/design.md)

## [ライブラリ類](./docs/libs.md)

## [テスト](./docs/test.md)

## 参考

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [The Algorithms - Go](https://github.com/TheAlgorithms/Go)
- [go-algorithms](https://github.com/0xAX/go-algorithms)