build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/foobarman ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/foobarman.exe ./cmd/app

run:
	go run ./cmd/app

readmeupdate:
	rm -rf bin/assets
	cp -r assets bin/
	rm bin/Readme.md
	cp Readme.md bin/Readme.md
