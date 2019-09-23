package module

import (
	"web-layout/utils/mysql"
	"web-layout/utils/redis"
)

// Config : config items read form config file
type Config struct {
	Mysql *mysql.Config
	Redis *redis.Config
}
