.PHONY:
.SILENT:

build: 
	go build -o ./.bin/sso cmd/sso/main.go

ssoRun:
	go run ./sso/cmd/sso/main.go --config=./sso/config/local.yaml

generate:
	cd ./protos && tasks generate

ssoMigrate:
	go run ./sso/cmd/migrator --storage-path=./sso/storage/sso.db --migrations-path=./sso/migrations

ssoTestMigrate:
	go run ./sso/cmd/migrator/main.go --storage-path=./sso/storage/sso.db --migrations-path=./sso/tests/migrations --migrations-table=migrations_test


