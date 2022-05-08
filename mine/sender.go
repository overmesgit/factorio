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
			err := s.sendItemFromStore(adjNode)
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

func (s *server) getPrevNode() *pb.Node {
	nextNode := s.getNextNode()
	prevRowOff, prevColOff := MyNode.Row-nextNode.Row, MyNode.Col-nextNode.Col
	prevRow, prevCol := MyNode.Row+prevRowOff, MyNode.Col+prevColOff
	return &pb.Node{Row: prevRow, Col: prevCol}
}

func (s *server) getNextNode() *pb.Node {
	offset, ok := directionIndex[MyNode.Direction]
	if !ok {
		return nil
	}
	adjRow, adjCol := MyNode.Row+offset[0], MyNode.Col+offset[1]
	return &pb.Node{Row: adjRow, Col: adjCol}
}

func (s *server) sendItemFromStore(adjNode *pb.Node) error {
	sugar.Infof("Send items. Current store. %v forSend %v", MyStorage.GetItemCount())

	var forSend *pb.Item
	if MyNode.Type == string(nodemap.Furnace) {
		forSend = MyStorage.Get(nodemap.IronPlate)
	} else {
		forSend = MyStorage.GetItemForSend()
	}

	if forSend == nil {
		sugar.Infof("Nothing to send.")
		return nil
	}
	err := s.sendItem(adjNode, forSend)
	if err != nil {
		err := MyStorage.Add(nodemap.ItemType(forSend.Type))
		if err != nil {
			sugar.Warnf("Could not stack item back %v %v", forSend, err)
		}
		return err
	}

	return nil
}

func (s *server) sendItem(adjNode *pb.Node, forSend *pb.Item) error {
	sugar.Infof("Send items. Current store. %v forSend %v", MyStorage.GetItemCount(), forSend)

	conn, err := grpc.Dial(fmt.Sprintf("r%vc%v:8080", adjNode.Row, adjNode.Col), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendResource(ctx, forSend)
	if err != nil {
		return err
	}

	sugar.Infof("Sent item %v. Resp %v.", forSend, r)
	return nil
}

func (s *server) askForItem(prevNode *pb.Node, itemType nodemap.ItemType, store bool) (*pb.Item, error) {
	sugar.Infof("Ask for item %v %v", prevNode, itemType)

	conn, err := grpc.Dial(fmt.Sprintf("r%vc%v:8080", prevNode.Row, prevNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GiveResource(ctx, &pb.Item{Type: string(itemType)})
	if err != nil {
		return nil, err
	}
	if store {
		err := MyStorage.Add(nodemap.ItemType(r.Type))
		if err != nil {
			sugar.Warnf("can't store aquared item %v %v", r, err)
		}
	}

	return r, nil
}

func (s *server) askForNeedItem(nextNode *pb.Node) (*pb.Item, error) {
	sugar.Infof("Ask for needed item %v", nextNode)

	conn, err := grpc.Dial(fmt.Sprintf("r%vc%v:8080", nextNode.Row, nextNode.Col),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	defer conn.Close()
	c := pb.NewMineClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return c.NeededResource(ctx, &pb.Empty{})
}
