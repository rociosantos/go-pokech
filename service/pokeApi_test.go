package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/rociosantos/go-pokech/model"
)

var (
	testPokemon, _ = ioutil.ReadFile("./testdata/testResponse.json")
	testType, _ = ioutil.ReadFile("./testdata/testTypeResponse.json")
)



func TestService_GetPokemon(t *testing.T) {
	timeout := 2 * time.Second
	poke := model.Pokemon{}
	json.Unmarshal(testPokemon, &poke)

	tests := []struct {
		name           string
		pokeName       string
		statusResponse int
		rawResponse    []byte
		want           *model.Pokemon
		wantErr        bool
	}{
		{
			name:           "Happy path",
			pokeName:       "pikachu",
			statusResponse: http.StatusOK,
			rawResponse:    testPokemon,
			want: &poke,
			wantErr: false,
		},
		{
			name:           "Unhappy path",
			pokeName:       "badpikachu",
			statusResponse: http.StatusNotFound,
			rawResponse:    []byte("error"),
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(tt.rawResponse))
				}))
			defer serv.Close()

			s := NewPokeAPI(serv.URL, timeout)

			got, err := s.GetPokemon(tt.pokeName)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}


func TestService_GetPokemonType(t *testing.T) {
	timeout := 2 * time.Second
	pokeType := model.PokemonDamage{}
	json.Unmarshal(testType, &pokeType)

	tests := []struct {
		name           string
		pokeName       string
		statusResponse int
		rawResponse    []byte
		want           *model.PokemonDamage
		wantErr        bool
	}{
		{
			name:           "Happy path",
			pokeName:       "pikachu",
			statusResponse: http.StatusOK,
			rawResponse:    testType,
			want: &pokeType,
			wantErr: false,
		},
		{
			name:           "Unhappy path",
			pokeName:       "badpikachu",
			statusResponse: http.StatusNotFound,
			rawResponse:    []byte("error"),
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(tt.rawResponse))
				}))
			defer serv.Close()

			s := NewPokeAPI(serv.URL, timeout)
			
			got, err := s.GetPokemonType(tt.pokeName)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

