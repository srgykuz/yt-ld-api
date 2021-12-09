build:
	go build .

test:
	go test ./...

migrate:
	go run ./migrations --envFile $(ENV_FILE)

dev:
	go run . --host 127.0.0.1 --port 8080 --infoLog /dev/stderr --debugLog /dev/stderr --envFile .env.development
