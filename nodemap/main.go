package nodemap

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedMapServer
}

func (s *server) NotifyIp(ctx context.Context, in *pb.IpRequest) (*pb.IpReply, error) {
	log.Printf("Received: %v", in)
	err := RegisterServer(in)
	if err != nil {
		return nil, err
	}

	return &pb.IpReply{}, nil
}

func (s *server) UpdateMap(ctx context.Context, in *pb.MapRequest) (*pb.MapReply, error) {
	log.Printf("Received: %v", in)
	nodeMap := UpdateMap(in)

	return &pb.MapReply{Nodes: nodeMap}, nil
}

func RunServer() {
	port := "8080"
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMapServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
