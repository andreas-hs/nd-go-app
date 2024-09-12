package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
	"sync"
)

type Config struct {
	MySQLDSN    string
	RabbitMQURL string
}

var once sync.Once
var config *Config

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		config, err = loadConfig()
	})
	return config, err
}

func loadConfig() (*Config, error) {
	_, filename, _, _ := runtime.Caller(0)
	projectDir := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	viper.SetConfigFile(filepath.Join(projectDir, "local.env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	cfg := &Config{
		MySQLDSN:    viper.GetString("MYSQL_DSN"),
		RabbitMQURL: viper.GetString("RABBITMQ_URL"),
	}

	if cfg.MySQLDSN == "" || cfg.RabbitMQURL == "" {
		return nil, fmt.Errorf("required environment variables are missing")
	}

	return cfg, nil
}
