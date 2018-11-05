package models

import (
	"github.com/go-harpist/harpist/pkg/config"
	"github.com/jinzhu/gorm"
	"os"
	"testing"
	//"log"
)

var testLogConfig = config.LogConfig{
	LogMode: false,
	LogFile: nil, // log.New(os.Stdout, "\r\n", 0),
}

func initTestDB(dbFilename string) *gorm.DB {
	cf := config.Config{
		DBConfig: config.DBConfig{
			Dialect: "sqlite3",
			Args:    dbFilename,
		},
		DefaultRoles: config.GetInitialRoles(),
		LogConfig:    testLogConfig}
	db := GetDB(&cf)
	db.AutoMigrate(User{}, Character{}, Game{}, GameProfile{}, Role{}, AuthProfile{})
	RolesInit(cf, db)
	return db
}

func TestDatabaseInit(t *testing.T) {

	t.Run("db:init", func(t *testing.T) {
		t.Run("db1:create", func(t *testing.T) {
			db := initTestDB("./test1.db")
			if _, err := os.Stat("./test1.db"); os.IsNotExist(err) {
				t.Fatal("Error: ", err)
			}
			defer db.Close()
		})
		t.Run("db2:assert", func(t *testing.T) {
			db := initTestDB("./test2.db")
			if _, err := os.Stat("./test3.db"); os.IsExist(err) {
				t.Fatal("Error: ", "Created the test.db3 file when it shouldn't have!")
			}
			defer db.Close()
		})
		t.Run("db3:adminRole", func(t *testing.T) {
			db := initTestDB("./test3.db")
			role := Role{}
			db.First(&role)
			if role.Name != "admin" {
				t.Fatal("Error: ", "Default first role is not `admin`!")
			}
			defer db.Close()
		})

	})

	os.Remove("./test1.db")
	os.Remove("./test2.db")
	os.Remove("./test3.db")

}
