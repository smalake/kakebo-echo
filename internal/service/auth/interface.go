package auth

import "database/sql"

type JWT struct {
	Token string `json:"token"`
}

type LoginCheck struct {
	GroupAdmin sql.NullInt32 `db:"group_admin"`
}

type ParentName struct {
	Name string `json:"name"`
}
