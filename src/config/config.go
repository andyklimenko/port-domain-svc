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
	viper.SetDefault("Port", ":50051")
	c.Port = viper.GetString("Port")
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (p *Postgres) Load() {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.Port", "5432")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "postgres")
	viper.SetDefault("db.dbname", "postgres")

	p.Host = viper.GetString("db.host")
	p.Port = viper.GetInt("db.Port")
	p.User = viper.GetString("db.user")
	p.Password = viper.GetString("db.password")
	p.DbName = viper.GetString("db.dbname")
}
