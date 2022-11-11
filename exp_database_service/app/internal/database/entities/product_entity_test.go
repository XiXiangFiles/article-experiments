package entities_test

import (
	factory "app/internal/database/collection_factory"
	"app/internal/database/entities"
	"context"
	"sync"

	"testing"
)

func TestCase(t *testing.T) {
	dbConn, _ := factory.NewMySQLHandler(context.Background(), "root:studyCircle@tcp(database:3306)", "cosmetic")
	collection := dbConn.GetCollection(factory.Example.ToString())
	var wg sync.WaitGroup
	dest := []*entities.ProductRaw{}
	for r := 0; r < 3; r++ {
		func() {
			defer wg.Done()
			wg.Add(1)
			for i := 0; i < 150000; i++ {
				collection.Query(context.TODO(), &entities.ProductFilter{
					Id: &i,
				}, &dest)
			}
		}()
	}
	wg.Wait()
}
