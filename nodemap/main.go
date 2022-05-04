package nodemap

import (
	"context"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedMapServer
	logger *zap.SugaredLogger
}

func (s *server) NotifyIp(ctx context.Context, in *pb.IpRequest) (*pb.IpReply, error) {
	p, ok := peer.FromContext(ctx)
	addr := "unknown"
	if ok {
		addr = p.Addr.String()
	}
	s.logger.Infow("Received message",
		"message", in,
		"ip", addr,
	)
	err := s.RegisterServer(in)
	if err != nil {
		return nil, err
	}
	resp := s.GetAdjustedNodes(in)
	s.logger.Infow("Send adjusted nodes",
		"nodes", resp,
	)
	return &pb.IpReply{AdjustedNodes: resp}, nil
}

func (s *server) UpdateMap(ctx context.Context, in *pb.MapRequest) (*pb.MapReply, error) {
	log.Printf("Received: %v", in)
	nodeMap := UpdateMap(in)

	return &pb.MapReply{Nodes: nodeMap}, nil
}

func RunServer() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	port := "8080"
	sugar := logger.Sugar()
	sugar.Infow("Starting map server",
		"port", port,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		sugar.Fatalw("Failed to listen",
			"error", err)
	}
	s := grpc.NewServer()
	pb.RegisterMapServer(s, &server{logger: sugar})
	sugar.Infow("server started",
		"port", port,
		"addr", lis.Addr(),
	)
	if err := s.Serve(lis); err != nil {
		sugar.Fatalw("Server failed",
			"error", err)
	}
}
