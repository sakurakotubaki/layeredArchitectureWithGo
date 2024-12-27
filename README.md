# レイヤードアーキテクチャ サンプルアプリケーション

作業進捗管理アプリケーションのサンプルコードです。
レイヤードアーキテクチャの学習用教材として作成されています。

## アーキテクチャ構成

- プレゼンテーション層 (interfaces/handlers)
  - HTTPリクエストの受け付けとレスポンスの返却
  - 入力値のバリデーション
  
- アプリケーション層 (usecase)
  - ユースケースの実装
  - トランザクション制御
  
- ドメイン層 (domain)
  - ビジネスロジック
  - ドメインモデル
  
- インフラストラクチャ層 (infrastructure)
  - 永続化の実装（このサンプルではインメモリ）

## 主な学習ポイント

1. レイヤードアーキテクチャの基本構造
2. Contextの利用方法
3. インターフェースを活用した依存関係の管理
4. クリーンアーキテクチャの考え方
5. GoのEchoフレームワークの基本的な使い方

## 起動方法

```bash
go run main.go
```

## API エンドポイント

- GET /api/works - 作業一覧の取得
- GET /api/works/:id - 特定の作業の取得
- POST /api/works - 新規作業の作成
- PUT /api/works/:id - 作業の更新
- DELETE /api/works/:id - 作業の削除

## システムの全体像
### インフラストラクチャ層 (internal/infrastructure/memory/)

- **インメモリ**とは、プログラムのメモリー上にデータを保存することを指します。
  - データベースやファイルシステムにアクセスする必要がなく、高速にデータを保存・読み込みすることができます。
  - しかし、プログラムが終了するとデータは消失します。

- ドメイン層 (internal/domain/)
  - Work モデルの定義
  - リポジトリインターフェースの定義
  - バリデーションロジック
- アプリケーション層 (internal/usecase/)
  - ビジネスロジックの実装
  - トランザクション制御
  - ドメインオブジェクトの操作
- インターフェース層 (internal/interfaces/handlers/)
  - HTTPリクエストの受け付け
  - レスポンスの返却
  - 入力値のバリデーション
- インフラストラクチャ層 (internal/infrastructure/memory/)
  - インメモリデータストアの実装
  - リポジトリインターフェースの実装

## 実行するcurlコマンド

POST:
```shell
curl -X POST -H "Content-Type: application/json" -d '{"title": "作業A", "progress": 30}' http://localhost:8080/api/works
```

GET:
```shell
curl http://localhost:8080/api/works
```

PUT:
```shell
curl -X PUT -H "Content-Type: application/json" -d '{"title": "作業A更新", "progress": 50}' http://localhost:8080/api/works/1
```

DELETE:
```shell
curl -X DELETE http://localhost:8080/api/works/1