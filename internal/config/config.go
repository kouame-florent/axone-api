package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/kouame-florent/axone-api/internal/store"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const DefaultConfig = `
AXONE:
  DEBUG:
    DSN: "root:axone8!@tcp(127.0.0.1:3306)/axone_test?charset=utf8mb4&parseTime=True&loc=Local"
  TEST:
    DSN: "root:axone8!@tcp(127.0.0.1:3306)/axone_test?charset=utf8mb4&parseTime=True&loc=Local"
  PROD:
    DSN: "axone:axone@tcp(127.0.0.1:3306)/axone?charset=utf8mb4&parseTime=True&loc=Local"
`

type environment string

const (
	DEBUG      environment = "debug"
	TEST       environment = "test"
	PRODUCTION environment = "prod"
)

func createConfigFile(home string) {
	content := []byte(strings.TrimPrefix(DefaultConfig, "\n"))
	cfgFile := path.Join(home, ".axone.yaml")
	if _, err := os.Stat(cfgFile); err != nil {
		err = os.WriteFile(cfgFile, content, 0664)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func InitConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	createConfigFile(home)

	// Search config in home directory with name ".icens" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".axone")

	//set default env
	viper.SetDefault("AXONE_ENV", DEBUG)

	viper.AutomaticEnv() // read in environment variables that match

	viper.BindEnv("AXONE_ENV", "AXONE_ENV")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}
}

func DataSourceName() string {
	InitConfig()

	env := viper.GetString("AXONE_ENV")
	if env == "" {
		log.Fatalf("got empty AXONE_ENV")
	}
	log.Printf("AXONE_ENV=%s", env)
	var dsn string

	switch env {
	case string(DEBUG):
		dsn = viper.GetString("AXONE.DEBUG.DSN")
	case string(TEST):
		dsn = viper.GetString("AXONE.TEST.DSN")
	case string(PRODUCTION):
		dsn = viper.GetString("AXONE.PROD.DSN")
	}

	return dsn
}

//Read from environment viriable ICENS_DEFAULT_LOCALITE
func DefaultLocaliteCode() string {
	return "cm-360"
}

func DefaultCentreCode() string {
	return "ct-01"
}

func GetDB() *gorm.DB {
	dsn := DataSourceName()
	return store.OpenDB(dsn)
}