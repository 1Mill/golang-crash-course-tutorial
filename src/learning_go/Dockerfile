FROM golang:1-alpine

RUN apk update && apk add git

WORKDIR /src
RUN mkdir /app
WORKDIR /app

COPY . .
