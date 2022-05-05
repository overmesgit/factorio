package mine

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
	"log"
	"net"
)

var MyType localmap.Type
var MyDir localmap.Direction
var MyItems = make(map[localmap.ItemType]*pb.Item)

var AdjustedNodes []*pb.Node

type server struct {
	pb.UnimplementedMineServer
	logger *zap.SugaredLogger
}

func (s *server) SendResource(ctx context.Context, request *pb.ItemRequest) (*pb.ItemReply, error) {
	panic("implement me")
}

func RunServer() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	port := "8080"
	sugar := logger.Sugar()
	sugar.Infow("Starting mine server",
		"port", port,
	)

	conn, err := grpc.Dial(localmap.MapServer+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()

	server := &server{logger: sugar}
	server.RunMapper(conn)
	server.RunWorker()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		sugar.Fatalw("Failed to listen",
			"error", err)
	}
	s := grpc.NewServer()
	pb.RegisterMineServer(s, server)
	sugar.Infow("server started",
		"port", port,
		"addr", lis.Addr(),
	)
	if err := s.Serve(lis); err != nil {
		sugar.Fatalw("Server failed",
			"error", err)
	}
}
