package main

import (
	"log"
	"net"

	anypb "github.com/golang/protobuf/ptypes/any"
	pb "github.com/reinventer/grpc-status-test/status"
	"golang.org/x/net/context"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) GetError(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	ss := &spb.Status{
		Code:    int32(codes.InvalidArgument),
		Message: `st message`,
		Details: []*anypb.Any{{
			TypeUrl: `type`,
			Value:   []byte(`test`),
		}}}
	return &pb.Empty{}, status.FromProto(ss).Err()
}

func main() {
	lis, err := net.Listen(`tcp`, `:10888`)
	if err != nil {
		log.Fatalf(`failed to listen: %v`, err)
	}
	s := grpc.NewServer()
	pb.RegisterTestErrorServer(s, &server{})
	log.Println(`Server started`)
	if err := s.Serve(lis); err != nil {
		log.Fatalf(`failed to serve: %v`, err)
	}
}
