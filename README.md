## 概要
Docker & Go言語を用いたJSON形式のREST API

|| Tool |
|:----------:|:-----------:|
| Server | golang |
| RDB | PostgreSQL |
| Framework | gin |
| O/RM | gorm |


## 設計
クリーンアーキテクチャを採用し、ディレクトリ構成は以下のようにした

```
├── README.md
├── docker-compose.yml
└── src/
    └── app/
        ├── Dockerfile
        ├── server.go
        ├── infrastructure/
        ├── interfaces/
        │   ├── controllers/
        │   └── database/
        ├── models/
        └── usecase/
```

|ディレクトリ| 役割 |
|:----------:|:-----------:|
| infrastructure | データベース接続周りとルーティング処理 |
| interfaces/controllers | 一般的なコントローラでルーティングからの処理を担当 |
| interfaces/database | データベースに対しての処理を管理 |
| models | データモデルを定義 |
| usecase | ビジネスロジックを担当しデータの流れを組み立てる |


## API仕様

user モデル
 - id
 - name
 - email
 - created_at
 - updated_at
 
|HTTPメソッド| エンドポイント |　機能 | ステータスコード | レスポンス |
|:----------|:-----------|:----------|:-----------|:-----------|
| GET | / | Hello World!! | 200 | { "message": "Hello World!!" } |
| GET | /users | user の一覧を表示 | 200 | [ user1, user2, ... ] |
| GET | /users/:id | 指定した id の user を表示 | 200 | user{id} |
| POST | /users | user を追加 | 201 | created_user |
| PUT | /users:id | user を更新 | 200 | updated_user |
| DELETE | /users:id | user を削除 | 204 |  |


## 使用方法

- `docker-compose up`で`localhost:8080`でAPIサーバーが立ち上がりアクセス可能
- dinghyを利用している場合はバーチャルホストとして`go-api.docker`でアクセス可能
- PostgreSQLのクライアントツールである`pgweb`のコンテナも用意してあり、使う場合は`docker-compose up -d pgweb`で起動して`localhost:8081`でDBをGUIで確認・操作可能

