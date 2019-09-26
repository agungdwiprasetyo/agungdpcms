package main

import (
	"fmt"
	"log"
	"net"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"google.golang.org/grpc"
)

func (s *service) ServeGRPC() {
	grpcPort := fmt.Sprintf(":%d", config.GlobalEnv.GRPCPort)

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	fmt.Println("GRPC Server Run at port", grpcPort)

	// register resume module
	s.resumeModule.GRPCHandler.Register(server)

	err = server.Serve(listener)
	if err != nil {
		log.Println("Unexpected Error", err)
	}
}
