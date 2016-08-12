package models

import "time"

const (
	STAGE_TYPE_START = iota
	STAGE_TYPE_END
	STAGE_TYPE_RUN
)

//Pipeline is DevOps workflow definition unit.
type Pipeline struct {
	ID          int64      `json:"id" gorm:"primary_key"`
	Namespace   string     `json:"namespace" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_component"`
	Component   string     `json:"component" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_component"`
	Pipeline    string     `json:"pipeline" sql:"not null"`
	Manifests   string     `json:"manifests" sql:"null"`
	Description string     `json:"description" sql:"null"`
	CreatedAt   time.Time  `json:"created" sql:""`
	UpdatedAt   time.Time  `json:"updated" sql:""`
	DeletedAt   *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Pipeline in MySQL database.
func (p *Pipeline) TableName() string {
	return "pipeline"
}

//
type Stage struct {
	ID         int64      `json:"id" gorm:"primary_key"`
	PipelineID int64      `json:"pipeline_id" sql:"not null"`
	Type       int64      `json:"type" sql:"not null"`
	Sequence   int64      `json:"sequence" sql:"not null;default:0"`
	Stage      string     `json:"stage" sql:"not null;type:varchar(255)"`
	CreatedAt  time.Time  `json:"created" sql:""`
	UpdatedAt  time.Time  `json:"updated" sql:""`
	DeletedAt  *time.Time `json:"deleted" sql:"index"`
}

//
func (s *Stage) TableName() string {
	return "stage"
}

//
type Action struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	StageID   int64      `json:"stage_id" sql:"not null"`
	Action    string     `json:"action" sql:"not null;varchar(255)"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}

//
func (a *Action) TableName() string {
	return "action"
}

//
type Outcome struct {
	Id         int64      `json:"id" gorm:"primary_key"`
	PipelineID int64      `json:"pipeline_id" sql:"not null"`
	ActionID   int64      `json:"action_id" sql:"not null"`
	Sequence   int64      `json:"sequence" sql:"not null;default:1"`
	Status     string     `json:"status" sql:"null;varchar(255)"`
	Result     bool       `json:"result" sql:"null;default:true"`
	Output     string     `json:"output" sql:"null;type:text"`
	CreatedAt  time.Time  `json:"created" sql:""`
	UpdatedAt  time.Time  `json:"updated" sql:""`
	DeletedAt  *time.Time `json:"deleted" sql:"index"`
}

//
func (o *Outcome) TableName() string {
	return "outcome"
}
