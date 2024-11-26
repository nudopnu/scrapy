# Scrapy

This is a prototype full-stack application. It lets you create search agents to search

## Modify database

Install sqlc with 
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Then you can edit the `schema` and `queries` in the `server/src/sql/` subfolder. Then you run:

```bash
cd server/src
sqlc generate
```