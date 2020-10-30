package server

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/selesy/hnjobs"
	log "github.com/sirupsen/logrus"
)

var port = flag.Int("-p", 8080, "")

func Serve() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	pb.RegisterImportServer(grpcServer, &ImportServer{})
	pb.RegisterPostServer(grpcServer, &PostServer{})

	grpcServer.Serve(lis)
}
