package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Pg Postgres

	Port string
}

func (c *Config) Load() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	c.Pg.Load()
	viper.SetDefault("svc.port", ":50051")
	c.Port = viper.GetString("svc.port")
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (p *Postgres) Load() {
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.user", "postgres")
	viper.SetDefault("postgres.password", "postgres")
	viper.SetDefault("postgres.dbname", "postgres")

	p.Host = viper.GetString("postgres.host")
	p.Port = viper.GetInt("postgres.Port")
	p.User = viper.GetString("postgres.user")
	p.Password = viper.GetString("postgres.password")
	p.DbName = viper.GetString("postgres.dbname")
}
