package main

import (
	"context"
	"os"

	"github.com/micheledinelli/aculei-be/api"
	"github.com/micheledinelli/aculei-be/api/archive"
	"github.com/micheledinelli/aculei-be/db"
	_ "github.com/micheledinelli/aculei-be/docs"
	"github.com/micheledinelli/aculei-be/models"
)

// @title aculei-be
// @version 0.0.1
// @description Live to serve aculei.xyz
// @contact.email dinellimichele00@gmail.com
// @contact.name Michele Dinelli
// @host localhost:8888
// @schemes http https
// @BasePath /

func main() {
	var err error

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	configuration := models.NewConfiguration()

	mongo, err := db.InitDatabase(ctx, configuration.DB.MongoUri)
	if err != nil {
		os.Exit(1)
	}

	repos := mongo.InitRepositories()

	archiveService := archive.NewService(configuration, mongo, &repos.Archive)

	if err = api.NewServer(
		configuration,
		archiveService,
	).Listen(); err != nil {
		os.Exit(1)
	}
}
