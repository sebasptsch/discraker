package structs

type AccessRefreshJWTParams struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessRefreshJWT struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Source   string `json:"source"`
	Action   string `json:"action"`
}
