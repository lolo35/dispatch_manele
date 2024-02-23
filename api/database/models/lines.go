package models

type LineRequest struct {
	Success bool         `json:"success"`
	Offset  int          `json:"offset"`
	Limit   int          `json:"limit"`
	Line    []LineStruct `json:"data"`
}

type LineStruct struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}
