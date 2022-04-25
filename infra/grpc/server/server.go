package server

import (
	"context"

	pb "github.com/guil95/grpc-golang/api/genpb"
	"github.com/guil95/grpc-golang/domain"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	uc *domain.UseCase
}

func NewServer(uc *domain.UseCase) *server {
	return &server{
		uc: uc,
	}
}

func (s *server) ListAllCharacters(ctx context.Context, request *emptypb.Empty) (*pb.CharactersResponse, error) {
	characters, err := s.uc.FindAllCharacters(ctx)
	if err != nil {
		return nil, err
	}

	return characterListToResponse(characters), nil
}

func (s *server) ListAllCharactersStream(req *emptypb.Empty, srv pb.RickMortyService_ListAllCharactersStreamServer) error {
	characterInfo, err := s.uc.GetCharacterInfo(srv.Context())
	if err != nil {
		return err
	}

	for i := 1; int32(i) <= characterInfo.Pages; i++ {
		characters, err := s.uc.FindCharacterByPage(srv.Context(), int32(i))
		if err != nil {
			return err
		}

		err = srv.Send(characterListToResponse(characters))
		if err != nil {
			return err
		}
	}

	srv.Context().Done()

	return nil
}

func characterListToResponse(characters []domain.Character) *pb.CharactersResponse {
	var charactersResponse []*pb.Character
	for _, c := range characters {
		charactersResponse = append(charactersResponse, &pb.Character{Name: c.Name, Species: c.Species})
	}

	return &pb.CharactersResponse{
		Characters: charactersResponse,
	}
}
