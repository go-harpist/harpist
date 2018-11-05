package models

import (
	_ "github.com/go-harpist/harpist/pkg/config"
	"github.com/jinzhu/gorm"
)

type Character struct {
	gorm.Model
	Name      string
	OwnerID   int
	OwnerType string
}

type Game struct {
	gorm.Model
	Name           string
	Slug           string
	OwnerID        int
	OwnerType      string
	GameCharacters []*Character
}
