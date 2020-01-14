package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Sy9d struct {
	Id    bson.ObjectId `bson:_id;json:"id"`
	Title string        `json:"title"`
	Url   string        `json:"url"`
	Time  time.Time     `json:"time"`
}
