cleardata:
	ENVIRONMENT=generate GENERATE=empty go run ./cmd/generatedata

generatedata:
	ENVIRONMENT=generate go run ./cmd/generatedata

build_full:
	make cleardata
	make generatedata
	GOOS=linux   GOARCH=amd64 go build -o ./bin/foobarman-linux       ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/foobarman-windows.exe ./cmd/app

build:
	make cleardata
	make generatedata	
	go build -o ./bin/foobarman ./cmd/app

run:
	make cleardata
	ENVIRONMENT=development go run ./cmd/app

