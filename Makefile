cleardata:
	ENVIRONMENT=generate GENERATE=empty go run ./cmd/generatedata

generatedata:
	ENVIRONMENT=generate go run ./cmd/generatedata

build-prod:
	make cleardata
	make generatedata
	make build	

build:
	GOOS=linux   GOARCH=amd64 go build -o ./bin/foobarman-linux       ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/foobarman-windows.exe ./cmd/app

buildonce:
	make cleardata
	make generatedata	
	go build -o ./bin/foobarman ./cmd/app

run-prod:
	make cleardata
	make run

run:
	ENVIRONMENT=development go run ./cmd/app

