init:
	curl -sSf https://atlasgo.sh | sh

build:
	export GOOS=linux && \
	go build -C ./cmd -o ../bin

run:
	make build && \
	docker compose up --build

run-db:
	docker compose up postgres

migration:
	atlas migrate diff backend \
		--dir "file://database/migrations" \
		--to "file://database/database.hcl" \
		--dev-url "docker://postgres/17/username@password?search_path=public"

migrate:
	atlas migrate apply \
		--dir "file://database/migrations" \
		--url "postgres://username:password@:5432/database?search_path=public&sslmode=disable" \

queries:
	sqlc generate
