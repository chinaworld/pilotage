package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	//StageTypeStart is the Stage type being the start the pipeline.
	StageTypeStart = iota
	//StageTypeEnd is the Stage type being the end of the pipeline.
	StageTypeEnd
	//StageTypeRun is the Stage type being the running stage of pipeline.
	StageTypeRun
)

const (
	// PipelineStateStop is the state that pipeline stop work
	PipelineStateStop = iota
	// PipelineStateRunning is the state that pipeline is working
	PipelineStateRunning
)

//Pipeline is DevOps workflow definition unit.
type Pipeline struct {
	ID          int64      `json:"id" gorm:"primary_key"`                                                             //
	Namespace   string     `json:"namespace" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_pipeline"` //Username or organization
	Pipeline    string     `json:"pipeline" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_pipeline"`  //pipeline name
	Event       int64      `json:"event" sql:"null;default:0"`                                                        //
	Version     string     `json:"version" sql:"null;type:varchar(255)"`                                              //User define Pipeline version
	VersionCode int64      `json:"versionCode" sql:"null;type:varchar(255)"`                                          //System define Pipeline version,unique,for query
	State       int64      `json:"state" sql:"null;type:bigint"`                                                      //pipeline state
	Manifest    string     `json:"manifest" sql:"null;type:text"`                                                     //
	Description string     `json:"description" sql:"null;type:text"`                                                  //
	CreatedAt   time.Time  `json:"created" sql:""`                                                                    //
	UpdatedAt   time.Time  `json:"updated" sql:""`                                                                    //
	DeletedAt   *time.Time `json:"deleted" sql:"index"`                                                               //
}

//TableName is return the table name of Pipeline in MySQL database.
func (p *Pipeline) TableName() string {
	return "pipeline"
}

func (p *Pipeline) GetPipeline() *gorm.DB {
	return db.Model(&Component{})
}

//Stage is Pipeline unit.
type Stage struct {
	ID          int64      `json:"id" gorm:"primary_key"`                  //
	Pipeline    int64      `json:"pipeline" sql:"not null;default:0"`      //Pipeline's ID.
	Istag       bool       `json:"isTag" sql:"null;default:false"`         //if is true,means this stage can't be update,delete
	Type        int64      `json:"type" sql:"not null;default:0"`          //StageTypeStart, StageTypeEnd or StageTypeRun
	PreStage    int64      `json:"preStage" sql:"not null;default:0"`      //Pre stage ID ,first stage is -1
	Stage       string     `json:"stage" sql:"not null;type:varchar(255)"` //Stage name for query.
	Title       string     `json:"title" sql:"not null;type:varchar(255)"` //Stage title for display
	Description string     `json:"description" sql:"null;type:text"`       //
	Event       int64      `json:"event" sql:"null;default:0"`             //
	Manifest    string     `json:"manifest" sql:"null;type:text"`          //
	CreatedAt   time.Time  `json:"created" sql:""`                         //
	UpdatedAt   time.Time  `json:"updated" sql:""`                         //
	DeletedAt   *time.Time `json:"deleted" sql:"index"`                    //
}

//TableName is return the table name of Stage in MySQL database.
func (s *Stage) TableName() string {
	return "stage"
}

func (s *Stage) GetStage() *gorm.DB {
	return db.Model(&Stage{})
}

//Action is Stage unit.
type Action struct {
	ID          int64      `json:"id" gorm:"primary_key"`                  //
	Stage       int64      `json:"stage" sql:"not null;default:0"`         //
	Istag       bool       `json:"isTag" sql:"null;default:false"`         //if is true,means this action can't be update,delete
	Component   int64      `json:"component" sql:"not null;default:0"`     //
	Service     int64      `json:"service" sql:"not null;default:0"`       //
	Action      string     `json:"action" sql:"not null;varchar(255)"`     //
	Title       string     `json:"title" sql:"not null;type:varchar(255)"` //
	Description string     `json:"description" sql:"null;type:text"`       //
	Event       int64      `json:"event" sql:"null;default:0"`             //
	Manifest    string     `json:"manifest" sql:"null;type:text"`          //
	CreatedAt   time.Time  `json:"created" sql:""`                         //
	UpdatedAt   time.Time  `json:"updated" sql:""`                         //
	DeletedAt   *time.Time `json:"deleted" sql:"index"`                    //
}

//TableName is return the name of Action in MySQL database.
func (a *Action) TableName() string {
	return "action"
}

func (a *Action) GetAction() *gorm.DB {
	return db.Model(&Action{})
}

//Outcome is Pipeline running results.
//When StageID point to the StageTypeStart or StageTypeEnd, the Action ID is 0.
type Outcome struct {
	ID        int64      `json:"id" gorm:"primary_key"`             //
	Pipeline  int64      `json:"pipeline" sql:"not null;default:0"` //Pipeline id
	Stage     int64      `json:"stage" sql:"not null;default:0"`    //stage id
	Action    int64      `json:"action" sql:"not null;default:0"`   //action id
	Event     int64      `json:"event" sql:"null;default:0"`        //event id
	Sequence  int64      `json:"sequence" sql:"not null;default:0"` //pipeline run sequence
	Status    string     `json:"status" sql:"null;varchar(255)"`    //
	Result    bool       `json:"result" sql:"null;default:true"`    //
	Output    string     `json:"output" sql:"null;type:text"`       //
	CreatedAt time.Time  `json:"created" sql:""`                    //
	UpdatedAt time.Time  `json:"updated" sql:""`                    //
	DeletedAt *time.Time `json:"deleted" sql:"index"`               //
}

//TableName is return the name of Outcome in MySQL database.
func (o *Outcome) TableName() string {
	return "outcome"
}

func (o *Outcome) GetOutcome() *gorm.DB {
	return db.Model(&Outcome{})
}
