package mine

import (
	"context"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine/sugar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func (s *server) RunMapper() {
	go func() {
		for {
			time.Sleep(time.Second)
			s.UpdateMapState()
		}
	}()
}

func (s *server) UpdateMapState() {
	url := "map:8080"
	if os.Getenv("local") != "" {
		url = "host.minikube.internal:8080"
	}
	sugar.Sugar.Infof("Map url %v", url)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sugar.Sugar.Errorw("failed to connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewMapClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	counter := MyWorker.GetItemCount()
	grpcCounter := make([]*pb.ItemCounter, 0, len(counter))
	for _, c := range counter {
		grpcCounter = append(
			grpcCounter, &pb.ItemCounter{
				Type:  c.Type,
				Count: c.Count,
			},
		)
	}
	r, err := c.UpdateNodeState(
		ctx, &pb.NodeState{
			Node: &pb.Node{
				Type:      string(MyNode.NodeType),
				Col:       MyNode.Col,
				Row:       MyNode.Row,
				Direction: string(MyNode.Direction),
			},
			Items: grpcCounter,
		},
	)
	if err != nil {
		sugar.Sugar.Errorf("Could not update my status: %v\n", err)
		return
	}
	sugar.Sugar.Infof("Response: %v\n", r)
}
