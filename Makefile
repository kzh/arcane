generate:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/arcanepb/arcane.proto

bin/arcanecli:
	go build -o ./bin/acli ./cmd/cli

bin/arcane:
	go build -o ./bin/arcane ./cmd/arcane
