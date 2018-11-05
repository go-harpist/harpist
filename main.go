package main

import (
	"github.com/go-harpist/harpist/pkg/config"
	"github.com/go-harpist/harpist/pkg/models"
)

func main() {
	_ = config.GetDefaultConfig()
	models.InitializeFirstRun()
}
