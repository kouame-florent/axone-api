package server

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/kouame-florent/axone-api/api/grpc/gen"
	"github.com/kouame-florent/axone-api/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	port = ":50051"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid credentials")
)

func (s *AxoneServer) StartGrpcServer(opts []grpc.ServerOption, db *gorm.DB) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	axSvr := NewAxoneServer(db)

	if err != nil {
		log.Fatal(err)
	}

	g := grpc.NewServer(opts...)
	gen.RegisterAxoneServer(g, axSvr)

	log.Printf("Starting gRPC listener on port " + port)
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *AxoneServer) ServerOptions() ([]grpc.ServerOption, error) {
	cert, err := tls.LoadX509KeyPair(config.ServerCertFile, config.ServerKeyFile)
	if err != nil {
		return []grpc.ServerOption{}, nil
	}
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.UnaryInterceptor(s.ensureValidBasicCredentials),
	}

	return opts, nil
}
