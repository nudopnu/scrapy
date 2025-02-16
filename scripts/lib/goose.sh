run_goose() {
  # Enable exporting all variables to the environment
  set -a

  # Define environment variables for Goose
  GOOSE_DRIVER="postgres"
  GOOSE_MIGRATION_DIR="./server/src/sql/schema"

  # Source the .env file for credentials
  . ./.env

  # Construct the Goose database connection string
  GOOSE_DBSTRING="postgres://${SCRAPY_DATABASE_USERNAME}:${SCRAPY_DATABASE_PASSWORD}@localhost:5432/${SCRAPY_DATABASE_NAME}"

  # Disable exporting variables
  set +a

  # Pass all arguments to Goose
  goose "$@"
}