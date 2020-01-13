package models

import "time"

type sy9d struct {
	Title string    `json:"title"`
	Url   string    `json:"url"`
	Time  time.Time `json:"time"`
}
