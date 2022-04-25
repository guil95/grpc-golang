package domain

import "context"

type RickMortyClient interface {
	GetCharactersInfo(ctx context.Context) (*CharactersInfo, error)
	FindAllByPage(ctx context.Context, page int32) ([]Character, error)
}
