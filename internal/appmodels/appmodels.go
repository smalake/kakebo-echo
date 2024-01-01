package appmodels

import (
	"kakebo-echo/pkg/mysql"
)

type AppModel struct {
	MysqlCli *mysql.Client
}

func New(mysqlCli *mysql.Client) *AppModel {
	return &AppModel{MysqlCli: mysqlCli}
}
