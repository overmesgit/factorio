package nodemap

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
	"time"
)

var sugar *zap.SugaredLogger
var updateCh chan *pb.NodeState

func init() {
	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()
	updateCh = make(chan *pb.NodeState, 100)
}

type server struct {
	pb.UnimplementedMapServer
}

func (s *server) UpdateNodeState(ctx context.Context, in *pb.NodeState) (*pb.Empty, error) {
	LogInput(ctx, "updateNodeState", in)

	select {
	case updateCh <- in:
	default:
		sugar.Infof("skip message from node %v", in.GetNode())
	}

	return &pb.Empty{}, nil
}

func updateNodeMap() {
	for {
		in := <-updateCh
		node := in.GetNode()
		key := Key{node.Row, node.Col}
		mapItems.nodes[key] = in
		mapItems.lastUpdate[key] = time.Now().Unix()
	}
}

func LogInput(ctx context.Context, name string, in interface{}) {
	p, ok := peer.FromContext(ctx)
	addr := "unknown"
	if ok {
		addr = p.Addr.String()
	}
	sugar.Infof(
		"Received message %v ip %v %v",
		name, addr, in,
	)
}

func RunServer() {
	port := "8080"

	go RunHttpServer()
	go CleanItems()
	go updateNodeMap()
	sugar.Infow(
		"Starting map server",
		"port", port,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		sugar.Fatalw(
			"Failed to listen",
			"error", err,
		)
	}
	s := grpc.NewServer()
	pb.RegisterMapServer(s, &server{})
	sugar.Infow(
		"server started",
		"port", port,
		"addr", lis.Addr(),
	)
	if err := s.Serve(lis); err != nil {
		sugar.Fatalw(
			"Server failed",
			"error", err,
		)
	}
}
