package server

import (
	"context"
	"errors"

	pb "github.com/guil95/grpc-golang/api/genpb"
	"github.com/guil95/grpc-golang/internal/domain"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
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

func (s *server) ListAllCharactersServerStream(req *emptypb.Empty, srv pb.RickMortyService_ListAllCharactersServerStreamServer) error {
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

func (s *server) ListByPageClientStream(req pb.RickMortyService_ListByPageClientStreamServer) error {
	var characters []domain.Character
	pages := map[int32][]domain.Character{}

	for {
		ctx := req.Context()
		request, err := req.Recv()

		if err != nil {
			if errors.Is(err, io.EOF) {
				return req.SendAndClose(characterListToResponse(characters))
			}

			return err
		}

		if len(pages[request.Page]) == 0 {
			c, err := s.uc.FindCharacterByPage(ctx, request.Page)
			if err != nil {
				return err
			}

			characters = append(characters, c...)
			pages[request.Page] = characters
		}
	}
}

func characterListToResponse(characters []domain.Character) *pb.CharactersResponse {
	var charactersResponse []*pb.Character
	for _, c := range characters {
		charactersResponse = append(charactersResponse, &pb.Character{Name: c.Name, Species: c.Species})
	}

	return &pb.CharactersResponse{
		Characters: charactersResponse,
		Total: int32(len(charactersResponse)),
	}
}
