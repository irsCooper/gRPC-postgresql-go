.PHONY:
.SILENT:

build: 
	go build -o ./.bin/sso cmd/sso/main.go

ssoRun:
	go run ./sso/cmd/sso/main.go --config=./sso/config/local.yaml

generate:
	cd ./protos && protoc -I proto proto/sso/sso.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative && protoc -I proto proto/notepad/notepad.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
	
ssoMigrate:
	go run ./sso/cmd/migrator --storage-path=./sso/storage/sso.db --migrations-path=./sso/migrations

ssoTestMigrate:
	go run ./sso/cmd/migrator/main.go --storage-path=./sso/storage/sso.db --migrations-path=./sso/tests/migrations --migrations-table=migrations_test

notepadRun:
	go run ./notepad/cmd/notepad/main.go --config=./notepad/config/local.yaml
