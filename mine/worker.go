package mine

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func (s *server) RunWorker() {
	go s.DoWork()
	go s.SendItems()
}

func (s *server) DoWork() {
	for {
		time.Sleep(time.Second)
		if MyNode == nil {
			s.logger.Infof("Waiting for my node %v\n", MyNode)
			continue
		}

		switch localmap.Type(MyNode.Type) {
		case localmap.IronMine:
			s.ironMine()
		default:
			s.logger.Warnf("Waiting for my node %v\n", MyNode)

		}
	}
}

func (s *server) ironMine() {
	mineType := localmap.Iron

	MyItems.Lock()
	defer MyItems.Unlock()

	localStore, ok := MyItems.items[mineType]
	if !ok {
		localStore = &pb.Item{
			Type:  string(mineType),
			Count: 0,
		}
		MyItems.items[mineType] = localStore
	}

	s.logger.Infof("LocalStore %v\n", localStore)
	if localStore.Count < 100 {
		localStore.Count++
	}
}

func (s *server) SendItems() {
	for {
		time.Sleep(time.Second)
		adjNode := s.getAdjNode()
		if adjNode != nil {
			err := s.sendItem(adjNode)
			if err != nil {
				s.logger.Errorf("err: %v", err)
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

func (s *server) getAdjNode() *pb.Node {
	if len(AdjustedNodes) == 0 {
		return nil
	}
	offset, ok := directionIndex[MyNode.Direction]
	if !ok {
		return nil
	}
	adjRow, adjCol := MyNode.Row+offset[0], MyNode.Col+offset[1]
	for _, node := range AdjustedNodes {
		if node.Row == adjRow && node.Col == adjCol {
			return node
		}
	}
	return nil
}

func (s *server) sendItem(adjNode *pb.Node) error {
	if adjNode.Ip == "" {
		return errors.New(fmt.Sprintf("Adj node does not have ip %v", adjNode))
	}

	MyItems.Lock()
	defer MyItems.Unlock()
	var forSend *pb.Item
	for _, i := range MyItems.items {
		if i.Count > 0 {
			forSend = i
			break
		}
	}
	if forSend == nil {
		s.logger.Infof("Nothing to send.")
		return nil
	}

	conn, err := grpc.Dial(adjNode.Ip+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendResource(ctx, &pb.ItemRequest{
		Item: &pb.Item{
			Type:  forSend.Type,
			Count: 1,
		},
	})
	if err != nil {
		s.logger.Warnf("Could not send item to %v: %v", adjNode, err)
	}
	forSend.Count -= 1
	s.logger.Infof("Sended item %v. Resp %v.", forSend, r)
	return nil
}
