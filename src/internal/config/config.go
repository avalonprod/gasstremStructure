package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTPConfig  HTTPConfig
		MongoConfig MongoConfig
		AuthConfig  AuthConfig
	}

	HTTPConfig struct {
		Host           string        `mapstructure:"host"`
		Port           string        `mapstructure:"port"`
		ReadTimeout    time.Duration `mapstructure:"readTimeout"`
		WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderBytes int           `mapstructure:"maxHeaderBytes"`
	}

	MongoConfig struct {
		URI      string
		DBName   string
		Username string
		Password string
	}

	AuthConfig struct {
		PasswordSalt string
	}
)

func Init(configsDir string) (*Config, error) {
	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)
	return &cfg, nil
}

func parseConfigFile(configsDir string) error {
	viper.AddConfigPath(configsDir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("mongo", &cfg.MongoConfig); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTPConfig); err != nil {
		return err
	}
	return nil
}

func setFromEnv(cfg *Config) {
	cfg.MongoConfig.URI = os.Getenv("MONGODB_URI")
	cfg.MongoConfig.Username = os.Getenv("MONGODB_USERNAME")
	cfg.MongoConfig.Password = os.Getenv("MONGODB_PASSWORD")
	cfg.MongoConfig.DBName = os.Getenv("MONGODB_DBNAME")
	cfg.HTTPConfig.Host = os.Getenv("HTTP_HOST")

	cfg.AuthConfig.PasswordSalt = os.Getenv("PASSWORD_SALT")
}
