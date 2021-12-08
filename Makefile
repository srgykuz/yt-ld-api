build:
	go build .

migrate:
	go run ./migrations

dev:
	go run . --host 127.0.0.1 --port 8080 --infoLog /dev/stderr --debugLog /dev/stderr --envFile .env.development
