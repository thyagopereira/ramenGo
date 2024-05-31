package util

import "os"

type Configs struct {
	MYSQL_ROOT_PASSWORD string
	MYSQL_DATABASE      string
	MYSQL_PASSWORD      string
	MYSQL_PORT          string
	API_KEY             string
}

func GetConfigs() *Configs {
	return &Configs{
		MYSQL_ROOT_PASSWORD: os.Getenv("MYSQL_ROOT_PASSWORD"),
		MYSQL_DATABASE:      os.Getenv("MYSQL_DATABASE"),
		MYSQL_PASSWORD:      os.Getenv("MYSQL_PASSWORD"),
		MYSQL_PORT:          os.Getenv("MYSQL_PORT"),
		API_KEY:             os.Getenv("API_KEY"),
	}
}
