package structs

type ServerConnectionIdentifyParams struct {
	ClientName  string `json:"client_name"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	AccessToken string `json:"access_token"`
	APIKey      string `json:"api_key"`
}

type ServerConnectionIdentify struct {
	ConnectionID int `json:"connection_id"`
}
