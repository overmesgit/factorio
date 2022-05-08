package mine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/nodemap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

var Sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	Sugar = logger.Sugar()

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

	if nodemap.Type(MyNode.Type) == nodemap.Manipulator {
		return nil, errors.New("i'm an manipulator dumb dumb")
	}

	err := MyStorage.Add(nodemap.ItemType(request.Type))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *server) NeededResource(ctx context.Context, request *pb.Empty) (*pb.Item, error) {
	nodemap.LogInput(ctx, "NeededResource", request)
	if nodemap.Type(MyNode.Type) == nodemap.Furnace {
		if MyStorage.GetCount(nodemap.Iron) > MyStorage.GetCount(nodemap.Coal) {
			return &pb.Item{Type: string(nodemap.Coal)}, nil
		} else {
			return &pb.Item{Type: string(nodemap.Iron)}, nil
		}
	}

	return nil, errors.New("nothing needed")
}

func (s *server) GiveResource(ctx context.Context, request *pb.Item) (*pb.Item, error) {
	nodemap.LogInput(ctx, "GiveResource", request)
	item := MyStorage.Get(nodemap.ItemType(request.GetType()))
	if item == nil {
		Sugar.Infof("Nothing to give %v.", MyStorage.GetItemCount())
		return nil, errors.New("nothing to give")
	}
	return item, nil
}

func RunServer() {
	port := "8080"
	Sugar.Infow(
		"Starting mine server",
		"port", port,
	)

	server := &server{}
	server.RunMapper()
	server.RunWorker()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		Sugar.Fatalw(
			"Failed to listen",
			"error", err,
		)
	}
	s := grpc.NewServer()
	pb.RegisterMineServer(s, server)
	Sugar.Infow(
		"server started",
		"port", port,
		"addr", lis.Addr(),
	)
	if err := s.Serve(lis); err != nil {
		Sugar.Fatalw(
			"Server failed",
			"error", err,
		)
	}
}
