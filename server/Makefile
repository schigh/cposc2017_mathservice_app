BIN = server

all: install build

build:
	go build -o ./$(BIN) server.go

install:
	go get -u github.com/codegangsta/gin
	go get -u github.com/sirupsen/logrus
	go get -u google.golang.org/grpc
	go get -u github.com/schigh/cposc2017_mathservice