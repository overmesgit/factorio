package mine

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"log"
	"net"
)

var MyType localmap.Type
var AdjustedNodes []*pb.Node

type server struct {
	pb.UnimplementedMineServer
}

func (s *server) UpdateMap(ctx context.Context, in *pb.MapRequest) (*pb.MapReply, error) {

	return &pb.MapReply{}, nil
}

func RunServer() {
	conn, err := grpc.Dial(localmap.MapServer+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	RunMapper(conn)
	RunWorker()

	port := "8080"
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMineServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
