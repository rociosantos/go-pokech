package controller

import (
	"net/http"

	"github.com/rociosantos/go-pokech/model"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// PokesUseCase interface, handles every key flow
type PokesUseCase interface {
	GetDamages(string) (*model.PokemonDamage, error)
}

// Pokes controller struct
type Pokes struct {
	useCase  PokesUseCase
	logger   *logrus.Logger
	render   *render.Render
}

// NewPokes returns a pokes controller
func NewPokes(
	u PokesUseCase,
	logger *logrus.Logger,
	r *render.Render,
) *Pokes {
	return &Pokes{u, logger, r}
}

func (p *Pokes) GetDamages(w http.ResponseWriter, r *http.Request){
	pathParams := mux.Vars(r)
	poke1 := pathParams["poke1"]
	// poke2 := pathParams["poke2"]

	p.logger.WithField("func","Get damages").Info("in")

	damage, err := p.useCase.GetDamages(poke1)
	if err != nil {
		p.logger.WithError(err).Error("getting damage")
		p.render.JSON(w, http.StatusNotFound, nil)
		return
	}

	p.render.JSON(w, http.StatusOK, damage)

}