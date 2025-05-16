init:
	curl -sSf https://atlasgo.sh | sh

build:
	export GOOS=linux && \
	go build -C ./cmd -o ../bin

es:
	docker compose up localstack postgres

run:
	make build && \
	docker compose up api --build

run-db:
	docker compose up postgres

migration:
	atlas migrate diff backend \
		--dir "file://database/migrations" \
		--to "file://database/database.hcl" \
		--dev-url "docker://postgres/17/username@password?search_path=public"

migration-rehash:
	atlas migrate hash --dir "file://database/migrations"

migrate:
	atlas migrate apply \
		--dir "file://database/migrations" \
		--url "postgres://username:password@:5432/database?search_path=public&sslmode=disable" \

queries:
	sqlc generate

make deploy:
	aws cloudformation deploy \
		--stack-name dev-rolesejogos-vpc \
		--template-file "./cloudformation/vpc.yaml" \
	&& \
	aws cloudformation deploy \
		--stack-name dev-rolesejogos-s3 \
		--template-file "./cloudformation/s3.yaml" \
	&& \
	aws cloudformation deploy \
		--stack-name dev-rolesejogos-cloudfront \
		--template-file "./cloudformation/cloudfront.yaml"


