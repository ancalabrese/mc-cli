package data

type AccessToken struct {
	Token        string
	RefreshToken string
	ExpiresIn    int64
}
