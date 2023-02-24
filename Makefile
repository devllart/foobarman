cleardata:
	ENVIRONMENT=generate GENERATE=empty go run ./cmd/generate_data

generatedata:
	ENVIRONMENT=generate go run ./cmd/generate_data

build:
	make cleardata
	make generatedata
	
	GOOS=linux   GOARCH=amd64 go build -o ./bin/foobarman-linux       ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/foobarman-windows.exe ./cmd/app

buildonce:
	make cleardata
	make generatedata	
	go build -o ./bin/foobarman ./cmd/app

run:
	make cleardata
	ENVIRONMENT=development go run ./cmd/app

