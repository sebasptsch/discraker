package structs

type AccessLoginParams struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Source   *string `json:"source,omitempty"`
}

type AccessLogin struct {
	Username     *string `json:"username"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
	Action       *string `json:"action"`
	Source       *string `json:"source"`
}
