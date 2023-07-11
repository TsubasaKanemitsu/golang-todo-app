package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config
type Config struct {
	Prd AppConfig `mapstructure:"prd"`
	Dev AppConfig `mapstructure:"dev"`
}

// AppConfig
type AppConfig struct {
	Postgres Postgres `mapstructure:"postgres"`
}

// Porstgress
type Postgres struct {
	DBName  string `mapstructure:"dbname" validate:"required"`
	Host    string `mapstructure:"host" validate:"required"`
	Pass    string `mapstructure:"pass" validate:"required"`
	Port    string `mapstructure:"port" validate:"required"`
	Sslmode string `mapstructure:"sslmode" validate:"required"`
	User    string `mapstructure:"user" validate:"required"`
}

func Prepare() AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configFilePath, err := getConfigFilePath()
	if err != nil {
		panic(err)
	}
	viper.AddConfigPath(configFilePath)
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	var appConfig AppConfig
	env := viper.GetString("env.name")
	switch env {
	case "prd":
		appConfig = c.Prd
	case "dev":
		appConfig = c.Dev
	default:
		panic(fmt.Sprintf("Unknown env: %s", env))
	}
	return appConfig
}

func getConfigFilePath() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get config file path")
	}
	configDir := filepath.Dir(file)
	pkgDir := filepath.Dir(configDir)
	rootDir := filepath.Dir(pkgDir)
	return rootDir, nil
}
