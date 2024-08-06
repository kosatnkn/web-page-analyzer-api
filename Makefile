# Running and Testing
run:
	# update metadata
	./metadata.sh
	go run main.go

test:
	go test -v ./...

# Mocking
mock:
	docker run --init --name mock-web-page-analyzer-api -it --rm -v $(PWD)/docs/api:/tmp -p 8000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"

# Containerizing
docker_build:
	# update metadata
	./metadata.sh
	docker build -t kosatnkn/web-page-analyzer-api:latest .


# Go dependancy management
dep_upgrade_list:
	go list -u -m all

dep_upgrade_all:
	go get -t -u ./... && go mod
