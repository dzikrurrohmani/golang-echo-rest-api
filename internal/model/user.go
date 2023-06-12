package model

type User struct {
	ID       string `json:"id"`
	Username string `gorm:"unique" json:"username"`
	Hash     string `json:"-"`
}

type RegisterRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	JWTToken string `json:"jwt_token"`
}
