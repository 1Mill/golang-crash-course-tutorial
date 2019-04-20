FROM golang:1-alpine

RUN mkdir /app
WORKDIR /app

COPY . .
