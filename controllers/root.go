package controllers

import (
	"os"
	"ulfr/models"
)

var (
	dbi, dbiErr = models.NewDatabase()
)

func Init() {
	if dbiErr != nil {
		panic(dbiErr)
	}

	dbi.Migrate(models.Path{})
	dbi.Migrate(models.Domain{})
	dbi.Migrate(models.Fire{})

	CreateDataFolder()
}

func CreateDataFolder() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}
}
