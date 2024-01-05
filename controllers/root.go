package controllers

import (
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
}
