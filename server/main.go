package server

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc/crud"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	crud.UnimplementedCRUDServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", *port))
	if err != nil {
		log.Fatalf("Couldn't listen on: %v", err)
	}
	s := grpc.NewServer()
	crud.RegisterCRUDServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
