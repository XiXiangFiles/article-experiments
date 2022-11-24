package test

import (
	"app/cmd"
	"app/internal/global"
	"app/pb"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const address = "0.0.0.0:50051"

func TestMain(m *testing.M) {
	_ = godotenv.Load()

	global.SetupEnv()
	go cmd.RunService()
	m.Run()
}

func GetConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

// 1 3 6 12 24 48 100 worker
func TestCase(t *testing.T) {
	var wg sync.WaitGroup
	conn := GetConnection()
	client := pb.NewExpServiceClient(conn)
	f, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	var sum float64
	var maxCount int = 1
	for count := 0; count < maxCount; count++ {
		fmt.Println("round := ", count+1)
		now := time.Now().Unix()
		for r := 0; r < 48; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 150000; i++ {
					client.QueryData(context.TODO(), &pb.QueryDataParameter{
						Id: int32(i),
					})
				}
			}()
		}
		wg.Wait()
		end := time.Now().Unix()
		sum += float64(end - now)
	}
	f.WriteString(fmt.Sprintf("%v,", sum/float64(maxCount)))
}

// 1 3 6 12 24 48 100 worker
func TestCase2(t *testing.T) {
	var wg sync.WaitGroup
	conn := GetConnection()
	conn2 := GetConnection()
	conn3 := GetConnection()
	client := pb.NewExpServiceClient(conn)
	client2 := pb.NewExpServiceClient(conn2)
	client3 := pb.NewExpServiceClient(conn3)
	f, _ := os.OpenFile("data2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	var sum float64
	var maxCount int = 1
	for count := 0; count < maxCount; count++ {
		fmt.Println("round := ", count+1)
		now := time.Now().Unix()
		for r := 0; r < 6; r++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 50000; i++ {
					client.QueryData(context.TODO(), &pb.QueryDataParameter{
						Id: int32(i),
					})
				}
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 50000; i < 100000; i++ {
					client2.QueryData(context.TODO(), &pb.QueryDataParameter{
						Id: int32(i),
					})
				}
			}()
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 100000; i < 150000; i++ {
					client3.QueryData(context.TODO(), &pb.QueryDataParameter{
						Id: int32(i),
					})
				}
			}()
		}
		wg.Wait()
		end := time.Now().Unix()
		sum += float64(end - now)
	}
	f.WriteString(fmt.Sprintf("%v,", sum/float64(maxCount)))
}
