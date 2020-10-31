generate-init:
	go run github.com/99designs/gqlgen generate

generate:
	go generate ./...

test:
	go test ./...
