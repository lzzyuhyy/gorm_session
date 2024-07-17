package session

import (
	"gorm.io/gorm"
)

func SessionAction(db *gorm.DB, actions func(db *gorm.DB) *gorm.DB) {
	tx := db.Begin()

	res := actions(tx)
	if res.Error != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

}
