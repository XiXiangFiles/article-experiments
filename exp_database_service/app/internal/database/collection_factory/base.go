package collectionfactory

import "app/internal/database/entities"

type CollectionFactory interface {
	Initialize() error
	GetCollection(name string) entities.Collection
}

type entityName string

const (
	Example entityName = "Example"
)

func (n entityName) ToString() string {
	return string(n)
}
