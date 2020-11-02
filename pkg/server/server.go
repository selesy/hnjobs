package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/selesy/hnjobs"
	"github.com/selesy/hnjobs/pkg/store"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Hostname   string
	Port       int
	Reflection bool
}

type Server struct {
	Cfg   Config
	Store store.Store
}

func (s Server) Serve() error {
	lis, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.Cfg.Hostname, s.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

		return err
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)

	if s.Cfg.Reflection {
		reflection.Register(grpcServer)
	}

	pb.RegisterImportServer(grpcServer, &ImportServer{})
	pb.RegisterPostServer(grpcServer, &PostServer{})

	return grpcServer.Serve(lis)
}
