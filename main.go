package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/rociosantos/go-pokech/config"
	"github.com/rociosantos/go-pokech/controller"
	"github.com/rociosantos/go-pokech/router"
	"github.com/rociosantos/go-pokech/service"
	"github.com/rociosantos/go-pokech/usecase"
	"github.com/unrolled/render"

	"github.com/sirupsen/logrus"
)

func main() {
	var publicConfigFile string
	flag.StringVar(&publicConfigFile, "public-config-file",
		"config.yml", "Path to public config file")
	flag.Parse()

	cfg, err := config.LoadConfiguration(publicConfigFile)
	if err != nil {
		log.Fatal("Failed to load config: %w", err)
	}

	// logger setup
	logger, err := createLogger(cfg)
	if err != nil || logger == nil {
		log.Fatal("creating logger: %w", err)
	}

	pokeService := service.NewPokeAPI(cfg.PokeAPIHost, cfg.PokeAPITimeout)

	pokeUseCase := usecase.PokeUseCaseNew(pokeService)

	// Controllers
	healthController := controller.NewHealthController(logger, render.New())
	pokeController := controller.NewPokes(
		pokeUseCase, logger, render.New())

	// Setup application routes
	httpRouter := router.Setup(
		healthController,
		pokeController,
		cfg,
	)

	logger.
		WithField("bind_address", cfg.HTTPPort).
		Info("starting server")
	err = http.ListenAndServe(":"+cfg.HTTPPort, httpRouter)
	if err != nil {
		logger.
			WithError(err).
			Fatal("starting server")
	}

}

func createLogger(cfg *config.Configuration) (*logrus.Logger, error) {
	logLevel := cfg.LogLevel
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_level": logLevel,
		}).Error("parsing log_level")

		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	if cfg.Env != "development" {
		logger.Formatter = &logrus.JSONFormatter{}
	}
	return logger, nil
}
