package usecase

import "github.com/rociosantos/go-pokech/model"

type PokesUseCase struct {
	storage PokeAPIService
}

type PokeAPIService interface {
	GetPokemon(string) (*model.Pokemon, error)
	GetPokemonType(string) (*model.PokemonDamage, error)
}

// PokeUseCaseNew authorization UseCase
func PokeUseCaseNew(
	pokeService PokeAPIService,
) *PokesUseCase {
	return &PokesUseCase{
		storage:           pokeService,
	}
}

// GetDamages - 
func (u *PokesUseCase) GetDamages(name string) (*model.PokemonDamage, error) {
	poke, err := u.storage.GetPokemon(name)
	damage, err := u.storage.GetPokemonType(poke.Types[0].Type.Name)

	return damage, err
}