package models

import "time"

const (
	//StageTypeStart is the Stage type being the start the pipeline.
	StageTypeStart = iota
	//StageTypeEnd is the Stage type being the end of the pipeline.
	StageTypeEnd
	//StageTypeRun is the Stage type being the running stage of pipeline.
	StageTypeRun
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

//Stage is Pipeline unit.
type Stage struct {
	ID             int64      `json:"id" gorm:"primary_key"`
	PipelineID     int64      `json:"pipeline_id" sql:"not null;default:0"`
	Type           int64      `json:"type" sql:"not null;default:0"`
	Sequence       int64      `json:"sequence" sql:"not null;default:0"`
	Stage          string     `json:"stage" sql:"not null;type:varchar(255)"`
	Title          string     `json:"title" sql:"null;type:varchar(255)"`
	Description    string     `json:"description" sql:"null;type:text"`
	Configurations string     `json:"configurations" sql:"null;type:text"`
	CreatedAt      time.Time  `json:"created" sql:""`
	UpdatedAt      time.Time  `json:"updated" sql:""`
	DeletedAt      *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name of Stage in MySQL database.
func (s *Stage) TableName() string {
	return "stage"
}

//Action is Stage unit.
type Action struct {
	ID             int64      `json:"id" gorm:"primary_key"`
	StageID        int64      `json:"stage_id" sql:"not null;default:0"`
	ComponentID    int64      `json:"component_id" sql:"not null;default:0"`
	ServiceID      int64      `json:"service_id" sql:"not null;default:0"`
	Action         string     `json:"action" sql:"not null;varchar(255)"`
	Title          string     `json:"title" sql:"null;type:varchar(255)"`
	Description    string     `json:"description" sql:"null;type:text"`
	Configurations string     `json:"configurations" sql:"null;type:text"`
	CreatedAt      time.Time  `json:"created" sql:""`
	UpdatedAt      time.Time  `json:"updated" sql:""`
	DeletedAt      *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the name of Action in MySQL database.
func (a *Action) TableName() string {
	return "action"
}

//Outcome is Pipeline running results.
//When StageID point to the StageTypeStart or StageTypeEnd, the Action ID is 0.
type Outcome struct {
	ID         int64      `json:"id" gorm:"primary_key"`
	PipelineID int64      `json:"pipeline_id" sql:"not null;default:0"`
	StageID    int64      `json:"stage_id" sql:"not null;default:0"`
	ActionID   int64      `json:"action_id" sql:"not null;default:0"`
	Sequence   int64      `json:"sequence" sql:"not null;default:0"`
	Status     string     `json:"status" sql:"null;varchar(255)"`
	Result     bool       `json:"result" sql:"null;default:true"`
	Output     string     `json:"output" sql:"null;type:text"`
	CreatedAt  time.Time  `json:"created" sql:""`
	UpdatedAt  time.Time  `json:"updated" sql:""`
	DeletedAt  *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the name of Outcome in MySQL database.
func (o *Outcome) TableName() string {
	return "outcome"
}
