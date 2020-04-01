[![Build Status](https://travis-ci.org/moemoe89/practicing-grpc-server-golang.svg?branch=master)](https://travis-ci.org/moemoe89/practicing-grpc-server-golang)
[![codecov](https://codecov.io/gh/moemoe89/practicing-grpc-server-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/practicing-grpc-server-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/practicing-grpc-server-golang)](https://goreportcard.com/report/github.com/moemoe89/practicing-grpc-server-golang)

# practicing-grpc-server-golang #

Practicing GRPC Using Golang (Gin Gonic Framework) as Programming Language, PostgreSQL as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ practicing-grpc-server-golang/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup PostgreSQL <https://www.postgresql.org/>
* Setup BloomRPC <https://github.com/uw-labs/bloomrpc/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> practicing-grpc-server-golang
```

## Running Migration
* Copy `config-sample.json` to `config.json` and changes the value based on your configurations.
* Create PostgreSQL database for example named `simple_api` and do migration with `Goose` <https://bitbucket.org/liamstask/goose/>
* Change database configuration on dbconf.yml like `dialect` and `dsn` for each environtment
* Do the following command
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ goose -env=development up
```

## Running Application with Makefile
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ make build
```
Run
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ make run
```
Stop
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ make stop
```
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Docker Build
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ make docker-build
```
Docker Up
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ make docker-up
```
Docker Down
```
$ cd $GOPATH/src/github.com/moemoe89/practicing-grpc-server-golang
$ make docker-down
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## How to Test gRPC
* Open BloomRPC
* Import proto file
* Fill server address & port
* Do testing

## How to Run Unit Test
Run
```
$ go test ./...
```
Run with cover
```
$ go test ./... -cover
```
Run with HTML output
```
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```

## Reference

Thanks to this medium [link](https://toolbox.kurio.co.id/implementing-grpc-service-in-golang-afb9e05c0064) for sharing the great article

## License

MIT