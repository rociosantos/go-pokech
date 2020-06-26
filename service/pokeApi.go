package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/rociosantos/go-pokech/model"

	"gopkg.in/resty.v1"
)

type Service struct {
	client   *resty.Client
}

const (
	pokemonEndpoint = "/pokemon/{name}"
	typeEndpoint = "/type/{name}"
)

// NewPokeApi - Creates a new Pokes API Client
func NewPokeApi(
	host string,
	timeout time.Duration,
) *Service {
	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout).
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			if r.IsSuccess() {
				return nil
			}
			return errors.New(strconv.Itoa(r.StatusCode()))
			
		})

	return &Service{client}
}

//GetPokemon - 
func (s *Service) GetPokemon(name string) (*model.Pokemon, error) {
	response, err := s.client.R().
		SetPathParams(map[string]string{"name": name}).
		Get(pokemonEndpoint)

	if err != nil {
		return nil, errors.New("getting pokemon")
	}

	poke := model.Pokemon{}
	body := response.Body()
	if err := json.Unmarshal(body, &poke); err != nil {
		return nil, errors.New("unmarshaling response")
	}

	return &poke, nil
}

//GetPokemonType - 
func (s *Service) GetPokemonType(typeName string) (*model.PokemonDamage, error) {
	response, err := s.client.R().
		SetPathParams(map[string]string{"name": typeName}).
		Get(typeEndpoint)

	if err != nil {
		return nil, errors.New("getting pokemon")
	}

	pokeType := model.PokemonDamage{}
	body := response.Body()
	if err := json.Unmarshal(body, &pokeType); err != nil {
		return nil, errors.New("unmarshaling response")
	}

	return &pokeType, nil
}