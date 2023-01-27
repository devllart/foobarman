build:
	go build -o ./bin/foobarman ./cmd/app
	go build -o ./bin/foobarman.exe ./cmd/app

run:
	go run ./cmd/app
