package usecase

import (
	"errors"
	"poke/domain"
	pokeUsecase "poke/feature/poke/usecase"
)

type contestUsecase struct {
	repo domain.ContestRepository
}

func NewContestUsecase(repo domain.ContestRepository) domain.ContestUsecase {
	return &contestUsecase{
		repo: repo,
	}
}

func (u contestUsecase) GetContest(contest_id string) (*domain.Contest, error) {
	return u.repo.GetContest(contest_id)
}

func (u contestUsecase) GetAllContest() ([]domain.Contest, error) {
	return u.repo.GetAllContest()
}

func (u contestUsecase) Contest(cornerRed, cornerBlue string) (map[string]interface{}, error) {
	bluePoke, err := pokeUsecase.PokeUsecaseInstance.GetPoke(cornerBlue)
	if err != nil {
		return nil, err
	} else if bluePoke["name"] == "" {
		return nil, errors.New("blue corner does not exist")
	}

	redPoke, err := pokeUsecase.PokeUsecaseInstance.GetPoke(cornerRed)
	if err != nil {
		return nil, err
	} else if redPoke["name"] == "" {
		return nil, errors.New("blue corner does not exist")
	}

	contest_result := map[string]interface{}{
		"result": "draw",
	}
	if redPoke["damage"].(float32) > bluePoke["damage"].(float32) {
		contest_result["winner"] = redPoke["name"]
		pokeUsecase.PokeUsecaseInstance.UpdatePoke(bluePoke["poke_id"].(string),
			map[string]interface{}{
				"health": bluePoke["health"].(float32) - redPoke["damage"].(float32),
			})

	} else if redPoke["damage"].(float32) < bluePoke["damage"].(float32) {
		contest_result["winner"] = bluePoke["name"]
		pokeUsecase.PokeUsecaseInstance.UpdatePoke(redPoke["poke_id"].(string),
			map[string]interface{}{
				"health": redPoke["health"].(float32) - bluePoke["damage"].(float32),
			})
	}

	return contest_result, nil
}
