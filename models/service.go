package models

import "time"

//ServiceDefinition is the list of DevOps service already integration the Pilotage.
type ServiceDefinition struct {
	ID            int64      `json:"id" gorm:"primary_key"`
	Service       string     `json:"service" sql:"unique;not null;type:varchar(255)"` //Service's name for query
	Title         string     `json:"title" sql:"null;type:varchar(255)"`              //Service's name for display
	Gravatar      string     `json:"gravatar" sql:"null;type:text"`                   //Service's logo.
	Endpoints     string     `json:"endpoints" sql:"null;type:text"`                  //Service's endpoints include multi endpoints url with JSON format.
	Status        string     `json:"status" sql:"null;type:text"`                     //Service's service status endpoint.
	Environment   string     `json:"environment" sql:"null;type:text"`                //
	Authorization string     `json:"authorization" sql:"null;type:text"`              //
	Configuration string     `json:"configuration" sql:"null;type:text"`              //
	Description   string     `json:"description" sql:"null;type:text"`                //
	CreatedAt     time.Time  `json:"created" sql:""`                                  //
	UpdatedAt     time.Time  `json:"updated" sql:""`                                  //
	DeletedAt     *time.Time `json:"deleted" sql:"index"`                             //
}

//TableName is return the table name of Component in MySQL database.
func (sd *ServiceDefinition) TableName() string {
	return "serivce_definition"
}

//Service is third DevOps service outside the system, is must be the one of Service List.
type Service struct {
	ID            int64      `json:"id" gorm:"primary_key"`                                                            //
	Namespace     string     `json:"namespace" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_service"` //User name or organization name
	Service       string     `json:"service" sql:"not null;type:varchar(255)" gorm:"unique_index:namespace_service"`   //Service's name for query.
	Title         string     `json:"title" sql:"null;type:varchar(255)"`                                               //
	Gravatar      string     `json:"gravatar" sql:"null;type:text"`                                                    //
	Endpoints     string     `json:"endpoints" sql:"null;type:text"`                                                   //
	Environment   string     `json:"environment" sql:"null;type:text"`                                                 //
	Authorization string     `json:"authorization" sql:"null;type:text"`                                               //
	Configuration string     `json:"configuration" sql:"null;type:text"`                                               //
	Description   string     `json:"description" sql:"null;type:text"`                                                 //
	CreatedAt     time.Time  `json:"created" sql:""`                                                                   //
	UpdatedAt     time.Time  `json:"updated" sql:""`                                                                   //
	DeletedAt     *time.Time `json:"deleted" sql:"index"`                                                              //
}

//TableName is return the table name of Component in MySQL database.
func (s *Service) TableName() string {
	return "service"
}
