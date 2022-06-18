.PHONY: run build init

run:
	go run src/main.go

build:
	go build src/main.go

init:
	go version
	go get github.com/gin-gonic/gin
	go mod download github.com/gin-gonic/gin