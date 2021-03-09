# Kafka

## 用語

|名前|役割|
|:---|:---|
|Kafka Cluster|Kafkaが実行されているサーバ(Broker)をグループ化したもの|
|Broker|Kafkaの単一サーバ|
|Zookeeper|Kafkaを管理するサーバ|
|Producer|Kafkaへメッセージを送信するアプリケーション|
|Consumer|Kafkaからメッセージを取得するアプリケーション|
|Topic|メッセージを整理するためのカテゴリー|
|Partition|Topic内のメッセージはパーティションという単位で分散させています|
|Replica|各Partitionは複数のBrokerに複製(Replica)されています|
|Leader|複製されているReplicaのうち唯一読み書きが許可されているReplica|
|Consumer Group|複数のConsumerを同一グループとして扱うためのもの。グループ化することで分散したConsumer間で同一メッセージを重複せずに読み込むことが可能です|
|Offset|Partition単位でメッセージをどこまで読んだか管理するためのもの|

## 全体像

<img src="./img/kafka_overview.png" />

## 設計

今回はサーバサイドとして Consumer で待ち受けるアプリを作る。