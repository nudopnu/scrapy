CONTAINER_NAME="scrapy_postgres"

ensure_postgres() {
    if docker ps -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "A container with the name '${CONTAINER_NAME}' already exists."
    # Start the container
    docker start "${CONTAINER_NAME}"
    echo "The container '${CONTAINER_NAME}' has been started."
    else
    # Run postgres docker container
    docker run --name "${CONTAINER_NAME}" \
        -p 5432:5432 \
        --env POSTGRES_PASSWORD=${SCRAPY_DATABASE_PASSWORD} \
        --env POSTGRES_USER=${SCRAPY_DATABASE_USERNAME} \
        --env POSTGRES_DB=${SCRAPY_DATABASE_NAME} \
        --health-cmd="pg_isready -U ${SCRAPY_DATABASE_USERNAME}" \
        --health-interval=10s \
        --health-timeout=5s \
        --health-retries=5 \
        -d postgres:17.2
    fi

    echo "Waiting for PostgreSQL container to be healthy..."
    until [ "$(docker inspect -f '{{.State.Health.Status}}' ${CONTAINER_NAME})" == "healthy" ]; do
        sleep 1
    done
    echo "PostgreSQL container is running on port 5432"
}
