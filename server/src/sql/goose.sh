#!/bin/sh

set -a
GOOSE_DRIVER="postgres"
GOOSE_MIGRATION_DIR="./schema"
GOOSE_DBSTRING="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable"
set +a

goose "$@"