package models

type Part struct {
	Code string `json:"code"`
}

type PartResponse struct {
	Success bool   `json:"success"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Error   string `json:"error,omitempty"`
	Parts   Part   `json:"data"`
}
