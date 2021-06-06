# Mongo DB

## Mongod のオプションについて

- dbpath
- auth
- bind_ip

### dbpath

- dbpath は、data ファイルが保存される場所
- ジャーナリングのログも保存される
- デフォルトは /data/db

### auth

- auth は、client が mongod に接続する際に認証を求める設定

### bind_ip

- bind_ip で指定された ip の client は mongod に接続できる
- , 区切りで複数設定できる

## Mongo DB のファイル構成について

- デフォルトでは /data/db 配下に Mongo DB が管理するファイルが格納される
    - 実際のデータや index など
        - collection-0-xxxxx.wt
        - index-0-xxxx.wt
    - 診断用の metrics など
        - diagnostic.data/metrics.interim

- サーバーが予期せず落ちた場合は、lock ファイルに値が残り続ける為、手動でのクリアが必要になる場合がある
    - /data/db/WiredTiger.lock
    - /data/db/mongod.lock
    - /tmp/mongod.sock
- 基本的に /data/db 配下のファイルの更新は、コマンド実行時にインストラクションが合ったり、ドキュメントに記載されていない場合は行わないようにする。
