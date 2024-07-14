package postgresql

import (
	"fmt"
	"kakebo-echo/internal/env"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client struct {
	DB *sqlx.DB
}

func (c *Client) GetDB() *sqlx.DB {
	return c.DB
}

func NewClient() (*Client, error) {
	err := env.SetPostgresConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Pc.Host,
		env.Pc.UserName,
		env.Pc.Password,
		env.Pc.DBName,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Client{DB: db}, nil
}

func (c *Client) Close() error {
	return c.DB.Close()
}
