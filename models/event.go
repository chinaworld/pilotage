package models

import (
	"time"
)

const (
	SYSTEM_EVENT = iota
	USER_EVENT
)

//Event is
type Event struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	ComponentID int64      `json:"component_id" sql:"not null" gorm:"unique_index:component_event"`
	Event       string     `json:"event" sql:"not null;type:varchar(255)" gorm:"unique_index:component_event"`
	Type        int64      `json:"type" sql:"not null"`
	Definition  string     `json:"type" sql:"null;type:text"`
	CreatedAt   time.Time  `json:"created" sql:""`
	UpdatedAt   time.Time  `json:"updated" sql:""`
	DeletedAt   *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Event in MySQL database.
func (e *Event) TableName() string {
	return "event"
}
