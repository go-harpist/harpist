package config

import (
	"log"
	"os"
)

type DBConfig struct {
	Dialect string
	Args    string
}

type roleConfig struct {
	RoleName        string
	RolePermissions string
}

type LogConfig struct {
	LogMode bool
	LogFile *log.Logger
}

type Config struct {
	DBConfig     DBConfig
	DefaultRoles []roleConfig
	LogConfig    LogConfig
}

var GetDefaultConfig = getDefaultConfig
var GetInitialRoles = getInitialRoles
var initialRoles []roleConfig

func getDefaultDbConfig() DBConfig {
	dbDialect, dbArgs := func() (string, string) {
		dbD, ok := os.LookupEnv("HARPIST_DB_DIALECT")
		if !ok {
			dbD = "sqlite3"
		}
		dbA, ok := os.LookupEnv("HARPIST_DB_ARGS")
		if !ok {
			dbA = "/tmp/gorm.db"
		}
		return dbD, dbA
	}()
	return DBConfig{dbDialect, dbArgs}
}

func getInitialRoles() []roleConfig {
	return append(initialRoles, roleConfig{"admin", ""})
}

func getDefaultConfig() Config {
	return Config{
		getDefaultDbConfig(),
		getInitialRoles(),
		LogConfig{true,
			log.New(os.Stdout, "\r\n", 0)}}
}

func main() {

}
