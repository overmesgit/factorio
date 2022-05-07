package nodemap

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
)

var sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	sugar = logger.Sugar()

}

type server struct {
	pb.UnimplementedMapServer
}

func (s *server) NotifyIp(ctx context.Context, in *pb.IpRequest) (*pb.IpReply, error) {
	LogInput(ctx, "NotifyIp", in, sugar)
	err := s.RegisterServer(in)
	if err != nil {
		return nil, err
	}
	resp := s.GetAdjustedNodes(in)
	sugar.Infow("Send adjusted nodes",
		"nodes", resp,
	)
	return &pb.IpReply{AdjustedNodes: resp}, nil
}

func LogInput(ctx context.Context, name string, in interface{}, logger *zap.SugaredLogger) {
	p, ok := peer.FromContext(ctx)
	addr := "unknown"
	if ok {
		addr = p.Addr.String()
	}
	logger.Infof("Received message %v ip %v %v",
		name, addr, in,
	)
}

func (s *server) UpdateMap(ctx context.Context, in *pb.MapRequest) (*pb.MapReply, error) {
	LogInput(ctx, "UpdateMap", in, sugar)
	nodeMap := s.RunUpdateMap(in)
	return &pb.MapReply{Nodes: nodeMap}, nil
}

func RunServer() {
	port := "8080"

	go RunHttpServer()
	sugar.Infow("Starting map server",
		"port", port,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		sugar.Fatalw("Failed to listen",
			"error", err)
	}
	s := grpc.NewServer()
	pb.RegisterMapServer(s, &server{})
	sugar.Infow("server started",
		"port", port,
		"addr", lis.Addr(),
	)
	if err := s.Serve(lis); err != nil {
		sugar.Fatalw("Server failed",
			"error", err)
	}
}
