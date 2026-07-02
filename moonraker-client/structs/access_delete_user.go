package structs

type AccessDeleteUserParams struct {
	Username string `json:"username"`
}

type AccessDeleteUser struct {
	Username string `json:"username"`
	Action   string `json:"action"`
}
