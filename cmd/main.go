package main

import (
	"context"

	"github.com/Beknur1003/clean-architecture.git/internal/composites"
	"github.com/Beknur1003/clean-architecture.git/internal/config"
	"github.com/Beknur1003/clean-architecture.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// entry point
	// logging.Init()
	logger := logging.GetLogger("")

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoDBComposite, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(mongoDBComposite)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	authorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBComposite)
	if err != nil {
		logger.Fatal("book composite failed")
	}
	bookComposite.Handler.Register(router)

	// start
}
