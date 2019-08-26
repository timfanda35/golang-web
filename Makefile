run:
	go run main.go

test:
	go test -v golang-web/...

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v