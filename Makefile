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

run-db:
	# This is for the testing not prod
	docker run --name go-101-test-db -e POSTGRES_PASSWORD=roach -p 5432:5432 -d postgres