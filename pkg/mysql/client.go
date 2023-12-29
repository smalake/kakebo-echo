package mysql

import (
	"fmt"
	"kakebo-echo/internal/env"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	DB *sqlx.DB
}

func NewClient() (*Client, error) {
	// MySQLの設定
	err := env.SetMysqlConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.Mc.Username,
		env.Mc.Password,
		env.Mc.DBname,
	)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	c := &Client{
		DB: db,
	}
	return c, nil
}
