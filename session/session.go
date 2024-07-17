package session

import (
	"gorm.io/gorm"
)

func SessionAction(db *gorm.DB, actions func(db *gorm.DB) error) {
	db.Begin()

	err := actions(db)
	if err != nil {
		db.Rollback()
	}

	db.Commit()

}
