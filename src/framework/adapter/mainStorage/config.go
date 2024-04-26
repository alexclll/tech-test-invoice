package mainStorage

import "os"

type config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

func Config() config {
	return config{
		Host:         os.Getenv("POSTGRES_HOST"),
		Port:         os.Getenv("POSTGRES_PORT"),
		User:         os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName: os.Getenv("POSTGRES_DATABASE"),
	}
}
