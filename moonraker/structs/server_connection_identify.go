package structs

type ServerConnectionIdentifyParams struct {
	ClientName  string `json:"client_name"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	AccessToken string `json:"access_token,omitempty"`
	APIKey      string `json:"api_key,omitempty"`
}

type ServerConnectionIdentify struct {
	ConnectionID int `json:"connection_id"`
}
