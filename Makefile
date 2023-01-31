build:
	go build -o ./bin/foobarman ./cmd/app
	go build -o ./bin/foobarman.exe ./cmd/app

run:
	go run ./cmd/app

readmeupdate:
	rm -rf bin/assets
	cp -r assets bin/
	rm bin/Readme.md
	cp Readme.md bin/Readme.md
