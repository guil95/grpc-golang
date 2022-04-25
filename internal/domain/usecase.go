package domain

import (
	"context"
	"fmt"
	"sync"
)

type UseCase struct {
	client RickMortyClient
}

func NewUseCase(client RickMortyClient) *UseCase {
	return &UseCase{
		client: client,
	}
}

func (uc *UseCase) FindAllCharacters(ctx context.Context) ([]Character, error) {
	info, err := uc.client.GetCharactersInfo(ctx)
	if err != nil {
		return nil, err
	}

	var characters []Character
	var wg sync.WaitGroup

	for i := 1; i <= int(info.Pages); i++ {
		wg.Add(1)
		go func(page int32) {
			c, err := uc.client.FindAllByPage(ctx, page)
			if err != nil {
				fmt.Println(err)
				return
			}

			if c[0].Name != "" {
				characters = append(characters, c...)
			}
			wg.Done()
		}(int32(i))
	}
	wg.Wait()

	return characters, nil
}

func (uc *UseCase) GetCharacterInfo(ctx context.Context) (*CharactersInfo, error) {
	characterInfo, err := uc.client.GetCharactersInfo(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return characterInfo, nil
}

func (uc *UseCase) FindCharacterByPage(ctx context.Context, page int32) ([]Character, error) {
	var characters []Character

	characters, err := uc.client.FindAllByPage(ctx, page)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return characters, nil
}
