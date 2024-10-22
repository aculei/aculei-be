package main

import (
	"context"
	"os"

	"github.com/micheledinelli/aculei-be/api"
	"github.com/micheledinelli/aculei-be/api/dataset"
	"github.com/micheledinelli/aculei-be/db"
	_ "github.com/micheledinelli/aculei-be/docs"
	"github.com/micheledinelli/aculei-be/models"
)

// @title aculei-be
// @version 0.0.1
// @description Live to serve aculei.xyz
// @contact.email dinellimichele00@gmail.com
// @contact.name Michele Dinelli

// @host      localhost:8080

func main() {
	var err error

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	configuration := models.NewConfiguration()

	database, err := db.NewDb(ctx, configuration)
	repository := database.InitRepositories()

	datasetService := dataset.NewService(configuration, repository)

	if err = api.NewServer(
		configuration,
		datasetService,
	).Listen(); err != nil {
		os.Exit(1)
	}
}
