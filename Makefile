run-cli:
	go run clients/cli/main.go -v

build-cli:
	GOARCH=amd64 GOOS=linux go build -o dist/dicegame-cli clients/cli/main.go
	GOARCH=amd64 GOOS=windows go build -o dist/dicegame-cli.exe clients/cli/main.go
	GOARCH=arm GOOS=linux go build -o dist/dicegame-cli-android clients/cli/main.go

proto:
	protoc --go_out=./server/models/proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false:./server/models/proto \
		--go-grpc_opt=paths=source_relative ./server/game.proto