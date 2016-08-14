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
	//CharacterServiceEvent is
	CharacterServiceEvent
	//CharacterComponentEvent is
	CharacterComponentEvent
)

//EventDefinition is the event type, source and definition in the customized DevOps workflow processing.
//And the system events will be initlization with pilotage system.
type EventDefinition struct {
	ID         int64      `json:"id" gorm:"primary_key"`
	Character  int64      `json:"character" sql:"not null;default:0"`
	Type       int64      `json:"type" sql:"not null;default:0"`
	Source     int64      `json:"source" sql:"not null;default:0"`
	Event      string     `json:"event" sql:"not null;type:varchar(255)"`
	Title      string     `json:"title" sql:"null:type:varchar(255)"`
	Definition string     `json:"type" sql:"null;type:text"`
	CreatedAt  time.Time  `json:"created" sql:""`
	UpdatedAt  time.Time  `json:"updated" sql:""`
	DeletedAt  *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Event in MySQL database.
func (e *EventDefinition) TableName() string {
	return "event_definition"
}

//Event is execute events in the system.
type Event struct {
	ID             int64      `json:"id" gorm:"primary_key"`
	DefinitionID   int64      `json:"definition_id" sql:"not null;default:0"`
	Header         string     `json:"header" sql:"not null;type:text"`
	Payload        string     `json:"payload" sql:"not null;type:text"`
	Authorizations string     `json:"authorization" sql:"null;type:text"`
	Type           int64      `json:"type" sql:"not null;default:0"`
	Source         int64      `json:"source" sql:"not null;default:0"`
	PipelineID     int64      `json:"pipeline_id" sql:"not null;default:0"`
	StageID        int64      `json:"stage_id" sql:"not null;default:0"`
	ActionID       int64      `json:"action_id" sql:"not null;default:0"`
	Sequence       int64      `json:"sequence" sql:"not null;default:0"`
	CreatedAt      time.Time  `json:"created" sql:""`
	UpdatedAt      time.Time  `json:"updated" sql:""`
	DeletedAt      *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Event in MySQL database.
func (e *Event) TableName() string {
	return "event"
}

//Environment is Pipeline environments.
type Environment struct {
	ID         int64      `json:"id" gorm:"primary_key"`
	PipelineID int64      `json:"pipeline_id" sql:"not null;default:0"`
	Sequence   int64      `json:"sequence" sql:"not null;default:0"`
	Key        string     `json:"key" sql:"not null;type:varchar(255)"`
	Value      string     `json:"value" sql:"not null;type:text"`
	CreatedAt  time.Time  `json:"created" sql:""`
	UpdatedAt  time.Time  `json:"updated" sql:""`
	DeletedAt  *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the name of Outcome in MySQL database.
func (e *Environment) TableName() string {
	return "environment"
}
