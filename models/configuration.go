package models

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type EnvironmentType string

const (
	Development EnvironmentType = "development"
	Production  EnvironmentType = "production"
)

type Configuration struct {
	Environment          EnvironmentType
	HTTPHost             string
	HTTPPort             int
	CORS                 CORSConfiguration
	GracefulExitDuration time.Duration
}

type CORSConfiguration struct {
	AllowOrigins []string
	AllowHeaders []string
}

func NewConfiguration() Configuration {
	var env EnvironmentType

	os.Setenv("GIN_MODE", "development")
	os.Setenv("ACULEI_BE_HTTP_HOST", "0.0.0.0")
	os.Setenv("ACULEI_BE_HTTP_PORT", "8080")
	string_environment := stringOrPanic("GIN_MODE")

	if string_environment == "production" {
		env = Production
	} else {
		env = Development
	}

	httpHost := stringOrPanic("ACULEI_BE_HTTP_HOST")
	httpPort := intOrPanic("ACULEI_BE_HTTP_PORT")

	return Configuration{
		Environment: env,
		HTTPHost:    httpHost,
		HTTPPort:    httpPort,
		CORS: CORSConfiguration{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{},
		},
	}
}

func stringOrPanic(key string) string {
	var result, found = os.LookupEnv(key)
	if !found {
		panic(errors.New("configuration value not set for key: " + key))
	}
	return result
}

func intOrPanic(key string) int {
	var result, found = os.LookupEnv(key)

	if !found {
		panic(errors.New("configuration value not set for key: " + key))
	}

	intResult, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		panic(errors.New("configuration value for key: " + key + " is not a int"))
	}

	return int(intResult)
}
