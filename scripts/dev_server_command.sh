load_env
check_docker
ensure_postgres
run_goose up

echo "Starting scrapy server..."
$(cd server/src && go run .)