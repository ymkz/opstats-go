macos:
	GOOS=darwin GOARCH=amd64 go build -o ./release/macos-amd64/opstats ./main.go
linux:
	GOOS=linux GOARCH=amd64 go build -o ./release/linux-amd64/opstats ./main.go