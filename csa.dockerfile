FROM golang:latest

MAINTAINER YXH

RUN mkdir -p /app/CSA
WORKDIR /app/CSA
COPY . /app/CSA
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod download

RUN go build main.go

EXPOSE 8081

ENTRYPOINT  ["./main"]