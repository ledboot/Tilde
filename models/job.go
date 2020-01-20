package models

import "time"

type Job struct {
	ID        int        `gorm:"primary_key;auto_increment;size:11" json:"id"`
	Title     string     `gorm:"type:varchar(255)" json:"title"`
	Desc      string     `gorm:"type:varchar(500)" json:"desc"`
	Status    uint       `gorm:"default:0" json:"status"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
