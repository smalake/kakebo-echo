package user

import "kakebo-echo/internal/appmodels"

type User struct {
	Email    string
	Password string
	Name     string
	GroupId  int
	Rtype    int
}

type Service struct {
	appModel appmodels.AppModel
}

func New(appModel appmodels.AppModel) *Service {
	return &Service{appModel: appModel}
}
