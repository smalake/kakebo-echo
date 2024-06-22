package user

type UserID struct {
	ID int `json:"id" db:"id"`
}

type LoginRequest struct {
	Uid string `json:"uid" db:"uid"`
}

type LoginGoogleRequest struct {
	Email string `json:"email" db:"email"`
}

type RegisterUserRequest struct {
	Uid  string `json:"uid" db:"uid"`
	Name string `json:"name" db:"name"`
	Type int    `json:"type" db:"type"`
}
