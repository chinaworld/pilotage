package models

import "time"

//
type Service struct {
	Id        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created" sql:""`
	UpdatedAt time.Time  `json:"updated" sql:""`
	DeletedAt *time.Time `json:"deleted" sql:"index"`
}

//
func (*Service) TableName() string {
	return "service"
}
