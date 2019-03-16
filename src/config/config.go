package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (p *Postgres) Load() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "postgres")
	viper.SetDefault("db.dbname", "postgres")
	viper.AutomaticEnv()

	p.Host = viper.GetString("db.host")
	p.Port = viper.GetInt("db.port")
	p.User = viper.GetString("db.user")
	p.Password = viper.GetString("db.password")
	p.DbName = viper.GetString("db.dbname")
}
