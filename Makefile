build:
	GOOS=linux GOARCH=amd64 go build -o shipcon-user-cli
	docker build -t shipcon-user-cli .
	go clean

run:
	docker run --net="host" \
	-e MICRO_REGISTRY=mdns \
	shipcon-user-cli