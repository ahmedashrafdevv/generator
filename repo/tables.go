package repo

import (
	"github.com/jinzhu/gorm"
)

type TablesRepo struct {
	db *gorm.DB
}

func NewTablesRepo(db *gorm.DB) TablesRepo {
	return TablesRepo{
		db: db,
	}
}

func (ur *TablesRepo) ListReceiptsByPosted(posted *bool) error {

	return nil
}
