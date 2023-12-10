.PHONY:
.SILENT:

build: 
	go build -o ./.bin/sso cmd/sso/main.go

run:
	go run ./sso/cmd/sso/main.go --config=./sso/config/local.yaml

generete:
	cd ./protos && tasks generate

migrate:
	go run ./sso/cmd/migrator --storage-path=./sso/storage/sso.db --migrations-path=./sso/migrations
