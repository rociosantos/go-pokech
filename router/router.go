package router

import (
	"net/http"

	"github.com/rociosantos/go-pokech/config"

	mux "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

// HealthController for monitoring health of our app
type HealthController interface {
	IsHealthy(w http.ResponseWriter, r *http.Request)
}

// PokesController for monitoring health of our app
type PokesController interface {
	GetDamages(w http.ResponseWriter, r *http.Request)
	GetMoves(w http.ResponseWriter, r *http.Request)
}

// Setup returns router instance which is used in main package to register handlers.
func Setup(
	healthController HealthController,
	pokesController PokesController,
	cfg *config.Configuration,
) *mux.Router {
	r := mux.NewRouter(mux.WithServiceName(cfg.AppName))

	r.HandleFunc("/healthz", healthController.IsHealthy).Methods(http.MethodGet).Name("healthz")
	r.HandleFunc("/pokes/{poke1}/damages/{poke2}", pokesController.GetDamages).
		Methods("GET").Name("get-damages")
	r.HandleFunc("/pokes/{poke1}/moves/{poke2}", pokesController.GetMoves).Methods("GET").
		Name("get-damages")

	return r
}
