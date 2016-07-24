package models

import "time"

//
type Pipeline struct {
	Id          int64      `json:"id" gorm:"primary_key"`
	Name        string     `json:"name" sql:"not null"`
	Manifests   string     `json:"manifests" sql:"null"`
	Description string     `json:"description" sql:"null"`
	CreatedAt   time.Time  `json:"created" sql:""`
	UpdatedAt   time.Time  `json:"updated" sql:""`
	DeletedAt   *time.Time `json:"deleted" sql:"index"`
}

//
func (*Pipeline) TableName() string {
	return "pipeline"
}

//
type Stage struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}

//
func (*Stage) TableName() string {
	return "stage"
}

//
type Movement struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}

//
func (*Movement) TableName() string {
	return "movement"
}

//
type Outcome struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	Pipeline  int64      `json:"pipeline_id" sql:"not null"`
	Movement  int64      `json:"movement_id" sql:"not null"`
	Status    string     `json:"status" sql:"null;varchar(255)"`
	Result    bool       `json:"result" sql:"null;default:true"`
	Output    string     `json:"output" sql:"null;type:text"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}
