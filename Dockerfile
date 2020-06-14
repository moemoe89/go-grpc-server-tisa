FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-grpc-server-tisa

WORKDIR /go/src/github.com/moemoe89/go-grpc-server-tisa

COPY . /go/src/github.com/moemoe89/go-grpc-server-tisa

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/go-grpc-server-tisa
