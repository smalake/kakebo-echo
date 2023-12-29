package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	DBname   string
	Username string
	Password string
	Host     string
	Port     int
}

// func Load() Environment error {
// 	var e Environment
// 	e.Mysql = setMysqlConfig()
// 	return e
// }

var Mc MysqlConfig

func SetMysqlConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	Mc.DBname = os.Getenv("MYSQL_NAME")
	Mc.Username = os.Getenv("MYSQL_USER")
	Mc.Password = os.Getenv("MYSQL_PASSWORD")
	Mc.Host = os.Getenv("MYSQL_HOST")
	Mc.Port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	return nil
}
