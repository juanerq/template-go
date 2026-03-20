package config

import env "github.com/caarlos0/env/v11"

type EnvConfig struct {
	Port     int    `env:"PORT" envDefault:"8080"`
	BasePath string `env:"BASE_PATH,required,notEmpty"`
}

func LoadConfig() (*EnvConfig, error) {
	cfg := EnvConfig{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
