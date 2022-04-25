package domain

type CharacterResult struct {
	Results []Character `json:"results"`
}

type Character struct {
	Name    string `json:"name"`
	Species string `json:"species"`
}

type ApiInfo struct {
	Info CharactersInfo `json:"info"`
}

type CharactersInfo struct {
	Pages int32 `json:"pages"`
	Count int32 `json:"count"`
}
