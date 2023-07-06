package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBName  string
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	WebPort string
	WebAddr string
}

func GetConfig() Config {
	v := viper.New()

	v.SetEnvPrefix("APP_")
	v.SetDefault("DBNAME", "postgres")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "")
	v.SetDefault("DBHOST", "127.0.0.1")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("WEBPORT", "8080")
	v.SetDefault("WEBADDR", "127.0.0.1")
	v.AutomaticEnv()

	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func (c *Config) GetDbConn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
}
