package appmodels

import (
	"kakebo-echo/pkg/postgresql"
)

type AppModel struct {
	PsgrCli *postgresql.Client
}

func New(psgrCli *postgresql.Client) *AppModel {
	return &AppModel{PsgrCli: psgrCli}
}
