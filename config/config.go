package config

import (
	"strconv"

	"github.com/danish45007/go-dynamodb/utils/env"
)

type Config struct {
	Port        int
	Timeout     int
	Dialect     string
	DatabaseURI string
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(env.GetEnvVariable(envName, defaultValue))
	if err != nil {
		return 0
	}
	return num
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInt("PORT", "8080"),
		Timeout:     parseEnvToInt("TIMEOUT", "30"),
		Dialect:     env.GetEnvVariable("DIALECT", "sqlite"),
		DatabaseURI: env.GetEnvVariable("DATABASE_URI", ":memory:"),
	}
}
