package prot

type ClientAddRequest struct {
	ClientID int64  `json:"ClientID"`
	Name     string `json:"Name"`
}

type ClientAddResponse struct {
}
