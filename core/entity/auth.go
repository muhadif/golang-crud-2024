package entity

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken *AccessToken `json:"access_token"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string
}

type AccessToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AtExpires    int64  `json:"atExpires"`
	RtExpires    int64  `json:"rtExpires"`
}

type UserToken struct {
	UserSerial   string
	RefreshToken string
}

type LogoutRequest struct {
	UserSerial string
}
