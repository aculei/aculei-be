package models

import (
	"errors"
	"os"
	"strconv"
)

type EnvironmentType string

const (
	Development EnvironmentType = "development"
	Production  EnvironmentType = "production"
)

type DBConfiguration struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Configuration struct {
	Environment EnvironmentType
	HTTPHost    string
	HTTPPort    int
	CORS        CORSConfiguration
	DB          DBConfiguration
}

type CORSConfiguration struct {
	AllowOrigins []string
	AllowHeaders []string
}

func NewConfiguration() Configuration {
	var env EnvironmentType
	string_environment := stringOrPanic("GIN_MODE")

	os.Setenv("GIN_MODE", "development")
	os.Setenv("ACULEI_BE_HTTP_HOST", "0.0.0.0")
	os.Setenv("ACULEI_BE_HTTP_PORT", "8080")

	os.Setenv("ACULEI_BE_DB_HOST", "localhost")
	os.Setenv("ACULEI_BE_DB_PORT", "5432")
	os.Setenv("ACULEI_BE_DB_USER", "admin")
	os.Setenv("ACULEI_BE_DB_PASS", "admin")
	os.Setenv("ACULEI_BE_DB_NAME", "aculei")
	httpHost := stringOrPanic("ACULEI_BE_HTTP_HOST")
	httpPort := intOrPanic("ACULEI_BE_HTTP_PORT")

	if string_environment == "production" {
		env = Production
	} else {
		env = Development
	}

	return Configuration{
		Environment: env,
		HTTPHost:    httpHost,
		HTTPPort:    httpPort,
		CORS: CORSConfiguration{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{},
		},
		DB: DBConfiguration{
			Host: stringOrPanic("ACULEI_BE_DB_HOST"),
			Port: stringOrPanic("ACULEI_BE_DB_PORT"),
			User: stringOrPanic("ACULEI_BE_DB_USER"),
			Pass: stringOrPanic("ACULEI_BE_DB_PASS"),
			Name: stringOrPanic("ACULEI_BE_DB_NAME"),
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
