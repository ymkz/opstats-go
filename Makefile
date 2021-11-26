dev-file:
	go run ./main.go ./mock/mini
	go run ./main.go ./mock/huge
	go run ./main.go ./mock/empty
dev-pipe:
	seq 10 | awk '{ print rand() }' | deno run ./src/main.ts
lint:
	gofmt -l .
test:
	go test
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o ./release/macos-amd64/opstats ./main.go
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./release/linux-amd64/opstats ./main.go