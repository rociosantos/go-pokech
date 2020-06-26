package usecase

import "github.com/rociosantos/go-pokech/model"

type PokesUseCase struct {
	storage PokeAPIService
}

type DamageResponse struct {
	Pokemon1       PokemonResponse
	Pokemon2       PokemonResponse
	Damage 		DamageResult
}

type PokemonResponse struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type DamageResult struct{
	DoubleDamageTo bool `json:"double_damage_to`
	HalfDamageFrom bool `json:"half_damage_from`
	NoDamageFrom   bool `json:"no_damage_from`
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
		storage: pokeService,
	}
}

// GetDamages -
func (u *PokesUseCase) GetDamages(name1 string, name2 string) (*DamageResponse, error) {
	poke1, err := u.storage.GetPokemon(name1)
	poke2, err := u.storage.GetPokemon(name2)
	poke1type := poke1.Types[0].Type.Name
	poke2type := poke2.Types[0].Type.Name

	damage, err := u.storage.GetPokemonType(poke1.Types[0].Type.Name)
	damageResponse := DamageResponse{
		Pokemon1: PokemonResponse{
			Name: name1,
			Type: poke1type,
		},
		Pokemon2: PokemonResponse{
			Name: name2,
			Type: poke2type,
		},
	}

	for _, m := range damage.DamageRelations.DoubleDamageTo {
		if m.Name == poke2type {
			damageResponse.Damage.DoubleDamageTo = true
			break
		}
	}
	for _, m := range damage.DamageRelations.HalfDamageFrom {
		if m.Name == poke2type {
			damageResponse.Damage.HalfDamageFrom = true
			break
		}
	}
	for _, m := range damage.DamageRelations.NoDamageFrom {
		if m.Name == poke2type {
			damageResponse.Damage.HalfDamageFrom = true
			break
		}
	}
	

	return &damageResponse, err
}
