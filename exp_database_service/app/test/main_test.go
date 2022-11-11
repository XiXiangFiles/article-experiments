package test

import (
	"app/cmd"
	"app/internal/global"
	"app/pb"
	"context"
	"log"
	"sync"
	"testing"

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

func TestCase(t *testing.T) {
	var wg sync.WaitGroup
	conn := GetConnection()
	client := pb.NewExpServiceClient(conn)

	for r := 0; r < 3; r++ {
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
}
