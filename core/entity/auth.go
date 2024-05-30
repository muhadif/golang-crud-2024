package entity

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken *AccessToken `json:"access_token"`
}

type RegisterRequest struct {
	Email    string
	FullName string
	Password string
	Username string
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
