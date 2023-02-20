cleardata:
	ENVIRONMENT=generate GENERATE=empty go run ./cmd/generate_data

generatedata:
	ENVIRONMENT=generate go run ./cmd/generate_data

build:
	make cleardata
	make generatedata
	
	GOOS=linux   GOARCH=386   go build -o ./bin/foobarman-386-linux         ./cmd/app
	GOOS=linux   GOARCH=amd64 go build -o ./bin/foobarman-amd64-linux       ./cmd/app
	GOOS=linux   GOARCH=arm   go build -o ./bin/foobarman-arm-linux         ./cmd/app
	GOOS=linux   GOARCH=arm64 go build -o ./bin/foobarman-arm64-linux       ./cmd/app
	
	# GOOS=darwin  GOARCH=amd64 go build -o ./bin/foobarman-amd64-darwin      ./cmd/app
	# GOOS=darwin  GOARCH=arm64 go build -o ./bin/foobarman-arm64-darwin      ./cmd/app
	
	GOOS=windows GOARCH=386   go build -o ./bin/foobarman-386-windows.exe   ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/foobarman-amd64-windows.exe ./cmd/app
	GOOS=windows GOARCH=arm   go build -o ./bin/foobarman-arm-windows.exe   ./cmd/app
	GOOS=windows GOARCH=arm64 go build -o ./bin/foobarman-arm64-windows.exe ./cmd/app

	# GOOS=android GOARCH=386   go build -o ./bin/foobarman-386-android       ./cmd/app
	# GOOS=android GOARCH=amd64 go build -o ./bin/foobarman-amd64-android     ./cmd/app
	# GOOS=android GOARCH=arm   go build -o ./bin/foobarman-arm-android       ./cmd/app
	# GOOS=android GOARCH=arm64 go build -o ./bin/foobarman-arm64-android     ./cmd/app

buildonce:
	make cleardata
	make generatedata	
	go build -o ./bin/foobarman ./cmd/app

run:
	make cleardata
	ENVIRONMENT=development go run ./cmd/app

