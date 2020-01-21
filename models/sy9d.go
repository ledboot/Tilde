package models

import "time"

type Sy9d struct {
	Hash        string    `json:"hash"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	PublishTime string    `json:"publish_time"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
