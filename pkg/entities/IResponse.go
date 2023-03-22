package entities

import "time"

type IResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
