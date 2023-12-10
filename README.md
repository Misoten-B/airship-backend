# AIRship-backend

## リンク


| Page         | URL                                      |
| ------------ | ---------------------------------------- |
| API          | http://localhost:8080                    |
| swagger      | http://localhost:8080/swagger/index.html |
| swagger json | http://localhost:8080/swagger/doc.json   |


## 実行手順

**初回実行**

```bash
docker run --name postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=airship -p 5432:5432 postgres:latest
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
psql -h 127.0.0.1 -p 5432 -U postgres -d airship
```

**Swagger更新**

```bash
go install github.com/google/wire/cmd/wire@latest
go generate
```

