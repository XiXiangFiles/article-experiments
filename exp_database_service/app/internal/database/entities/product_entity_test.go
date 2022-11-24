package entities_test

import (
	factory "app/internal/database/collection_factory"
	"app/internal/database/entities"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"testing"
)

// 1 3 6 12 24 48 100 worker
func TestCase(t *testing.T) {
	dbConn, _ := factory.NewMySQLHandler(context.Background(), "root:studyCircle@tcp(database:3306)", "cosmetic")
	collection := dbConn.GetCollection(factory.Example.ToString())
	var wg sync.WaitGroup
	dest := []*entities.ProductRaw{}
	f, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	var sum float64
	var maxCount int = 1
	for count := 0; count < maxCount; count++ {
		fmt.Println("round := ", count+1)
		now := time.Now().Unix()
		for r := 0; r < 100; r++ {
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
		end := time.Now().Unix()
		sum += float64(end - now)
	}
	f.WriteString(fmt.Sprintf("%v,", sum/float64(maxCount)))
}
