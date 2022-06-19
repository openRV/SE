.PHONY: run build init

run:
	go run src/*.go

build:
	go build src/*.go

init:
	go version
	go get github.com/gin-gonic/gin
	go get github.com/mattn/go-sqlite3
	go get github.com/BurntSushi/toml
	go get github.com/juju/ratelimit
	go get github.com/sirupsen/logrus
	go mod download github.com/gin-gonic/gin