# Build and start the container
up:
    docker compose up -d

# Stop and remove the container
down:
    docker compose down

# Rebuild the container
rebuild:
    docker compose down
    docker compose up -d --build

# Get a shell in the running container
shell:
    docker compose exec magic-alias bash

# Run tests in the container
test:
    docker compose exec magic-alias go test ./...
    docker run --rm -it -v "$PWD:/code" bats/bats:latest /code/test

# Run the application
run:
    docker compose exec magic-alias go run main.go

# View logs
logs:
    docker compose logs -f

# Format code
fmt:
    go fmt ./...

# Exec in container
exec *args:
    docker compose exec magic-alias go run main.go {{args}}
