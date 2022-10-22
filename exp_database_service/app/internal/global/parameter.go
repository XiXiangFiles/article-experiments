package global

import (
	"os"
)

type env string

var (
	ServeAddress  env
	MySQL_DSN     env
	MySQL_DATA_DB env
	LOG_LEVEL     env
)

func SetupEnv() {
	ServeAddress = env(os.Getenv("SERVE_ADDRESS"))
	MySQL_DSN = env(os.Getenv("MYSQL_DSN"))
	MySQL_DATA_DB = env(os.Getenv("MYSQL_DATA_DB"))
	LOG_LEVEL = env(os.Getenv("LOG_LEVEL"))
}

func (s env) ToString() string {
	return string(s)
}
