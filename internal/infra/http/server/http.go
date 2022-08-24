package server

import (
	"github.com/gin-gonic/gin"
	"github.com/guil95/grpc-golang/internal/domain"
	"net/http"
)

type httpServer struct {
	uc      *domain.UseCase
	handler *gin.Engine
}

func NewHTTPServer(uc *domain.UseCase, handler *gin.Engine) *httpServer {
	return &httpServer{uc, handler}
}

func (server httpServer) CreateApi() {
	server.handler.GET("/", func(ctx *gin.Context) {
		server.listAllCharacters(ctx)
	})
}

type ListAllResponse struct {
	Characters []domain.Character `json:"characters"`
	Total      int                `json:"total"`
}

func (server httpServer) listAllCharacters(ctx *gin.Context) {
	characters, err := server.uc.FindAllCharacters(ctx)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, ListAllResponse{characters, len(characters)})
}
