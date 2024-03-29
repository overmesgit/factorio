package mine

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine/grpcsender"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"github.com/overmesgit/factorio/nodemap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func init() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := loggerConfig.Build()
	sugar.Sugar = logger.Sugar()

	col, err := strconv.Atoi(os.Getenv("COL"))
	if err != nil {
		panic(err)
	}
	row, err := strconv.Atoi(os.Getenv("ROW"))
	if err != nil {
		panic(err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	nodeType := basic.Type(os.Getenv("TYPE"))
	nodeProduction := basic.Type(os.Getenv("PRODUCTION"))
	MyNode = basic.NewNode(
		int32(row),
		int32(col),
		nodeType,
		basic.Direction(os.Getenv("DIRECTION")),
		hostname,
	)

	sender := grpcsender.NewSender()

	var workerNode basic.WorkerNode
	nextNode := MyNode.GetNextNode()
	prevNode := MyNode.GetPrevNode()
	switch nodeType {
	case basic.Mine:
		// TODO: move get next MyNode into constructor
		workerNode = workers.NewMine(MyNode, basic.ItemType(nodeProduction), sender)
	case basic.Furnace:
		workerNode = workers.NewFurnaceNode(MyNode, sender)
	case basic.Manipulator:
		workerNode = workers.NewManipulator(nextNode, prevNode, sender)
	case basic.Belt:
		workerNode = workers.NewBelt(nextNode, sender)
	case basic.AssemblingMachine:
		workerNode = workers.NewAssemblingMachine(
			MyNode, basic.ItemType(nodeProduction), sender,
		)
	}
	MyWorker = workerNode

}

var MyWorker basic.WorkerNode
var MyNode basic.Node

type server struct {
	pb.UnimplementedMineServer
}

func (s *server) ReceiveResource(ctx context.Context, request *pb.Item) (*pb.Empty, error) {
	nodemap.LogInput(ctx, "SendResource", request)

	err := MyWorker.ReceiveResource(grpcsender.GrpcItemToItem(request))
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

	return &pb.Item{Type: string(item)}, nil
}

func (s *server) GetResource(ctx context.Context, request *pb.Item) (*pb.Item, error) {
	nodemap.LogInput(ctx, "GiveResource", request)

	item, err := MyWorker.GetResourceForSend(basic.ItemType(request.Type))
	if err != nil {
		return nil, err
	}

	grpcItem := grpcsender.ItemToGrpcItem(item)
	return &grpcItem, nil
}

func RunServer() {
	port := "8080"
	sugar.Sugar.Infow(
		"Starting mine server",
		"port", port,
	)

	server := &server{}
	RunMapper()
	MyWorker.StartWorker()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		sugar.Sugar.Fatalw(
			"Failed to listen",
			"error", err,
		)
	}
	s := grpc.NewServer()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		s.GracefulStop()
	}()

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
