package main

import (
	"context"
	"os"

	"github.com/micheledinelli/aculei-be/api"
	"github.com/micheledinelli/aculei-be/api/archive"
	"github.com/micheledinelli/aculei-be/api/experience"
	"github.com/micheledinelli/aculei-be/api/filters"
	"github.com/micheledinelli/aculei-be/db"
	_ "github.com/micheledinelli/aculei-be/docs"
	"github.com/micheledinelli/aculei-be/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
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
		log.Panic().Err(err).Msg("Failed to initialize MongoDB connection")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     configuration.RedisHost,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	repos := mongo.InitRepositories()
	archiveService := archive.NewService(configuration, mongo, &repos.Archive, rdb)
	experienceService := experience.NewService(configuration, mongo, &repos.Experience)
	filtersService := filters.NewService(configuration, mongo, &repos.Filters)

	if err = api.NewServer(
		configuration,
		archiveService,
		experienceService,
		filtersService,
	).Listen(); err != nil {
		os.Exit(1)
	}
}
