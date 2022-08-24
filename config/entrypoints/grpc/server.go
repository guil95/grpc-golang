package grpc

import (
	"github.com/guil95/grpc-golang/internal/domain"
	"github.com/guil95/grpc-golang/internal/infra/grpc/server"
	"github.com/guil95/grpc-golang/internal/infra/http/clients"
	"log"
	"net"

	pb "github.com/guil95/grpc-golang/api/genpb/rickmorty"
	"google.golang.org/grpc"
)

func RunGrpcServer(quit chan error) {
	lis, err := net.Listen("tcp", "0.0.0.0:50002")
	if err != nil {
		quit <- err
	}
	client := clients.NewRickMortyClient()
	uc := domain.NewUseCase(client)
	s := server.NewServer(uc)
	grpcServer := grpc.NewServer()
	pb.RegisterRickMortyServiceServer(grpcServer, s)

	log.Println("Server running on port 50002")
	if err := grpcServer.Serve(lis); err != nil {
		quit <- err
	}
}
