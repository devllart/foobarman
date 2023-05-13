cleardata:
	ENVIRONMENT=generate GENERATE=empty go run ./cmd/generatedata

generatedata:
	ENVIRONMENT=generate go run ./cmd/generatedata

prebuild:
	make cleardata
	make generatedata

build_all_platform:
	GOOS=linux   GOARCH=amd64 go build -o ./.dist/bin/foobarman-linux       ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./.dist/bin/foobarman-windows.exe ./cmd/app

build_full:
	make prebuild
	make build_all_platform
	make build_web_app

build_web_app:
	gopherjs build ./cmd/web_app -o ./.dist/web_app/index.js

build:
	make prebuild
	go build -o ./.dist/bin/foobarman ./cmd/app

run-prod:
	make cleardata
	make run

run:
	ENVIRONMENT=development go run ./cmd/app

