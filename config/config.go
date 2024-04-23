package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		Port string `yaml:"port"`
	} `yaml:"listen"`
	Repository RepoConfig `yaml:"db"`
}

type RepoConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"passwordDB"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {

		instance = &Config{}
		if err := cleanenv.ReadConfig("./config/config.yaml", instance); err != nil {
			panic(err.Error())
		}
	})
	return instance
}
