generate:
	sqlc generate -f ./db/sqlc.yaml

build-mac:
	env GOOS=darwin GOARCH=arm64 go build -o ./tmp/bin
