check_docker() {
    # Check if Docker is running by attempting to communicate with the Docker daemon
    echo "Checking whether docker is running..."
    if docker info >/dev/null 2>&1; then
        echo "Docker is running."
    else
        echo "Docker is not running or not installed."
        exit 1
    fi
}