package application

import (
	"fmt"
	"message_api/config"
	"message_api/domain/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *MessageDB
)

func init() {
	if db != nil {
		return
	}

	var err error
	db, err = connectDatabase()
	if err != nil {
		fmt.Println(err)
	}
	db.migrate()
}

// MessageDB satella database
type MessageDB struct {
	DB *gorm.DB
}

func (db *MessageDB) close() error {
	return db.DB.Close()
}

func connectDatabase() (*MessageDB, error) {
	database, err := gorm.Open("mysql", config.SQLConnectionName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if database == nil {
		return nil, fmt.Errorf("not connect data base")
	}

	database = database.Set("gorm:auto_preload", true)
	database.LogMode(true)

	masterPartyDB := MessageDB{DB: database}

	return &masterPartyDB, nil
}

func (db *MessageDB) migrate() {
	db.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&entity.DirectMessage{},
		&entity.GroupMessage{},
		&entity.GroupUser{},
		&entity.Group{},
		&entity.User{},
	)
}
