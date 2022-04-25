package main

import "github.com/guil95/grpc-golang/pkg/grpc"

func main() {
	//client := clients.NewRickMortyClient()
	//uc := domain.NewUseCase(client)

	//characters, err := uc.FindAllCharacters()
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(characters)

	//characterInfo, err := uc.GetCharacterInfo()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(characterInfo)

	//characterByPage, err := uc.FindCharacterByPage(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(characterByPage)

	grpc.RunGrpcServer()
}
