# AIRship-backend

## リンク


| Page         | URL                                      |
| ------------ | ---------------------------------------- |
| API          | http://localhost:8080                    |
| swagger      | http://localhost:8080/swagger/index.html |
| swagger json | http://localhost:8080/swagger/doc.json   |


## 実行手順

**初回実行**

db
```bash
docker run --detach --name mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=airship -e MYSQL_USER=user -e MYSQL_PASSWORD=password -p 3306:3306 mysql:latest
```

go
```bash
# .env.sampleをコピペして.envに
# .serviceAccountKey.jsonをルートに
# ローカル環境にインストールされます
make init
```

**データベース初期化**

```bash
go run scripts/migrate/main.go
go run scripts/seed/main.go
```

**アプリケーション実行**

```bash
go run .
```

**linter実行**
```bash
docker compose -f compose.yml -f compose.lint.yml up
```

**データベース接続**

```bash
mysql -uuser -p
```

**Swagger更新**

```bash
go generate
```

