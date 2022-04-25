package grpc

import (
	"log"
	"net"

	pb "github.com/guil95/grpc-golang/api/genpb"
	"github.com/guil95/grpc-golang/domain"
	"github.com/guil95/grpc-golang/infra/grpc/server"
	"github.com/guil95/grpc-golang/infra/http/clients"
	"google.golang.org/grpc"
)

func RunGrpcServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:50002")
	if err != nil {
		log.Fatal(err)
	}
	client := clients.NewRickMortyClient()
	uc := domain.NewUseCase(client)
	s := server.NewServer(uc)
	grpcServer := grpc.NewServer()
	pb.RegisterRickMortyServiceServer(grpcServer, s)

	log.Println("Server running on port 50002")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
