package structs

type AccessUserPasswordParams struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type AccessUserPassword struct {
	Username string `json:"username"`
	Action   string `json:"action"`
}
