FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-grpc-server-golang

WORKDIR /go/src/github.com/moemoe89/practicing-grpc-server-golang

COPY . /go/src/github.com/moemoe89/practicing-grpc-server-golang

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/practicing-grpc-server-golang
