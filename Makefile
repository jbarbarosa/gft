default:
	go build -o bin/gft .

test:
	go test ./...
