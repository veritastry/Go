# docker build -t 17199911145/gin-demo .
FROM golang:latest

WORKDIR /server
COPY . /server

ENV GOPROXY="https://goproxy.io"
RUN go mod download
RUN go build -x -o /server/build/server

EXPOSE 8080

ENTRYPOINT ["/server/build/server"]
