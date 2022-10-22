package entities

import (
	"gorm.io/gorm"
)

func WhereNotDeleted(db *gorm.DB) *gorm.DB {
	exp := "deleted_at = ?"
	return db.Where(exp, 0)
}

func StringToSortType(s string) sortType {
	switch s {
	case "DESC":
		return Desc
	default:
		return Asc
	}
}
