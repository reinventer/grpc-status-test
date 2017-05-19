package main

import (
	"errors"
	"log"
	"net"

	anypb "github.com/golang/protobuf/ptypes/any"
	spb "github.com/google/go-genproto/googleapis/rpc/status"
	pb "github.com/reinventer/grpc-status-test/status"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetError(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, &status{}
}

type status struct{}

func (s *status) Error() string {
	return `error`
}

func (s *status) Code() codes.Code {
	return codes.InvalidArgument
}

func (s *status) Proto() *spb.Status {
	return &spb.Status{
		Code:    int32(codes.InvalidArgument),
		Message: `status message`,
		Details: []*anypb.Any{{
			TypeUrl: `type`,
			Value:   []byte(`test`),
		}},
	}
}

func (s *status) Message() string {
	return `message`
}

func (s *status) Err() error {
	return errors.New(`err`)
}

func main() {
	lis, err := net.Listen(`tcp`, `:10888`)
	if err != nil {
		log.Fatalf(`failed to listen: %v`, err)
	}
	s := grpc.NewServer()
	pb.RegisterStatusServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf(`failed to serve: %v`, err)
	}
}
