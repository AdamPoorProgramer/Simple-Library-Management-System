package configs

import "github.com/caarlos0/env/v6"

type Sqlite struct {
	Path string `env:"PATH"`
}
type Server struct {
	URL string `env:"URL"`
}

type Config struct {
	Sqlite
	Server
}

func LoadConfig() *Config {
	var config *Config = new(Config)
	if err := env.Parse(config); err != nil {
		panic(err)
	}
	return config
}
