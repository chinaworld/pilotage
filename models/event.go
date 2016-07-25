package models

import (
	"time"
)

//
type EventType struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}

//
func (*EventType) TableName() string {
	return "event_type"
}

//
type Event struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}

//
func (*Event) TableName() string {
	return "event"
}
