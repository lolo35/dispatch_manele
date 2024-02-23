package models

type Pitch struct {
	ID            int    `json:"id"`
	PitchStart    string `json:"pitch_start"`
	PitchEnd      string `json:"pitch_end"`
	ActualProduct int    `json:"actual_product"`
}

type PitchResponse struct {
	Success bool    `json:"success"`
	Offset  int     `json:"offset"`
	Limit   int     `json:"limit"`
	Error   string  `json:"error,omitempty"`
	Pitches []Pitch `json:"data"`
}

type UpdatePitchResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
