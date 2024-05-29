package entity

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken *AccessToken
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
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type UserToken struct {
	UserSerial   string
	RefreshToken string
}

type LogoutRequest struct {
	UserSerial string
}
