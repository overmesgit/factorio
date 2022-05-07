package mine

import (
	"context"
	pb "github.com/overmesgit/factorio/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conn, err := grpc.Dial("map:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		sugar.Errorw("failed to connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewMapClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.UpdateNodeState(ctx, &pb.NodeState{
		Node:  MyNode,
		Items: MyStorage.GetItemCount(),
	})
	if err != nil {
		sugar.Errorf("Could not update my status: %v\n", err)
		return
	}
	sugar.Infof("Response: %v\n", r)
}
