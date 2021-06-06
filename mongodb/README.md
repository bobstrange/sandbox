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
