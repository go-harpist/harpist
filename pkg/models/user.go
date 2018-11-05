package models

import (
	_ "github.com/go-harpist/harpist/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name          string
	GameProfile   GameProfile
	GameProfileID int
	AuthProfile   AuthProfile
	AuthProfileID int
}

type AuthProfile struct {
	gorm.Model
	Roles    []*Role `gorm:"many2many:user_roles;"`
	Username string
	password string
}

func (ap AuthProfile) Authenticate(username string, password string) (bool, error) {
	var err error = nil
	return ap.password == password, err
}

type GameProfile struct {
	gorm.Model
	Games      []*Game
	Characters []*Character `gorm:"polymorphic:Owner;"`
}

func init() {

}

func (u User) hasRole(r *Role) bool {
	for _, _role := range u.AuthProfile.Roles {
		if _role.ID == r.ID {
			return true
		}
	}
	return false
}

func (u User) isAdmin() bool {
	return u.hasRole(&Role{Name: "admin"})
}
