package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Port string `yaml:"port" env:"APP_PORT" env-default:"8080"`
}

type DB struct {
	Host   string `yaml:"host" env:"DB_HOST" env-default:"postgres"`
	Port   string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name   string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
	User   string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Pass   string `yaml:"pass" env:"DB_PASS" env-default:"test"`
	Schema string `yaml:"schema" env:"DB_SCHEMA" env-default:"public"`
}

type Log struct {
	Log string `yaml:"level" env:"LOG_LEVEL" env-default:"INFO"`
}

type Config struct {
	DB  `yaml:"db"`
	Log `yaml:"log"`
	App `yaml:"app"`
}

func (cfg Config) GetDbConfig() DB {
	return cfg.DB
}

func (c DB) GetDsn() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
	)
}

func (c DB) GetMigrateDsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
		c.Schema,
	)
}

func New() (Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
