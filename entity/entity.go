package entity

import "time"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
}

type BodyResponse struct {
	ID          int            `json:"id"`
	Category    int            `json:"category"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Footer      string         `json:"footer"`
	Tags        []string       `json:"tags"`
	Items       []BodyResponse `json:"items,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
}

type BodyResponseForQuestionThree struct {
	ID          int            `json:"id"`
	Category    int            `json:"category"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Footer      string         `json:"footer"`
	Items       []BodyResponse `json:"items,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
}
