package client_exec

type ClientGetRequest struct {
}

type ClientGetResponse struct {
	ClientID int64  `json:"ClientID"`
	Name     string `json:"Name"`
}
