package mine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/nodemap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func (s *server) SendItems() {
	for {
		time.Sleep(200 * time.Millisecond)
		adjNode := s.getNextNode()
		if adjNode != nil {
			err := s.sendItem(adjNode)
			if err != nil {
				sugar.Errorf("err: %v", err)
			}
		}
	}
}

var directionIndex = map[string][]int32{
	//  ROW / COL
	"A": {-1, 0},
	"V": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

func (s *server) getNextNode() *pb.Node {
	offset, ok := directionIndex[MyNode.Direction]
	if !ok {
		return nil
	}
	adjRow, adjCol := MyNode.Row+offset[0], MyNode.Col+offset[1]
	return &pb.Node{Row: adjRow, Col: adjCol}
}

func (s *server) sendItem(adjNode *pb.Node) error {
	sugar.Infof("Send items. Current store. %v", MyStorage.GetItemCount())

	conn, err := grpc.Dial(fmt.Sprintf("r%vc%v:8080", adjNode.Row, adjNode.Col), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	forSend := MyStorage.GetItemForSend()
	if forSend == nil {
		sugar.Infof("Nothing to send.")
		return nil
	}

	r, err := c.SendResource(ctx, forSend)
	if err != nil {
		err := MyStorage.Add(nodemap.ItemType(forSend.Type))
		if err != nil {
			sugar.Warnf("Could not stack item back %v %v", forSend, err)
		}
		return err
	}

	sugar.Infof("Sent item %v. Resp %v.", forSend, r)
	return nil
}
