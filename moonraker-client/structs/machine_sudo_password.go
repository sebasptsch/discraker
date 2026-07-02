package structs

type MachineSudoPasswordParams struct {
	Password string `json:"password"`
}

type MachineSudoPassword struct {
	SudoResponses []string `json:"sudo_responses"`
	IsRestarting  bool     `json:"is_restarting"`
}
