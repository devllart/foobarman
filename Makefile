build:
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

run:
	go run ./cmd/app

readmeupdate:
	rm -rf bin/assets
	cp -r assets bin/
	rm -f bin/Readme.md
	cp Readme.md bin/Readme.md

buildfull:
	make readmeupdate
	make build
