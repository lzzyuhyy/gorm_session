package session

import (
	"gorm.io/gorm"
)

func SessionAction(db *gorm.DB, actions func(db *gorm.DB) *gorm.DB) {
	db.Begin()

	res := actions(db)
	if res.Error != nil {
		db.Rollback()
		return
	}

	db.Commit()

}
