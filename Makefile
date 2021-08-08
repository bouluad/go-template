get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	/Users/bouluad/go/src/github.com/swaggo/swag/swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o bin/restapi cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./test/...

build-docker: build
	docker build . -t api-go

run-docker: build-docker
	docker run -p 3000:3000 api-go
