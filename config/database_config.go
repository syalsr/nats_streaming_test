package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type config_db struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
}

func ConfigDatabase() *config_db {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &config_db{User: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Host:     os.Getenv("HOST"),
		Port:     port,
		Dbname:   os.Getenv("DATABASE")}
}
