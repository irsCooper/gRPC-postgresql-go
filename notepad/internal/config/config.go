package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string    `yaml:"env" env-default:"local"`
	StoragePath string    `yaml:"storage_path" env-required:"true"`
	GRPC        GRPConfig `yaml:"grpc"`
}

type GRPConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	return MostLoadByPath(path)
}

func MostLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does'n exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("notepad/config/local.yaml")
	}

	return res
}
