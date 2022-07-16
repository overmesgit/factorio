package mine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers"
	"github.com/overmesgit/factorio/nodemap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

func init() {
	logger, _ := zap.NewProduction()
	sugar.Sugar = logger.Sugar()

	col, err := strconv.Atoi(os.Getenv("COL"))
	if err != nil {
		panic(err)
	}
	row, err := strconv.Atoi(os.Getenv("ROW"))
	if err != nil {
		panic(err)
	}

	nodeType := workers.Type(os.Getenv("TYPE"))
	node := workers.NewNode(
		int32(col),
		int32(row),
		nodeType,
		workers.Direction(os.Getenv("DIRECTION")),
	)

	var workerNode workers.WorkerNode
	switch nodeType {
	case workers.IronMine:
		workerNode = workers.NewMine(node.GetNextNode(), workers.Iron)
	case workers.CoalMine:
		workerNode = workers.NewMine(node.GetNextNode(), workers.Coal)
	case workers.Furnace:
		workerNode = workers.NewFurnaceNode(node.GetNextNode())
	case workers.Manipulator:
		workerNode = workers.NewManipulator(node.GetNextNode(), node.GetPrevNode())
	}
	MyWorker = workerNode

}

var MyWorker workers.WorkerNode
var MyNode workers.Node

type server struct {
	pb.UnimplementedMineServer
}

func (s *server) ReceiveResource(ctx context.Context, request *pb.Item) (*pb.Empty, error) {
	nodemap.LogInput(ctx, "SendResource", request)

	err := MyWorker.ReceiveResource(workers.ItemType(request.Type))
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *server) NeededResource(ctx context.Context, request *pb.Empty) (*pb.Item, error) {
	nodemap.LogInput(ctx, "NeededResource", request)
	item, err := MyWorker.GetNeededResource()
	if err != nil {
		return nil, err
	}

	return &pb.Item{Type: string(item)}, errors.New("nothing needed")
}

func (s *server) GetResource(ctx context.Context, request *pb.Item) (*pb.Item, error) {
	nodemap.LogInput(ctx, "GiveResource", request)

	item, err := MyWorker.GetResourceForSend()
	if err != nil {
		return nil, err
	}

	return &pb.Item{Type: string(item)}, nil
}

func RunServer() {
	port := "8080"
	sugar.Sugar.Infow(
		"Starting mine server",
		"port", port,
	)

	server := &server{}
	MyWorker.StartWorker()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		sugar.Sugar.Fatalw(
			"Failed to listen",
			"error", err,
		)
	}
	s := grpc.NewServer()
	pb.RegisterMineServer(s, server)
	sugar.Sugar.Infow(
		"server started",
		"port", port,
		"addr", lis.Addr(),
	)
	if err := s.Serve(lis); err != nil {
		sugar.Sugar.Fatalw(
			"Server failed",
			"error", err,
		)
	}
}
