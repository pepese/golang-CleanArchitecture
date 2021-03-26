# golang-CleanArchitecture
![go version](https://img.shields.io/badge/Go-v1.15.6-blue)

---

Go 言語でクリーアーキテクチャを実現する構成を作ってみる。

---

## 実行

Docker Compose で DB 込みで起動する。

### Gin

```zsh
# 起動
% docker-compose -f deployments/docker-compose-gin.yml up --build -d
# 停止
% docker-compose -f deployments/docker-compose-gin.yml down
```

### gRPC

```zsh
# 起動
% docker-compose -f deployments/docker-compose-grpc.yml up --build -d
# 停止
% docker-compose -f deployments/docker-compose-grpc.yml down
```

### Kafka

```zsh
# 起動
% docker-compose -f deployments/docker-compose-kafka.yml up --build -d
# 停止
% docker-compose -f deployments/docker-compose-kafka.yml down
```

## 疎通

### Gin

```zsh
% curl localhost:8080/api/v1/hello
Hello Go!

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

### gRPC

```zsh
% evans -p 8080 api/user.proto

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


user.v1.UserService@127.0.0.1:50051> show service
+-------------+-------------+--------------+---------------+
|   SERVICE   |     RPC     | REQUEST TYPE | RESPONSE TYPE |
+-------------+-------------+--------------+---------------+
| UserService | ListUsers   | UserRequest  | UsersResponse |
| UserService | GetUser     | UserRequest  | UserResponse  |
| UserService | CreateUser  | UserRequest  | UserResponse  |
| UserService | UpdateUsers | UserRequest  | UserResponse  |
| UserService | DeleteUsers | UserRequest  | UserResponse  |
+-------------+-------------+--------------+---------------+

user.v1.UserService@127.0.0.1:8080> call ListUsers
ID (TYPE_INT64) => 
FirstName (TYPE_STRING) => 
LastName (TYPE_STRING) => 
{}

user.v1.UserService@127.0.0.1:50051> exit
Good Bye :)
```

### Kafka

```zsh
# topic 作成
% docker-compose -f deployments/docker-compose-kafka.yml exec kafka /opt/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic users
Created topic users.

# topic 一覧取得
% docker-compose -f deployments/docker-compose-kafka.yml exec kafka /opt/kafka/bin/kafka-topics.sh --list --zookeeper zookeeper:2181
users

# message 送信
% docker-compose -f deployments/docker-compose-kafka.yml exec kafka /opt/kafka/bin/kafka-console-producer.sh --broker-list kafka:9092 --topic users
>{"method":"create","message":{"first_name":"first","last_name":"last"}}
>message1
>^C

# ログを見て message の受信を確認
% docker logs <container_id>
Waiting for mysql
............MySQL is up - executing command
2021/03/15 16:30:10 consumer created
2021/03/15 16:30:10 commence consuming
# （省略）
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

>以下、 DB が無い場合動かないがメモ程度に。
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
>% docker build -t ca-app -f build/package/Dockerfile-gin .
>% docker run -d -p 8080:8080 --name ca-app ca-app
>% curl localhost:8080/api/v1/hello
>Hello Go!
>% docker container stop ca-app
>```

## [命名規則](https://tanakakns.github.io/go/naming/)

## [設計](./docs/design.md)

## [gRPC](./docs/grpc.md)

## [Kafka](./docs/kafka.md)

## [ライブラリ類](./docs/libs.md)

## [テスト](./docs/test.md)

## [バッジ](./docs/badge.md)

## 参考

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [The Algorithms - Go](https://github.com/TheAlgorithms/Go)
- [go-algorithms](https://github.com/0xAX/go-algorithms)