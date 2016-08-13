package models

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"

	"github.com/containerops/pilotage/setting"
)

var (
	db *gorm.DB
)

//
func init() {
	var err error

	if db, err = gorm.Open(setting.DatabaseDriver, setting.DatabaseURI); err != nil {
		log.Fatal("Initlization database connection error.")
		os.Exit(1)
	} else {
		db.DB()
		db.DB().Ping()
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.SingularTable(true)
	}
}

//Sync is
func Sync() error {
	log.Info("Sync database structs")

	db.AutoMigrate(&ServiceList{}, &Service{}, &Component{})
	db.AutoMigrate(&Pipeline{}, &Stage{}, &Action{}, &Outcome{})
	db.AutoMigrate(&EventDefinition{}, &Event{}, &Environment{})

	return nil
}
