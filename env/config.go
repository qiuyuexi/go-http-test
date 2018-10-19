package env

import (
	"os"
)

const (
	DEV = "dev"
)

func GetMysqlConfig(fileName string) MysqlConf {
	mysqlConf := new(MysqlConf)
	path := "env/" + getEnv() + "/" + "mysql/" + fileName
	return mysqlConf.GetConfig(path)
}

func getEnv() string {
	env := os.Getenv("env")
	if env == "" {
		env = DEV
	}
	return env
}
