FROM golang:1-alpine

RUN apk update && apk add git

COPY . /go

RUN go get github.com/pilu/fresh

WORKDIR /go/src/app
CMD fresh
