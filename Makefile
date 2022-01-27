run-cli:
	go run clients/cli/main.go

build-cli:
	go build -o dist/dicegame-cli clients/cli/main.go
	GOARCH=amd64 GOOS=windows go build -o dist/dicegame-cli.exe clients/cli/main.go