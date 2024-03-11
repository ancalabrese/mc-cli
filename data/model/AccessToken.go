package model

type AccessToken struct {
	Token string
	RefreshToken string
	ExpiresIn int64
}