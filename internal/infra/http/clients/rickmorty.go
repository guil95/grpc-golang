package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/guil95/grpc-golang/internal/domain"
	"io"
	"net/http"
)

type rickMortyClient struct{}

func NewRickMortyClient() *rickMortyClient {
	return &rickMortyClient{}
}

const (
	URI = "http://rickandmortyapi.com/api"
	PATH = "character"
)

func (r *rickMortyClient) GetCharactersInfo(ctx context.Context) (*domain.CharactersInfo, error) {
	url := fmt.Sprintf("%s/%s", URI, PATH)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var charactersInfo domain.ApiInfo

	err = json.Unmarshal(body, &charactersInfo)
	if err != nil {
		return nil, err
	}

	return &charactersInfo.Info, nil
}

func (r *rickMortyClient) FindAllByPage(ctx context.Context, page int32) ([]domain.Character, error) {
	url := fmt.Sprintf("%s/%s?page=%d", URI, PATH, page)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var characters domain.CharacterResult

	err = json.Unmarshal(body, &characters)
	if err != nil {
		return nil, err
	}

	return characters.Results, nil
}
