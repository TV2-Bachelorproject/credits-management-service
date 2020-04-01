package config

import "github.com/kelseyhightower/envconfig"

type DB struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     uint   `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASSWORD" default:"root"`
	Database string `envconfig:"DB_DATABASE" default:"root"`
}

type Config struct {
	DB DB
}

var config Config

func init() {
	envconfig.Process("", &config)
}

func Get() Config {
	return config
}
