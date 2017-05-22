package main

import (
	"log"

	pb "github.com/reinventer/grpc-status-test/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial(`:10888`, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewTestErrorClient(conn)
	_, err = c.GetError(context.Background(), &pb.Empty{})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			log.Fatalf(`%#v`, s.Proto().Details[0])
		}
		log.Fatal(err)
	}
	log.Println(`Where is my error, dude?`)
}
