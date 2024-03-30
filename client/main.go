package client

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/crud"
	"log"
	"time"
)

const defaultName = "asd"

var (
	addr = flag.String("addr", "localhost:8080", "address to connect to")
	name = flag.String("name", defaultName, "Name of user")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := crud.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateTask(ctx, &crud.Task{
		UserId:      defaultName,
		Title:       defaultName,
		Description: defaultName,
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("User created.")
}
