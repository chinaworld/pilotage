package models

import "time"

//Component is customized container(docker or rkt) for executing DevOps tasks.
type Component struct {
	ID             int64      `json:"id" gorm:"primary_key"`
	Namespace      string     `json:"namespace" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_component"`
	Component      string     `json:"component" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_component"`
	Location       string     `json:"location" sql:"not null;type:text"`
	Tag            string     `json:"tag" sql:"null;type:varchar(255)"`
	Description    string     `json:"description" sql:"null;type:text"`
	VolumeLocation string     `json:"volume_location" sql:"null;type:text"`
	VolumeData     string     `json:"volume_data" sql:"null;type:text"`
	Resource       string     `json:"resource" sql:"null;type:varchar(255)"`
	Dockerfile     string     `json:"dockerfile" sql:"null;type:text"`
	Kubernetes     string     `json:"kubernetes" sql:"null;type:text"`
	Swarm          string     `json:"swarm" sql:"null;type:text"`
	CreatedAt      time.Time  `json:"created" sql:""`
	UpdatedAt      time.Time  `json:"updated" sql:""`
	DeletedAt      *time.Time `json:"deleted" sql:"index"`
}

//TableName is return the table name in MySQL database.
func (c *Component) TableName() string {
	return "component"
}
