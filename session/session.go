package session

import (
	"gorm.io/gorm"
)

func SessionAction(db *gorm.DB, actions func() error) {
	db.Begin()

	err := actions()
	if err != nil {
		db.Rollback()
	}

	db.Commit()

}
