# go-template
Go (Golang) API REST Template with Gin Framework



## 1. Run with Docker

1. **Build**

```shell script
make build
docker build . -t goapi
```

2. **Run**

```shell script
docker run -p 3000:3000 goapi
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## 2. Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --dir cmd/api --parseDependency --output docs
```

Run and go to **http://localhost:3000/docs/index.html**
