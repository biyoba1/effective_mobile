package initializer

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loadin .env variables")
	}
}
