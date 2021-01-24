# setup

## messageデータベースの作成

MySQLを使っています。

```
CREATE DATABASE message CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

## シードデータを入れる

main.goと同じ階層にmessage_dumpファイルを用意しています。

```
mysql -u USER_NAME -p message < message_dump
```

## .envを各自用意

main.go と同じ階層でお願いします

```
DATABASE_NAME=message
DATABASE_USER_NAME=
DATABASE_PASSWORD=
DATABASE_HOST=
DATABASE_PORT=
```

例
```
DATABASE_NAME=message
DATABASE_USER_NAME=root
DATABASE_PASSWORD=password
DATABASE_HOST=127.0.0.1
DATABASE_PORT=3306
```


## 実行

```
go run main.go
```

# 仕様説明

## ルーティング

main.go を実行後にコンソールにも出力されます

```
ユーザー作成
POST /v1/auth/signup

グループ作成
POST /v1/groups

指定したグループに投稿されたメッセージ取得
GET /v1/groups/{group_id}/messages

指定したグループにメッセージを投稿
POST /v1/groups/{group_id}/messages

グループに投稿したメッセージを削除
DELETE /v1/groups/{group_id}/messages/{group_message_id}

指定したユーザーにDMを送信
POST /v1/users/{user_id}/direct_messages

指定したユーザーとの会話を取得
GET /v1/users/{user_id}/direct_messages

指定したユーザーとの特定のDMを削除
DELETE /v1/users/{user_id}/direct_messages/{direct_message_id}
```

## HTTP CLIENT

https://qiita.com/toshi0607/items/c4440d3fbfa72eac840c

ごめんなさいVSCodeでのみなのですが

{project_root}/test/http/test.http

にAPIを実行しやすいようなファイルを用意しています。

こちらからだと簡単にAPI叩けます。上記Qiitaを参照してください。

こちらのファイルをAPIドキュメントとしても考えていただけると嬉しいので、

VSCodeのHTTP Clientを利用しなくても見て欲しいです！！

## 使用技術

Go 1.14.13

go-chi https://github.com/go-chi/chi

gorm https://github.com/go-gorm/gorm

godotenv https://github.com/joho/godotenv

jwt-go https://github.com/dgrijalva/jwt-go
