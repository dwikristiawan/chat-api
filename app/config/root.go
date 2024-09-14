package config

import (
	"github.com/Valgard/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Root struct {
	Server   Server
	Postgres Postgres
}

func Load(filenames ...string) *Root {
	_ = godotenv.Overload(".env")
	r := Root{
		Server:   Server{},
		Postgres: Postgres{},
	}
	mustLoad("SERVER", &r.Server)
	mustLoad("POSTGRES", &r.Postgres)
	return &r
}

func mustLoad(prefix string, spec interface{}) {
	err := envconfig.Process(prefix, spec)
	if err != nil {
		panic(err)
	}
}
