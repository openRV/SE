.PHONY: run build init

run:
	go run src/main.go

build:
	go build src/main.go

init:
	go version
	go get github.com/gin-gonic/gin
	go get github.com/mattn/go-sqlite3
	go mod download github.com/gin-gonic/gin