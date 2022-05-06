package mine

import (
	"context"
	"errors"
	"fmt"
	"github.com/overmesgit/factorio/grpc"
	grpc2 "google.golang.org/grpc"
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

func (s *server) getNextNode() *grpc.Node {
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

func (s *server) sendItem(adjNode *grpc.Node) error {
	s.logger.Infof("Send items. Current store. %v", MyItems.items)

	if adjNode.Ip == "" {
		return errors.New(fmt.Sprintf("Adj node does not have ip %v", adjNode))
	}

	MyItems.Lock()
	defer MyItems.Unlock()
	var forSend *grpc.Item
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

	conn, err := grpc2.Dial(adjNode.Ip+":8080", grpc2.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := grpc.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendResource(ctx, &grpc.ItemRequest{
		Item: &grpc.Item{
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
