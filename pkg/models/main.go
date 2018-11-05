package models

import (
	"fmt"
	"github.com/go-harpist/harpist/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

var defaultConfig = config.GetDefaultConfig

func RolesInit(cf config.Config, db *gorm.DB) {
	var initialRoles = []Role{}

	for _, _roleConfig := range cf.DefaultRoles {
		initialRoles = append(initialRoles, Role{Name: _roleConfig.RoleName, Permissions: _roleConfig.RolePermissions})
	}

	for _, _role := range initialRoles {
		_ = db.FirstOrCreate(&_role, _role)
	}
}
func GetDB(cf *config.Config) *gorm.DB {
	db, err := gorm.Open(cf.DBConfig.Dialect, cf.DBConfig.Args)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	db.LogMode(cf.LogConfig.LogMode)
	db.SetLogger(cf.LogConfig.LogFile)
	return db
}

func runInit() {
	cf := defaultConfig()
	db := GetDB(&cf)
	db.AutoMigrate(User{}, Character{}, Game{}, GameProfile{}, Role{}, AuthProfile{})
	RolesInit(cf, db)
	defer db.Close()
}

var InitializeFirstRun = runInit

func main() {

}
