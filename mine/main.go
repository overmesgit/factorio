package mine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"github.com/overmesgit/factorio/nodemap"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials/insecure"
	"sync"

	"google.golang.org/grpc"
	"log"
	"net"
)

var MyNode *pb.Node
var MyItems struct {
	items map[localmap.ItemType]*pb.Item
	sync.Mutex
}

var AdjustedNodes []*pb.Node

type server struct {
	pb.UnimplementedMineServer
	logger *zap.SugaredLogger
}

func (s *server) SendResource(ctx context.Context, request *pb.ItemRequest) (*pb.ItemReply, error) {
	nodemap.LogInput(ctx, "SendResource", request, s.logger)

	MyItems.Lock()
	defer MyItems.Unlock()
	localStore, ok := MyItems.items[localmap.ItemType(request.Item.Type)]
	if !ok {
		MyItems.items[localmap.ItemType(request.Item.Type)] = request.Item
		return &pb.ItemReply{}, nil

	}

	if localStore.Count < 100 {
		localStore.Count += request.Item.Count
	} else {
		return nil, errors.New("don't have space")
	}
	return &pb.ItemReply{}, nil
}

func RunServer() {
	MyItems.items = map[localmap.ItemType]*pb.Item{}

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
