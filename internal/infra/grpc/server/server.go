package server

import (
	"context"
	domain2 "github.com/guil95/grpc-golang/internal/domain"

	pb "github.com/guil95/grpc-golang/api/genpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	uc *domain2.UseCase
}

func NewServer(uc *domain2.UseCase) *server {
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

func characterListToResponse(characters []domain2.Character) *pb.CharactersResponse {
	var charactersResponse []*pb.Character
	for _, c := range characters {
		charactersResponse = append(charactersResponse, &pb.Character{Name: c.Name, Species: c.Species})
	}

	return &pb.CharactersResponse{
		Characters: charactersResponse,
	}
}
