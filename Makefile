.PHONY:
.SILENT:

build: 
	go build -o ./.bin/sso cmd/sso/main.go

run:
	go run ./sso/cmd/sso/main.go --config=./sso/config/local.yaml

generete:
	cd ./protos && tasks generate