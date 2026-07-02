package structs

type AccessGetUser struct {
	Username  string  `json:"username"`
	Source    string  `json:"source"`
	CreatedOn float64 `json:"created_on"`
}
