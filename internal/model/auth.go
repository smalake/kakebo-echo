package model

import "database/sql"

type RegisterRequest struct {
	Uid  string `json:"uid" db:"uid"`
	Name string `json:"name" db:"name"`
	Type int    `json:"type" db:"type"`
}

type LoginCheck struct {
	GroupAdmin sql.NullInt32 `db:"group_admin"`
}
