package auth

type JWT struct {
	Token string `json:"token"`
}

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type App interface {
// 	LoginMail(ctx echo.Context) error
// }

type Register struct {
	Mail     string `json:"mail"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

type ParentName struct {
	Name string `json:"name"`
}

type AuthCode struct {
	Code  int    `json:"code"`
	Email string `json:"email"`
}
