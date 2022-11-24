package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"exp/pb"

	"sync"
	"time"

	"google.golang.org/grpc"
)

func GetConnection(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

func main() {
	var wg sync.WaitGroup
	conn := GetConnection("db-service:50051")
	client := pb.NewExpServiceClient(conn)
	conn2 := GetConnection("db-service2:50051")
	client2 := pb.NewExpServiceClient(conn2)
	conn3 := GetConnection("db-service3:50051")
	client3 := pb.NewExpServiceClient(conn3)
	f, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	var sum float64
	var maxCount int = 5
	for count := 0; count < maxCount; count++ {
		fmt.Println("round := ", count+1)
		now := time.Now().Unix()
		for r := 0; r < 1; r++ {
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
