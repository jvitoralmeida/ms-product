package config

import "github.com/joho/godotenv"

func LoadEnvs() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
