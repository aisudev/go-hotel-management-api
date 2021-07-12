package usecase

import (
	"errors"
	"math/rand"
	"poke/domain"

	"github.com/gofrs/uuid"
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

func (u *pokeUsecase) CreatePoke(uid string, specie_id uint, name string) error {
	poke_id, _ := uuid.NewV1()
	poke := domain.Poke{
		Poke_id:   poke_id.String(),
		UUID:      uid,
		Specie_id: specie_id,
		Name:      name,
		Exp:       0,
		Health:    100,
		MaxHealth: 100,
		Damage:    float32(rand.Intn(10) + 5),
	}

	return u.repo.CreatePoke(&poke)
}

func (u *pokeUsecase) GetPoke(poke_id string) (map[string]interface{}, error) {
	poke, err := u.repo.GetPoke(poke_id)
	if err != nil {
		return nil, err
	}

	if poke.Poke_id == "" {
		return nil, errors.New("empty pokes.")
	}

	return map[string]interface{}{
		"poke_id":   poke.Poke_id,
		"specie_id": poke.Specie_id,
		"name":      poke.Name,
		"health":    poke.Health,
		"damage":    poke.Damage,
	}, nil
}

func (u *pokeUsecase) GetAllPoke(uuid string) ([]map[string]interface{}, error) {
	pokes, err := u.repo.GetAllPoke(uuid)
	if err != nil {
		return nil, err
	}

	result := []map[string]interface{}{}
	for _, poke := range pokes {
		result = append(result, map[string]interface{}{
			"poke_id":    poke.Poke_id,
			"specie_id":  poke.Specie_id,
			"name":       poke.Name,
			"health":     poke.Health,
			"damage":     poke.Damage,
			"max_health": poke.MaxHealth,
		})
	}

	return result, nil
}

func (u *pokeUsecase) ChangeNamePoke(poke_id string, name string) error {
	return u.repo.UpdatePoke(poke_id, map[string]interface{}{
		"name": name,
	})
}

// TODO: ADD WITHDRAW BALANCE
func (u *pokeUsecase) TreatPoke(poke_id string) error {
	poke, err := u.GetPoke(poke_id)
	if err != nil {
		return err
	}

	return u.repo.UpdatePoke(poke_id, map[string]interface{}{
		"health": poke["max_health"],
	})
}

func (u *pokeUsecase) DeletePoke(poke_id string) error {
	return u.repo.DeletePoke(poke_id)
}

func (u *pokeUsecase) VerifyPoke(poke_id string) error {
	return nil
}
