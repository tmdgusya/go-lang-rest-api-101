remove-build-cache:
	if [ -d "./build" ]; then \
		rm -rf ./build; \
	fi

build:
	@go build -o ./build/server go-rest-api-101

run: remove-build-cache build server

server:
	./build/server

test:
	go test -v go-rest-api-101