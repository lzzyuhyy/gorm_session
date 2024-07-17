package session

import (
	"gorm.io/gorm"
)

func SessionAction(db *gorm.DB, actions func(db *gorm.DB) *gorm.DB) {
	tx := db.Begin()

	err := actions(tx)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

}
