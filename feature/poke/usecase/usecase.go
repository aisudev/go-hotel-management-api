package usecase

import (
	"poke/domain"

	"github.com/mtslzr/pokeapi-go"
)

type pokeUsecase struct {
	repo domain.PokeRepository
}

func NewPokeUsecase(repo domain.PokeRepository) domain.PokeUsecase {
	return &pokeUsecase{
		repo: repo,
	}
}

func (u *pokeUsecase) GetMorePokeAPI(offset int, limit int) ([]map[string]interface{}, error) {
	resp, err := pokeapi.Resource("pokemon", offset, limit)

	if err != nil {
		return nil, err
	}

	pokemons := []map[string]interface{}{}

	for _, poke := range resp.Results {
		poke, _ := pokeapi.Pokemon(poke.Name)

		pokemons = append(pokemons, map[string]interface{}{
			"specie_id": poke.ID,
			"name":      poke.Name,
			"imgUrls":   poke.Sprites,
		})
	}

	return pokemons, nil
}

func (u *pokeUsecase) GetPokeAPI(name string) ([]map[string]interface{}, error) {
	resp, err := pokeapi.Search("pokemon", name)

	if err != nil {
		return nil, err
	}

	pokemons := []map[string]interface{}{}

	for _, poke := range resp.Results {
		poke, _ := pokeapi.Pokemon(poke.Name)

		pokemons = append(pokemons, map[string]interface{}{
			"specie_id": poke.ID,
			"name":      poke.Name,
			"imgUrls":   poke.Sprites,
		})
	}

	return pokemons, nil
}

func (u *pokeUsecase) GetPokeImageAPI(name string) (map[string]interface{}, error) {

	resp, err := pokeapi.Pokemon(name)

	if err != nil {
		return nil, err
	}

	images := map[string]interface{}{
		"back_default":       resp.Sprites.BackDefault,
		"back_female":        resp.Sprites.BackFemale,
		"back_shiny":         resp.Sprites.BackShiny,
		"back_shiny_female":  resp.Sprites.BackShinyFemale,
		"front_default":      resp.Sprites.FrontDefault,
		"front_female":       resp.Sprites.FrontFemale,
		"front_shiny":        resp.Sprites.FrontShiny,
		"front_shiny_female": resp.Sprites.FrontShinyFemale,
	}

	return images, nil
}

func (u *pokeUsecase) CreatePoke(specie_id uint, name string) error {
	return nil
}

func (u *pokeUsecase) GetPoke(poke_id string) (map[string]interface{}, error) {
	return nil, nil
}

func (u *pokeUsecase) GetAllPoke() ([]map[string]interface{}, error) {
	return nil, nil
}

func (u *pokeUsecase) UpdatePoke(poke_id string, newPoke map[string]interface{}) error {
	return nil
}

func (u *pokeUsecase) DeletePoke(poke_id string) error {
	return nil
}

func (u *pokeUsecase) VerifyPoke(poke_id string) error {
	return nil
}
