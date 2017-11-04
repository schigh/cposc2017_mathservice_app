FROM golang:1.9-alpine
COPY ./ /go/src/github.com/schigh/cposc2017_app/client
WORKDIR /go/src/github.com/schigh/cposc2017_app/client

RUN apk add --no-cache git make bash
RUN make

CMD gin --bin ./client -i run