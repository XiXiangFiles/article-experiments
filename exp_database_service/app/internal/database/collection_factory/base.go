package collectionfactory

import "app/internal/database/entities"

type CollectionFactory interface {
	Initialize() error
	GetCollection(name string) entities.Collection
}

type entityName string

const (
	EnterpriseInfo entityName = "EnterpriseInfo"
)

func (n entityName) ToString() string {
	return string(n)
}
