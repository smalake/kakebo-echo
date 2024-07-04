package appmodel

import (
	"kakebo-echo/pkg/database/postgresql"
)

type AppModel struct {
	PsgrCli *postgresql.Client
}

func New(psgrCli *postgresql.Client) *AppModel {
	return &AppModel{PsgrCli: psgrCli}
}
