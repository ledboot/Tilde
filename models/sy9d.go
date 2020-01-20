package models

import (
	"time"
)

type Sy9d struct {
	Hash  string    `json:"hash"`
	Title string    `json:"title"`
	Url   string    `json:"url"`
	Time  time.Time `json:"time"`
}
