//
//  Practicing gRPC
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	"github.com/moemoe89/go-grpc-server-tisa/api/v1/user"
	usrGrpc "github.com/moemoe89/go-grpc-server-tisa/api/v1/user/delivery/grpc"
	conf "github.com/moemoe89/go-grpc-server-tisa/config"

	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	dbR, dbW, err := conf.InitDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := dbR.Close()
		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		err := dbW.Close()
		if err != nil {
			panic(err)
		}
	}()

	log := conf.InitLog()

	userRepo := user.NewPostgresRepository(dbR, dbW)
	userSvc := user.NewService(log, userRepo)

	list, err := net.Listen("tcp", ":" + conf.Configuration.Port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	usrGrpc.NewAUserServerGrpc(server, userSvc)

	fmt.Printf("Listening gRPC server on: %s\n", conf.Configuration.Port)
	err = server.Serve(list)
	if err != nil {
		panic(err)
	}

}
