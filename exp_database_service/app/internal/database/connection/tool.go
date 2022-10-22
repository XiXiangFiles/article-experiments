package connectionhandler

import (
	"strings"

	"gorm.io/gorm/schema"
)

type IndexIXNamingNamingStrategy struct {
	schema.NamingStrategy
}

func (ns IndexIXNamingNamingStrategy) IndexName(table, column string) string {
	var name = ns.NamingStrategy.IndexName(table, column)
	if strings.HasPrefix(name, "idx") {
		name = strings.Replace(name, "idx", "ix", 1)
	}
	return name
}
