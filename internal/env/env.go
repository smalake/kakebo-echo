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

type PostgresConfig struct {
	DBName   string
	UserName string
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

var Pc PostgresConfig

func SetPostgresConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	Pc.DBName = os.Getenv("POSTGRES_NAME")
	Pc.UserName = os.Getenv("POSTGRES_USER")
	Pc.Password = os.Getenv("POSTGRES_PASSWORD")
	Pc.Host = os.Getenv("POSTGRES_HOST")
	Pc.Port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	return nil
}
