package mine

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/nodemap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

var sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()

	col, err := strconv.Atoi(os.Getenv("COL"))
	if err != nil {
		panic(err)
	}
	row, err := strconv.Atoi(os.Getenv("ROW"))
	if err != nil {
		panic(err)
	}
	MyNode = &pb.Node{
		Type:      os.Getenv("TYPE"),
		Col:       int32(col),
		Row:       int32(row),
		Direction: os.Getenv("DIRECTION"),
	}
}

var MyNode *pb.Node

type server struct {
	pb.UnimplementedMineServer
}

func (s *server) SendResource(ctx context.Context, request *pb.Item) (*pb.Empty, error) {
	nodemap.LogInput(ctx, "SendResource", request)

	//localStore, ok := MyItems.items[nodemap.ItemType(request.Item.Type)]
	//if !ok {
	//	MyItems.items[nodemap.ItemType(request.Item.Type)] = request.Item
	//	return &pb.ItemReply{}, nil
	//
	//}
	//
	//if localStore.Count < 100 {
	//	localStore.Count += request.Item.Count
	//} else {
	//	return nil, errors.New("don't have space")
	//}
	return &pb.Empty{}, nil
}

func (s *server) GiveResource(ctx context.Context, request *pb.Item) (*pb.Item, error) {
	nodemap.LogInput(ctx, "GiveResource", request)
	return nil, nil
}

func RunServer() {
	port := "8080"
	sugar.Infow("Starting mine server",
		"port", port,
	)

	server := &server{}
	server.RunMapper()
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
