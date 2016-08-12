package models

import (
	"time"
)

const (
	//TypeSystemEvent is Event type, is definition and maintaince by the system.
	TypeSystemEvent = iota
	//TypeUserEvent is Event Type. TypeUserEvent is definition by the user, and maintaince by the system.
	TypeUserEvent
	//SourceInnerEvent is Event source type. It's cerate in the system.
	SourceInnerEvent
	//SourceOutsideEvent is Event source type. It's create from outside, like POST/PUT to a REST API interface.
	SourceOutsideEvent
)

//EventDefinition is the event type, source and definition in the customized DevOps workflow processing.
//And the system events will be initlization with pilotage system.
type EventDefinition struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	ComponentID int64      `json:"component_id" sql:"not null" gorm:"unique_index:component_event"`
	Event       string     `json:"event" sql:"not null;type:varchar(255)" gorm:"unique_index:component_event"`
	Title       string     `json:"title" sql:"null:type:varchar(255)"`
	Type        int64      `json:"type" sql:"not null"`
	Source      int64      `json:"source" sql:"not null"`
	Definition  string     `json:"type" sql:"null;type:text"`
	CreatedAt   time.Time  `json:"created" sql:""`
	UpdatedAt   time.Time  `json:"updated" sql:""`
	DeletedAt   *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Event in MySQL database.
func (e *EventDefinition) TableName() string {
	return "event_definition"
}

//Event is execute events in the system.
type Event struct {
	ID             int64      `json:"id" gorm:"primary_key"`
	Source         string     `json:"source" sql:"not null;type:text"`
	Header         string     `json:"header" sql:"not null;type:text"`
	Payload        string     `json:"payload" sql:"not null;type:text"`
	Authorizations string     `json:"authorization" sql:"null;type:text"`
	CreatedAt      time.Time  `json:"created" sql:""`
	UpdatedAt      time.Time  `json:"updated" sql:""`
	DeletedAt      *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Event in MySQL database.
func (e *Event) TableName() string {
	return "event"
}
