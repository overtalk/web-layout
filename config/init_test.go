package config_test

import (
	"os"
	"testing"

	"web-layout/config"
)

type MySQL struct {
	Host     string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	Port     int    `env:"MYSQL_PORT" envDefault:"3306"`
	Password string `env:"MYSQL_PASSWORD,required"`
}

type Redis struct {
	Addr     []string `env:"REDIS_ADDR" envSeparator:":" envDefault:"127.0.0.1:6379"`
	Password string   `env:"REDIS_PORT"`
}

type GlobalConfig struct {
	Mysql *MySQL
	Redis *Redis
}

func TestNewConfig(t *testing.T) {
	os.Setenv("MYSQL_PASSWORD", "TEST_MYSQL_PORT")

	c := GlobalConfig{
		Mysql: new(MySQL),
		Redis: new(Redis),
	}

	if err := config.NewConfig(config.PROD, &c); err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", c.Mysql)
}
