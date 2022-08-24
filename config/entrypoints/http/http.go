package http

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/grpc-golang/internal/domain"
	"github.com/guil95/grpc-golang/internal/infra/http/clients"
	"github.com/guil95/grpc-golang/internal/infra/http/server"
	"log"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func RunHTTPServer(quit chan error) {
	handler := gin.Default()

	client := clients.NewRickMortyClient()
	uc := domain.NewUseCase(client)
	s := server.NewHTTPServer(uc, handler)
	s.CreateApi()

	log.Println("Running http server on :8080 port")
	
	if err := handler.Run(":8080"); err != nil {
		quit <- err
	}
}
