package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"

	"google.golang.org/grpc"
	"log"
	"net"
)

type Type string

const (
	MINE Type = "MINE"
)

type Service struct {
	row, col    int32
	serviceType Type
}

type server struct {
	pb.UnimplementedMapperServer
}

var (
	Nodes  []*pb.Node
	UrlMap []*pb.URL
)

var dataStore *datastore.Client

type ItemType string

const (
	Iron ItemType = "iron"
)

type Item struct {
	ItemType ItemType
	Count    int
}

type Node struct {
	NodeType string
}

func (s *server) DoMineWork() {
}

func (s *server) UpdateMap(ctx context.Context, in *pb.MapRequest) (*pb.MapReply, error) {
	Nodes = in.GetNodes()
	log.Printf("Received6: %v", Nodes)

	UrlMap = in.GetUrlMap()
	log.Printf("Received6: %v", UrlMap)

	s.DoMineWork()

	return &pb.MapReply{}, nil
}

func main() {
	myIp, err := GetIp()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	port := "8080"
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMapperServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
