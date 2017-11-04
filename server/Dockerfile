FROM golang:1.9-alpine
COPY ./ /go/src/github.com/schigh/cposc2017_app/server
WORKDIR /go/src/github.com/schigh/cposc2017_app/server

RUN apk add --no-cache git make bash
RUN make

CMD gin --bin ./server -i run